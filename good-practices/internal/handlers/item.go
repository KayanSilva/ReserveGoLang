package handlers

import (
	"encoding/json"
	"myapi/internal/repositories"
	"myapi/internal/services"
	"myapi/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListItems(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewItemRepository()
	items, err := repository.ListAll()
	if err != nil {
		utils.RespondWithError(w, "Erro ao listar os itens", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		utils.RespondWithError(w, "Erro ao codigar os itens", http.StatusInternalServerError)
		return
	}
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.RespondWithError(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	item, err := repository.GetByID(id)
	if err != nil {
		utils.RespondWithError(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		utils.RespondWithError(w, "Erro ao codigar os itens", http.StatusInternalServerError)
		return
	}
}

func GetItemByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["codigo"]

	if code == "" {
		utils.RespondWithError(w, "Código não fornecido", http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	item, err := repository.GetByCode(code)
	if err != nil {
		utils.RespondWithError(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		utils.RespondWithError(w, "Erro ao codigar os itens", http.StatusInternalServerError)
		return
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidateItem(r)

	repository := repositories.NewItemRepository()
	createdItem, err := repository.Create(item)
	if err != nil {
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(createdItem)
	if err != nil {
		utils.RespondWithError(w, "Erro ao codigar os itens", http.StatusInternalServerError)
		return
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	item, err := services.DecodeAndValidateItem(r)

	repository := repositories.NewItemRepository()
	if err := repository.Update(item); err != nil {
		utils.RespondWithError(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		utils.RespondWithError(w, "Erro ao codigar os itens", http.StatusInternalServerError)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		utils.RespondWithError(w, "ID não fornecido", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	repository := repositories.NewItemRepository()
	if err := repository.Delete(id); err != nil {
		utils.RespondWithError(w, "Erro ao deletar o items", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write([]byte("Item deletado com sucesso")); err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
	}
}
