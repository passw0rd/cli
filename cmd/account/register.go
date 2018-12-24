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
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/passw0rd/cli/utils"

	"github.com/passw0rd/cli/models"

	"github.com/passw0rd/cli/cmd/app"
	"github.com/passw0rd/phe-go"

	"github.com/passw0rd/cli/client"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func Register(client *client.VirgilHttpClient) *cli.Command {
	return &cli.Command{
		Name:      "register",
		Aliases:   []string{"reg"},
		ArgsUsage: "email",
		Usage:     "Registers a new account",
		Action: func(context *cli.Context) error {
			return registerFunc(context, client)
		},
	}
}
func registerFunc(context *cli.Context, vcli *client.VirgilHttpClient) error {

	if context.NArg() < 1 {
		return errors.New("invalid number of arguments")
	}

	email := context.Args().First()

	req := &models.RegisterRequest{Email: email}

	resp := &models.RegisterResponse{}

	_, err := vcli.Send(http.MethodPost, "", "accounts/v1/account", req, resp)

	if err != nil {
		return err
	}

	fmt.Println("Enter confirmation code:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	code := scanner.Text()

	if code == "" {
		fmt.Println("Your session url2fa:", resp.ConfirmationSessionToken)
		return errors.New("you did not enter confirmation code. Try again with confirm command")
	}

	pwd, url2fa, err := confirmFunc(email, resp.ConfirmationSessionToken, code, vcli)

	if err != nil {
		fmt.Println("Your registration token:", resp.ConfirmationSessionToken)
		return errors.Wrap(err, "error while trying to confirm account. Try again with confirm command")
	}

	fmt.Printf("Your two-factor QR code URL (Authy or Google Auth):\n%s\n", url2fa)

	fmt.Println("Would you like to create a new default app and a client secret key right now? [y]")

	scanner.Scan()
	text := scanner.Text()

	if text == "" || text == "y" {

		err := utils.Login(email, pwd, vcli)

		if err != nil {
			return err
		}

		token, err := utils.LoadAccessToken()
		if err != nil {
			return err
		}

		appName := make([]byte, 4)
		_, err = rand.Read(appName)
		if err != nil {
			return err
		}

		pub, appToken, err := app.CreateFunc(token, "My_Default_App_"+hex.EncodeToString(appName), vcli)
		if err != nil {
			return err
		}

		fmt.Println("Your credentials:")
		fmt.Println("public_key:", pub)
		fmt.Println("access_token:", appToken)
		key := phe.GenerateClientKey()
		fmt.Println("secret_key:", "SK.1."+base64.StdEncoding.EncodeToString(key))
		return nil
	}

	if resp != nil {
		fmt.Println("Your registration session url2fa:", resp.ConfirmationSessionToken)
	}

	return nil
}
