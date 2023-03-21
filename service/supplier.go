package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maxsol/model"
	"maxsol/utility"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Supplier struct {
}

func (input *Supplier) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		response  model.Response
		supplier  model.Supplier
		suppliers []model.Supplier
	)

	db := utility.Connect()
	defer db.Close()

	supplier = model.Supplier{}

	suppliers = supplier.GetAll(db)

	if len(suppliers) == 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Supplier{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil seluruh data supplier"
		response.Data = suppliers
	}

	fmt.Println("Endpoint Hit: get all suppliers")
	json.NewEncoder(w).Encode(response)
}

func (input *Supplier) GetById(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		supplier model.Supplier
	)

	db := utility.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	supplier = model.Supplier{}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	supplier = supplier.GetById(db, idInt)

	if supplier.Id <= 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Supplier{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil data supplier berdasarkan id"
		response.Data = []model.Supplier{supplier}
	}

	fmt.Println("Endpoint Hit: get supplier by id")
	json.NewEncoder(w).Encode(response)
}

func (input *Supplier) Create(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		supplier model.Supplier
	)

	db := utility.Connect()
	defer db.Close()

	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &supplier)

	message := supplier.ValidateInput()

	if message != "" {
		response.Status = "Gagal"
		response.Message = message
		response.Data = []model.Supplier{}
	} else {

		err := supplier.Create(db)
		if err != nil {
			log.Print(err)
		}

		response.Status = "Berhasil"
		response.Message = "Buat data supplier berhasil"
		response.Data = []model.Supplier{supplier}
	}

	fmt.Println("Endpoint Hit: create supplier")
	json.NewEncoder(w).Encode(response)
}

func (input *Supplier) Update(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		supplier model.Supplier
	)

	db := utility.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	supplier = model.Supplier{
		Id: idInt,
	}

	exist := supplier.ValidateId(db)
	if exist {
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &supplier)

		message := supplier.ValidateInput()

		if message != "" {
			response.Status = "Gagal"
			response.Message = message
			response.Data = []model.Supplier{}
		} else {
			err = supplier.Update(db, idInt)
			if err != nil {
				log.Print(err)
			}

			response.Status = "Berhasil"
			response.Message = "Ubah data supplier berhasil"
			response.Data = []model.Supplier{supplier}
		}
	} else {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Supplier{}
	}

	fmt.Println("Endpoint Hit: update supplier")
	json.NewEncoder(w).Encode(response)
}

func (input *Supplier) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		supplier model.Supplier
	)

	vars := mux.Vars(r)
	id := vars["id"]

	db := utility.Connect()
	defer db.Close()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	supplier = model.Supplier{
		Id: idInt,
	}

	exist := supplier.ValidateId(db)
	if exist {
		err = supplier.Delete(db)
		if err != nil {
			log.Print(err)
		}

		response.Status = "Berhasil"
		response.Message = "Hapus data supplier berhasil"
		response.Data = []model.Supplier{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Supplier{}
	}

	fmt.Println("Endpoint Hit: delete supplier")
	json.NewEncoder(w).Encode(response)
}

func (input *Supplier) FindByName(w http.ResponseWriter, r *http.Request) {
	var (
		response  model.Response
		supplier  model.Supplier
		suppliers []model.Supplier
	)

	vars := mux.Vars(r)
	name := vars["name"]

	db := utility.Connect()
	defer db.Close()

	supplier = model.Supplier{}

	suppliers = supplier.FindByName(db, name)

	if len(suppliers) == 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Supplier{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil data supplier berdasarkan nama"
		response.Data = suppliers
	}

	fmt.Println("Endpoint Hit: get suppliers by name")
	json.NewEncoder(w).Encode(response)
}
