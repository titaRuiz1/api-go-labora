package controller

import(
	"net/http"
	"labora-api/model"
	"labora-api/service"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {

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

	var itemsToReturn []model.Item
	if start >= len(model.Items) {
		itemsToReturn = []model.Item{}
	} else if end >= len(model.Items) {
		itemsToReturn = model.Items[start:]
	} else {
		itemsToReturn = model.Items[start:end]
	}

	json.NewEncoder(w).Encode(itemsToReturn)
}

func ObtenerItems(w http.ResponseWriter, _ *http.Request) {
	items, err := service.GetItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al obtener los items"))
		return
	}

	json.NewEncoder(w).Encode(items)
}

func CreateNewItem(w http.ResponseWriter, r *http.Request) {

	var newItem model.Item

	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	newID := len(model.Items) + 1
	newItem = model.Item{
		CustomerName: newItem.CustomerName,
	}
	model.Items = append(model.Items, newItem)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Item created with ID %d", newID)
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, item := range model.Items {
		if fmt.Sprint(item.ID) == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func SearchItems(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("name")

	var results []model.Item
	for _, item := range model.Items {
		if strings.Contains(item.CustomerName, searchTerm) {
			results = append(results, item)
		}
	}

	json.NewEncoder(w).Encode(results)
}



func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	var updatedItem model.Item

	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if id < 1 || id > len(model.Items) {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if id > 0 && id <= len(model.Items) {
		model.Items[id-1].CustomerName = updatedItem.CustomerName
		json.NewEncoder(w).Encode(model.Items[id-1])
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Item update ok")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if id < 1 || id > len(model.Items) {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Eliminar el item del slice de items
	model.Items = append(model.Items[:id-1], model.Items[id:]...)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item with ID %d has been deleted", id)
}
