package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/midorigreen/goapi-architecture/form"
	"github.com/midorigreen/goapi-architecture/hub"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	f, err := form.NewTask(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("parameter error\n"))
		return
	}
	m, err := hub.TaskHub(f)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
		return
	}

	j, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "json marshal error")
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
