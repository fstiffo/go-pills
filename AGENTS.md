# Repository Guidelines

## Project Structure & Module Organization
The module `github.com/fstiffo/go-pills` targets Go 1.23 with the Go 1.24 toolchain. `main.go` boots the CLI, while command routing lives in `control/`, data models and migrations in `model/`, terminal views in `view/`, and cross-cutting helpers in `utils/`. Validation logic stays isolated in `validation/`. SQLite artifacts such as `pills.db` and `test.db` are local state only; omit regenerated binaries and dumps from commits.

## Build, Test, and Development Commands
Run `go run ./...` to exercise the interactive CLI against the bundled SQLite store. `go build ./...` produces a binary for packaging or manual distribution. Validate changes with `go test ./...`, which spins up in-memory SQLite fixtures, and keep formatting consistent with `go fmt ./...`. Add `go vet ./...` (or `go test -race ./...` when touching concurrency) before opening a PR that modifies database or control flows.

## Coding Style & Naming Conventions
Rely on `go fmt` for indentation (tabs) and import order. Exported identifiers use CamelCase with doc comments; scoped helpers stay lowercase and favor concise naming (`updateStockedUnits`, `newOverviewTable`). Keep side effects in the control layer, pass context-rich structs rather than primitive lists, and inject clocks via `utils.TimeProvider` when work depends on `time.Now()`.

## Testing Guidelines
Unit tests sit beside their code (`model/*_test.go`, `utils/time_test.go`) and use the standard `testing` package with `testify/assert`. Use the provided in-memory SQLite setup helpers instead of touching `pills.db`; reset state within `t.Cleanup` when tests alter globals. Name subtests after business scenarios (`"SinglePrescriptionNoLogs"`) so failures map cleanly to CLI behaviour. Expand coverage whenever adding dosing, stock, or validation rules.

## Commit & Pull Request Guidelines
Commits follow short, sentence-case imperatives (`Add today's date to stock overview table header`). Keep bodies concise: what changed, why, and how it’s verified. Pull requests should include a user-facing summary, verification notes (`go test ./...`, manual CLI screenshots if UI changes), and links to issues or specs. Call out schema updates or data migrations explicitly so reviewers can refresh local databases.

## Local Data & Security Notes
SQLite files may contain personal dosing records in real deployments. Keep them out of version control and redact terminal output or screenshots before sharing publicly. Store environment-specific secrets and configuration outside the repository.
