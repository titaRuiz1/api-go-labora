package main

import (
	"fmt"
	"labora-api/controller"
	"labora-api/service"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

)



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
