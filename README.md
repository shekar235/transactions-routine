# Transactions Routine

## Description

This Go project provides an HTTP API for managing accounts and transactions, with a simple in-memory storage.

## Endpoints

- **POST /accounts** - Creates a new account
- **GET /accounts/{accountId}** - Retrieves account details
- **POST /transactions** - Creates a new transaction for an account

## Running the Project

### Prerequisites
- Go (for running locally)
- Docker (for running in a container)
- `make` (for Makefile commands)

### 1. **Run Locally**:
      make run
   - **(OR)**
### 2. **Run using docker Locally**:
      make docker-build
      make docker-run
