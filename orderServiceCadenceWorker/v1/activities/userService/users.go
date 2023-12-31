package userService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"orderServiceCadenceWorker/v1/env"
	"orderServiceCadenceWorker/v1/log"
	"orderServiceCadenceWorker/v1/models"
	"orderServiceCadenceWorker/v1/utils"
)

const UserServiceBaseURL = "/api/v1/"

func CreateAUser(_ context.Context, user models.User) (models.User, error) {
	logger := log.GetLogger()
	restyClient := utils.GetRestyClient()

	URL := fmt.Sprintf("%s/users/", UserServiceBaseURL)

	endpointUrl := utils.GenerateURL(URL, "", env.Env.UserServiceScheme, env.Env.UserServiceHost).
		String()

	logger.Info(fmt.Sprintf("Calling External Endpoint: %s", endpointUrl))

	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "/").
		SetBody(user).
		Post(endpointUrl)

	var response models.User

	if err != nil {
		return response, err
	} else if resp.IsError() {
		err = errors.New(string(resp.Body()))

		return response, err
	}

	responseData := gjson.Get(string(resp.Body()), "data").String()

	if err = json.Unmarshal([]byte(responseData), &response); err != nil {
		return response, err
	}

	return response, nil
}
