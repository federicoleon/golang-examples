package services

import (
	"fmt"
)

const (
	pong = "pong"
)

var (
	PingService pingService = pingServiceImpl{}
)

type pingService interface {
	HandlePing() (string, error)
}

type pingServiceImpl struct{}

func (service pingServiceImpl) HandlePing() (string, error) {
	fmt.Println("doing some complex things...")
	return pong, nil
}
