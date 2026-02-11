package base

import (
	"context"
	"log"
)

type BaseService struct {
	Logger *log.Logger
}

func (service *BaseService) WithContext(ctx context.Context) context.Context {
	return ctx
}

func (service *BaseService) Log(message string) {
	service.Logger.Println(message)
}
