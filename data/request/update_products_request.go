package request

type UpdateProductsRequest struct {
	Id            int    `validate:"required"`
	Name          string `validate:"max=255,min=3" json:"name"`
	Category      string `validate:"max=255,min=3" json:"category"`
	OriginalPrice string `validate:"max=255,min=3" json:"original_price"`
	DiscountPrice string `validate:"max=255,min=3" json:"discount_price"`
	Supermarket   string `validate:"required,max=255,min=3" json:"supermarket"`
}
