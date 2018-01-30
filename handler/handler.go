package handler

import (
	"log"
	"net/http"

	"github.com/midorigreen/goapi-architecture/form"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	f, err := form.NewTask(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("parameter error\n"))
		return
	}
	log.Println(f)
	w.Write([]byte("Hello Wolrd\n"))
}
