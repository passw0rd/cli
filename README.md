# Passw0rd CLI

[![Production](https://travis-ci.org/passw0rd/cli.svg?branch=master)](https://travis-ci.org/passw0rd/cli)
[![GitHub license](https://img.shields.io/badge/license-BSD%203--Clause-blue.svg)](https://github.com/VirgilSecurity/virgil/blob/master/LICENSE)

[Installation](#installation) | [Launching CLI](#launching-cli) | [Commands](#commands) | [Usage Examples](#usage-examples) | [Support](#support)

<a href="https://passw0rd.io/"><img width="260px" src="https://cdn.virgilsecurity.com/assets/images/github/logos/passw0rd.png" align="left" hspace="0" vspace="0"></a>[Virgil Security](https://virgilsecurity.com) introduces to developers a **Passw0rd CLI** – an open source tool that provides commands for interacting with the [Passw0rd Service](https://passw0rd.io/). With minimal configuration, you can start using all of the functionality provided by the Passw0rd from your favorite terminal program.
- **Linux shells** – Use common shell programs such as Bash, Zsh, and tsch to run commands in Linux, macOS, or Unix.
- **Windows command line** – On Microsoft Windows, run commands in either PowerShell or the Windows Command Processor.


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

## Commands
Using the passw0rd CLI you can:
  * register and manage your **FREE** account
  * register and manage your passw0rd application
  * get your passw0rd App's credentials, such as: App Token, Server Public Key, Account Secret Key.
  * get your Application Token
  * try passw0rd Demo

To get more information, run the passw0rd CLI or its command with the `--help` or `-h` option that displays full help list and available commands.


## Usage Example
The passw0rd CLI has the following usage syntax:
`passw0rd [global options] command [command options] [arguments...]`

The examples below are given for the FreeBSD/Linux/Mac OS platforms.

#### Register Your Account and Set Up 2FA 

**First**, register your account with the following comand:
```bash
./passw0rd account register my@email.com
```
You have to confirm your account by entering a confirmation code you got in the email and create your own account password.

> !!! Once you've confirmed your account and created an account password, the password CLI immediately asks you to set up two factor authentication and to create a new passw0rd App (with a default App's name) and a Secret Key. You can:
- accept the offer and get all the necessary credentials (app_token, public_key, secret_key) to start working with passw0rd service or
- come back later. If you choose this option, you get only your Session url2fa - store it somewhere in a safe place. Other credentials (app_token, public_key, secret_key) you'll be able to generate using your app_token later.

**Second**, set up two factor authentication (2FA).

To set up 2FA you have to:
- Copy a QR link that you got in passw0rd CLI
- Open the QR link in a brouser or some application 
- Scan the QR using your 2FA application
- Enter a 2FA code into a passw0rd CLI

> To set up 2FA you have to use an application that generates secure 2 step verification tokens on your device. For an exanple, you can download and install [Google Authenticator](https://support.google.com/accounts/answer/1066447?co=GENIE.Platform%3DAndroid&hl=en) or [Authy](https://authy.com/download/).  

As a result you get your passw0rd's application credentials:
- application `name`
- application `app_token`
- server `service_public_key`
- client `app_secret_key`


#### Create a new passw0rd's application

To create a new passw0rd application:
- be sure you're logged in your account. To log in the account use the following coomand (2FA is required): 

```
./passw0rd login my@email.com
```
- then, use the `create` command:
```bash
./passw0rd application create my_new_passw0rd_app
```

where:
- my_new_passw0rd_app - is a name of your new passw0rd application.

#### Generate a new client app_secret_key

```bash
./passw0rd keygen
```

>! Remember, it's not possible to restore a secret key at all.

#### Get an update_token
```bash
./passw0rd --access_token 8Bw0003U000000000000000q6n5dKIlz --app_id 857111111111111111111c app rotate
```

where:
- 8Bw0003U000000000000000q6n5dKIlz - is an access token that you got at registration step (it's not possible to restore or generate a new access token in this version);
- 857111111111111111111c - is an ID of your application.


Use next two examples to try out as a demo the Passw0rd technologies and see how the functions work without setting up the SDK. 



## Passw0rd Demo

Use next examples to try out Passw0rd technologies and see how the functions work without setting up a passw0rd SDK. 

#### Enroll User passw0rd
The demo command allows you to create user's passwOrd record:

```bash
./passw0rd --config passw0rd.yaml demo enroll user_password 
```
where:
- passw0rd.yaml - a config file that contains your account credentials: access_token, app_id, public_key, secret_key. This file is not created by default. So, create passw0rd.yaml file, paste your account credentials into it and specify the pass to it.
- user_password - user password that he or she uses to sign in to your server side. 
- user_passw0rd_record - database passw0rd's record that is associated with the user.


#### Verify User password
The demo command allows you to verify user password:
```bash
./passw0rd --config passw0rd.yaml demo verify user_password user_passw0rd_record
```

where:
- passw0rd.yaml - a config file that contains your account credentials: access_token, app_id, public_key, secret_key. This file is not created by default. So, create passw0rd.yaml file, paste your account credentials into it and specify the pass to it.
- user_password - user password that he or she uses to sign in to your server side. 
- user_passw0rd_record - database passw0rd's record that is associated with the user.


## License
See [LICENSE](https://github.com/VirgilSecurity/virgil-cli/tree/master/LICENSE) for details.

## Support
Our developer support team is here to help you. Find out more information on our [Help Center](https://help.virgilsecurity.com/).

Also, get extra help from our support team: support@VirgilSecurity.com.
