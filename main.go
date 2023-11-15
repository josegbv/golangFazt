package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josegbv/go-apirest-fazt/db"
	"github.com/josegbv/go-apirest-fazt/models"
	"github.com/josegbv/go-apirest-fazt/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//task routes
	r.HandleFunc("/task", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
