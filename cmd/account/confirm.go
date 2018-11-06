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

package account

import (
	"fmt"
	"log"
	"net/http"

	"github.com/passw0rd/cli/client"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func Confirm(client *client.VirgilHttpClient) *cli.Command {
	return &cli.Command{
		Name:      "confirm",
		Aliases:   []string{"c"},
		ArgsUsage: "email session_token confirmation_code",
		Usage:     "Registers a new account",
		Action: func(context *cli.Context) error {

			if context.NArg() < 3 {
				return errors.New("invalid number of arguments")
			}

			email := context.Args().First()
			sessionToken := context.Args().Get(1)
			confirmationCode := context.Args().Get(2)

			if email == "" {
				log.Fatal("please specify email")
			}
			if sessionToken == "" {
				log.Fatal("please specify session token")
			}
			if confirmationCode == "" {
				log.Fatal("please specify confirmation code")
			}

			token, err := confirmFunc(email, sessionToken, confirmationCode, client)
			if err != nil {
				return err
			}
			fmt.Println("Your access token:", token)
			return nil
		},
	}
}
func confirmFunc(email, sessionToken, confirmationCode string, vcli *client.VirgilHttpClient) (string, error) {

	req := &ConfirmAccountReq{
		Email: email,
		ConfirmationSessionToken: sessionToken,
		ConfirmationCode:         confirmationCode,
	}

	var resp *ConfirmAccountResp

	_, err := vcli.Send(http.MethodPost, "", "accounts/v1/confirm-account", req, &resp)

	if err != nil {
		return "", err
	}

	if resp != nil {
		return resp.AccessToken, nil
	}

	return "", errors.New("empty response")
}
