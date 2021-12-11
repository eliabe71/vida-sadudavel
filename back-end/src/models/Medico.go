package models
type Medico struct{
	Name string `json:"Name"`
	LastName string `json:"LastName"`
	Id int `json:"Id"`
	City string `json:"City"`
	State string `json:"State"`
	Crm string `json:"Cpf"`
	AreaOfOcupation string `json:"AreaOfOcupation"`
}