# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Running the Application
```bash
go run main.go                    # Run the interactive CLI with default database (pills.db)
PILLS_DB_PATH=custom.db go run main.go  # Run with custom database path
go run main.go -reset            # Reset database and seed with fresh data
```

### Build and Test
```bash
go build ./...                   # Build the project
go test ./...                    # Run all tests with in-memory SQLite
go test ./model -v               # Run specific package tests
go test -race ./...              # Run tests with race detection
go fmt ./...                     # Format code
go vet ./...                     # Run static analysis
```

## Architecture Overview

The project follows a clean architecture pattern with clear separation of concerns:

### Core Components
- **main.go**: Application entry point, handles database initialization, migration, and launches the main view loop
- **control/**: Command routing and application state management (singletons via `control.SetDB()`)
- **model/**: Data models, database migrations, and population logic using GORM with SQLite
- **view/**: Terminal UI implementation using pterm library for interactive forms and tables
- **validation/**: Input validation with business rules for ATC/AIC codes, dosages, and units
- **utils/**: Cross-cutting utilities, primarily time/date helpers for prescription calculations

### Key Design Patterns
- **Singleton DB Access**: Database connection managed globally through `control.SetDB()` and `control.GetDB()`
- **State Machine**: Application flow controlled via screen states in `control/appState`
- **Command Pattern**: Single-letter commands ('s', 'p', 'r', 'a', 'f', 'q') drive navigation
- **Decimal Arithmetic**: Uses `shopspring/decimal` for precise dosage and stock calculations

### Data Model Highlights
- **ActiveIngredient**: Central entity tracking stock with ATC codes, units (mg/ml/UI), and automatic intake calculations
- **Medicine**: Commercial products linked to active ingredients via ATC codes with AIC identification
- **Prescription**: Time-bound dosing regimens with frequency-based stock depletion
- **StockLog**: Audit trail for all restocking events with medicine and ATC references

## Testing Guidelines

Tests use in-memory SQLite fixtures and should never touch the local `pills.db`. Key patterns:
- Unit tests live alongside source files (`*_test.go`)
- Use `testify/assert` for assertions
- Reset global state in `t.Cleanup()` when modifying data
- Name subtests after business scenarios for clear failure mapping

## Database Schema Notes

- SQLite database file location controlled by `PILLS_DB_PATH` environment variable (defaults to `pills.db`)
- Schema auto-migrates on startup; use `-reset` flag to reseed fresh data
- All decimal fields use `numeric` type for precision (dosage, stock calculations)
- ATC codes are 7-character identifiers with format validation
- AIC codes are 9-digit numeric strings for medicine identification

## Development Workflow

1. Make changes following Go conventions and existing patterns
2. Run `go fmt ./...` and `go vet ./...` locally
3. Execute `go test ./...` to verify functionality
4. Test CLI manually for UI changes: `go run main.go`
5. Database changes require explicit migration handling in `model/`