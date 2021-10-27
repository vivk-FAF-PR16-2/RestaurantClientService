package dto

type ClientInData struct {
	ClientId int           `json:"client_id"`
	Orders   []OrderInData `json:"orders"`
}
