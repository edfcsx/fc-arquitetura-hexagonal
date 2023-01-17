package cli_test

import (
	"fmt"
	"github.com/edfcsx/fc-arquitetura-hexagonal/adapters/cli"
	mock_application "github.com/edfcsx/fc-arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 25.99
	status := "enabled"
	id := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s.",
		id, name, price, status)

	result, err := cli.Run(service, "create", "", name, price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpectedEnable := fmt.Sprintf("Product %s has been enabled.", name)
	result, err = cli.Run(service, "enable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpectedEnable, result)

	resultExpectedDisable := fmt.Sprintf("Product %s has been disabled.", name)
	result, err = cli.Run(service, "disable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpectedDisable, result)

	resultDefault := fmt.Sprintf("Product Id: %s\nName: %s\nPrice: %f\nStatus: %s",
		id, name, price, status)
	result, err = cli.Run(service, "get", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultDefault, result)

}
