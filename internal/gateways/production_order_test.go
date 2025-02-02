package gateways

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestProductionOrderGateway_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
