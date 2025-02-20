package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/authorizerdev/authorizer/server/constants"
	"github.com/authorizerdev/authorizer/server/db"
	"github.com/authorizerdev/authorizer/server/enum"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/robertkrimen/otto"
)

func CreateAuthToken(user db.User, tokenType enum.TokenType, roles []string) (string, int64, error) {
	t := jwt.New(jwt.GetSigningMethod(constants.JWT_TYPE))
	expiryBound := time.Hour
	if tokenType == enum.RefreshToken {
		// expires in 1 year
		expiryBound = time.Hour * 8760
	}

	expiresAt := time.Now().Add(expiryBound).Unix()

	customClaims := jwt.MapClaims{
		"exp":                    expiresAt,
		"iat":                    time.Now().Unix(),
		"token_type":             tokenType.String(),
		"email":                  user.Email,
		"id":                     user.ID,
		"allowed_roles":          strings.Split(user.Roles, ","),
		constants.JWT_ROLE_CLAIM: roles,
	}

	// check for the extra access token script

	accessTokenScript := os.Getenv("CUSTOM_ACCESS_TOKEN_SCRIPT")
	if accessTokenScript != "" {
		userInfo := map[string]interface{}{
			"id":            user.ID,
			"email":         user.Email,
			"firstName":     user.FirstName,
			"lastName":      user.LastName,
			"image":         user.Image,
			"roles":         strings.Split(user.Roles, ","),
			"signUpMethods": strings.Split(user.SignupMethod, ","),
		}

		vm := otto.New()
		userBytes, _ := json.Marshal(userInfo)
		claimBytes, _ := json.Marshal(customClaims)

		vm.Run(fmt.Sprintf(`
			var user = %s;
			var tokenPayload = %s;
			var customFunction = %s;
			var functionRes = JSON.stringify(customFunction(user, tokenPayload));
		`, string(userBytes), string(claimBytes), accessTokenScript))

		val, err := vm.Get("functionRes")

		if err != nil {
			log.Println("error getting custom access token script:", err)
		} else {
			extraPayload := make(map[string]interface{})
			err = json.Unmarshal([]byte(fmt.Sprintf("%s", val)), &extraPayload)
			if err != nil {
				log.Println("error converting accessTokenScript response to map:", err)
			} else {
				for k, v := range extraPayload {
					customClaims[k] = v
				}
			}
		}
	}

	t.Claims = customClaims

	token, err := t.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		return "", 0, err
	}

	return token, expiresAt, nil
}

func GetAuthToken(gc *gin.Context) (string, error) {
	token, err := GetCookie(gc)
	if err != nil || token == "" {
		// try to check in auth header for cookie
		log.Println("cookie not found checking headers")
		auth := gc.Request.Header.Get("Authorization")
		if auth == "" {
			return "", fmt.Errorf(`unauthorized`)
		}

		token = strings.TrimPrefix(auth, "Bearer ")
	}
	return token, nil
}

func VerifyAuthToken(token string) (map[string]interface{}, error) {
	var res map[string]interface{}
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SECRET), nil
	})
	if err != nil {
		return res, err
	}

	// claim parses exp & iat into float 64 with e^10,
	// but we expect it to be int64
	// hence we need to assert interface and convert to int64
	intExp := int64(claims["exp"].(float64))
	intIat := int64(claims["iat"].(float64))

	data, _ := json.Marshal(claims)
	json.Unmarshal(data, &res)
	res["exp"] = intExp
	res["iat"] = intIat

	return res, nil
}
