<p align="center">
  <a href="https://authorizer.dev">
    <img alt="Logo" src="https://github.com/authorizerdev/authorizer/blob/main/assets/logo.png" width="60" />
  </a>
</p>
<h1 align="center">
  Authorizer
</h1>

**Authorizer** is an open-source authentication and authorization solution for your applications. Bring your database and have complete control over the user information. You can self-host authorizer instances and connect to any SQL database.

## Table of contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Contributing](https://github.com/authorizerdev/authorizer/blob/main/.github/CONTRIBUTING.md)
- [Docs](http://docs.authorizer.dev/)
- [Join Community](https://discord.gg/Zv2D5h6kkK)

# Introduction

<img src="https://github.com/authorizerdev/authorizer/blob/main/assets/authorizer-architecture.png" style="height:20em"/>

#### We offer the following functionality

- ✅ Sign-in / Sign-up with email ID and password
- ✅ Secure session management
- ✅ Email verification
- ✅ APIs to update profile securely
- ✅ Forgot password flow using email
- ✅ Social logins (Google, Github, Facebook, more coming soon)
- ✅ Role-based access management
- ✅ Password-less login with email and magic link

## Roadmap

- Support more JWT encryption algorithms (Currently supporting HS256)
- 2 Factor authentication
- Back office (Admin dashboard to manage user)
- Support more database
- VueJS SDK
- Svelte SDK
- React Native SDK
- Flutter SDK
- Android Native SDK
- iOS native SDK
- Golang SDK
- Python SDK
- PHP SDK
- WordPress plugin
- Kubernetes Helm Chart
- [Local Stack](https://github.com/localstack/localstack)
- AMI
- Digital Ocean Droplet
- Azure
- Render
- Edge Deployment using Fly.io
- Password-less login with mobile number and OTP SMS

# Getting Started

## Trying out Authorizer

This guide helps you practice using Authorizer to evaluate it before you use it in a production environment. It includes instructions for installing the Authorizer server in local or standalone mode.

- [Install using source code](#install-using-source-code)
- [Install using binaries](#install-using-binaries)
- [Install instance on heroku](#install-instance-on-Heroku)
- [Install instance on railway.app](#install-instance-on-railway)

## Install using source code

### Prerequisites

- OS: Linux or macOS or windows
- Go: (Golang)(https://golang.org/dl/) >= v1.15

### Project Setup

1. Fork the [authorizer](https://github.com/authorizerdev/authorizer) repository (**Skip this step if you have access to repo**)
2. `git clone https://github.com/authorizerdev/authorizer.git`
3. `cd authorizer`
4. `cp .env.sample .env`. Check all the supported env [here](https://docs.authorizer.dev/core/env/)
5. Build the code `make clean && make`
   > Note: if you don't have [`make`](https://www.ibm.com/docs/en/aix/7.2?topic=concepts-make-command), you can `cd` into `server` dir and build using the `go build` command
6. Run binary `./build/server`

## Install using binaries

Deploy / Try Authorizer using binaries. With each [Authorizer Release](https://github.com/authorizerdev/authorizer/releases)
binaries are baked with required deployment files and bundled. You can download a specific version of it for the following operating systems:

- Mac OSX
- Linux

### Step 1: Download and unzip bundle

- Download the Bundle for the specific OS from the [release page](https://github.com/authorizerdev/authorizer/releases)

> Note: For windows, we recommend running using docker image to run authorizer.

- Unzip using following command

  - Mac / Linux

  ```sh
  tar -zxf AUTHORIZER_VERSION -c authorizer
  ```

- Change directory to `authorizer`

  ```sh
  cd authorizer
  ```

### Step 2: Configure environment variables

Required environment variables are pre-configured in `.env` file. But based on the production requirements, please configure more environment variables. You can refer to [environment variables docs](/core/env) for more information.

### Step 3: Start Authorizer

- Run following command to start authorizer

  - For Mac / Linux users

  ```sh
  ./build/server
  ```

> Note: For mac users, you might have to give binary the permission to execute. Here is the command you can use to grant permission `xattr -d com.apple.quarantine build/server`

## Install instance on Heroku

Deploy Authorizer using [heroku](https://github.com/authorizerdev/authorizer-heroku) and quickly play with it in 30seconds
<br/><br/>
[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/authorizerdev/authorizer-heroku)

# Install instance on railway

Deploy production ready Authorizer instance using [railway.app](https://github.com/authorizerdev/authorizer-railway) with postgres and redis for free and build with it in 30seconds
<br/>

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template?template=https%3A%2F%2Fgithub.com%2Fauthorizerdev%2Fauthorizer-railway&plugins=postgresql%2Credis&envs=ENV%2CDATABASE_TYPE%2CADMIN_SECRET%2CCOOKIE_NAME%2CJWT_ROLE_CLAIM%2CJWT_TYPE%2CJWT_SECRET%2CFACEBOOK_CLIENT_ID%2CFACEBOOK_CLIENT_SECRET%2CGOOGLE_CLIENT_ID%2CGOOGLE_CLIENT_SECRET%2CGITHUB_CLIENT_ID%2CGITHUB_CLIENT_SECRET%2CALLOWED_ORIGINS%2CROLES%2CPROTECTED_ROLES%2CDEFAULT_ROLES&optionalEnvs=FACEBOOK_CLIENT_ID%2CFACEBOOK_CLIENT_SECRET%2CGOOGLE_CLIENT_ID%2CGOOGLE_CLIENT_SECRET%2CGITHUB_CLIENT_ID%2CGITHUB_CLIENT_SECRET%2CALLOWED_ORIGINS%2CROLES%2CPROTECTED_ROLES%2CDEFAULT_ROLES&ENVDesc=Deployment+environment&DATABASE_TYPEDesc=With+railway+we+are+deploying+postgres+db&ADMIN_SECRETDesc=Secret+to+access+the+admin+apis&COOKIE_NAMEDesc=Name+of+http+only+cookie+that+will+be+used+as+session&FACEBOOK_CLIENT_IDDesc=Facebook+client+ID+for+facebook+login&FACEBOOK_CLIENT_SECRETDesc=Facebook+client+secret+for+facebook+login&GOOGLE_CLIENT_IDDesc=Google+client+ID+for+google+login&GOOGLE_CLIENT_SECRETDesc=Google+client+secret+for+google+login&GITHUB_CLIENT_IDDesc=Github+client+ID+for+github+login&GITHUB_CLIENT_SECRETDesc=Github+client+secret+for+github+login&ALLOWED_ORIGINSDesc=Whitelist+the+URL+for+which+this+instance+of+authorizer+is+allowed&ROLESDesc=Comma+separated+list+of+roles+that+platform+supports.+Default+role+is+user&PROTECTED_ROLESDesc=Comma+separated+list+of+protected+roles+for+which+sign-up+is+disabled&DEFAULT_ROLESDesc=Default+role+that+should+be+assigned+to+user.+It+should+be+one+from+the+list+of+%60ROLES%60+env.+Default+role+is+user&JWT_ROLE_CLAIMDesc=JWT+key+to+be+used+to+validate+the+role+field.&JWT_TYPEDesc=JWT+encryption+type&JWT_SECRETDesc=Random+string+that+will+be+used+for+encrypting+the+JWT+token&ENVDefault=PRODUCTION&DATABASE_TYPEDefault=postgres&COOKIE_NAMEDefault=authorizer&JWT_TYPEDefault=HS256&JWT_ROLE_CLAIMDefault=role)

### Things to consider

- For social logins, you will need respective social platform key and secret
- For having verified users, you will need an SMTP server with an email address and password using which system can send emails. The system will send a verification link to an email address. Once an email is verified then, only able to access it.
  > Note: One can always disable the email verification to allow open sign up, which is not recommended for production as anyone can use anyone's email address 😅
- For persisting user sessions, you will need Redis URL (not in case of railway.app). If you do not configure a Redis server, sessions will be persisted until the instance is up or not restarted. For better response time on authorization requests/middleware, we recommend deploying Redis on the same infra/network as your authorizer server.

## Integrating into your website

This example demonstrates how you can use [`@authorizerdev/authorizer-js`](/authorizer-js/getting-started) CDN version and have login ready for your site in few seconds. You can also use the ES module version of [`@authorizerdev/authorizer-js`](/authorizer-js/getting-started) or framework-specific versions like [`@authorizerdev/authorizer-react`](/authorizer-react/getting-started)

### Copy the following code in `html` file

> **Note:** Change AUTHORIZER_URL in the below code with your authorizer URL. Also, you can change the logout button component

```html
<script src="https://unpkg.com/@authorizerdev/authorizer-js/lib/authorizer.min.js"></script>

<script type="text/javascript">
	const authorizerRef = new authorizerdev.Authorizer({
		authorizerURL: `AUTHORIZER_URL`,
		redirectURL: window.location.origin,
	});

	// use the button selector as per your application
	const logoutBtn = document.getElementById('logout');
	logoutBtn.addEventListener('click', async function () {
		await authorizerRef.logout();
		window.location.href = '/';
	});

	async function onLoad() {
		const res = await authorizerRef.browserLogin();
		if (res && res.user) {
			// you can use user information here, eg:
			/**
      const userSection = document.getElementById('user');
      const logoutSection = document.getElementById('logout-section');
      logoutSection.classList.toggle('hide');
      userSection.innerHTML = `Welcome, ${res.user.email}`;
      */
		}
	}
	onLoad();
</script>
```

---

### Support my work

<a href="https://www.buymeacoffee.com/lakhansamani" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>
