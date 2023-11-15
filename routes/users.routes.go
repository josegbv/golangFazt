package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josegbv/go-apirest-fazt/db"
	"github.com/josegbv/go-apirest-fazt/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	for i, _ := range users {
		db.DB.Model(&users[i]).Association("Tasks").Find(&users[i].Tasks)
	}

	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	error := createdUser.Error

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	json.NewEncoder(w).Encode(&user)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
