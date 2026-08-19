# Restructure Plan: Alex Edwards "Let's Go" / "Let's Go Further" Architecture

## Target Directory Structure

```
sabify/
├── cmd/
│   └── web/
│       ├── main.go              # Entry point: config, DB, DI, graceful shutdown
│       └── helpers.go           # serverError, clientError, notFound, render helpers
├── internal/
│   ├── models/
│   │   ├── users.go            # User struct + DB methods (Create, FindByEmail, FindByID)
│   │   ├── courses.go          # Course struct + DB methods
│   │   ├── quizzes.go          # Quiz struct + DB methods
│   │   ├── questions.go        # Question struct + DB methods
│   │   ├── materials.go        # Material struct + DB methods
│   │   ├── submissions.go      # Submission struct + DB methods
│   │   └── studygroups.go      # StudyGroup struct + DB methods
│   ├── validator/
│   │   └── validator.go        # NonFieldErrors map, ValidateField, Valid
│   ├── sessions/
│   │   └── sessions.go         # Session manager (Let's Go Further)
│   ├── templates/
│   │   └── functions.go        # Template helper functions (date formatting, etc.)
│   └── middleware/
│       └── middleware.go        # SetHeaders, LogRequest, RecoverPanic, Authenticate
├── ui/
│   ├── static/
│   │   ├── css/                # (all existing CSS files, unchanged)
│   │   └── js/                 # (all existing JS files, unchanged)
│   └── html/
│       ├── pages/
│       │   └── home/
│       │       └── index.html
│       ├── layouts/
│       │   ├── base.html
│       │   └── public.html
│       └── components/
│           ├── navbar.html
│           └── footer.html
├── migrations/
│   └── 001_initial_schema.sql  # Moved from internal/database/migrations/
├── Makefile
├── .env
├── .env.example                # Committed template (no secrets)
├── docker-compose.yml
├── Dockerfile                  # Multi-stage build
├── go.mod
├── go.sum
├── doc/
├── README.md
├── .gitignore
└── .github/workflows/
    └── pr-check.yml
```

## Key Architectural Principles Applied

### 1. Central `application` struct (DI hub)

No more separate handler/service/repository constructors with manual wiring. One struct holds everything:

```go
type application struct {
    config        config
    logger        *slog.Logger
    models        models
    templateCache map[string]*template.Template
    session       *scs.SessionManager
}
```

### 2. Models own DB operations (no repositories package)

Each model file contains the struct AND its database methods. This is Alex Edwards' explicit pattern. The `repositories/` package is deleted entirely. Methods take `*pgxpool.Pool` as a receiver or are called on a model-specific helper.

### 3. No services layer

Business logic (validation, bcrypt, etc.) lives directly in handlers. The `services/` package is deleted entirely.

### 4. Handlers are methods on `*application`

```go
func (app *application) register(w http.ResponseWriter, r *http.Request) { ... }
```

No more separate handler structs with injected services.

### 5. Template caching at startup

Templates parsed once into `map[string]*template.Template`, not per-request.

### 6. Graceful shutdown

Signal handling with `context.Context`, `server.Shutdown(ctx)`.

### 7. Structured logging

`log/slog` with JSON or text handler, replacing `log.Println`.

### 8. Proper error handling

`serverError()` logs the stack trace, returns generic 500. `clientError()` returns 4xx. `notFound()` returns 404. `render()` handles template execution with error logging.

### 9. Middleware on `application`

```go
func (app *application) recoverPanic(next http.Handler) http.Handler { ... }
```

### 10. Health check with DB verification

`GET /health` pings the database, returns JSON status.

## Files Deleted

| File | Reason |
|------|--------|
| `internal/repositories/` (entire dir) | DB methods merge into models |
| `internal/services/` (entire dir) | Business logic moves to handlers |
| `internal/routes/routes.go` | Routes defined in `cmd/web/main.go` |
| `internal/config/config.go` | Config struct lives in `cmd/web/main.go` or `internal/config/config.go` |
| `internal/database/postgres.go` | DB connection moves to `cmd/web/main.go` |
| `internal/handlers/` (entire dir) | Handlers become methods on `*application` |
| `static/` (root level) | Moved to `ui/static/` |
| `templates/` (root level) | Moved to `ui/html/` |
| `internal/database/migrations/` | Moved to `migrations/` |

