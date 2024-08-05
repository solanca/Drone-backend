package handler

import (
	"drone-backend/smart-contracts/Policy"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type PolicyHandler struct {
	Instance *Policy.Policy
	Auth     *bind.TransactOpts
}

func NewPolicyHandler(instance *Policy.Policy, auth *bind.TransactOpts) *PolicyHandler {
	return &PolicyHandler{
		Instance: instance,
		Auth: auth,
	}
}

func (h *PolicyHandler) GetPolicies() ([]Policy.PolicyContractPolicy, error) {
	policies, err := h.Instance.GetPolicies(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return policies, nil
}

func (h *PolicyHandler) CreatePolicy(zone int, startTime, endTime string) (string, error) {
    tx, err := h.Instance.CreatePolicy(h.Auth, big.NewInt(int64(zone)), startTime, endTime)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *PolicyHandler) UpdatePolicy(id uint, zone int, startTime, endTime string) (string, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.UpdatePolicy(h.Auth, bigID, big.NewInt(int64(zone)), startTime, endTime)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *PolicyHandler) RemovePolicy(id uint) (string, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.DeletePolicy(h.Auth, bigID)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *PolicyHandler) GetPolicyByZone(zone int) (Policy.PolicyContractPolicy, error) {
    return h.Instance.GetPolicyByZone(&bind.CallOpts{}, big.NewInt(int64(zone)))
}