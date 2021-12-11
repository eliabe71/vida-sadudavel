package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/eliabe71/vida-saudavel/back-end/src/database/database"
	"github.com/eliabe71/vida-saudavel/back-end/src/models"
)

var db *database.Db

func handleMedics(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		stmt, err := db.DB.Query("SELECT name, lastname, city, state,crm,areaofocupation From Medico")
		if err != nil {
			log.Fatal(err)
		}
		resp := models.Response{}
		resp.Type = "Medics"
		for stmt.Next() {
			medics := models.Medico{}
			stmt.Scan(&medics.Name, &medics.LastName, &medics.City, &medics.State, &medics.Crm, &medics.AreaOfOcupation)
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
func Router() {

	db = new(database.Db)
	db.OpenDb()
	http.HandleFunc("/medics", handleMedics)
	http.HandleFunc("/consultas/medic/", handleConsultasM)
	http.HandleFunc("/consultas/client/", handleConsultasC)
	http.ListenAndServe(":8084", nil)

}
