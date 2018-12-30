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

	"github.com/passw0rd/phe-go"
	"github.com/passw0rd/sdk-go"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func UpdateKeys() *cli.Command {
	return &cli.Command{
		Name:      "update-keys",
		Aliases:   []string{"u"},
		ArgsUsage: "public_key service_secret_key update_token",
		Usage:     "update secret key and public key using update token",
		Action: func(context *cli.Context) error {
			return updateFunc(context)
		},
	}
}
func updateFunc(context *cli.Context) error {

	if context.NArg() < 3 {
		return errors.New("invalid number of arguments")
	}

	pkStr := context.Args().First()
	skStr := context.Args().Get(1)
	tokenStr := context.Args().Get(2)

	pkVersion, pk, err := passw0rd.ParseVersionAndContent("PK", pkStr)
	if err != nil {
		return err
	}
	skVersion, sk, err := passw0rd.ParseVersionAndContent("SK", skStr)
	if err != nil {
		return err
	}
	tokenVersion, updateToken, err := passw0rd.ParseVersionAndContent("UT", tokenStr)

	if err != nil {
		return err
	}

	if (pkVersion+1) != tokenVersion || (skVersion+1) != tokenVersion {
		return errors.New("Key version must be 1 less than token version")
	}

	newSk, newPk, err := phe.RotateClientKeys(sk, pk, updateToken)
	if err != nil {
		return err
	}

	fmt.Printf("New server public key:\nPK.%d.%s\nNew client private key:\nSK.%d.%s\n", tokenVersion, base64.StdEncoding.EncodeToString(newPk), tokenVersion, base64.StdEncoding.EncodeToString(newSk))

	return nil
}
