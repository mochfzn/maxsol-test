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

type Product struct {
}

func (input *Product) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		product  model.Product
		products []model.Product
	)

	db := utility.Connect()
	defer db.Close()

	product = model.Product{}

	products = product.GetAll(db)

	if len(products) == 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Product{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil seluruh data category"
		response.Data = products
	}

	fmt.Println("Endpoint Hit: get all products")
	json.NewEncoder(w).Encode(response)
}

func (input *Product) GetById(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		product  model.Product
	)

	db := utility.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	product = model.Product{}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	product = product.GetById(db, idInt)

	if product.Id <= 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil data product berdasarkan id"
		response.Data = []model.Product{product}
	}

	fmt.Println("Endpoint Hit: get product by id")
	json.NewEncoder(w).Encode(response)
}

func (input *Product) Create(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		product  model.Product
	)

	db := utility.Connect()
	defer db.Close()

	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &product)

	message := product.ValidateInput()

	if message != "" {
		response.Status = "Gagal"
		response.Message = message
		response.Data = []model.Product{}
	} else {
		err := product.Create(db)
		if err != nil {
			log.Print(err)
		}

		response.Status = "Berhasil"
		response.Message = "Buat data product berhasil"
		response.Data = []model.Product{product}
	}

	fmt.Println("Endpoint Hit: create product")
	json.NewEncoder(w).Encode(response)
}

func (input *Product) Update(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		product  model.Product
	)

	db := utility.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	product = model.Product{
		Id: idInt,
	}

	exist := product.ValidateId(db)
	if exist {
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &product)

		message := product.ValidateInput()

		if message != "" {
			response.Status = "Gagal"
			response.Message = message
			response.Data = []model.Product{}
		} else {
			err = product.Update(db, idInt)
			if err != nil {
				log.Print(err)
			}

			response.Status = "Berhasil"
			response.Message = "Ubah data product berhasil"
			response.Data = []model.Product{product}
		}
	} else {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Product{}
	}

	fmt.Println("Endpoint Hit: update product")
	json.NewEncoder(w).Encode(response)
}

func (input *Product) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		product  model.Product
	)

	vars := mux.Vars(r)
	id := vars["id"]

	db := utility.Connect()
	defer db.Close()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	product = model.Product{
		Id: idInt,
	}

	exist := product.ValidateId(db)
	if exist {
		err = product.Delete(db)
		if err != nil {
			log.Print(err)
		}

		response.Status = "Berhasil"
		response.Message = "Hapus data product berhasil"
		response.Data = []model.Product{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Product{}
	}

	fmt.Println("Endpoint Hit: delete product")
	json.NewEncoder(w).Encode(response)
}

func (input *Product) FindByName(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		product  model.Product
		products []model.Product
	)

	vars := mux.Vars(r)
	name := vars["name"]

	db := utility.Connect()
	defer db.Close()

	product = model.Product{}

	products = product.FindByName(db, name)

	if len(products) == 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil data product berdasarkan nama"
		response.Data = products
	}

	fmt.Println("Endpoint Hit: get product by name")
	json.NewEncoder(w).Encode(response)
}
