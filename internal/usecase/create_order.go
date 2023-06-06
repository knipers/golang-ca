package usecase

import entity "github.com/knipers/golang-ca/internal/entity/order"

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	Repository entity.OrderRepositoryInterface
}

func NewCreateOrderUseCase(repository entity.OrderRepositoryInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		Repository: repository,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	// create order entity
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}

	err := order.CalculateFinalPrice()
	if err != nil {
		return OrderOutputDTO{}, err
	}

	err = c.Repository.Save(&order)

	if err != nil {
		return OrderOutputDTO{}, err
	}
	// return output
	output := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	return output, nil
}
