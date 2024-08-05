package api

import (
	handler "drone-backend/internal/handlers/on-chain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DroneAPI struct {
	Handler *handler.DroneHandler
}

func NewDroneAPI(h *handler.DroneHandler) *DroneAPI {
	return &DroneAPI{Handler: h}
}

type DroneResponse struct {
	ID uint 	`json:"id"`
	Model string	`json:"model"`
	Zone int `json:"zone"`
}

type CreateDroneRequest struct {
	Model string `json:"model"`
	Zone int `json:"zone"`
}

type CreateDroneResponse struct {
    TransactionHash string `json:"transaction_hash"`
}

// UpdateDroneRequest represents the JSON structure of the update drone request
type UpdateDroneRequest struct {
    ID    uint   `json:"id"`
    Model string `json:"model"`
    Zone  int    `json:"zone"`
}

// UpdateDroneResponse represents the JSON structure of the update drone response
type UpdateDroneResponse struct {
    TransactionHash string `json:"transaction_hash"`
}

// RemoveDroneRequest represents the JSON structure of the remove drone request
type RemoveDroneRequest struct {
    ID uint `json:"id"`
}

// RemoveDroneResponse represents the JSON structure of the remove drone response
type RemoveDroneResponse struct {
    TransactionHash string `json:"transaction_hash"`
}

// GetDroneResponse represents the JSON structure of the get drone response
type GetDroneResponse struct {
    ID    uint   `json:"id"`
    Model string `json:"model"`
    Zone  int    `json:"zone"`
}

func (api *DroneAPI) GetDronesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    drones, err := api.Handler.GetDrones()
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drones: %v", err), http.StatusInternalServerError)
        return
    }

    var response []DroneResponse
    for _, drone := range drones {
        response = append(response, DroneResponse{
            ID:    uint(drone.Id.Uint64()),
            Model: drone.Model,
            Zone:  int(drone.Zone.Int64()),
        })
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}

// CreateDroneHandler handles the /createDrone endpoint
func (api *DroneAPI) CreateDroneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req CreateDroneRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    txHash, err := api.Handler.CreateDrone(req.Model, req.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := CreateDroneResponse{
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *DroneAPI) UpdateDroneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req UpdateDroneRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    txHash, err := api.Handler.UpdateDrone(req.ID, req.Model, req.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to update drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := UpdateDroneResponse{
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

// RemoveDroneHandler handles the /removeDrone endpoint
func (api *DroneAPI) RemoveDroneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req RemoveDroneRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    txHash, err := api.Handler.RemoveDrone(req.ID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to remove drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := RemoveDroneResponse{
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

// GetDroneHandler handles the /getDrone endpoint
func (api *DroneAPI) GetDroneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Extract the 'id' query parameter from the URL
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
        return
    }

    // Convert 'id' to uint
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
        return
    }

    drone, err := api.Handler.GetDrone(uint(id))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := GetDroneResponse{
        ID:    uint(drone.Id.Uint64()),
        Model: drone.Model,
        Zone:  int(drone.Zone.Int64()),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}