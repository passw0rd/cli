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
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"github.com/passw0rd/sdk-go"

	"github.com/passw0rd/cli/client"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func Rotate(client *client.VirgilHttpClient) *cli.Command {
	return &cli.Command{
		Name:      "rotate",
		Aliases:   []string{"r"},
		ArgsUsage: "",
		Usage:     "rotate secret on server and issue an update token",
		Action: func(context *cli.Context) error {
			return RotateFunc(context, client)
		},
	}
}
func RotateFunc(context *cli.Context, vcli *client.VirgilHttpClient) error {

	if context.NArg() > 0 {
		return errors.New("invalid number of arguments")
	}

	token := context.String("token")
	appId := context.String("appid")

	var resp *RotateResponse

	_, err := vcli.Send(http.MethodPost, token, "phe/v1/"+appId+"/rotate", nil, &resp)

	if err != nil {
		return err
	}

	if resp != nil {

		version := strconv.Itoa(resp.Version)
		token := passw0rd.MarshalUpdateToken(resp.UpdateToken.A, resp.UpdateToken.B)
		fmt.Println("Your update token:", "UT."+version+"."+base64.StdEncoding.EncodeToString(token))
	}

	return nil
}
