package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vivekganesan01/go-rest-memcached/internal/pkg/util"
)

type Error struct {
	Message string `json:"error"`
}

func init() {
	log.Println("EnvReady: Booted.")
}

func main() {
	db, err := util.NewPostgreSQL()
	if err != nil {
		log.Fatalf("Not able to initialize psgl connection", err)
	}
	defer db.Close()
	m, err := util.NewMemCached()
	if err != nil {
		log.Fatalf("Not able to initialize mem cache connection", err)
	}
	router := mux.NewRouter()
	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}
	router.HandleFunc("/names/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		val, err := m.Get(id)
		if err == nil {
			log.Printf("got the key from memcache")
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(val)
			return
		}
		name, err := db.FindByValue(id)
		if err != nil {
			renderJSON(rw, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		log.Printf("got the key from db")
		_ = m.Set(name)
		renderJSON(rw, &name, http.StatusOK)
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server :8080")
	log.Fatal(srv.ListenAndServe())
}
