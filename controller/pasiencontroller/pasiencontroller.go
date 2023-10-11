package pasiencontroller

import (
	"crud-back/entities"
	"crud-back/libraries"
	"crud-back/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var validation = libraries.NewValidation()
var pasienModel = models.NewPasienModel()

func Index(w http.ResponseWriter, r *http.Request) { //menampilkan semua data
	pasien, err := pasienModel.FindAll()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(pasien)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var pasien entities.Pasien
	err := json.NewDecoder(r.Body).Decode(&pasien)
	if err != nil {
		fmt.Println(err)
	}

	vErrors := validation.Struct(pasien)

	if vErrors != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := libraries.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  vErrors,
		}
		json.NewEncoder(w).Encode(webResponse)
	} else {
		response := pasienModel.Create(pasien)
		json.NewEncoder(w).Encode(response)
	}

}
