package models

type Medico struct {
	Name            string `json:"name"`
	LastName        string `json:"lastName"`
	Id              int    `json:"id"`
	City            string `json:"city"`
	State           string `json:"state"`
	Crm             string `json:"crm"`
	AreaOfOcupation string `json:"areaOfOcupation"`
	HourInit        string `json:"hourinit"`
	HourEnd         string `json:"hourend"`
	Price           string `json:"price"`
}
