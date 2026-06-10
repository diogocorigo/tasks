# Tasks

Tasks is a simple task manager written in Go and Webview.

## Prerequisites
 
Make sure you have the following installed:
 
- [Go](https://go.dev/dl/) 1.26+
- [Node.js](https://nodejs.org/) 24+
- A C/C++ compiler (required by WebView)
    - **Windows:** [MSYS2](https://www.msys2.org/) with `mingw-w64-x86_64-gcc`
    - **macOS:** Xcode Command Line Tools (`xcode-select --install`)
    - **Linux:** `gcc` via your package manager (`apt install gcc`)

---
 
## Project Structure

```
myapp/
├── main.go
├── .env
├── Makefile
└── frontend/
    ├── package.json
    ├── vite.config.ts
    ├── src/
    └── dist/          ← generated on build, do not edit manually
```

---
 
## Development
 
In dev mode, Vite runs its dev server (with hot reload) and Go proxies requests to it.

### Option A — Using Make
 
```bash
make dev
```
 
### Option B — Without Make
 
```bash
cd frontend && npm run dev:app
```
 
Both options start two processes in parallel:
- **Vite** on `http://localhost:5173` with hot reload
- **Go** app with `APP_DEBUG=true`, proxying the WebView to Vite
> If either process crashes or is stopped, the other one is automatically killed too (`--kill-others`).

---
 
## Production Build
 
In production, Vite compiles the frontend into a static `dist/` folder which Go then embeds directly into the binary — no external files needed.
 
### Option A — Using Make
 
```bash
make build
```
 
### Option B — Without Make
 
```bash
# Step 1 — build the frontend
cd frontend && npm run build
 
# Step 2 — compile the Go binary (from project root)
cd .. && go build -o .build/tasks.exe .
```
 
This produces a single `tasks` executable (or `tasks.exe` on Windows) with everything bundled inside.

---
 
## Environment Variables
 
Configuration lives in the `.env` file at the project root:
 
| Variable    | Values          | Description                          |
|-------------|-----------------|--------------------------------------|
| `APP_DEBUG` | `true` / `false` | Switches between dev proxy and embedded file serving |
 
> In dev mode `APP_DEBUG` is set to `true` automatically by the `dev:app` script.
> For production builds leave it as `false` or omit it entirely.

---
 
## Installing Frontend Dependencies
 
If you are setting up the project for the first time or after cloning:
 
```bash
cd frontend && npm install
```

And copy the `.env_example` file to `.env` and edit it to your liking.
