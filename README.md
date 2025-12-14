# Shepherd (Auth Demo)

This project demonstrates a Gin (Go) backend with an embedded React (Vite) frontend using TailwindCSS + daisyUI. It includes:

- Email/password signup and login (simple in-memory demo user store)
- Google OAuth2 sign-in/up (server-side OAuth flow)
- Session via JWT stored in an HttpOnly cookie

## Setup

Set required environment variables (example):

```bash
export GOOGLE_CLIENT_ID=your-google-client-id
export GOOGLE_CLIENT_SECRET=your-google-client-secret
export APP_SECRET=some-secret-for-jwt
```

## Build & Run

From the repository root:

```bash
# install node deps and build UI, then build Go binary
make build

# run server
./bin/shepherd
```

Open http://localhost:8080/ to see the app. The Login and Signup pages are available at `/login` and `/signup`.

Note: This is a demo. Replace the in-memory user store with a proper database for production use, and secure cookies (Secure, SameSite, HTTPS).

## Development

To run the frontend in dev mode (hot reload) while proxying API calls to the backend:

```bash
# Start the Go backend (in one terminal)
go run ./

# In another terminal, start the Vite dev server
cd ui && npm install && npm run dev
```

The Vite dev server proxies `/api` and `/auth` to the local backend by default.
