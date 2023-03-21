package model

import (
	"database/sql"
	"log"
)

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	SupplierId int     `json:"supplier_id"`
	CategoryId int     `json:"category_id"`
	Unit       string  `json:"unit"`
	Price      float64 `json:"price"`
}

func (input *Product) GetAll(db *sql.DB) []Product {
	var (
		product  Product
		products []Product
	)

	rows, err := db.Query("SELECT id, name, supplier_id, category_id, unit, price FROM products")
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.SupplierId, &product.CategoryId, &product.Unit, &product.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			products = append(products, product)
		}
	}

	return products
}

func (input *Product) GetById(db *sql.DB, id int) Product {
	var (
		product Product
	)

	err := db.QueryRow("SELECT id, name, supplier_id, category_id, unit, price FROM products WHERE id = ?", id).Scan(&product.Id, &product.Name, &product.SupplierId, &product.CategoryId, &product.Unit, &product.Price)
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	return product
}

func (input *Product) Create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO products(name, supplier_id, category_id, unit, price) VALUES (?,?,?,?,?)", input.Name, input.SupplierId, input.CategoryId, input.Unit, input.Price)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Product) Update(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE products SET name = ?, supplier_id = ?, category_id = ?, unit = ?, price = ? WHERE id = ?", input.Name, input.SupplierId, input.CategoryId, input.Unit, input.Price, id)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Product) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id = ?", input.Id)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Product) FindByName(db *sql.DB, name string) []Product {
	var (
		product  Product
		products []Product
	)

	rows, err := db.Query("SELECT id, name, supplier_id, category_id, unit, price FROM products WHERE name like '%" + name + "%'")
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.SupplierId, &product.CategoryId, &product.Unit, &product.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			products = append(products, product)
		}
	}

	return products
}

func (input *Product) ValidateId(db *sql.DB) bool {
	var (
		id int
	)

	err := db.QueryRow("SELECT id FROM products WHERE id = ?", input.Id).Scan(&id)
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

func (input *Product) ValidateInput() string {

	// if input.Id <= 0 {
	// 	return "ID tidak boleh kosong!"
	// }

	if input.Name == "" {
		return "Name tidak boleh kosong"
	}

	if input.SupplierId <= 0 {
		return "Supplier ID tidak boleh kosong"
	}

	if input.CategoryId <= 0 {
		return "Category ID tidak boleh kosong"
	}

	if input.Unit == "" {
		return "Unit tidak boleh kosong"
	}

	if input.Price <= 0 {
		return "Price tidak boleh kosong"
	}

	return ""
}
