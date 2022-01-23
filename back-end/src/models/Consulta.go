package models

type Consulta struct {
	Status   bool    `json :"Status"`
	Effected bool    `json:"Effected"`
	MedicoID string  `json: "MedicId"`
	ClientID string  `json:"ClientId"`
	Id       int     `json:"Id"`
	Price    float32 `json:"Price"`
	Day      string  `json:"Day"`
	HourInit string  `json:"Hour-Init"`
	HourEnd  string  `json:"Hour-End"`
}
