/*
 * Copyright (C) 2015-2019 Virgil Security Inc.
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
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/passw0rd/cli/client"
	"github.com/passw0rd/cli/models"
	"github.com/passw0rd/cli/utils"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func FetchUpdateToken(vcli *client.VirgilHttpClient) *cli.Command {
	return &cli.Command{
		Name:      "fetch-update-token",
		Aliases:   []string{"ft"},
		ArgsUsage: "app_token",
		Usage:     "Fetches the latest update token available for this app",
		Action: func(context *cli.Context) (err error) {

			if context.NArg() < 1 {
				return errors.New("invalid number of arguments")
			}

			token, err := LoadAccessTokenOrLogin(vcli)

			if err != nil {
				return err
			}

			for err == nil {
				err = fetchFunc(token, context.Args().First(), vcli)
				if err == nil {
					break
				}
				token, err = utils.CheckRetry(err, vcli)
			}

			return err
		},
	}
}

func fetchFunc(token, appToken string, vcli *client.VirgilHttpClient) (err error) {

	resp := &models.UpdateTokenResponse{}

	_, err = vcli.Send(http.MethodGet, token, "phe/v1/fetch-update-token", nil, resp, appToken)

	if err != nil {
		return
	}

	if resp != nil && resp.Version > 1 {
		fmt.Printf("Your update token:\nUT.%d.%s\n", resp.Version, base64.StdEncoding.EncodeToString(resp.UpdateToken))
		return nil
	}

	return errors.New("no update token found.")
}
