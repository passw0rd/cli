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
	"github.com/passw0rd/cli/models"
	"github.com/passw0rd/cli/utils"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func List(vcli *client.VirgilHttpClient) *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Lists your apps",
		Action: func(context *cli.Context) (err error) {

			token, err := LoadAccessTokenOrLogin(vcli)

			if err != nil {
				return err
			}

			var apps []*models.Application
			for err == nil {
				apps, err = listFunc(token, vcli)
				if err == nil {
					break
				}
				token, err = utils.CheckRetry(err, vcli)
			}

			if err != nil {
				return err
			}

			for i, app := range apps {
				if i > 0 {
					fmt.Println()
				}
				fmt.Printf("=====%s=====\n", app.Name)
				fmt.Println("ID:", app.Id)
				fmt.Println("Public key:", app.PublicKey)
				fmt.Println("Access token:", app.AppToken)
			}
			return nil
		},
	}
}

func listFunc(token string, vcli *client.VirgilHttpClient) (apps []*models.Application, err error) {

	resp := &models.AccountApplicationsResponse{}

	_, err = vcli.Send(http.MethodGet, token, "accounts/v1/account/applications", nil, resp)

	if err != nil {
		return
	}

	if resp != nil {
		return resp.App, nil
	}

	return nil, errors.New("empty response")
}
