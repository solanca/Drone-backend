package handler

import (
	"drone-backend/smart-contracts/Attribute"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type AttributeHandler struct {
    Instance *Attribute.Attribute
    Auth     *bind.TransactOpts
}

func NewAttributeHandler(instance *Attribute.Attribute, auth *bind.TransactOpts) *AttributeHandler {
    return &AttributeHandler{
        Instance: instance,
        Auth:     auth,
    }
}

func (h *AttributeHandler) GetAttribute(id uint) (Attribute.AttributeContractAttribute, error) {
    bigID := big.NewInt(int64(id))
    return h.Instance.GetAttribute(&bind.CallOpts{}, bigID)
}

func (h *AttributeHandler) CreateAttribute(name string, values []uint) (string, error) {
	bigIntValues := make([]*big.Int, len(values))
	for i, val := range values {
		bigIntValues[i] = big.NewInt(int64(val))
	}

	tx, err := h.Instance.CreateAttribute(h.Auth, name, bigIntValues)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (h *AttributeHandler) RemoveAttribute(id uint) (string, error) {
	bigID := big.NewInt(int64(id))
	tx, err := h.Instance.DeleteAttribute(h.Auth, bigID)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (h *AttributeHandler) UpdateAttribtue(id uint, name string, values []uint) (string, error) {
	bigID := big.NewInt(int64(id))
	bigIntValues := make([]*big.Int, len(values))
	for i, val := range values {
		bigIntValues[i] = big.NewInt(int64(val))
	}

	tx, err := h.Instance.UpdateAttribute(h.Auth, bigID, name, bigIntValues)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (h *AttributeHandler) GetAttributes() ([]Attribute.AttributeContractAttribute, error) {
	var attrs []Attribute.AttributeContractAttribute

	attrs, err := h.Instance.GetAttributes(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return attrs, nil
}
// Implement other methods for attribute actions (Create, Update, Delete)
