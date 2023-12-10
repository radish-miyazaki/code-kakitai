package order

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/code-kakitai/application/order"
	"github.com/radish-miyazaki/code-kakitai/presentation/settings"
	"github.com/radish-miyazaki/go-pkg/validator"
)

type Handler struct {
	saveOrderUseCase *order.SaveOrderUseCase
}

func NewHandler(saveOrderUseCase *order.SaveOrderUseCase) *Handler {
	return &Handler{
		saveOrderUseCase: saveOrderUseCase,
	}
}

func (h *Handler) PostOrders(ctx *gin.Context) {
	var params []*PostOrdersParams
	if err := ctx.ShouldBind(&params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	// 本来は Session に入っている UserID を取得するが、本質では無いので省略
	userID := "01HCNYK0PKYZWB0ZT1KR0EPWGP"
	dtos := make([]order.SaveOrderUseCaseInputDto, 0, len(params))

	for _, param := range params {
		dtos = append(dtos, order.SaveOrderUseCaseInputDto{
			ProductID: param.ProductID,
			Quantity:  param.Quantity,
		})
	}
	id, err := h.saveOrderUseCase.Run(
		ctx.Request.Context(),
		userID,
		dtos,
		time.Now(),
	)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
		return
	}

	settings.ReturnStatusCreated(ctx, id)
}
