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

package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/howeyc/gopass"
	"github.com/passw0rd/cli/client"
	"github.com/passw0rd/cli/models"
)

//Login obtains temporary account access token. Email and password may be empty
func Login(email, password string, vcli *client.VirgilHttpClient) (err error) {

	if email == "" {
		fmt.Println("Enter your email:")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		email = scanner.Text()
	}

	if password == "" {
		pwd, err := gopass.GetPasswdPrompt("Enter account password:\n", false, os.Stdin, os.Stdout)
		if err != nil {
			return err
		}
		password = string(pwd)
	}

	code, err := gopass.GetPasswdPrompt("Enter 2-factor code:\n", true, os.Stdin, os.Stdout)
	if err != nil {
		return
	}

	req := &models.LoginRequest{
		Email:    email,
		Password: password,
		Totp:     string(code),
	}

	resp := &models.LoginResponse{}

	_, err = vcli.Send(http.MethodPost, "", "accounts/v1/login", req, resp)

	if err != nil {
		return err
	}

	return SaveAccessToken(resp.AccountToken)
}
