package handler

import (
	"context"
	"drone-backend/smart-contracts/PDP"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

func (h *PDPHandler) EvaluateAccess(entityID uint, zone int, startTime string, endTime string) (bool, string, error) {
    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return false, "", fmt.Errorf("Ethereum client is not initialized")
    }

    if h.Instance == nil {
        log.Printf("PDP instance is nil")
        return false, "", fmt.Errorf("PDP instance is not initialized")
    }
    
    bigID := big.NewInt(int64(entityID))
    log.Printf("Starting EvaluateAccess for entityID: %d", entityID)
    
    bigZone := big.NewInt(int64(zone))

    tx, err := h.Instance.EvaluateAccess(h.Auth, bigID, bigZone, startTime, endTime)
    if err != nil {
        log.Printf("Failed to call EvaluateAccess: %v", err)
        return false, "", err
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return false, "", err
    }

    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("ActionLogged(address,string,string,string)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)
    accessGranted := false

    // var policy Policy.PolicyContractPolicy
    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var event struct {
            User    *big.Int
            Action string
            Timestamp  string
            Data string
        }
        err := h.parseAccessDecision(vLog, &event)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return false, "", fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        // Log event details
        // log.Printf("Event parsed successfully: EntityId=%s, Granted=%t", event.EntityId.String(), event.Granted)

        // bigInt := new(big.Int)
        parts := strings.Split(event.Data, ",")
        log.Printf(event.Data)
        // id, _ := bigInt.SetString(strings.TrimSpace(strings.Split(parts[0], ":")[1]), 10)
        // zone, _ := bigInt.SetString(strings.TrimSpace(strings.Split(parts[1], ":")[1]), 10)
        // curretnTime := strings.TrimSpace(strings.SplitN(parts[2], ":", 2)[1])
        granted := strings.TrimSpace(strings.Split(parts[3], ":")[1]);

        if granted == "Progressed" {
            accessGranted = true
        } else {
            accessGranted = false
        }
        // Check if the event matches the given entityId

    }

    log.Printf("EvaluateAccess completed: Granted=%t", accessGranted)
    return accessGranted, tx.Hash().Hex(), nil
}

func (h *PDPHandler) parseAccessDecision(logEntry *types.Log, event *struct {
    User    *big.Int
    Action string
    Timestamp  string
    Data string
}) error {
    abiDefinition := `[{
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "user",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "action",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "timestamp",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "data",
        "type": "string"
      }
    ],
    "name": "ActionLogged",
    "type": "event"
  }]`
    parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
    if err != nil {
        return err
    }

    err = parsedABI.UnpackIntoInterface(event, "ActionLogged", logEntry.Data)
    if err != nil {
        return err
    }

    if len(logEntry.Topics) > 1 {
        event.User = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
    } else {
        return fmt.Errorf("insufficient topics in log entry")
    }

    return nil
}
