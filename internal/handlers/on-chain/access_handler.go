package handler

import (
	"drone-backend/smart-contracts/PDP"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PDPHandler struct {
    Instance *PDP.PDP
    Auth     *bind.TransactOpts
    Client   *ethclient.Client
}

func NewPDPHandler(instance *PDP.PDP, auth *bind.TransactOpts, client *ethclient.Client) *PDPHandler {
    return &PDPHandler{
        Instance: instance,
        Auth:     auth,
        Client:   client,
    }
}

func (h *PDPHandler) EvaluateAccess(entityID uint) (bool, string, error) {
    bigID := big.NewInt(int64(entityID))
    tx, err := h.Instance.EvaluateAccess(h.Auth, bigID)
    if err != nil {
        return false, "", err
    }

    // Wait for the transaction to be mined
    receipt, err := bind.WaitMined(nil, h.Client, tx)
    if err != nil {
        return false, "", err
    }

    // Parse the logs to find out if access was granted
    var granted bool
    for _, log := range receipt.Logs {
        if len(log.Data) > 0 && log.Data[0] == 1 {
            granted = true
            break
        }
    }

    return granted, tx.Hash().Hex(), nil
}