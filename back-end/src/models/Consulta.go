package models

type Consulta struct {
	Id       int     `json:"id"`
	Status   bool    `json:"status"`
	Effected bool    `json:"effected"`
	MedicoID int     `json:"medicId"`
	ClientID int     `json:"clientId"`
	ClientName string `json:"clientName"`
	MedicName string `json:"medicName"`
	State	  string `json:"state"`
	City      string `json:"city"`
	Price    string `json:"price"`
	Day      string  `json:"day"`
	HourInit string  `json:"hourInit"`
	HourEnd  string  `json:"hourEnd"`
}
