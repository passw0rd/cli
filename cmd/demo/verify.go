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

package demo

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/passw0rd/sdk-go"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func Verify() *cli.Command {
	return &cli.Command{
		Name:      "verify",
		Aliases:   []string{"v"},
		ArgsUsage: "password record",
		Usage:     "verify password against a record",
		Action: func(context *cli.Context) error {
			return verifyFunc(context)
		},
	}
}
func verifyFunc(context *cli.Context) error {

	if context.NArg() < 2 {
		return errors.New("invalid number of arguments")
	}

	token := context.String("access_token")
	appId := context.String("app_id")
	pub := context.String("pk")
	sk := context.String("sk")
	pwd := context.Args().First()
	recStr := context.Args().Get(1)

	if token == "" {
		log.Fatal("please specify your access token")
	}
	if appId == "" {
		log.Fatal("please specify app id")
	}
	if pub == "" {
		log.Fatal("please specify server public key")
	}
	if sk == "" {
		log.Fatal("please specify your secret key")
	}
	if pwd == "" {
		log.Fatal("please specify your password")
	}

	if pwd == "" {
		log.Fatal("please specify record")
	}

	rec, err := base64.StdEncoding.DecodeString(recStr)
	if err != nil {
		return err
	}

	ctx, err := passw0rd.CreateContext(token, appId, sk, pub)
	if err != nil {
		return err
	}

	prot, err := passw0rd.NewProtocol(ctx)
	if err != nil {
		return err
	}

	key, err := prot.VerifyPassword(pwd, rec)
	if err != nil {
		return err
	}

	fmt.Printf("Password is correct. Encryption key: %x\n", key)

	return nil
}
