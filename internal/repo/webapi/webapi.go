package webapi

import "log"

type WebApiRepo struct {
	serverAddress string
	logger        *log.Logger
}

func New(address string, logger *log.Logger) *WebApiRepo {
	return &WebApiRepo{
		serverAddress: address,
		logger:        logger,
	}
}
