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
  * get your passw0rd App's credentials, such as: App ID, Access token, Server Public Key, Client Secret Key.
  * get your Access Token

To get more information, run the passw0rd CLI or its command with the `--help` or `-h` option that displays full help list and available commands.


## Usage Examples
The passw0rd CLI has the following usage syntax:
`passw0rd [global options] command [command options] [arguments...]`

The examples below are given for the FreeBSD/Linux/Mac OS platforms.

#### Register Your Account
```bash
./passw0rd account register my@email.com
```
Then, you have to confirm your account by entering a confirmation code you got in the email. Now, your account is confirmed and registered.

> !!! Once you've confirmed your account, the password CLI immediately asks you to create a new passw0rd App (with a default App's name) and a Secret Key. You can:
- accept the offer and get all the necessary credentials (access_token, app_id, public_key, secret_key) to start working with passw0rd service or
- come back later. If you choose this option, you get only your Access Token - store it somewhere in a safe place. Other credentials (app_id, public_key, secret_key) you'll be able to generate using your access_token later.

Remember, if you lose your access token it won't be possible to restore it or generate a new one in this version of CLI.



#### Register passw0rd's Application
```bash
./passw0rd --access_token 8Bw0003U000000000000000q6n5dKIlz application create my_passw0rd_app
```

where:
- 8Bw0003U000000000000000q6n5dKIlz - is an access token that you got at registration step (it's not possible to restore or generate a new access token in this version);
- my_passw0rd_app - is a name of your application.

#### Generate a new Secret Key
```bash
./passw0rd keygen
```

>! Remember, it's not possible to restore a secret key at all.

#### Get an Update Token
```bash
./passw0rd --access_token 8Bw0003U000000000000000q6n5dKIlz --app_id 857111111111111111111c app rotate
```

where:
- 8Bw0003U000000000000000q6n5dKIlz - is an access token that you got at registration step (it's not possible to restore or generate a new access token in this version);
- 857111111111111111111c - is an ID of your application.


Use next two examples to try out as a demo the Passw0rd technologies and see how the functions work without setting up the SDK. 

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
