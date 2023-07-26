package service

import (
	"context"
	"github.com/walbety/payment-system/portal-service/internal/canonical"
)

type Service struct{}

func New() Service {
	return Service{}
}

func (svc *Service) welcome() {
	return
}

func (svc *Service) ListUsers(ctx context.Context) (canonical.User, error) {
	return canonical.User{
		Name: "jao",
		Code: 123,
	}, nil
}
