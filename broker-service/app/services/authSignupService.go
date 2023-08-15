package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/shaikrasheed99/broker-service/app/models/requests"
	"github.com/shaikrasheed99/broker-service/app/models/responses"
	"github.com/shaikrasheed99/broker-service/utils/constants"
)

func (bs *brokerService) HandleAuthSignup(req *requests.AuthSignupPayload) (string, error) {
	log.Println("[HandleAuthSignupService] Hitting handle auth signup service")

	body, err := json.Marshal(&req)
	if err != nil {
		log.Println("[HandleAuthSignupService]", err.Error())
		return "", err
	}

	res, err := http.Post(constants.AuthSignupServiceAPI, "", bytes.NewBuffer(body))
	if err != nil {
		log.Println("[HandleAuthSignupService]", err.Error())
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println("[HandleAuthSignupService]", err.Error())
		return "", err
	}

	if res.StatusCode != http.StatusCreated {
		var errRes responses.ErrorResponse
		json.Unmarshal(resBody, &errRes)
		err := errors.New(errRes.Message)
		log.Println("[HandleAuthSignupService]", err.Error())
		return "", err
	}

	var successRes responses.SuccessResponse
	json.Unmarshal(resBody, &successRes)

	return successRes.Message, nil
}
