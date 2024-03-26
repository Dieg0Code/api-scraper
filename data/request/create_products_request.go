package request

type CreateProductsRequest struct {
	Name          string `validate:"required,max=255,min=3" json:"name"`
	Category      string `validate:"required,max=255,min=3" json:"category"`
	OriginalPrice string `validate:"required,max=255,min=3" json:"original_price"`
	DiscountPrice string `validate:"required,max=255,min=3" json:"discount_price"`
	Supermarket   string `validate:"required,max=255,min=3" json:"supermarket"`
}
