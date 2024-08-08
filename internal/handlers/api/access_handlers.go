package api

import (
	"drone-backend/internal/database"
	handler "drone-backend/internal/handlers/on-chain"
	"drone-backend/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"gorm.io/datatypes"
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

var mu sync.Mutex

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

func (api *AccessAPI) Layer2AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    
    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    uint64ID, err := strconv.ParseUint(req.EntityID, 10, 0)
    if err != nil {
        http.Error(w, "Invalid entity ID", http.StatusBadRequest)
        return
    }
    
    requestSentence := fmt.Sprintf("Small Drone %s sent the access request", req.EntityID)
	accessJSON, err := json.Marshal(models.Access{
        Request: requestSentence,
		Status:  "Received",
	})
	if err != nil {
        log.Printf("Error marshaling Access: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
    
    history := models.History{
        Drone:  req.EntityID,
		Access: datatypes.JSON(accessJSON),
		Status: "In Progress",
	}

	if err := database.DB.Create(&history).Error; err != nil {
        log.Printf("Error creating history record: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

    // PIP
    drone, err := findDroneByID(req.EntityID)
    if err != nil {
		log.Printf("Error finding drone: %v", err)
		http.Error(w, "Drone does not exist.", http.StatusBadRequest)
	}

	pipContent := fmt.Sprintf("Drone %d attributes gathered: Model Type: %s, Zone: %d",
		drone.ID, drone.ModelType, drone.Zone)

	pipJSON, err := json.Marshal(models.PIP{
		Content: pipContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PIP: %v", err)
	}

	history.PIP = datatypes.JSON(pipJSON)

	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PIP: %v", err)
	}

    // PRP
    var policy models.Policy
    result := database.DB.First(&policy, "zone = ?", drone.Zone)
    if result.Error != nil {
        http.Error(w, "Failed to retrieve policies", http.StatusInternalServerError)
        return
    }

    prpContent := fmt.Sprintf("Policies retrieved for zone %d policies found.", policy.Zone)

    prpJSON, err := json.Marshal(models.PRP{
		Content: prpContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PRP: %v", err)
	}

	history.PRP = datatypes.JSON(prpJSON)
	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PRP: %v", err)
	}

    // PDP
    granted, txHash, err := api.PDPHandler.EvaluateAccess(uint(uint64ID), policy.Zone, policy.StartTime.String(), policy.EndTime.String())
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    pdpDecision := "Denied"
	if granted {
		pdpDecision = "Granted"
	}

	pdpJSON, err := json.Marshal(models.PDP{
		Decision: pdpDecision,
		Status:   "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PDP: %v", err)
	}

	history.PDP = datatypes.JSON(pdpJSON)
	history.Status = pdpDecision

	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PDP: %v", err)
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
    mu.Unlock()
}

func (api *AccessAPI) Layer3AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    requestSentence := fmt.Sprintf("Small Drone %s sent the access request", req.EntityID)
	accessJSON, err := json.Marshal(models.Access{
        Request: requestSentence,
		Status:  "Received",
	})
	if err != nil {
        log.Printf("Error marshaling Access: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
    
    history := models.History{
        Drone:  req.EntityID,
		Access: datatypes.JSON(accessJSON),
		Status: "In Progress",
	}

	if err := database.DB.Create(&history).Error; err != nil {
        log.Printf("Error creating history record: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

    uint64ID, _ := strconv.ParseUint(req.EntityID, 10, 0)

    // PIP

    drone, err := findDroneByID(req.EntityID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }

	pipContent := fmt.Sprintf("Drone %d attributes gathered: Model Type: %s, Zone: %d",
		drone.ID, drone.ModelType, drone.Zone)

	pipJSON, err := json.Marshal(models.PIP{
		Content: pipContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PIP: %v", err)
	}

	history.PIP = datatypes.JSON(pipJSON)

	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PIP: %v", err)
	}

    // PRP
    policy, err := api.PolicyAPI.Handler.GetPolicyByZone(drone.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get policy by zone: %v", err), http.StatusInternalServerError)
        return
    }

    prpContent := fmt.Sprintf("Policies retrieved for zone %d policies found.", policy.Zone)

    prpJSON, err := json.Marshal(models.PRP{
		Content: prpContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PRP: %v", err)
	}

	history.PRP = datatypes.JSON(prpJSON)
	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PRP: %v", err)
	}

    policyResponse := PolicyResponse{
        ID: uint(policy.Id.Uint64()),
        Zone: int(policy.Id.Int64()),
        StartTime: policy.StartTime,
        EndTime: policy.EndTime,
    }

    // PDP
    granted, txHash, err := api.PDPHandler.EvaluateAccess(uint(uint64ID), policyResponse.Zone, policyResponse.StartTime, policyResponse.EndTime)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    pdpDecision := "Denied"
	if granted {
		pdpDecision = "Granted"
	}

	pdpJSON, err := json.Marshal(models.PDP{
		Decision: pdpDecision,
		Status:   "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PDP: %v", err)
	}

	history.PDP = datatypes.JSON(pdpJSON)
	history.Status = pdpDecision

	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PDP: %v", err)
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
    mu.Unlock()
}
func (api *AccessAPI) AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
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
    mu.Unlock()
}

func findDroneByID(droneID string) (*models.Drone, error) {
	var drone models.Drone
	result := database.DB.First(&drone, "ID = ?", droneID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &drone, nil
}