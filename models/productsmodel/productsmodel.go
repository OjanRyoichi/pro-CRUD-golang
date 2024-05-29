package productsmodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll() []entities.Products {
	rows, err := config.DB.Query(`SELECT products.id, products.name, categories.name as categories_name, products.stock, products.description FROM products JOIN categories ON products.category_id = categories.id`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []entities.Products

	for rows.Next() {
		var product entities.Products
		err = rows.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products
}

func Create(product entities.Products) bool {
	result, err := config.DB.Exec(`INSERT INTO products (name, category_id, stock, description) VALUES (?, ?, ?, ?)`, product.Name, product.Category.Id, product.Stock, product.Description)
	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func Detail(id int) entities.Products {
	row := config.DB.QueryRow(`SELECT products.id, products.name, categories.name as categories_name, products.stock, products.description FROM products JOIN categories ON products.category_id = categories.id WHERE products.id = ?`, id)

	var product entities.Products

	// err := row.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description)
	// if err != nil {
	// 	panic(err)
	// }
	if err := row.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description); err != nil{
		panic(err.Error())
	}

	return product
}

func Update(id int, product entities.Products) bool {
	query, err := config.DB.Exec(`UPDATE products SET name = ?, category_id = ?, stock = ?, description = ? WHERE id = ?`, product.Name, product.Category.Id, product.Stock, product.Description, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error{
	_, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}
