package main

import (
	"fmt"
	"labora-api/controller"
	"labora-api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	// "time"
	// "sync"
)

// func getItemDetails(w http.ResponseWriter, r *http.Request) {
// 	// Obtenemos todos los items
// 	allItems := items

// 	// Creamos un WaitGroup para esperar a que todas las gorutinas terminen
// 	var wg sync.WaitGroup

// 	// Creamos un canal para recibir los resultados de las gorutinas
// 	results := make(chan ItemDetails, len(allItems))

// 	// Por cada item, iniciamos una gorutina que busca información adicional
// 	// y la almacena en una estructura de datos
// 	for _, item := range allItems {
// 		wg.Add(1)
// 		go func(item Item) {
// 			defer wg.Done()

// 			// Simulamos la búsqueda de información adicional
// 			time.Sleep(100 * time.Millisecond)
// 			details := "Details for " + item.Name

// 			// Almacenamos el resultado en el canal
// 			results <- ItemDetails{Item: item, Details: details}
// 		}(item)
// 	}

// 	// Esperamos a que todas las gorutinas terminen
// 	wg.Wait()

// 	// Cerramos el canal de resultados para que la función range a continuación
// 	// termine cuando todos los resultados hayan sido recibidos
// 	close(results)

// 	// Creamos un slice de ItemDetails para almacenar los resultados
// 	var itemsDetails []ItemDetails

// 	// Recorremos el canal de resultados y almacenamos los elementos en el slice
// 	for res := range results {
// 		itemsDetails = append(itemsDetails, res)
// 	}

// 	// Codificamos la respuesta en formato JSON y la enviamos al cliente
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(itemsDetails)
// }

func main() {
    service.UpDb()
	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetAllItems).Methods("GET")
  router.HandleFunc("/get_items", controller.ObtenerItems).Methods("GET")
	router.HandleFunc("/items/{id}", controller.GetItemByID).Methods("GET")
	router.HandleFunc("/items/{name}", controller.SearchItems).Methods("GET")
	router.HandleFunc("/items", controller.CreateNewItem).Methods("POST")
	router.HandleFunc("/items/{id}", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")
	// router.HandleFunc("/items/details", getItemDetails).Methods("GET")

	// Configurar el middleware CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// Agregar el middleware CORS a todas las rutas
	handler := corsOptions.Handler(router)

    service.Db.PingOrDie()
	fmt.Println("Server is listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))


}
