package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/eliabe71/vida-saudavel/back-end/src/database/database"
	"github.com/eliabe71/vida-saudavel/back-end/src/models"
	"github.com/eliabe71/vida-saudavel/back-end/src/utils"
)

var db *database.Db

func handleMedics(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		stmt, err := db.DB.Query("SELECT name, lastname, city, state,crm,areaofocupation,hourend,hourinit From Medico")
		if err != nil {
			log.Fatal(err)
		}
		resp := models.Response{}
		resp.Type = "Medics"
		for stmt.Next() {
			medics := models.Medico{}
			stmt.Scan(&medics.Name, &medics.LastName, &medics.City, &medics.State, &medics.Crm, &medics.AreaOfOcupation, &medics.HourEnd, &medics.HourInit)
			resp.Data = append(resp.Data, medics)
		}

		bJson, _ := json.Marshal(resp)
		w.Write(bJson)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleConsultasM(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	re := regexp.MustCompile(`\/[a-z]+\/[a-z]+\/[0-9]+$`)
	if re.MatchString(r.URL.Path) {
		re = regexp.MustCompile(`\/[a-z]+\/[a-z]+\/`)
		numberS := re.Split(r.URL.Path, 2)[1]
		re = regexp.MustCompile(`[0-9]+`)
		byt := re.FindIndex([]byte(numberS))
		b := []byte(numberS)
		numberS = string(b[byt[0]:byt[1]])
		stmt, err := db.DB.Query("SELECT status,effected,medicid,clienteid,day,price,hourend,hourinit From Consulta where medicId = $1", numberS)

		if err == nil {
			resp := models.Response{}
			resp.Type = "Consultas"
			for stmt.Next() {
				cons := models.Consulta{}
				stmt.Scan(cons.Status, cons.Effected, cons.MedicoID, cons.ClientID, cons.Day, cons.Price, cons.HourEnd, cons.HourInit)
				resp.Data = append(resp.Data, cons)
			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
func handleConsultasC(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	re := regexp.MustCompile(`\/[a-z]+\/[a-z]+\/[0-9]+$`)
	if re.MatchString(r.URL.Path) {
		re = regexp.MustCompile(`\/[a-z]+\/[a-z]+\/`)
		numberS := re.Split(r.URL.Path, 2)[1]
		re = regexp.MustCompile(`[0-9]+`)
		byt := re.FindIndex([]byte(numberS))
		b := []byte(numberS)
		numberS = string(b[byt[0]:byt[1]])
		stmt, err := db.DB.Query("SELECT status,effected,medicid,clienteid,day,price,hourend,hourinit From Consulta where clienteId = $1", numberS)

		if err == nil {
			resp := models.Response{}
			resp.Type = "Consultas"
			for stmt.Next() {
				cons := models.Consulta{}
				stmt.Scan(cons.Status, cons.Effected, cons.MedicoID, cons.ClientID, cons.Day, cons.Price, cons.HourEnd, cons.HourInit)
				resp.Data = append(resp.Data, cons)
			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func handleRecepcionista(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	re := regexp.MustCompile(`\/[a-z]+\/[0-9]+$`)
	if re.MatchString(r.URL.Path) {
		re = regexp.MustCompile(`\/[a-z]+\/`)
		numberS := re.Split(r.URL.Path, 2)[1]
		re = regexp.MustCompile(`[0-9]+`)
		byt := re.FindIndex([]byte(numberS))
		b := []byte(numberS)
		numberS = string(b[byt[0]:byt[1]])
		stmt, err := db.DB.Query("SELECT name,lastname, id, medicoId From  recepcionista where cpf = $1", numberS)

		if err == nil {
			resp := models.Response{}
			resp.Type = "Recepcionista"
			for stmt.Next() {
				cons := models.Recepcionista{}
				stmt.Scan(cons.Name, cons.LastName, cons.Id, cons.MedicId)
				resp.Data = append(resp.Data, cons)
			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func handleConsultasSingup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var consulta models.Consulta
		dec := json.NewDecoder(r.Body)
		dec.Decode(&consulta)
		stmt1, _ := db.DB.Query(`SELECT houred From Medico where crm=$1`, consulta.MedicoID)
		stmt1.Next()
		var hourBuffer string
		stmt1.Scan(&hourBuffer)
		r, flag := utils.BiggerThen(consulta.HourInit, hourBuffer)
		if flag {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if strings.Compare(r, consulta.HourEnd) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		r, flag = utils.BiggerThen(consulta.HourEnd, hourBuffer)
		if strings.Compare(r, consulta.HourEnd) == 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		stmt, err := db.DB.Query(`SELECT Count(*) From consulta Where medicId=$1 and hourinit>=CAST($2 AS TIME) and hourinit < CAST($2 AS TIME) + interval '25 minutes'`, consulta.MedicoID, consulta.HourInit)
		if err == nil {
			//and hourinit < time $2 + interval '25 minutes' or hourend < $2 + interval '25 minutes' and hourend >= $2
			for stmt.Next() {
				var count int
				stmt.Scan(&count)
				if count == 0 {
					_, err := db.DB.Query("Insert into consulta (medicid, clientid, status, effected,price,day,hourinit, hourend) values($1, $2,$3,$4, $5,$6, $7 ,CAST($7 AS TIME)+interval '25 minutes')", consulta.MedicoID, consulta.ClientID, consulta.Status, consulta.Effected, consulta.Price, consulta.Day, consulta.HourInit)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
					}

				} else {
					w.WriteHeader(http.StatusBadRequest)
				}

			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
func handleMedicoSingup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var consulta models.Medico
		dec := json.NewDecoder(r.Body)
		dec.Decode(&consulta)

		stmt, err := db.DB.Query(`SELECT Count(*) From Medico Where crm=$1`, consulta.Crm)
		if err == nil {
			for stmt.Next() {
				var count int
				stmt.Scan(&count)
				if count == 0 {
					_, err := db.DB.Query("Insert into Medico (crm, city , state, areaofocupation, name, lastname, hourinit, hourend, price) values($1,$2,$3,$4,$5,$6,$7,$8,$9)", consulta.Crm, consulta.City, consulta.State, consulta.AreaOfOcupation, consulta.Name, consulta.LastName, consulta.HourInit, consulta.HourEnd, consulta.Price)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
func Router() {

	db = new(database.Db)
	db.OpenDb()
	http.HandleFunc("/medic", handleMedicoSingup)
	http.HandleFunc("/medics", handleMedics)
	http.HandleFunc("/consultas/medic/", handleConsultasM)
	http.HandleFunc("/consulta", handleConsultasSingup)
	http.HandleFunc("/consultas/client/", handleConsultasC)
	http.HandleFunc("/recepcionista/", handleRecepcionista)
	http.ListenAndServe(":8084", nil)

}
