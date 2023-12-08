package product

type SaveProductUseCase struct{}

type SaveProductUseCaseInputDTO struct {
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}

type SaveProductUseCaseOutputDTO struct {
	ID          string
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}
