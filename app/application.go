package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yessikaplata/go-api-rest/app/config"
	"github.com/yessikaplata/go-api-rest/db"
	"github.com/yessikaplata/go-api-rest/handler"
	"github.com/yessikaplata/go-api-rest/repository"
	"github.com/yessikaplata/go-api-rest/service"
)

func Start() {
	var (
		err      error
		database *sql.DB
		cnf      config.Config
	)
	func() {
		cnf, err = config.NewParsedConfig()
		if err != nil {
			return
		}
		dbConfig := db.DatabaseConfiguraton{
			Host:         cnf.Database.Host,
			Port:         cnf.Database.Port,
			User:         cnf.Database.User,
			Password:     cnf.Database.Password,
			DatabaseName: cnf.Database.DatabaseName,
		}

		database, err = db.Connect(dbConfig)
	}()
	if err != nil {
		panic(err)
	}
	mux := mux.NewRouter()
	registerEndpoint(mux, database)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cnf.ServerPort), mux))

}

func registerEndpoint(r *mux.Router, db *sql.DB) {
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	r.HandleFunc("/api/users", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/api/users/{id}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/api/users", handler.GetAll()).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{id}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/api/users/{id}", handler.Delete()).Methods(http.MethodDelete)
}
