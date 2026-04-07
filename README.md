# Inforce
Simple full-stack web application built with React and Go, backed by PostgreSQL.

## Tech Used
Backend:
- Go
- Gin
- pgxpool
- cron
- godotenv
- cors

Frontend:
- React (Vite)
- TypeScript
- Axios

Database:
- PostgreSQL

Infrastructure:
- Podman

## Install
1. Clone repo
```zsh
git clone https://github.com/theverysameliquidsnake/inforce.git
cd inforce
```
2. Create network
>Note:
>(Docker users can replace podman with docker in all commands below)
```zsh
podman network create inforce
```
3. Run PostgrSQL container and create tables
```zsh
podman run --name postgres \
    --network inforce \
    -e POSTGRES_USER=user \
    -e POSTGRES_PASSWORD=pass \
    -e POSTGRES_DB=inforce \
    -p 5432:5432
    postgres:latest
```
```sql
CREATE TABLE user_events (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    action TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metadata JSONB
);
CREATE TABLE user_event_summaries (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    events INT NOT NULL,
    start_time TIMESTAMPTZ datetime NOT NULL
);
```
4. Build and run Go
```zsh
podman build -t go-api ./go
podman run --name go-api \
    --network inforce \
    -e POSTGRES_URL=postgres://user:pass@postgres:5432/inforce \
    -e REACT_URL=http://localhost:5173 \
    -p 8000:8000 \
    go-api:latest
```
5. Build and run React
```zsh
podman build -t react ./react
podman run --name react \
    -e VITE_API_URL=http://localhost:8000 \
    -p 5173:5173 \
    react:latest
```

## Use
### UI:

Web UI for viewing and filtering events at http://localhost:5173

### REST API:
#### Create event
Endpoint `POST /event`

Request body:
```json
{
  "user_id": 42,
  "action": "page_view",
  "metadata": {
    "page": "/home"
  }
}
```
- `user_id` (integer, required)
- `action` (string, required)
- `metadata` (object, optional)

Example:
```zsh
curl -X POST http://localhost:8000/event \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 42,
    "action": "page_view",
    "metadata": {"page": "/home"}
  }'
```

#### Get events
Endpoint `GET /event`

Optional query parameters:
- `user_id`  (integer)
- `start_time` (RFC3339 datetime)
- `end_time` (RFC3339 datetime)

Example:
```zsh
curl "http://localhost:8000/event?user_id=42&start_time=2026-04-06T22:55:00Z&end_time=2026-04-06T23:00:00Z"
```
