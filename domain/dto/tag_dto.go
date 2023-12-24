package dto

type RequestCreateTag struct {
	Name string `json:"name"`
}

type RequestUpdateTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
