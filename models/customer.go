//models/customer.go

package models

//Customer used to retreive customers
type Customer struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
	Attention    string `json:"attention"`
	Address1     string `json:"address_1"`
	Address2     string `json:"address_2"`
}
