package models

type AccessRequest struct {
	DroneId string `json:"drone_id"`
}

type AccessResponse struct {
	Granted bool   `json:"granted"`
	Message string `json:"message"`
}