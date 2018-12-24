package utils

import (
	"github.com/passw0rd/cli/client"
	"github.com/passw0rd/cli/models"
)

func CheckRetry(errToCheck error, vcli *client.VirgilHttpClient) (token string, err error) {
	httpErr, ok := errToCheck.(*models.HttpError)

	if ok && httpErr.Code == 40404 {
		err = Login("", "", vcli)
		if err != nil {
			return
		}
		return LoadAccessToken()
	}
	return "", errToCheck
}
