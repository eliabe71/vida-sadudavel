package models

type Medico struct {
	Name            string `json:"name"`
	LastName        string `json:"lastName"`
	Id              int    `json:"id"`
	City            string `json:"city"`
	State           string `json:"state"`
	Crm             string `json:"cpf"`
	AreaOfOcupation string `json:"areaOfOcupation"`
	HourInit        string `json:"hour-init"`
	HourEnd         string `json:"hour-end"`
	Price           string `json:"price"`
}
