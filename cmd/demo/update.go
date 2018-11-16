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

	"github.com/passw0rd/phe-go"

	"github.com/passw0rd/sdk-go"

	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func Update() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Aliases:   []string{"u"},
		ArgsUsage: "record update_token",
		Usage:     "update a record using update token",
		Action: func(context *cli.Context) error {
			return updateFunc(context)
		},
	}
}
func updateFunc(context *cli.Context) error {

	if context.NArg() < 2 {
		return errors.New("invalid number of arguments")
	}

	recStr := context.Args().First()
	updateTokenStr := context.Args().Get(1)

	recb, err := base64.StdEncoding.DecodeString(recStr)
	if err != nil {
		return err
	}

	recVersion, rec, err := passw0rd.UnmarshalRecord(recb)
	if err != nil {
		return err
	}

	tokenVersion, updateToken, err := passw0rd.ParseVersionAndContent("UT", updateTokenStr)

	if (recVersion + 1) != tokenVersion {
		return errors.New("record version should be 1 less than update token version")
	}

	token, err := passw0rd.UnmarshalUpdateToken(updateToken)
	if err != nil {
		return err
	}

	newRec, err := phe.UpdateRecord(rec, token)
	if err != nil {
		return err
	}

	newRecBin, err := passw0rd.MarshalRecord(tokenVersion, newRec)
	if err != nil {
		return err
	}

	fmt.Println(base64.StdEncoding.EncodeToString(newRecBin))

	return nil
}