## Files Created / Rewritten

| File | Purpose |
|------|---------|
| `cmd/web/main.go` | Config loading, DB connection, template cache, DI, graceful shutdown, routes |
| `cmd/web/helpers.go` | `serverError`, `clientError`, `notFound`, `render` helpers |
| `cmd/web/routes.go` | `routes()` method on `*application`, returns `chi.Mux` |
| `internal/models/users.go` | User struct + `UserModel` with all DB methods |
| `internal/models/courses.go` | Course struct + `CourseModel` with all DB methods |
| `internal/models/quizzes.go` | Quiz struct + `QuizModel` with all DB methods |
| `internal/models/questions.go` | Question struct + `QuestionModel` with all DB methods |
| `internal/models/materials.go` | Material struct + `MaterialModel` (empty for now) |
| `internal/models/submissions.go` | Submission struct + `SubmissionModel` with DB methods |
| `internal/models/studygroups.go` | StudyGroup struct + `StudyGroupModel` (empty for now) |
| `internal/models/models.go` | `Models` struct aggregating all model types |
| `internal/validator/validator.go` | `Validator` with `NonFieldErrors`, `ValidateField`, `Valid` |
| `internal/sessions/sessions.go` | Session manager setup (cookie-based) |
| `internal/middleware/middleware.go` | `SetHeaders`, `LogRequest`, `RecoverPanic`, `Authenticate` |
| `Makefile` | `run`, `build`, `test`, `lint`, `db_up`, `db_down` |
| `.env.example` | Committed template with placeholder values |
| `Dockerfile` | Multi-stage build: build binary, copy to distroless/alpine |

## Migration Steps (execution order)

### Phase 1: Structural moves
1. Move `static/` → `ui/static/`
2. Move `templates/` → `ui/html/`
3. Move `internal/database/migrations/` → `migrations/`
4. Update CSS `url()` paths in `base.html` to use `/static/` prefix (no `../../`)

### Phase 2: Models (merge repositories into models)
5. Rewrite `internal/models/users.go` — add `UserModel{DB *pgxpool.Pool}` and all methods from `user_repository.go`
6. Rewrite `internal/models/courses.go` — add `CourseModel` with methods from `course_repository.go`
7. Rewrite `internal/models/quizzes.go` — add `QuizModel` with methods from `quiz_repository.go`
8. Rewrite `internal/models/questions.go` — add `QuestionModel` with methods from `quiz_repository.go` (questions part)
9. Rewrite `internal/models/submissions.go` — add `SubmissionModel` with methods from `submission_repository.go`
10. Create `internal/models/models.go` — `Models` struct aggregating all model types
11. Delete `internal/repositories/` entirely

### Phase 3: Delete services, create application struct
12. Delete `internal/services/` entirely
13. Create `internal/validator/validator.go`
14. Create `internal/sessions/sessions.go`
15. Rewrite `internal/middleware/middleware.go` — methods on `*application`

### Phase 4: Handlers and routes
16. Create `cmd/web/helpers.go` — `serverError`, `clientError`, `notFound`, `render`
17. Create `cmd/web/routes.go` — `routes()` method on `*application`
18. Create `cmd/web/main.go` — full entry point with config, DB, DI, shutdown, routes
19. Delete `internal/routes/routes.go`
20. Delete `internal/handlers/` entirely
21. Delete `internal/config/config.go`
22. Delete `internal/database/postgres.go`

### Phase 5: Build infrastructure
23. Create `Makefile`
24. Create `.env.example`
25. Write `Dockerfile` (multi-stage)
26. Update `.gitignore` (add `tmp/`, `*.exe`, etc.)

### Phase 6: Verify
27. `go build ./cmd/web`
28. `go vet ./...`
29. `go test ./...`
30. Verify the app starts and serves the homepage

## What stays the same

- All CSS/JS/HTML content (just moved, not rewritten)
- All SQL migrations (just moved)
- `go.mod` module name and dependencies
- `docker-compose.yml` (Postgres service unchanged)
- Database schema
- `.env` values
