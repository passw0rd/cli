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

package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/passw0rd/cli/models"

	"github.com/golang/protobuf/proto"

	"github.com/pkg/errors"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type VirgilHttpClient struct {
	Client  HttpClient
	Address string
}

func (vc *VirgilHttpClient) Send(method string, token string, urlPath string, payload proto.Message, respObj proto.Message, extraOptions ...interface{}) (headers http.Header, err error) {
	var body []byte
	if payload != nil {
		body, err = proto.Marshal(payload)
		if err != nil {
			return nil, errors.Wrap(err, "VirgilHTTPClient.Send: marshal payload")
		}
	}

	u, err := url.Parse(vc.Address)
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHTTPClient.Send: URL parse")
	}

	u.Path = path.Join(u.Path, urlPath)
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHTTPClient.Send: new request")
	}

	if token != "" {
		req.Header.Add("AccountToken", token)
	}

	if len(extraOptions) > 0 {
		appToken, ok := extraOptions[0].(string)
		if ok {
			req.Header.Add("AppToken", appToken)
		}
	}

	client := vc.getHttpClient()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHTTPClient.Send: send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		if respObj != nil {

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				return nil, errors.Wrap(err, "VirgilHTTPClient.Send: read body")
			}

			err = proto.Unmarshal(body, respObj)
			if err != nil {
				return nil, errors.Wrap(err, "VirgilHTTPClient.Send: unmarshal response object")
			}
		}
		return resp.Header, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHTTPClient.Send: read response body")
	}

	if len(respBody) > 0 {
		httpErr := &models.HttpError{}
		err = proto.Unmarshal(respBody, httpErr)
		if err == nil {

			return nil, httpErr
		}
	}

	return nil, fmt.Errorf("%d %s", resp.StatusCode, string(respBody))
}

func (vc *VirgilHttpClient) getHttpClient() HttpClient {
	if vc.Client != nil {
		return vc.Client
	}
	return http.DefaultClient
}
