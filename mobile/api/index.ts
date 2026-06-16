// API client for the Go backend.
//
// Import from anywhere in the app with:  import { getHello } from "@/api";
//
// The base URL comes from the EXPO_PUBLIC_API_URL env var (set it in
// mobile/.env). Pick a value the device/simulator can actually reach:
//   - Web / iOS simulator: http://localhost:8080
//   - Android emulator:    http://10.0.2.2:8080  (localhost alias)
//   - Physical device:     http://<your-computer-LAN-IP>:8080
//
// EXPO_PUBLIC_-prefixed vars are inlined into the app at build time by Expo,
// so restart the dev server after changing .env.
export const API_URL =
  process.env.EXPO_PUBLIC_API_URL ?? "http://localhost:8080";

export type HelloResponse = {
  message: string;
};

export type HealthResponse = {
  status: string;
  time: string;
};

/** Low-level GET helper that returns parsed JSON or throws on a non-2xx. */
async function getJSON<T>(path: string): Promise<T> {
  const res = await fetch(`${API_URL}${path}`);
  if (!res.ok) {
    throw new Error(`Request to ${path} failed: ${res.status}`);
  }
  return (await res.json()) as T;
}

/** GET /api/hello — sample endpoint that proves end-to-end connectivity. */
export function getHello(name?: string): Promise<HelloResponse> {
  const query = name ? `?name=${encodeURIComponent(name)}` : "";
  return getJSON<HelloResponse>(`/api/hello${query}`);
}

/** GET /healthz — backend liveness check. */
export function getHealth(): Promise<HealthResponse> {
  return getJSON<HealthResponse>("/healthz");
}
