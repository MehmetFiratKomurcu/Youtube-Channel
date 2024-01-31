package model

// CreateOrderRequest
// @Description Request about creating Order
type CreateOrderRequest struct {
	// shipment no of Order
	ShipmentNumber string `json:"shipmentNumber" validate:"required"`
	// country code like: tr, us
	CountryCode string `json:"countryCode" validate:"required,len=2"`
	// age to make sure you are young
	Age int `json:"age" validate:"required,oldAge"`
}
