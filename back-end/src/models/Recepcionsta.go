package models
type Recepcionista struct{
	Name string `json:"Name"`
	LastName string `json:"LastName"`
	Id int `json:"Id"`
	MedicId string `json:"MedicId"`
	Cpf string `json:"Cpf"`
}