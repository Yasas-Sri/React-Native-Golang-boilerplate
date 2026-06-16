# React Native (Expo) + Go Template

A clean starter template for building a mobile app with a Go backend. Clone it, configure it, and start building.

- **`mobile/`** — React Native app built with [Expo](https://expo.dev) + [expo-router](https://docs.expo.dev/router/introduction/)
- **`server/`** — Go HTTP server using the [chi](https://github.com/go-chi/chi) router

The base template ships with no database .

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [Node.js](https://nodejs.org/) 18+
- The [Expo Go](https://expo.dev/go) app on your phone, or an iOS/Android simulator

## Quick start

### 1. Run the Go server

```bash
cd server
cp .env.example .env        # optional — defaults work
go run .
```

The server listens on `http://localhost:8080`. Verify it:

```bash
curl http://localhost:8080/healthz
# {"status":"ok","time":"..."}
curl http://localhost:8080/api/hello?name=Dev
# {"message":"Hello, Dev!"}
```

### 2. Run the mobile app

```bash
cd mobile
npm install
cp .env.example .env        # set EXPO_PUBLIC_API_URL for your setup
npx expo start -c           # clear cache to ensure .env is correctly loaded
```

Then press `w` (web), `i` (iOS simulator), `a` (Android emulator), or scan the QR code with Expo Go.

Tap **"Call Go server"** on the home screen — it calls `/api/hello` and shows the response, proving the mobile ↔ backend wiring works.

### Pointing the app at the server

Expo only loads environment variables with the `EXPO_PUBLIC_` prefix, and requires the file to be named exactly `.env`. `EXPO_PUBLIC_API_URL` must be reachable from wherever the app runs:

| Running on            | Use                              |
| --------------------- | -------------------------------- |
| Web / iOS simulator   | `http://localhost:8080`          |
| Android emulator      | `http://10.0.2.2:8080`           |
| Physical device       | `http://<your-computer-LAN-IP>:8080` (e.g., `192.168.X.X`) |

**Troubleshooting Network Requests:**
If you get a "Network request failed" error on a physical device, ensure:
1. Your `.env` file exists in the `mobile` directory and points to your computer's local IP address.
2. Your phone and computer are connected to the exact same Wi-Fi network.
3. You restarted the Expo server with a cleared cache after creating/modifying `.env` (`npx expo start -c`).

## Project layout

```
.
├── mobile/                 # Expo app
│   ├── app/                # expo-router screens (file-based routing)
│   │   ├── _layout.tsx     # root navigation stack
│   │   └── index.tsx       # home screen — calls the Go server
│   ├── api/index.ts        # API client (one place for the server URL)
│   └── .env.example
├── server/                 # Go backend
│   ├── main.go             # entrypoint + graceful shutdown
│   ├── internal/api/api.go # router, middleware, handlers
│   └── .env.example

```

## future work

Add a database template. Each one implements a common `Store` interface in the server so business logic never changes when you swap databases. Database options (Supabase, Firebase, MongoDB, Postgres, …) are added later, one at a time, as swappable implementations of a common Go `Store` interface. 
 
