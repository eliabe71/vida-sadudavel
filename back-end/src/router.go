package router

import (
	"encoding/json"
	"io"
	"net/http"
)

func e(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	r.Response = new(http.Response)
	r.Response.Status = "Error"
	r.Response.StatusCode = 404
	js, _ := json.Marshal(r.Response)

	io.WriteString(w, string(js))

}
func Eliabe() {
	http.HandleFunc("/c", e)
	http.HandleFunc("/ca", e)
	http.ListenAndServe(":8084", nil)

}
