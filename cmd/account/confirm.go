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
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/howeyc/gopass"
	"github.com/passw0rd/cli/models"

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

			_, url2fa, err := confirmFunc(email, sessionToken, confirmationCode, client)
			if err != nil {
				return err
			}
			fmt.Printf("Your two-factor QR code url (Authy or Google auth):\n%s\nYou will need it to log in.\n", url2fa)
			return nil
		},
	}
}
func confirmFunc(email, sessionToken, confirmationCode string, vcli *client.VirgilHttpClient) (password, qrUrl string, err error) {

	pwd, err := gopass.GetPasswdPrompt("Enter account password:\r\n", false, os.Stdin, os.Stdout)
	if err != nil {
		return
	}
	pwdAgain, err := gopass.GetPasswdPrompt("Again:\r\n", false, os.Stdin, os.Stdout)
	if err != nil {
		return
	}

	if subtle.ConstantTimeCompare(pwd, pwdAgain) != 1 {
		err = errors.New("passwords do not match")
		return
	}

	req := &models.ConfirmAccountRequest{
		Email:                    email,
		ConfirmationSessionToken: sessionToken,
		ConfirmationCode:         confirmationCode,
		Password:                 string(pwd),
	}

	resp := &models.ConfirmAccountResponse{}

	_, err = vcli.Send(http.MethodPost, "", "accounts/v1/confirm-account", req, resp)

	if err != nil {
		return
	}

	if resp != nil {
		return string(pwd), resp.QrUrl, nil
	}

	return "", "", errors.New("empty response")
}
