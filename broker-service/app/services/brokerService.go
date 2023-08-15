package services

import "github.com/shaikrasheed99/broker-service/app/models/requests"

type IBrokerService interface {
	HandleAuthSignup(*requests.AuthSignupPayload) (string, error)
}

type brokerService struct{}

func NewBrokerService() IBrokerService {
	return &brokerService{}
}
