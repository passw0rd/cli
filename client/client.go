package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type VirgilHttpClient struct {
	Client  HttpClient
	Address string
}

func (vc *VirgilHttpClient) Send(method string, url string, payload interface{}, respObj interface{}) (headers http.Header, err error) {
	var body []byte
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, errors.Wrap(err, "VirgilHttpClient.Send: marshal payload")
		}
	}
	req, err := http.NewRequest(method, vc.Address+url, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHttpClient.Send: new request")
	}

	client := vc.getHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHttpClient.Send: send request")
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("not found")
	}

	if resp.StatusCode == http.StatusOK {
		if respObj != nil {

			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(respObj)
			if err != nil {
				return nil, errors.Wrap(err, "VirgilHttpClient.Send: unmarshal response object")
			}
		}
		return resp.Header, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "VirgilHttpClient.Send: read response body")
	}

	return nil, errors.New(fmt.Sprintf("%s\n", string(respBody)))
}

func (vc *VirgilHttpClient) getHttpClient() HttpClient {
	if vc.Client != nil {
		return vc.Client
	}
	return http.DefaultClient
}
