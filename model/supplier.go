package model

import (
	"database/sql"
	"log"
)

type Supplier struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ContactName string `json:"contact_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}

func (input *Supplier) GetAll(db *sql.DB) []Supplier {
	var (
		supplier  Supplier
		suppliers []Supplier
	)

	rows, err := db.Query("SELECT id, name, contact_name, address, city, postal_code, country, phone FROM suppliers")
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&supplier.Id, &supplier.Name, &supplier.ContactName, &supplier.Address, &supplier.City, &supplier.PostalCode, &supplier.Country, &supplier.Phone); err != nil {
			log.Fatal(err.Error())
		} else {
			suppliers = append(suppliers, supplier)
		}
	}

	return suppliers
}

func (input *Supplier) GetById(db *sql.DB, id int) Supplier {
	var (
		supplier Supplier
	)

	err := db.QueryRow("SELECT id, name, contact_name, address, city, postal_code, country, phone FROM suppliers WHERE id = ?", id).Scan(&supplier.Id, &supplier.Name, &supplier.ContactName, &supplier.Address, &supplier.City, &supplier.PostalCode, &supplier.Country, &supplier.Phone)
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	return supplier
}

func (input *Supplier) Create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO suppliers(name, contact_name, address, city, postal_code, country, phone) VALUES (?,?,?,?,?,?,?)", input.Name, input.ContactName, input.Address, input.City, input.PostalCode, input.Country, input.Phone)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Supplier) Update(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE suppliers SET name = ?, contact_name = ?, address = ?, city = ?, postal_code = ?, country = ?, phone = ? WHERE id = ?", input.Name, input.ContactName, input.Address, input.City, input.PostalCode, input.Country, input.Phone, id)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Supplier) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM suppliers WHERE id = ?", input.Id)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Supplier) FindByName(db *sql.DB, name string) []Supplier {
	var (
		supplier  Supplier
		suppliers []Supplier
	)

	rows, err := db.Query("SELECT id, name, contact_name, address, city, postal_code, country, phone FROM suppliers WHERE name like '%" + name + "%'")
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&supplier.Id, &supplier.Name, &supplier.ContactName, &supplier.Address, &supplier.City, &supplier.PostalCode, &supplier.Country, &supplier.Phone); err != nil {
			log.Fatal(err.Error())
		} else {
			suppliers = append(suppliers, supplier)
		}
	}

	return suppliers
}

func (input *Supplier) ValidateId(db *sql.DB) bool {
	var (
		id int
	)

	err := db.QueryRow("SELECT id FROM suppliers WHERE id = ?", input.Id).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
		return false
	}

	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

func (input *Supplier) ValidateInput() string {

	// if input.Id <= 0 {
	// 	return "ID tidak boleh kosong!"
	// }

	if input.Name == "" {
		return "Name tidak boleh kosong"
	}

	if input.ContactName == "" {
		return "Contact name tidak boleh kosong"
	}

	if input.Address == "" {
		return "Address tidak boleh kosong"
	}

	if input.City == "" {
		return "City tidak boleh kosong"
	}

	if input.PostalCode == "" {
		return "Postal code tidak boleh kosong"
	}

	if input.Country == "" {
		return "Country tidak boleh kosong"
	}

	if input.Phone == "" {
		return "Phone tidak boleh kosong"
	}

	return ""
}
