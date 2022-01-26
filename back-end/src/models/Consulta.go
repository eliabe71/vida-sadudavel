package models

type Consulta struct {
	Status   bool    `json:"status"`
	Effected bool    `json:"effected"`
	MedicoID int     `json:"medicId"`
	ClientID int     `json:"clientId"`
	Id       int     `json:"id"`
	Price    float32 `json:"price"`
	Day      string  `json:"day"`
	HourInit string  `json:"hourInit"`
	HourEnd  string  `json:"hourEnd"`
}
