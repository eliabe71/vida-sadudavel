package models

type Consulta struct {
	Status   bool    `json :"status"`
	Effected bool    `json:"effected"`
	MedicoID string  `json: "medicId"`
	ClientID string  `json:"clientId"`
	Id       int     `json:"id"`
	Price    float32 `json:"price"`
	Day      string  `json:"day"`
	HourInit string  `json:"hour-init"`
	HourEnd  string  `json:"hour-end"`
}
