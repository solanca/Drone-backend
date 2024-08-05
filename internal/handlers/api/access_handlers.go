package api

import (
	handler "drone-backend/internal/handlers/on-chain"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessAPI struct {
	Handler *handler.PDPHandler
}

func NewAccessAPI(h *handler.PDPHandler) *AccessAPI {
	return &AccessAPI{Handler: h}
}

type AccessRequest struct {
	EntityID uint `json:"entity_id"`
}

type AccessResponse struct {
    Granted          bool   `json:"granted"`
    TransactionHash  string `json:"transaction_hash"`
}

// AccessRequestHandler handles the /accessRequest endpoint
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

    granted, txHash, err := api.Handler.EvaluateAccess(req.EntityID)
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