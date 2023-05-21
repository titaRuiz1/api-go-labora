package service

import (
	"fmt"
	"labora-api/model"
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