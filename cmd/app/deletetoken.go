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

package app

import (
	"fmt"
	"net/http"

	"github.com/passw0rd/cli/client"
	"github.com/passw0rd/cli/utils"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func DeleteUpdateToken(vcli *client.VirgilHttpClient) *cli.Command {
	return &cli.Command{
		Name:      "delete-update-token",
		Aliases:   []string{"dt"},
		ArgsUsage: "dt <app_access_token>",
		Usage:     "Deletes the latest update token available for this app",
		Action: func(context *cli.Context) (err error) {

			if context.NArg() < 1 {
				return errors.New("invalid number of arguments")
			}

			token, err := LoadAccessTokenOrLogin(vcli)

			if err != nil {
				return err
			}

			for err == nil {
				err = deleteUpdateTokenFunc(token, context.Args().First(), vcli)
				if err == nil {
					break
				}
				token, err = utils.CheckRetry(err, vcli)
			}

			if err == nil {
				fmt.Println("Update token delete ok.")
			}

			return err
		},
	}
}

func deleteUpdateTokenFunc(token, appToken string, vcli *client.VirgilHttpClient) (err error) {

	_, err = vcli.Send(http.MethodDelete, token, "phe/v1/delete-update-token", nil, nil, appToken)

	return err
}
