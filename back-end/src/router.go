package router

import (
	"net/http"
)

func handleMedics(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodGet {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// 	w.Header().Add("Name", "Eliabe")
	// }
	// w.WriteHeader(http.StatusMethodNotAllowed)
	// w.Header().Add("Name", "Eliabe")
	// fmt.Print(w.Header().Get("Name"))
	// fmt.Print(r.Header.Get("Key"))
	// w.Header().Del("Data")
	// js, _ := json.Marshal(r.Response)

	// io.WriteString(w, string(js))

}
func Router() {
	http.HandleFunc("/medics", handleMedics)

	http.ListenAndServe(":8084", nil)

}
