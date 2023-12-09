package model

type Medicine struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_no"`
	MoleculeName string `json:"molecule_name"`
	Quantity     int    `json:"quantity"`
}
