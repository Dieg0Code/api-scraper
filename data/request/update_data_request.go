package request

type UpdateDataRequest struct {
	Update bool `validate:"required" json:"update"`
}
