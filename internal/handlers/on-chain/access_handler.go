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

    eventSignature := []byte("AccessDecision(uint256,bool)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)
    accessGranted := false

    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var event struct {
            EntityId *big.Int
            Granted  bool
        }
        err := h.parseAccessDecision(vLog, &event)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return false, "", fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        // Log event details
        log.Printf("Event parsed successfully: EntityId=%s, Granted=%t", event.EntityId.String(), event.Granted)

        // Check if the event matches the given entityId
        if event.EntityId.Cmp(bigID) == 0 {
            log.Printf("Event entityId matches: %s", event.EntityId.String())
            accessGranted = event.Granted
            break
        } else {
            log.Printf("Event entityId does not match: %s", event.EntityId.String())
        }
    }

    log.Printf("EvaluateAccess completed: Granted=%t", accessGranted)
    return accessGranted, tx.Hash().Hex(), nil
}

func (h *PDPHandler) parseAccessDecision(logEntry *types.Log, event *struct {
    EntityId *big.Int
    Granted  bool
}) error {
    abiDefinition := `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"entityId","type":"uint256"},{"indexed":false,"internalType":"bool","name":"granted","type":"bool"}],"name":"AccessDecision","type":"event"}]`
    parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
    if err != nil {
        log.Printf("Failed to parse ABI definition: %v", err)
        return err
    }

    // Ensure that the logEntry contains enough topics
    if len(logEntry.Topics) < 2 {
        log.Printf("Log entry has insufficient topics: %v", logEntry.Topics)
        return fmt.Errorf("insufficient topics in log entry")
    }

    err = parsedABI.UnpackIntoInterface(event, "AccessDecision", logEntry.Data)
    if err != nil {
        log.Printf("Failed to unpack event data: %v", err)
        return err
    }

    // Indexed fields are not part of the log.Data, so extract them from Topics
    event.EntityId = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
    log.Printf("Extracted indexed EntityId: %s", event.EntityId.String())

    return nil
}

