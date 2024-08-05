package handler

import (
	"drone-backend/smart-contracts/Drone" // Replace with your actual module name
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type DroneHandler struct {
    Instance *Drone.Drone
    Auth     *bind.TransactOpts
}

func NewDroneHandler(instance *Drone.Drone, auth *bind.TransactOpts) *DroneHandler {
    return &DroneHandler{
        Instance: instance,
        Auth:     auth,
    }
}

func (h *DroneHandler) GetDrones() ([]Drone.DroneContractDrone, error) {
    return h.Instance.GetDrones(&bind.CallOpts{})
}

func (h *DroneHandler) CreateDrone(model string, zone int) (string, error) {
    tx, err := h.Instance.CreateDrone(h.Auth, model, big.NewInt(int64(zone)))
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *DroneHandler) UpdateDrone(id uint, model string, zone int) (string, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.UpdateDrone(h.Auth, bigID, model, big.NewInt(int64(zone)))
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *DroneHandler) RemoveDrone(id uint) (string, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.DeleteDrone(h.Auth, bigID)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *DroneHandler) GetDrone(id uint) (Drone.DroneContractDrone, error) {
    bigID := big.NewInt(int64(id))
    return h.Instance.GetDrone(&bind.CallOpts{}, bigID)
}