package handlers

import (
	"encoding/json"
	"myapi/internal/config"
	"myapi/internal/models"
	"myapi/internal/repositories"
	"myapi/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListItensHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	repository := repositories.NewItemRepository()
	items, err := repository.ListAll()
	if err != nil {
		http.Error(w, "Erro ao buscar itens", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func GetItenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func GetItenByCodigoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	cod := vars["code"]
	if cod == "" {
		http.Error(w, "Código não fornecido", http.StatusBadRequest)
		return
	}
	var item models.Item
	if err := config.DB.Where("codigo = ?", cod).First(&item).Error; err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateItenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o item", http.StatusBadRequest)
		return
	}

	newItem, err := services.CreateItem(&item)
	if err != nil {
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}

	if err := config.DB.Create(&newItem).Error; err != nil {
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newItem)
}

func UpdateItenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Erro ao decodificar o item", http.StatusBadRequest)
		return
	}
	if err := config.DB.Save(&item).Error; err != nil {
		http.Error(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func DeleteItenHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		http.Error(w, "Erro ao deletar o item", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Item deletado com sucesso"))
}
