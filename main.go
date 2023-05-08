package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"github.com/gorilla/mux"
    "strconv"
    "strings"
    // "time"
    // "sync"
)

type Item struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {1, "Item 1"},
    {2, "Item 2"},
    {3, "Item 3"},
    {4, "Item 4"},
    {5, "Item 5"},
    {6, "Item 6"},
    {7, "Item 7"},
    {8, "Item 8"},
    {9, "Item 9"},
    {10, "Item 10"},
    {11, "Item 11"},
    {12, "Item 12"},
    {13, "Item 13"},
    {14, "Item 44"},
    {15, "Item 15"},
    {16, "Item 16"},
    {17, "Item 17"},
    {18, "Item 18"},
    {19, "Item 19"},
    {20, "Item 20"},
    
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
  	
    page := 1
	itemsPerPage := 5

    query := r.URL.Query()
	
	if itemsPage := query.Get("itemsPerPage"); itemsPage != "" {
		if itemsPerPageNum, err := strconv.Atoi(itemsPage); err == nil {
			itemsPerPage = itemsPerPageNum
		}
        http.Error(w, "Invalid number of items per page", http.StatusBadRequest)
        return
	}

    start := (page - 1) * itemsPerPage
	end := start + itemsPerPage

    var itemsToReturn []Item
	if start >= len(items) {
		itemsToReturn = []Item{}
	} else if end >= len(items) {
		itemsToReturn = items[start:]
	} else {
		itemsToReturn = items[start:end]
	}

    json.NewEncoder(w).Encode(itemsToReturn)
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
    for _, item := range items {
    if fmt.Sprint(item.Id) == id {
        json.NewEncoder(w).Encode(item)
        return
    }
}
    http.Error(w, "Item not found", http.StatusNotFound)
}

func searchItems(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.URL.Query().Get("name")

    var results []Item
    for _, item := range items {
        if strings.Contains(item.Name, searchTerm) {
            results = append(results, item)
        }
    }

    json.NewEncoder(w).Encode(results)
}

func createNewItem(w http.ResponseWriter, r *http.Request) {

    var newItem Item

    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }
    
    newID := len(items) + 1
    newItem = Item{
        Name: newItem.Name,
        }
    items = append(items, newItem)

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Item created with ID %d", newID)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }
    var updatedItem Item

    if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }
    if id < 1 || id > len(items) {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    if id > 0 && id <= len(items) {
        items[id-1].Name = updatedItem.Name
        json.NewEncoder(w).Encode(items[id-1])
    } else {
        http.Error(w, "Item not found", http.StatusNotFound)
    }

    fmt.Fprintf(w, "Item update ok")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if id < 1 || id > len(items) {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    // Eliminar el item del slice de items
    items = append(items[:id-1], items[id:]...)

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Item with ID %d has been deleted", id)
}

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
	router := mux.NewRouter()

    router.HandleFunc("/items", getAllItems).Methods("GET")
    router.HandleFunc("/items/{id}", getItemByID).Methods("GET")
    router.HandleFunc("/items/{name}", searchItems).Methods("GET")
    // router.HandleFunc("/items/details", getItemDetails).Methods("GET")
    router.HandleFunc("/items", createNewItem).Methods("POST")
    router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

    


    fmt.Println("Server is listening on port 8000...")
    http.ListenAndServe(":8000", router)
}

