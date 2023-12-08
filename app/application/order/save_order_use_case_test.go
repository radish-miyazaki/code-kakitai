package order

import (
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	cartDomain "github.com/radish-miyazaki/code-kakitai/domain/cart"
	orderDomain "github.com/radish-miyazaki/code-kakitai/domain/order"
	"github.com/radish-miyazaki/go-pkg/ulid"
)

func TestSaveOrderUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrderDomainService := orderDomain.NewMockOrderDomainService(ctrl)
	mockCartRepo := cartDomain.NewMockCartRepository(ctrl)
	uc := NewSaveOrderUseCase(mockOrderDomainService, mockCartRepo)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	userID := ulid.NewULID()
	dtos := []SaveOrderUseCaseInputDto{
		{
			ProductID: ulid.NewULID(),
			Quantity:  1,
		},
		{
			ProductID: ulid.NewULID(),
			Quantity:  3,
		},
	}
	cart, _ := cartDomain.NewCart(userID)
	for _, dto := range dtos {
		cart.AddProduct(dto.ProductID, dto.Quantity)
	}

	tests := []struct {
		name     string
		dtos     []SaveOrderUseCaseInputDto
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "Valid",
			dtos: dtos,
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), userID).Return(cart, nil),
					mockOrderDomainService.EXPECT().OrderProducts(gomock.Any(), cart, now).Return("", nil),
				)
			},
			wantErr: false,
		},
		{
			name: "Invalid: cart's quantity does not match with input",
			dtos: []SaveOrderUseCaseInputDto{
				{
					ProductID: ulid.NewULID(),
					Quantity:  1,
				},
			},
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), userID).Return(cart, nil),
				)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.mockFunc()
			_, err := uc.Run(context.Background(), userID, tt.dtos, now)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveOrderUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
