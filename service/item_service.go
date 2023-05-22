package service

import (
	"fmt"
	"labora-api/model"
	"log"
)

// GetItems obtiene todos los items de la tabla 'items' de la base de datos.
// Retorna una lista de struct 'models.Item' y un error en caso de que haya ocurrido alguno.
func GetItems() ([]model.Item, error) {
	items := make([]model.Item, 0)
	rows, err := Db.Query("SELECT * FROM items")
	if err != nil {
		
		fmt.Println(err)
		return nil, err
	}
	
	defer rows.Close()

	// Itera sobre cada fila en 'rows' y crea una instancia de 'models.Item' con los valores de cada columna.
	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}


func CreateItem(item model.Item) error {
	// Iniciamos una transacción
	tx, err := Db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Ejecutamos la consulta SQL para insertar un nuevo registro en la tabla 'items' dentro de la transacción.
	_, err = tx.Exec(SQLInsertNewItem,item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price)
	if err != nil {
		// Si algo salió mal, hacemos un rollback de la transacción
		tx.Rollback()
		log.Fatal(err)
	}

	// Si todo salió bien, hacemos commit de la transacción
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}


const SQLInsertNewItem =`INSERT INTO items
												(customer_name,
													order_date,
												 	product,
													quantity, 
													price) 
												 	VALUES (?, ?, ?, ?, ?)`
