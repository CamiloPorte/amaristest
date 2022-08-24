package client

type responseReq struct {
	Pokemon pokemon `json:"pokemon"`
}
type pokemon struct {
	Name string `json:"name"`
}
