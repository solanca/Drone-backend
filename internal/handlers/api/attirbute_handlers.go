package api

import (
	handler "drone-backend/internal/handlers/on-chain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AttributeAPI struct {
    Handler *handler.AttributeHandler
}

func NewAttributeAPI(h *handler.AttributeHandler) *AttributeAPI {
    return &AttributeAPI{Handler: h}
}

type AttributeResponse struct {
    ID    uint     `json:"id"`
    Name  string   `json:"name"`
    Value []string `json:"value"`
}

type CreateAttributeRequest struct {
	Name string	`json:"name"`
	Value []uint64	`json:"value"`
}

type CreateAttributeResponse struct {
	TransactionHash string `json:"transaction_hash"`
}

type UpdateAttributeRequest struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Value []uint64 `json:"value"`
}

type UpdateAttributeResponse struct {
	TransactionHash string `json:"transaction_hash"`
}

type RemoveAttributeResponse struct {
	TransactionHash string `json:"transaction_hash"`
}

func (api *AttributeAPI) GetAttributeHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
        return
    }

    idUint, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
        return
    }

    attr, err := api.Handler.GetAttribute(uint(idUint))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get attribute: %v", err), http.StatusInternalServerError)
        return
    }

    valueStrings := make([]string, len(attr.Value))
    for i, v := range attr.Value {
        valueStrings[i] = v.String()
    }

    response := AttributeResponse{
        ID:    uint(attr.Id.Uint64()),
        Name:  attr.Name,
        Value: valueStrings,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *AttributeAPI) CreateAttributeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req CreateAttributeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	values := make([]uint, len(req.Value))
	for i, v := range req.Value {
		values[i] = uint(v)
	}

	txHash, err := api.Handler.CreateAttribute(req.Name, values)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create attribute: %v", err), http.StatusInternalServerError)
		return
	}
	response := CreateAttributeResponse{
		TransactionHash: txHash,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func (api *AttributeAPI) RemoveAttributeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
        return
    }

    idUint, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
        return
    }

	txHash, err := api.Handler.RemoveAttribute(uint(idUint))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to remove attribute: %v", err), http.StatusInternalServerError)
	}

	response := RemoveAttributeResponse{
		TransactionHash: txHash,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func (api *AttributeAPI) UpdateAttributeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var req UpdateAttributeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	values := make([]uint, len(req.Value))
	for i, v := range req.Value {
		values[i] = uint(v)
	}

	txHash, err := api.Handler.UpdateAttribtue(uint(req.ID), req.Name, values)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update attribute: %v", err), http.StatusInternalServerError)
		return
	}

	response := UpdateAttributeResponse{
		TransactionHash: txHash,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func (api *AttributeAPI) GetAttributesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	attrs, err := api.Handler.GetAttributes()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get attributes: %v", err), http.StatusInternalServerError)
		return
	}

	var response []AttributeResponse
	for _, attr := range attrs {
		valueStrings := make([]string, len(attr.Value))
		for i, v := range attr.Value {
			valueStrings[i] = v.String()
		}
		response = append(response, AttributeResponse{
			ID: uint(attr.Id.Uint64()),
			Name: attr.Name,
			Value: valueStrings,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
// Implement other handlers for attributes (Create, Update, Delete)
