# AGENTS.md

## What this is

Go web app (MVP) following Alex Edwards' "Let's Go" pattern. PostgreSQL backend, Chi router, server-side HTML templates via `html/template`, cookie sessions via `scs/v2`. No frontend framework — vanilla CSS/JS in `ui/static/`.

## Quick commands

```bash
make build    # go build -o ./tmp/sabify ./cmd/web
make run      # go run ./cmd/web (port 4000)
make test     # go test -v ./...
make lint     # go vet ./...
make db_up    # docker compose up -d (Postgres 16)
make db_down  # docker compose down
make migrate  # psql ... -f migrations/001_initial_schema.sql
```

Database requires `make db_up` before running the app. Schema is in `migrations/001_initial_schema.sql`.

## Architecture

**Entry point:** `cmd/web/main.go` — wires config, DB pool, session manager, template cache into a central `application` struct. All handlers are methods on `*application`.

**Route groups** (`cmd/web/routes.go`):
- Public: `/`, `/health`, `/register`, `/login`
- Auth required: `/dashboard` (redirects by role)
- Teacher only: `/teacher/*`
- Student only: `/student/*`

**Models** (`internal/models/`) own their DB methods directly (no repository/service layers). Models are accessed via `app.models.<Model>`.

**Templates** (`ui/html/`): Cached at startup in `map[string]*template.Template`. Layout in `layouts/base.html`, pages in `pages/*/`, components in `components/`. **Changes to templates require app restart.**

**Validation** (`internal/validator/`): Use `validator.New()` + `CheckField` pattern.

## Key gotchas

- **Port mismatch:** `.env.example` says 8080, but `cmd/web/main.go` defaults to `:4000`. `APP_PORT` env var overrides the default. Makefile `run` target uses the Go default (`:4000`).
- **Makefile includes `.env`:** The Makefile has `include .env` at the top. The `.env` file must exist or `make` targets will fail. Copy `.env.example` to `.env` if missing.
- **Template cache:** Templates are parsed once at startup. Edit templates → restart the app.
- **No tests yet:** Zero `*_test.go` files. `make test` will pass vacuously.
- **CI is placeholder:** `.github/workflows/pr-check.yml` only echoes success — no real checks.
- **README is stale:** Describes planned Next.js/NestJS/Python stack. Actual stack is Go + Chi + pgx + Go templates.
- **Most handlers are stubs:** `student_handlers.go` and parts of `teacher_handlers.go` render templates with no data. Auth flow and course creation are the only implemented features.
- **Session cookie config:** `Secure: true` in production — will not work over plain HTTP unless you set `APP_ENV=development` or use HTTPS.
