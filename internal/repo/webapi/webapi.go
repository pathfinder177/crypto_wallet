package webapi

type WebApiRepo struct {
	serverAddress string
}

func New(address string) *WebApiRepo {
	return &WebApiRepo{serverAddress: address}
}
