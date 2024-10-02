# Supply Chain Management Chaincode

This repository contains the Hyperledger Fabric chaincode for managing the lifecycle of products in a supply chain system. The chaincode provides a set of functions to create, update, and query product information as it progresses through various stages in the supply chain, such as manufacturing, supply, and wholesale.

## Table of Contents

- [Overview](#overview)
- [Chaincode Functions](#chaincode-functions)
  - [Product Structure](#product-structure)
  - [Function Descriptions](#function-descriptions)
- [Helper Functions](#helper-functions)
  - [`putState`](#putstate)
  - [`getState`](#getstate)
  - [Benifits of helper functions](#benifits-of-helper-functions)

## Overview

The **Supply Chain Management** chaincode is designed to track the movement and status of products across different stakeholders, such as manufacturers, warehouses, wholesalers, and retailers. Each product has a unique ID, and its status is updated as it moves through the supply chain.

The chaincode supports the following functionalities:

- **Create a Product**: Add a new product to the ledger.
- **Update Product Status**: Modify the status of a product (e.g., from manufacturing to supply).
- **Query Product**: Retrieve product details based on its unique ID.
- **Supply and Wholesale Operations**: Record the supply and wholesale stages of the product lifecycle.

## Chaincode Functions

### Product Structure

Each product is represented by the following structure in the ledger:

```json
{
    "productID": "string",
    "name": "string",
    "description": "string",
    "manufacturingDate": "string",
    "batchNumber": "string",
    "status": "string",
    "supplyDate": "string (optional)",
    "warehouseLocation": "string (optional)",
    "wholesaleDate": "string (optional)",
    "wholesaleLocation": "string (optional)",
    "quantity": "int (optional)"
}
```

### Function Descriptions

1. InitLedger: Initializes the ledger with some sample products.
2. CreateProduct: Adds a new product to the ledger.
3. SupplyProduct: Updates the product status with the supply date and warehouse location.
4. WholesaleProduct: Updates the product with wholesale date, location, and quantity.
5. QueryProduct: Retrieves product details from the ledger based on the product ID.
6. UpdateProductStatus: Updates the product status to a new value, such as 'Sold' or 'Returned'

## Helper Functions

The helper functions are designed to streamline common operations for interacting with the ledger, such as saving and retrieving data in a structured way. These functions help reduce redundancy in the chaincode and improve maintainability.

### `putState`

The `putState` function is a utility to convert a data structure to JSON format and store it in the world state. This ensures that any data passed to the ledger is serialized correctly.

### `getState`

The `getState` function retrieves the data from the world state based on the key provided. It deserializes the JSON data and returns the original structure to the caller.

### Benifits of helper functions

1. Reusability: These functions can be used throughout the chaincode to avoid redundant code for serializing and deserializing data.
2. Maintainability: If changes are needed in how data is stored or retrieved, the modifications can be made in one place without needing to update multiple parts of the code.
3. Error Handling: These functions centralize error handling, making the chaincode easier to debug and more robust.

By using these helper functions, the chaincode becomes cleaner, more organized, and easier to extend or modify in the future.

