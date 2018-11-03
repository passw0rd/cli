# Passw0rd CLI

[![Production](https://travis-ci.org/passw0rd/cli.svg?branch=master)](https://travis-ci.org/passw0rd/cli)
[![GitHub license](https://img.shields.io/badge/license-BSD%203--Clause-blue.svg)](https://github.com/VirgilSecurity/virgil/blob/master/LICENSE)

[Installation](#installation) | [Commands](#commands) | [Using Example](#using-example) | [Support](#support)

<a href="https://passw0rd.io/"><img width="260px" src="https://cdn.virgilsecurity.com/assets/images/github/logos/passw0rd.png" align="left" hspace="0" vspace="0"></a>[Virgil Security](https://virgilsecurity.com) introduces to developers a Passw0rd CLI is an open source tool that provides commands for interacting with Passw0rd Service and Virgil Cloud. With minimal configuration, you can start using all of the functionality provided by the Passw0rd from your favorite terminal program.
- **Linux shells** – Use common shell programs such as Bash, Zsh, and tsch to run commands in Linux, macOS, or Unix.
- **Windows command line** – On Microsoft Windows, run commands in either PowerShell or the Windows Command Processor.


## Installation

The Passw0rd CLI is provided as a binary file, and it is available for Mac OS, Linux OS and Windows OS.


### Mac / Linux OS
Run the Passw0rd CLI with following command:
```bash
chmod +x ./cli
```
> or use `sudo chmod +x ./cli` when you need to run the command as an administrator

### Windows OS
Run the Passw0rd CLI with following command:
```bash
./cli
```


## Commands

Using the Passw0rd CLI you can:
  * register and manage your **FREE** Account at Virgil Cloud
  * register and manage your Passw0rd Project
  * get your Passw0rd Project's credentials, such as: App ID, API Key, Server Public Key, Client Secret Key.
  * get your access token

To get more information run the Passw0rd CLI or its command with the `--help` or `-h` option, that displays full help list and available commands.

> consequently, to get a help the run the following command: `./cli --help`. To get help using a command use the `--help` or `-h` option after the command, for example: `./cli account --help`.

## Using Example
The Passw0rd CLI has the following usage syntax:
`cli [global options] command [command options] [arguments...]`

### Register your account
```bash
./cli account register my@email.com
```

### Register Passw0rd's project
```bash
./cli application create MyApplication
```

### Get your updateToken
```bash
./cli --token
```

## License

See [LICENSE](https://github.com/VirgilSecurity/virgil-cli/tree/master/LICENSE) for details.

## Support
Our developer support team is here to help you. Find out more information on our [Help Center](https://help.virgilsecurity.com/).

Also, get extra help from our support team: support@VirgilSecurity.com.
