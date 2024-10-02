package main

import (
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateProduct adds a new product to the ledger
func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterface, productID, name, description, manufacturingDate, batchNumber string) error {
    existingProduct, err := s.QueryProduct(ctx, productID)
    if err == nil && existingProduct != nil {
        return fmt.Errorf("product with ID %s already exists", productID)
    }

    product := Product{
        ProductID:         productID,
        Name:              name,
        Description:       description,
        ManufacturingDate: manufacturingDate,
        BatchNumber:       batchNumber,
        Status:            "Created",
    }

    return putState(ctx, productID, product)
}

// SupplyProduct updates the product status with supply details
func (s *SmartContract) SupplyProduct(ctx contractapi.TransactionContextInterface, productID, supplyDate, warehouseLocation string) error {
    product, err := s.QueryProduct(ctx, productID)
    if err != nil {
        return err
    }

    product.SupplyDate = supplyDate
    product.WarehouseLocation = warehouseLocation
    product.Status = "Supplied"

    return putState(ctx, productID, *product)
}

// WholesaleProduct updates the product with wholesale details
func (s *SmartContract) WholesaleProduct(ctx contractapi.TransactionContextInterface, productID, wholesaleDate, wholesaleLocation string, quantity int) error {
    product, err := s.QueryProduct(ctx, productID)
    if err != nil {
        return err
    }

    product.WholesaleDate = wholesaleDate
    product.WholesaleLocation = wholesaleLocation
    product.Quantity = quantity
    product.Status = "Wholesaled"

    return putState(ctx, productID, *product)
}

// QueryProduct retrieves a product from the ledger by productID
func (s *SmartContract) QueryProduct(ctx contractapi.TransactionContextInterface, productID string) (*Product, error) {
    product, err := getState[Product](ctx, productID)
    if err != nil {
        return nil, err
    }
    return product, nil
}

// UpdateProductStatus updates the status of a product (e.g., sold)
func (s *SmartContract) UpdateProductStatus(ctx contractapi.TransactionContextInterface, productID, status string) error {
    product, err := s.QueryProduct(ctx, productID)
    if err != nil {
        return err
    }

    product.Status = status

    return putState(ctx, productID, *product)
}
