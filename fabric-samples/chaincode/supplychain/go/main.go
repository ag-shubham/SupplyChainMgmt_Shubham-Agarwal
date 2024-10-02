package main

import (
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing the product lifecycle
type SmartContract struct {
    contractapi.Contract
}

// Product describes the details of the product in the supply chain
type Product struct {
    ProductID          string `json:"productID"`
    Name               string `json:"name"`
    Description        string `json:"description"`
    ManufacturingDate  string `json:"manufacturingDate"`
    BatchNumber        string `json:"batchNumber"`
    Status             string `json:"status"`
    SupplyDate         string `json:"supplyDate,omitempty"`
    WarehouseLocation  string `json:"warehouseLocation,omitempty"`
    WholesaleDate      string `json:"wholesaleDate,omitempty"`
    WholesaleLocation  string `json:"wholesaleLocation,omitempty"`
    Quantity           int    `json:"quantity,omitempty"`
}

// InitLedger initializes the ledger with some sample products
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    initialProducts := []Product{
        {
            ProductID:         "P001",
            Name:              "Product1",
            Description:       "Description for Product1",
            ManufacturingDate: "2023-09-25",
            BatchNumber:       "B001",
            Status:            "Created",
        },
        {
            ProductID:         "P002",
            Name:              "Product2",
            Description:       "Description for Product2",
            ManufacturingDate: "2023-09-26",
            BatchNumber:       "B002",
            Status:            "Created",
        },
    }

    for _, product := range initialProducts {
        if err := putState(ctx, product.ProductID, product); err != nil {
            return fmt.Errorf("failed to initialize ledger for product %s: %w", product.ProductID, err)
        }
    }

    return nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&SmartContract{})
    if err != nil {
        fmt.Printf("Error creating supply chain smart contract: %s\n", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting supply chain smart contract: %s\n", err.Error())
    }
}
