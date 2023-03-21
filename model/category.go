package model

import (
	"database/sql"
	"fmt"
	"log"
)

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (input *Category) GetAll(db *sql.DB) []Category {
	var (
		category   Category
		categories []Category
	)

	rows, err := db.Query("SELECT id, name, description FROM categories")
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&category.Id, &category.Name, &category.Description); err != nil {
			log.Fatal(err.Error())
		} else {
			categories = append(categories, category)
		}
	}

	return categories
}

func (input *Category) GetById(db *sql.DB, id int) Category {
	var (
		category Category
	)

	err := db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id).Scan(&category.Id, &category.Name, &category.Description)
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	return category
}

func (input *Category) Create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO categories(name, description) VALUES (?,?)", input.Name, input.Description)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Category) Update(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE categories SET name = ?, description = ? WHERE id = ?", input.Name, input.Description, id)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Category) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM categories WHERE id = ?", input.Id)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (input *Category) FindByName(db *sql.DB, name string) []Category {
	var (
		category   Category
		categories []Category
	)

	fmt.Println(name)

	rows, err := db.Query("SELECT id, name, description FROM categories WHERE name like '%" + name + "%'")
	if err != nil && err != sql.ErrNoRows {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&category.Id, &category.Name, &category.Description); err != nil {
			log.Fatal(err.Error())
		} else {
			categories = append(categories, category)
		}
	}

	return categories
}

func (input *Category) ValidateId(db *sql.DB) bool {
	var (
		id int
	)

	err := db.QueryRow("SELECT id FROM categories WHERE id = ?", input.Id).Scan(&id)
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

func (input *Category) ValidateInput() string {

	// if input.Id <= 0 {
	// 	return "ID tidak boleh kosong!"
	// }

	if input.Name == "" {
		return "Name tidak boleh kosong"
	}

	if input.Description == "" {
		return "Description tidak boleh kosong"
	}

	return ""
}
