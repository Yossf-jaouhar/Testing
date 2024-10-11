package test

import (
	"net/http"
	"strconv"
)

type page struct {
	Text string
}
var data = &page{}
// Download handles file download requests.
func Download(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		
		return
	}

	
	result := r.FormValue("arttext")
	if result == "" {
		w.WriteHeader(http.StatusBadRequest)
		
		return
	}

	
	w.Header().Set("Content-Disposition", "attachment; filename=result.txt") 
	w.Header().Set("Content-Type", "text/plain") 
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))  


	w.Write([]byte(result))
}
