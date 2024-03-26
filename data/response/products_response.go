package response

type ProductsResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	OriginalPrice string `json:"original_price"`
	DiscountPrice string `json:"discount_price"`
	Supermarket   string `json:"supermarket"`
}
