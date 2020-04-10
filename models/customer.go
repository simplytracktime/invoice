//models/customer.go

package models

//Customer used to retreive customers
type Customer struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"unique_index:name_ndx"`
	EmailAddress string `json:"email_address" gorm:"unique_index:email_address_ndx"`
	Attention    string `json:"attention"`
	Address1     string `json:"address_1"`
	Address2     string `json:"address_2"`
}
