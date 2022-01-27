package models
type Cliente struct{
	Name string `json:"name"`
	LastName string `json:"lastName"`
	Id int `json:"id"`
	City string `json:"city"`
	State string `json:"state"`
	Cpf string `json:"cpf"`
}
