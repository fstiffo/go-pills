# go-pills

## Project Goals
- Track prescriptions and medicine stock from the command line.
- Provide a quick summary of inventory and dosing needs.

## Core Features
- Interactive terminal UI for managing medicines and prescriptions.
- Summary screen showing remaining stock and next doses.
- Forms to update pharmacy stock and prescriptions.
- Add new medicine boxes and refresh data at any time.

## Build and Run
```bash
go run main.go
```
This starts the application using an embedded SQLite database (`pills.db`).

## Text Commands
Commands are entered by typing a letter and pressing Enter rather than using function keys.

- `(S)ummary` – show remaining stock and next doses.
- `(P)harmacy` – update pharmacy stock.
- `P(R)escriptions` – update prescriptions.
- `(A)dd` – add medicine boxes.
- `Re(F)resh` – refresh summary.
- `(Q)uit` – exit the application.

## Data Model Overview
- **ActiveIngredient** – identified by an ATC code, tracks stock and units.
- **Medicine** – references an active ingredient, dosage, and packaging.
- **Prescription** – dosage and frequency for an active ingredient.
- **StockLog** – records restocking events.
- **PrescriptionLog** – tracks changes to prescriptions.

## Roadmap
- Export reports and summaries.
- Support for multiple users and profiles.
- Enhanced validation and logging.

## Contributing
1. Fork the repository and create a feature branch.
2. Run `go fmt` and `go test ./...` before submitting.
3. Open a pull request describing your changes.
4. Review the detailed contributor guide in [`AGENTS.md`](AGENTS.md) before filing your PR.
