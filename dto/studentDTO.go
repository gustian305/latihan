package dto

type RequestStudent struct {
	Name  string `json:"name" validate:"required"`
	Class string `json:"class" validate:"required"`
}

type ResponseStudent struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}
