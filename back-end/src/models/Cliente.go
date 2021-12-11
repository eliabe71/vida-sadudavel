package models
type Cliente struct{
	Name string `json:"Name"`
	LastName string `json:"LastName"`
	Id int `json:"Id"`
	City string `json:"City"`
	State string `json:"State"`
	Cpf string `json:"Cpf"`
}