package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"drone-backend/internal/database"
	"drone-backend/internal/handlers/api"
	handlers "drone-backend/internal/handlers/off-chain"
	handler "drone-backend/internal/handlers/on-chain"
	"drone-backend/smart-contracts/Attribute"
	"drone-backend/smart-contracts/Drone"
	"drone-backend/smart-contracts/PDP"
	"drone-backend/smart-contracts/Policy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	err := database.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	handlers.InitDB(database.DB)
	
	r := mux.NewRouter()
	// r.HandleFunc("/ping", handlers.PingHandler).Methods(http.MethodGet)

	r.HandleFunc("/drones", handlers.CreateDroneHandler).Methods(http.MethodPost)
	r.HandleFunc("/drones", handlers.GetDronesHandler).Methods(http.MethodGet)
	r.HandleFunc("/drones/{id:[0-9]+}", handlers.GetDroneHandler).Methods(http.MethodGet)
	r.HandleFunc("/drones/{id:[0-9]+}", handlers.UpdateDroneHandler).Methods(http.MethodPut)
	r.HandleFunc("/drones/{id:[0-9]+}", handlers.DeleteDroneHandler).Methods(http.MethodDelete)
	r.HandleFunc("/drones/zone/{zone}", handlers.GetDroneByZoneHandler).Methods(http.MethodGet)

	r.HandleFunc("/attributes", handlers.CreateAttributeHandler).Methods(http.MethodPost)
	r.HandleFunc("/attributes", handlers.GetAttributesHandler).Methods(http.MethodGet)
	r.HandleFunc("/attributes/{id:[0-9]+}", handlers.GetAttributeHandler).Methods(http.MethodGet)
	r.HandleFunc("/attributes/{id:[0-9]+}", handlers.UpdateAttributeHandler).Methods(http.MethodPut)
	r.HandleFunc("/attributes/{id:[0-9]+}", handlers.DeleteAttributeHandler).Methods(http.MethodDelete)
	r.HandleFunc("/attributes/{name}", handlers.GetAttributeByNameHandler).Methods(http.MethodGet)

	r.HandleFunc("/policies", handlers.CreatePolicyHandler).Methods(http.MethodPost)
	r.HandleFunc("/policies", handlers.GetPoliciesHandler).Methods(http.MethodGet)
	r.HandleFunc("/policies/{id:[0-9]+}", handlers.GetPolicyHandler).Methods(http.MethodGet)
	r.HandleFunc("/policies/{id:[0-9]+}", handlers.UpdatePolicyHandler).Methods(http.MethodPut)
	r.HandleFunc("/policies/{id:[0-9]+}", handlers.DeletePolicyHandler).Methods(http.MethodDelete)

	r.HandleFunc("/access-request", handlers.AccessRequestHandler).Methods(http.MethodPost)

	
	clientURL := "https://json-rpc.evm.testnet.iotaledger.net"
	keystoreFile := "./wallet.json"
	password := "Wahaha123!@#";
	attributeContractAddress := "0x98F1748e5396E7da251fA730521Fcd0e0622b1F5"
	policyContractAddress := "0xf8A9eD693ea899ed7a188Fc4dC5349868037e83b"
	droneContractAddress := "0xc1A2FF8FDF2085C249DFB1f1E50beC4A6Db0ee1e"
	pdpContractAddress := "0xdB920e56E7fEca196cAed694963cfdc5A9305320"

	client, err := ethclient.Dial(clientURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Load the keystore file
    keyJSON, err := ioutil.ReadFile(keystoreFile)
    if err != nil {
        log.Fatalf("Failed to read the keyfile: %v", err)
    }

    // Decrypt the keystore file
    key, err := keystore.DecryptKey(keyJSON, password)
    if err != nil {
        log.Fatalf("Failed to decrypt the key: %v", err)
    }

    // Create an authenticated transactor
    auth := bind.NewKeyedTransactor(key.PrivateKey)

    // Instantiate the smart contract instances
    attributeInstance, err := Attribute.NewAttribute(common.HexToAddress(attributeContractAddress), client)
    if err != nil {
        log.Fatalf("Failed to instantiate AttributeContract: %v", err)
    }

    policyInstance, err := Policy.NewPolicy(common.HexToAddress(policyContractAddress), client)
    if err != nil {
        log.Fatalf("Failed to instantiate PolicyContract: %v", err)
    }

    droneInstance, err := Drone.NewDrone(common.HexToAddress(droneContractAddress), client)
    if err != nil {
        log.Fatalf("Failed to instantiate DroneContract: %v", err)
    }

	pdpInstance, err := PDP.NewPDP(common.HexToAddress(pdpContractAddress), client)
    if err != nil {
        log.Fatalf("Failed to instantiate PDPContract: %v", err)
    }

    attributeHandler := handler.NewAttributeHandler(attributeInstance, auth)
    policyHandler := handler.NewPolicyHandler(policyInstance, auth)
    droneHandler := handler.NewDroneHandler(droneInstance, auth)
	pdpHandler := handler.NewPDPHandler(pdpInstance, auth, client)

	attributeAPI := api.NewAttributeAPI(attributeHandler)
    policyAPI := api.NewPolicyAPI(policyHandler)
    droneAPI := api.NewDroneAPI(droneHandler)
	pdpAPI := api.NewAccessAPI(pdpHandler)

	r.HandleFunc("/getAttribute", attributeAPI.GetAttributeHandler)
	r.HandleFunc("/createAttribute", attributeAPI.CreateAttributeHandler)
	r.HandleFunc("/removeAttribute", attributeAPI.RemoveAttributeHandler)
	r.HandleFunc("/updateAttribute", attributeAPI.UpdateAttributeHandler)
	r.HandleFunc("/getAttributes", attributeAPI.GetAttributesHandler)
	r.HandleFunc("/getPolicies", policyAPI.GetPoliciesHandler)
	r.HandleFunc("/createPolicy", policyAPI.CreatePolicyHandler)
	r.HandleFunc("/updatePolicy", policyAPI.UpdatePolicyHandler)
	r.HandleFunc("/removePolicy", policyAPI.RemovePolicyHandler)
	r.HandleFunc("/getPolicyByZone", policyAPI.GetPolicyByZoneHandler)
	r.HandleFunc("/getDrones", droneAPI.GetDronesHandler)
	r.HandleFunc("/getDrone", droneAPI.GetDroneHandler)
	r.HandleFunc("/createDrone", droneAPI.CreateDroneHandler)
	r.HandleFunc("/updateDrone", droneAPI.UpdateDroneHandler)
	r.HandleFunc("/removeDrone", droneAPI.RemoveDroneHandler)
	r.HandleFunc("/accessRequest", pdpAPI.AccessRequestHandler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	})

	// srv := server.NewServer()
	handler := c.Handler((r));
	log.Fatal(http.ListenAndServe(":8080", handler))
}