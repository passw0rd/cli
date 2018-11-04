/*
 * Copyright (C) 2015-2018 Virgil Security Inc.
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     (1) Redistributions of source code must retain the above copyright
 *     notice, this list of conditions and the following disclaimer.
 *
 *     (2) Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in
 *     the documentation and/or other materials provided with the
 *     distribution.
 *
 *     (3) Neither the name of the copyright holder nor the names of its
 *     contributors may be used to endorse or promote products derived from
 *     this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR ''AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 * Lead Maintainer: Virgil Security Inc. <support@virgilsecurity.com>
 */

package main

import (
	"log"
	"os"

	"github.com/passw0rd/cli/client"

	"github.com/passw0rd/cli/cmd"

	"gopkg.in/urfave/cli.v2"
)

var Version string

func main() {

	vcli := &client.VirgilHttpClient{
		Address: "https://api.passw0rd.io/",
	}

	app := &cli.App{
		Version:     Version,
		Description: "password.io client application",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"t"},
				Usage:   "Auth token",
				EnvVars: []string{"PASSW0RD_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "appid",
				Usage:   "App ID",
				EnvVars: []string{"PASSW0RD_APP_ID"},
			},
			&cli.StringFlag{
				Name:    "public_key",
				Aliases: []string{"pk"},
				Usage:   "Service public key",
				EnvVars: []string{"PASSW0RD_PUB"},
			},
			&cli.StringFlag{
				Name:    "secret_key",
				Aliases: []string{"sk"},
				Usage:   "Client private key",
				EnvVars: []string{"PASSW0RD_SECRET"},
			},
		},
		Commands: []*cli.Command{
			cmd.Account(vcli),
			cmd.Application(vcli),
			cmd.Keygen(),
			cmd.Demo(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
