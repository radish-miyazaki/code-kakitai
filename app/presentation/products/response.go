package products

type productResponseModel struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}

type GetProductResponse struct {
	*productResponseModel
	OwnerName string `json:"owner_name"`
}

type PostProductResponse struct {
	Product productResponseModel `json:"product"`
}
