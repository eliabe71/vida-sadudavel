package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/eliabe71/vida-saudavel/back-end/src/database/database"
	"github.com/eliabe71/vida-saudavel/back-end/src/models"
	"github.com/eliabe71/vida-saudavel/back-end/src/utils"
)

var db *database.Db

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleMedics(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		stmt, err := db.DB.Query("SELECT id,name, lastname, city, state,crm,areaofocupation,hourend,hourinit,price From Medico")
		if err != nil {
			log.Fatal(err)
		}
		resp := models.Response{}
		resp.Type = "Medics"
		for stmt.Next() {
			medics := models.Medico{}
			stmt.Scan(&medics.Id, &medics.Name, &medics.LastName, &medics.City, &medics.State, &medics.Crm, &medics.AreaOfOcupation, &medics.HourEnd, &medics.HourInit, &medics.Price)
			resp.Data = append(resp.Data, medics)
		}

		bJson, _ := json.Marshal(resp)
		w.Write(bJson)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleConsultasM(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println(r.URL.Path)
	re := regexp.MustCompile(`\/[a-z]+\/[a-z]+\/[0-9]+$`)
	if re.MatchString(r.URL.Path) {
		re = regexp.MustCompile(`\/[a-z]+\/[a-z]+\/`)
		numberS := re.Split(r.URL.Path, 2)[1]
		re = regexp.MustCompile(`[0-9]+`)
		byt := re.FindIndex([]byte(numberS))
		b := []byte(numberS)
		numberS = string(b[byt[0]:byt[1]])
		fmt.Println(numberS)
		stmt, err := db.DB.Query("SELECT id,status,effected,medicid,clientid,day,price,hourend,hourinit From Consulta where medicid = $1", numberS)

		if err == nil {
			resp := models.Response{}
			resp.Type = "Consultas"
			for stmt.Next() {
				cons := models.Consulta{}
				stmt.Scan(&cons.Id, &cons.Status, &cons.Effected, &cons.MedicoID, &cons.ClientID, &cons.Day, &cons.Price, &cons.HourEnd, &cons.HourInit)
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
	enableCors(&w)
	fmt.Println(r.URL.Path)
	re := regexp.MustCompile(`\/[a-z]+\/[a-z]+\/[0-9]+$`)
	if re.MatchString(r.URL.Path) {
		re = regexp.MustCompile(`\/[a-z]+\/[a-z]+\/`)
		numberS := re.Split(r.URL.Path, 2)[1]
		re = regexp.MustCompile(`[0-9]+`)
		byt := re.FindIndex([]byte(numberS))
		b := []byte(numberS)
		numberS = string(b[byt[0]:byt[1]])
		stmt, err := db.DB.Query("SELECT id,status,effected,medicid,clientid,day,price,hourend,hourinit From Consulta where clientId = $1", numberS)

		if err == nil {
			resp := models.Response{}
			resp.Type = "Consultas"
			for stmt.Next() {
				cons := models.Consulta{}
				stmt.Scan(&cons.Id, &cons.Status, &cons.Effected, &cons.MedicoID, &cons.ClientID, &cons.Day, &cons.Price, &cons.HourEnd, &cons.HourInit)
				resp.Data = append(resp.Data, cons)
			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handlegetClient(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		re := regexp.MustCompile(`\/[a-z]+\/[0-9]+$`)
		fmt.Println(r.URL.Path)
		if re.MatchString(r.URL.Path) {
			re = regexp.MustCompile(`\/[a-z]+\/`)
			numberS := re.Split(r.URL.Path, 2)[1]
			re = regexp.MustCompile(`[0-9]+`)
			byt := re.FindIndex([]byte(numberS))
			b := []byte(numberS)
			numberS = string(b[byt[0]:byt[1]])
			i, err := strconv.Atoi(numberS)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			stmt, errDb := db.DB.Query("SELECT id,name ,lastname ,cpf ,city ,state From cliente where id = $1", i)
			if errDb != nil {
				fmt.Println(errDb)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			resp := models.Response{}
			for stmt.Next() {
				cons := models.Cliente{}
				stmt.Scan(&cons.Id, &cons.Name, &cons.LastName, &cons.Cpf, &cons.City, &cons.State)
				cons.Id = i
				resp.Type = "Cliente"
				resp.Data = append(resp.Data, cons)
			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		} else {

			w.WriteHeader(http.StatusBadRequest)
		}

	}
}

func handleRemoveConsulta(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		re := regexp.MustCompile(`\/[a-z]+\/[0-9]+$`)
		fmt.Println(r.URL.Path)
		if re.MatchString(r.URL.Path) {
			re = regexp.MustCompile(`\/[a-z]+\/`)
			numberS := re.Split(r.URL.Path, 2)[1]
			re = regexp.MustCompile(`[0-9]+`)
			byt := re.FindIndex([]byte(numberS))
			b := []byte(numberS)
			numberS = string(b[byt[0]:byt[1]])
			i, err := strconv.Atoi(numberS)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			stmt, errDb := db.DB.Query("SELECT status, effected, medicid, clienteid, price, day, hourend, hourinit From consulta where id = $1", i)
			if errDb != nil {
				fmt.Println(errDb)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			resp := models.Response{}
			row := 0
			for stmt.Next() {
				row++
				cons := models.Consulta{}
				cons.Id = i
				stmt.Scan(&cons.Status, &cons.Effected, &cons.MedicoID, &cons.ClientID, &cons.Price, &cons.Day, &cons.HourEnd, &cons.HourInit)
				resp.Type = "Consulta Deletada"
				resp.Data = append(resp.Data, cons)
				_, errDb := db.DB.Query("Delete from consulta o where id = $1", i)
				if errDb != nil {
					fmt.Println(errDb)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handlegetMedic(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		re := regexp.MustCompile(`\/[a-z]+\/[0-9]+$`)
		fmt.Println(r.URL.Path)
		if re.MatchString(r.URL.Path) {
			re = regexp.MustCompile(`\/[a-z]+\/`)
			numberS := re.Split(r.URL.Path, 2)[1]
			re = regexp.MustCompile(`[0-9]+`)
			byt := re.FindIndex([]byte(numberS))
			b := []byte(numberS)
			numberS = string(b[byt[0]:byt[1]])
			i, err := strconv.Atoi(numberS)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			stmt, errDb := db.DB.Query("SELECT id,name,lastname, crm , hourend, hourinit,city,state,price, areaofocupation  From medico where id = $1", i)
			if errDb != nil {
				fmt.Println(errDb)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			resp := models.Response{}
			for stmt.Next() {
				cons := models.Medico{}
				cons.Id = i
				stmt.Scan(&cons.Id,&cons.Name, &cons.LastName, &cons.Crm, &cons.HourEnd, &cons.HourInit, &cons.City, &cons.State, &cons.Price, &cons.AreaOfOcupation)
				resp.Data = append(resp.Data, cons)

			}
			bJson, _ := json.Marshal(resp)
			w.Write(bJson)
		} else {

			w.WriteHeader(http.StatusBadRequest)
		}

	}
}

func handleRecepcionista(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
	if r.Method == "POST" {
		var consulta models.Consulta
		dec := json.NewDecoder(r.Body)
		dec.Decode(&consulta)
		fmt.Println(consulta.MedicoID)
		stmt1, _ := db.DB.Query(`SELECT hourend From Medico where id=$1`, consulta.MedicoID)
		stmt1.Next()
		var hourBuffer string
		stmt1.Scan(&hourBuffer)
		r, flag := utils.BiggerThen(consulta.HourInit, hourBuffer)
		if flag {
			fmt.Println("Op1")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if strings.Compare(r, consulta.HourEnd) == 0 {
			fmt.Println("Op2")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		r, flag = utils.BiggerThen(consulta.HourEnd, hourBuffer)
		if strings.Compare(r, consulta.HourEnd) == 1 {
			fmt.Println("Op3")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		stmt, err := db.DB.Query(`SELECT Count(*) From consulta Where medicId=$1 and day = $2 and hourinit>=CAST($3 AS TIME) and hourinit < CAST($3 AS TIME) + interval '25 minutes'`, consulta.MedicoID, consulta.Day, consulta.HourInit)
		if err == nil {
			for stmt.Next() {
				var count int
				stmt.Scan(&count)
				if count == 0 {
					fmt.Println(consulta.ClientID)
					_, err := db.DB.Query("Insert into Consulta (clientid,medicid, status, effected,price,day,hourinit, hourend) values($1, $2,$3,$4, $5,$6,CAST($7 AS TIME),CAST($8 AS TIME))", consulta.ClientID, consulta.MedicoID, consulta.Status, consulta.Effected, consulta.Price, consulta.Day, consulta.HourInit, consulta.HourEnd)
					if err != nil {
						fmt.Println(err)
						w.WriteHeader(http.StatusBadRequest)
					}
				} else {
					fmt.Println("Op5")
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			fmt.Println("Op6")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
func handleMedicoSingup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "POST" {
		var consulta models.Medico
		dec := json.NewDecoder(r.Body)
		dec.Decode(&consulta)
		fmt.Println("Aqui " + consulta.Crm)
		stmt, err := db.DB.Query(`SELECT Count(*) From Medico Where id=$1`, consulta.Id)

		if err == nil {
			for stmt.Next() {
				var count int
				stmt.Scan(&count)

				fmt.Println(count)
				if count == 0 {
					_, err := db.DB.Query("Insert into Medico (crm, city , state, areaofocupation, name, lastname, price, hourinit, hourend) values($1,$2,$3,$4,$5,$6,$7, CAST($8 AS TIME), CAST($9 AS TIME))", consulta.Crm, consulta.City, consulta.State, consulta.AreaOfOcupation, consulta.Name, consulta.LastName, consulta.Price, consulta.HourInit, consulta.HourEnd)
					if err != nil {
						fmt.Println(err)
						w.WriteHeader(http.StatusBadRequest)
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func handleUpdateConsulta(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "POST" {
		var consulta models.Consulta
		dec := json.NewDecoder(r.Body)
		dec.Decode(&consulta)
		stmt, err := db.DB.Query(`SELECT Count(*) From consulta Where id=$1`, &consulta.Id)
		if err == nil {
			for stmt.Next() {
				var count int
				stmt.Scan(&count)
				if count == 1 {
					_, err := db.DB.Query("Update consulta SET price=$1, day=CAST($2 AS DATE), hourend=CAST($3 AS TIME), hourinit=CAST($4 AS TIME), clienteid=$5, medicid=$6, effected=$7, status=$8 where id=$9", &consulta.Price, &consulta.Day, &consulta.HourEnd, &consulta.HourInit, &consulta.ClientID, &consulta.MedicoID, &consulta.Effected, &consulta.Status, &consulta.Id)
					if err != nil {
						fmt.Println(err)
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					fmt.Println("Query atualizada com Sucesso")
				} else {
					fmt.Println(err)
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		fmt.Print("Aqui")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleClienteSingup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "POST" {
		var consulta models.Cliente
		dec := json.NewDecoder(r.Body)
		dec.Decode(&consulta)
		stmt, err := db.DB.Query(`SELECT Count(*) From Medico Where id=$1`, consulta.Id)

		if err == nil {
			for stmt.Next() {
				var count int
				stmt.Scan(&count)

				fmt.Println(count)
				if count == 0 {
					_, err := db.DB.Query("Insert into Cliente (cpf, city , state, name, lastname) values($1,$2,$3,$4,$5)", &consulta.Cpf, &consulta.City, &consulta.State, &consulta.Name, &consulta.LastName)
					if err != nil {
						fmt.Println(err)
						w.WriteHeader(http.StatusBadRequest)
					}
				} else {
					fmt.Println(err)
					w.WriteHeader(http.StatusBadRequest)
				}
			}
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func Router() {

	db = new(database.Db)
	db.OpenDb()
	http.HandleFunc("/client", handleClienteSingup)
	http.HandleFunc("/medic", handleMedicoSingup)
	http.HandleFunc("/getmedic/", handlegetMedic)
	http.HandleFunc("/removeconsulta/", handleRemoveConsulta)
	http.HandleFunc("/updateconsulta", handleUpdateConsulta)
	http.HandleFunc("/getclient/", handlegetClient)
	http.HandleFunc("/medics", handleMedics)
	http.HandleFunc("/consultas/medic/", handleConsultasM)
	http.HandleFunc("/consulta", handleConsultasSingup)
	http.HandleFunc("/consultas/client/", handleConsultasC)
	http.HandleFunc("/recepcionista/", handleRecepcionista)
	http.ListenAndServe(":8084", nil)

}
