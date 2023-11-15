package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josegbv/go-apirest-fazt/db"
	"github.com/josegbv/go-apirest-fazt/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task

	db.DB.Find(&task)

	json.NewEncoder(w).Encode(&task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	CreateTask := db.DB.Create(&task)
	Error := CreateTask.Error

	if Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tarea no encontrada"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
