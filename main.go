package main

import (
	"waysbuck/database"
	"waysbuck/pkg/mysql"
	"waysbuck/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()
	
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

