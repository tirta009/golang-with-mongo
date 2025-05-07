package payload

type TransactionRequest struct {
	UserId          string          `json:"user_id"`
	TotalAmount     float64         `json:"total_amount"`
	ProductRequest  ProductRequest  `json:"product_request"`
	ShipmentRequest ShipmentRequest `json:"shipment_request"`
}

type ProductRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
}

type ShipmentRequest struct {
	Province   string `json:"province"`
	City       string `json:"city"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}
