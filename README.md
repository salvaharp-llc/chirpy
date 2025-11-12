# chirpy

Tiny microblog service written in Go. Serves static files, exposes a small REST API and uses PostgreSQL for storage.

## Quick start

Requirements:
- Go 1.20+ (or a recent stable Go)
- PostgreSQL

Create or start a PostgreSQL instance and set environment variables (example):

```bash
export DB_URL="postgres://postgres:password@localhost:5432/chirpy?sslmode=disable"
export PLATFORM="dev"
export JWT_SECRET="replace-me"
export POLKA_KEY="replace-me"
```

Run the server:

```bash
go run .
```

The server listens on port `8080` by default.

## Configuration / Environment variables

- `DB_URL` (required) — Postgres DSN used by the app
- `PLATFORM` (required) — short name for the running platform (used in logs)
- `JWT_SECRET` (required) — secret used to sign JWT tokens
- `POLKA_KEY` (required) — webhook/key for external service

## Database / Migrations

Schema files live in `sql/schema/` and queries in `sql/queries/`.

Quick, manual setup (one-off):

```bash
psql "$DB_URL" -f sql/schema/001_users.sql
psql "$DB_URL" -f sql/schema/002_chirps.sql
psql "$DB_URL" -f sql/schema/003_passwords.sql
psql "$DB_URL" -f sql/schema/004_refresh_tokens.sql
psql "$DB_URL" -f sql/schema/005_chirpy_red.sql
```

## API (selected)

- `GET /api/healthz` — health check
- `POST /api/validate_chirp` — validate chirp content
- `POST /api/login` — login/auth
- `POST /api/refresh` — refresh JWT
- `POST /api/revoke` — revoke refresh token
- `POST /api/chirps` — create chirp
- `GET /api/chirps` — list chirps (query params: `author_id`, `sort`)
- `GET /api/chirps/{chirpID}` — get a chirp
- `DELETE /api/chirps/{chirpID}` — delete a chirp
- `POST /api/users` — create user
- `PUT /api/users` — update user
- `GET /admin/metrics` — metrics
- `POST /admin/reset` — reset metrics/data

Example health check:

```bash
curl -i http://localhost:8080/api/healthz
```
