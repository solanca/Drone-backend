package api

import (
	"drone-backend/internal/database"
	handler "drone-backend/internal/handlers/on-chain"
	"drone-backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type AccessAPI struct {
	PDPHandler *handler.PDPHandler
    DroneAPI    *DroneAPI
    PolicyAPI   *PolicyAPI
}

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func NewAccessAPI(pdpHandler *handler.PDPHandler, droneAPI *DroneAPI, policyAPI *PolicyAPI) *AccessAPI {
	return &AccessAPI{PDPHandler: pdpHandler, DroneAPI: droneAPI, PolicyAPI: policyAPI}
}

type AccessRequest struct {
	EntityID string `json:"entity_id"`
    RequestTarget uint `json:"request_target"`
}

type AccessResponse struct {
    Granted          bool   `json:"granted"`
    TransactionHash  string `json:"transaction_hash"`
}

func findDroneByID(droneID string) (*models.Drone, error) {
	var drone models.Drone
	result := database.DB.First(&drone, "ID = ?", droneID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &drone, nil
}

func (api *AccessAPI) Layer2AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    uint64ID, _ := strconv.ParseUint(req.EntityID, 10, 0)

	drone, err := findDroneByID(req.EntityID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }
    var policy models.Policy
    result := db.Where("zone = ?", drone.Zone).First(&policy)
	if result.Error != nil {
        http.Error(w, "Failed to retrieve policies", http.StatusBadRequest)
        return
	}
    granted, txHash, err := api.PDPHandler.EvaluateAccess(uint(uint64ID), policy.Zone, policy.StartTime.String(), policy.EndTime.String())

    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}
func (api *AccessAPI) Layer3AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    uint64ID, _ := strconv.ParseUint(req.EntityID, 10, 0)

    drone, err := findDroneByID(req.EntityID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }
    policy, err := api.PolicyAPI.Handler.GetPolicyByZone(drone.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get policy by zone: %v", err), http.StatusInternalServerError)
        return
    }

    policyResponse := PolicyResponse{
        ID: uint(policy.Id.Uint64()),
        Zone: int(policy.Id.Int64()),
        StartTime: policy.StartTime,
        EndTime: policy.EndTime,
    }

    granted, txHash, err := api.PDPHandler.EvaluateAccess(uint(uint64ID), policyResponse.Zone, policyResponse.StartTime, policyResponse.EndTime)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}
func (api *AccessAPI) AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    uint64ID, _ := strconv.ParseUint(req.EntityID, 10, 0)

    drone, err := api.DroneAPI.Handler.GetDrone(uint(uint64ID))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }

    droneResponse := DroneResponse{
        ID: uint(drone.Id.Uint64()),
        ModelType: drone.ModelType,
        Zone: int(drone.Zone.Int64()),
    }

    policy, err := api.PolicyAPI.Handler.GetPolicyByZone(droneResponse.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get policy by zone: %v", err), http.StatusInternalServerError)
        return
    }

    policyResponse := PolicyResponse{
        ID: uint(policy.Id.Uint64()),
        Zone: int(policy.Id.Int64()),
        StartTime: policy.StartTime,
        EndTime: policy.EndTime,
    }

    granted, txHash, err := api.PDPHandler.EvaluateAccess(uint(uint64ID), policyResponse.Zone, policyResponse.StartTime, policyResponse.EndTime)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}