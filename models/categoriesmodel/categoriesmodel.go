package categoriesmodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Categories {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []entities.Categories
	for rows.Next() {
		var category entities.Categories
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Categories) bool {
	result, err := config.DB.Exec(`INSERT INTO categories (name) VALUES (?)`, category.Name)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Categories {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id=?`, id)

	var category entities.Categories
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}
	return category
}

func Update(id int, category entities.Categories) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name=? WHERE id=?`, category.Name, id)
	if err != nil {
		panic(err)
	}

	res, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return res > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id=?`, id)
	return err
}
