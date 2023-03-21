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

type Category struct {
}

func (input *Category) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		response   model.Response
		category   model.Category
		categories []model.Category
	)

	db := utility.Connect()
	defer db.Close()

	category = model.Category{}

	categories = category.GetAll(db)

	if len(categories) == 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil seluruh data category"
		response.Data = categories
	}

	fmt.Println("Endpoint Hit: get all categories")
	json.NewEncoder(w).Encode(response)
}

func (input *Category) GetById(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		category model.Category
	)

	db := utility.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	category = model.Category{}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	category = category.GetById(db, idInt)

	if category.Id <= 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil data category berdasarkan id"
		response.Data = []model.Category{category}
	}

	fmt.Println("Endpoint Hit: get category by id")
	json.NewEncoder(w).Encode(response)
}

func (input *Category) Create(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		category model.Category
	)

	db := utility.Connect()
	defer db.Close()

	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &category)

	message := category.ValidateInput()

	if message != "" {
		response.Status = "Gagal"
		response.Message = message
		response.Data = []model.Category{}
	} else {
		err := category.Create(db)
		if err != nil {
			log.Print(err)
		}

		response.Status = "Berhasil"
		response.Message = "Buat data category berhasil"
		response.Data = []model.Category{category}
	}

	fmt.Println("Endpoint Hit: create product")
	json.NewEncoder(w).Encode(response)
}

func (input *Category) Update(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		category model.Category
	)

	db := utility.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	category = model.Category{
		Id: idInt,
	}

	exist := category.ValidateId(db)
	if exist {
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &category)

		message := category.ValidateInput()

		if message != "" {
			response.Status = "Gagal"
			response.Message = message
			response.Data = []model.Category{}
		} else {
			err = category.Update(db, idInt)
			if err != nil {
				log.Print(err)
			}

			response.Status = "Berhasil"
			response.Message = "Ubah data category berhasil"
			response.Data = []model.Category{category}
		}
	} else {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	}

	fmt.Println("Endpoint Hit: update category")
	json.NewEncoder(w).Encode(response)
}

func (input *Category) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		response model.Response
		category model.Category
	)

	vars := mux.Vars(r)
	id := vars["id"]

	db := utility.Connect()
	defer db.Close()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}

	category = model.Category{
		Id: idInt,
	}

	exist := category.ValidateId(db)
	if exist {
		err = category.Delete(db)
		if err != nil {
			log.Print(err)
		}

		response.Status = "Berhasil"
		response.Message = "Hapus data product berhasil"
		response.Data = []model.Category{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	}

	fmt.Println("Endpoint Hit: delete product")
	json.NewEncoder(w).Encode(response)
}

func (input *Category) FindByName(w http.ResponseWriter, r *http.Request) {
	var (
		response   model.Response
		category   model.Category
		categories []model.Category
	)

	vars := mux.Vars(r)
	name := vars["name"]

	db := utility.Connect()
	defer db.Close()

	category = model.Category{}

	categories = category.FindByName(db, name)

	if len(categories) == 0 {
		response.Status = "Berhasil"
		response.Message = "Tidak ada data"
		response.Data = []model.Category{}
	} else {
		response.Status = "Berhasil"
		response.Message = "Ambil seluruh data category"
		response.Data = categories
	}

	fmt.Println("Endpoint Hit: get categories by name")
	json.NewEncoder(w).Encode(response)
}
