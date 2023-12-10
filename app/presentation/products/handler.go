package products

import (
	"github.com/gin-gonic/gin"

	"github.com/radish-miyazaki/code-kakitai/application/product"
	"github.com/radish-miyazaki/code-kakitai/presentation/settings"
	"github.com/radish-miyazaki/go-pkg/validator"
)

type Handler struct {
	saveProductUseCase  *product.SaveProductUseCase
	fetchProductUseCase *product.FetchProductUseCase
}

func NewHandler(
	saveproductUseCase *product.SaveProductUseCase,
	fetchProductUseCase *product.FetchProductUseCase,
) *Handler {
	return &Handler{
		saveProductUseCase:  saveproductUseCase,
		fetchProductUseCase: fetchProductUseCase,
	}
}

// PostProducts godoc
// @Summary 商品を保存する
// @Tags products
// @Accept json
// @Produce json
// @Param request body PostProductParams true "商品情報"
// @Success 201 {object} PostProductResponse
// @Router /v1/products [post]
func (h *Handler) PostProducts(ctx *gin.Context) {
	var params PostProductParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
	}

	validate := validator.GetValidator()
	err := validate.Struct(params)
	if err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
	}

	input := product.SaveProductUseCaseInputDTO{
		OwnerID:     params.OwnerID,
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
		Stock:       params.Stock,
	}

	dto, err := h.saveProductUseCase.Run(ctx, input)
	if err != nil {
		settings.ReturnError(ctx, err)
	}

	output := PostProductResponse{
		productResponseModel{
			ID:          dto.ID,
			OwnerID:     dto.OwnerID,
			Name:        dto.Name,
			Description: dto.Description,
			Price:       dto.Price,
			Stock:       dto.Stock,
		},
	}
	settings.ReturnStatusCreated(ctx, output)
}
