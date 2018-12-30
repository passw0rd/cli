# Passw0rd CLI

[![Production](https://travis-ci.org/passw0rd/cli.svg?branch=master)](https://travis-ci.org/passw0rd/cli)
[![GitHub license](https://img.shields.io/badge/license-BSD%203--Clause-blue.svg)](https://github.com/VirgilSecurity/virgil/blob/master/LICENSE)

<a href="https://passw0rd.io/"><img width="260px" src="https://cdn.virgilsecurity.com/assets/images/github/logos/passw0rd.png" align="left" hspace="0" vspace="0"></a>[Virgil Security](https://virgilsecurity.com) introduces to developers a **Passw0rd CLI** – an open source tool that provides commands for interacting with the [Passw0rd Service](https://passw0rd.io/). With minimal configuration, you can start using all of the functionality provided by the Passw0rd from your favorite terminal program.
- **Linux shells** – Use common shell programs such as Bash, Zsh, and tsch to run commands in Linux, macOS, or Unix.
- **Windows command line** – On Microsoft Windows, run commands in either PowerShell or the Windows Command Processor.

## Content
- [Installation](#installation)
- [Launching CLI](#launching-cli)
- [Features](#features)
- [Register your account and set up 2FA](#register-your-account-and-set-up-2FA)
- [Commands usage](#commands-usage)
  - [Confirm an account](#confirm-an-account)
  - [Log in the account](#log-in-the-account)
  - [Log out the account](#log-out-the-account)
  - [Create a new application](#create-a-new-application)
  - [Get application list](#get-application-list)
  - [Get an update token](#get-an-update-token)
  - [Fetch an update token](#fetch-an-update-token)
  - [Delete an update token](#delete-an-update-token)
  - [Update keys](#update-keys)
  - [Generate a secret key](#generate-a-secret-key)
- [Passw0rd Demo](#passw0rd-demo)
  - [Enroll user passw0rd record](#enroll-user-passw0rd-record)
  - [Verify user password](#verify-user-password)
  - [Update user passw0rd record](#update-user-passw0rd-record)
- [License](#license)
- [Support](#support)

## Installation
The passw0rd CLI is provided as a binary file, and it is available for Mac OS, FreeBSD,  Linux OS and Windows OS. Download the latest CLI package here: https://github.com/passw0rd/cli/releases.


## Launching CLI

#### FreeBSD / Linux / Mac OS
Run the passw0rd CLI with the following command:
```bash
./passw0rd
```
> or use `sudo ./passw0rd` when you need to run the command as an administrator

#### Windows OS
Run the passw0rd CLI with the following command:
```bash
passw0rd.exe
# or just `passw0rd`
```

## Features
Using the passw0rd CLI you can:
  * register and manage your **FREE** passw0rd account
  * register and manage your passw0rd application
  * get your passw0rd application credentials, such as: Application Token, Service Public Key, Application Secret Key
  * try a passw0rd Demo

To get more information, run the passw0rd CLI or its command with the `--help` or `-h` option that displays full help list and available commands.


## Register your account and set up 2FA

**First**, register your account with the following command:
```bash
./passw0rd account register my@email.com
```
You have to confirm your account by entering a confirmation code you got in the email and create your account password.

> !!! Once you've confirmed your account and created an account password, the passw0rd CLI immediately asks you to set up two-factor authentication and to create a new passw0rd App (with a default application name) and an `app_secret_Key`. Accept the offer to get all the necessary credentials (app_token, service_public_key, app_secret_key) to start working with passw0rd service.

**Second**, set up two-factor authentication (2FA):
- Copy a QR link that you got in a passw0rd CLI
- Open the QR link in a browser or some application
- Scan the QR using your 2FA application
- Enter a 2FA code into a passw0rd CLI

> To set up 2FA you have to use an application that generates secure 2 step verification tokens on your device. For example, you can download and install [Google Authenticator](https://support.google.com/accounts/answer/1066447?co=GENIE.Platform%3DAndroid&hl=en) or [Authy](https://authy.com/download/).  

As a result, you get your passw0rd's application credentials:
- application `name`
- application `app_token`
- server `service_public_key`
- client `app_secret_key`

## Commands usage

The passw0rd CLI has the following usage syntax:
`passw0rd [global options] command [command options] [arguments...]`

The examples below are given for the FreeBSD/Linux/Mac OS platforms.

### Confirm an account
This command allows you to confirm your passw0rd account in case you didn't confirm it at the registration step.

The command has the following structure:
```bash
./passw0rd account confirm my@email.com  <confirmation_code>
```

where `confirmation_code` is a code you got in a confirmation email at the registration step.

### Log in the account
To log into your prefer account use the following command structure:
```bash
./passw0rd login my@email.com
```

### Log out the account
To log out the account use the following command structure:
```bash
./passw0rd logout my@email.com
```

### Create a new application
To create a new passw0rd application:
- be sure you're logged in your account. To log in the account use the following command (2FA is required):

```bash
./passw0rd login my@email.com
```
- then, use the `create` command:
```bash
./passw0rd application create my_new_passw0rd_app
```

where:
- my_new_passw0rd_app - is the name of your new passw0rd application.

### Get application list
To show all your registered applications use the following command:
```bash
./passw0rd application list
```

### Get an update token
An update token is used to update a user's passw0rd record in your database and to get a new `app_secret_key` and `service_public_key` of a specific application.

To get an update token:
- be sure you're logged in your account. To log in the account use the following command (2FA is required):

```bash
./passw0rd login my@email.com
```

- then, use the `rotate` command;

```bash
./passw0rd application rotate <app_token>
```

where:
- <app_access_token> - is your application token.

as a result, you get your `update_token`.

### Fetch an update token
In case you forgot your `update_token` you can fetch the latest one from the Passw0rd service using the following command:

```bash
./passw0rd application fetch <app_token>
```

### Delete an update token
Delete the latest `update_token` available for current application:

```bash
./passw0rd application delete-update-token <app_token>
```

### Update keys
This command is used to update the `app_secret_key` and `service_public_key` of a specific application

```bash
./passw0rd application update-keys <service_public_key> <app_secret_key> <update_token>
```

### Generate a secret key
This command is used to generate a new `app_secret_key`:
```bash
./passw0rd keygen
```

## Passw0rd Demo
Passw0rd CLI provides you with a Demo mode that allows you to try out passw0rd technologies and see how the functions work without setting up a passw0rd SDK.

To start working with a passw0rd Demo you need to have a registed passw0rd account and created application.

#### Enroll user passw0rd record
The demo command allows you to create user's passw0rd record:

```bash
./passw0rd --config passw0rd.yaml demo enroll user_password
```
where:
- passw0rd.yaml - a config file that contains your account credentials: app_token, service_public_key, app_secret_key. This file is not created by default. So, create passw0rd.yaml file, paste your account credentials into it and specify the pass to it.
- user_password - user password that he or she uses to sign in to your server side.

as a result, you get:
- encryption key - secret key, that can be used to encrypt user data (for example, photos)
- record - database passw0rd's record that is associated with the user.


#### Verify user password
The demo command allows you to verify user password:
```bash
./passw0rd --config passw0rd.yaml demo verify user_password user_passw0rd_record
```

where:
- passw0rd.yaml - a config file that contains your account credentials: app_token, app_id, service_public_key, app_secret_key. This file is not created by default. So, create passw0rd.yaml file, paste your account credentials into it and specify the pass to it.
- user_password - user password that he or she uses to sign in to your server side.
- user_passw0rd_record - database passw0rd's record that is associated with the user.

As a result, you get an encryption key and information whether the password is correct or not.

#### Update user passw0rd record
This function allows you to use a special `update_token` to update the passw0rd record in your database.

Use this flow only if your database has been COMPROMISED! When a user only needs to change his or her own password, use the `enroll` function (step 5) to replace the user's old `record` value in your database.

to update user's `passw0rd record`:
- get [your `update_token` using passw0rd CLI](https://github.com/passw0rd/cli#get-an-update_token)
- then use the `update token` function to create a new password_record for your users (you don't need to ask your users to create a new password because the original password is not changing, just the protected record of it in the passw0rd system).
- then update the `record` with the following command:
```bash
./passw0rd --config passw0rd.yaml demo update user_passw0rd_record update_token
```

where:
- passw0rd.yaml - a config file that contains your account credentials: app_token, service_public_key, app_secret_key. This file is not created by default. So, create passw0rd.yaml file, paste your account credentials into it and specify the pass to it.
- user_passw0rd_record - database passw0rd's record that is going to be updated.
- update_token - update token that you got using the update_token command.

As a result, you get an **updated user's passw0rd record**.

Then, you have to update the `app_secret_key` and `service_public_key` of your application

```bash
./passw0rd application update-keys <service_public_key> <app_secret_key> <update_token>
```
As a result, you get **`app_secret_key` and `service_public_key` of your application**.

So, now upgrade the passw0rd.yaml file with your new application credentials and use the [verify user password](https://github.com/passw0rd/cli#verify-user-password) step to check whether the password is correct or not.

## License
See [LICENSE](https://github.com/VirgilSecurity/virgil-cli/tree/master/LICENSE) for details.

## Support
Our developer support team is here to help you. Find out more information on our [Help Center](https://help.virgilsecurity.com/).

Also, get extra help from our support team: support@VirgilSecurity.com.
