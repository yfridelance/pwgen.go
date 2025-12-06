# PWGen Go

A secure, high-performance Password and Passphrase Generator built with **Go** (Backend) and **React/Vite** (Frontend PWA).

![Frontend PWA](https://raw.githubusercontent.com/yfridelance/pwgen.go/main/web/public/vite.svg)

## Features

- **Secure Password Generation**: Cryptographically secure random passwords with configurable complexity (Uppercase, Digits, Special Characters).
- **Homoglyph Exclusion**: Option to exclude confusing characters (e.g., `l`, `1`, `I`, `0`, `O`).
- **Passphrase Generation**: Generate memorable passphrases (e.g., `correct-horse-battery-staple`) using English, Finnish, or French wordlists.
- **HIBP Integration**: Application automatically checks generated passwords against [HaveIBeenPwned](https://haveibeenpwned.com/) (privacy-preserving k-anonymity model) to warn if a password has been compromised.
- **Modern PWA Frontend**: Responsive web interface installable on mobile and desktop.
- **Dockerized**: Fully containerized setup with separate Backend and Frontend services.

## Quick Start

### Prerequisites
- Docker & Docker Compose

### Run via Docker
```bash
docker-compose up --build -d
```
Access the application at: **http://localhost:5070**

## Configuration

Configuration is handled via Environment Variables in `docker-compose.yml`.

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Backend service port | `5069` |
| `MULTI_GEN` | Allow generating multiple passwords at once (max 5) | `true` |
| `HIBP_ENABLED` | Enable HaveIBeenPwned security checks | `true` |
| `DEFAULT_LENGTH` | Default length for passwords | `16` |
| `CUSTOM_WORDLIST_URL` | Optional URL for custom wordlist | `-` |

### Ports
- **Frontend**: `5070` (Mapped to container port `80`)
- **Backend API**: `5069` (Mapped to container port `5069`)

## API Usage

The backend exposes a REST API at `/api/generate`.

### Generate Password
**POST** `/api/generate`
```json
{
  "type": "password",
  "count": 1,
  "length": 24,
  "include_uppercase": true,
  "include_digits": true,
  "include_special": true,
  "exclude_homoglyphs": true
}
```

### Generate Passphrase
**POST** `/api/generate`
```json
{
  "type": "passphrase",
  "word_count": 4,
  "separator": "-",
  "capitalize": true,
  "language": "en"
}
```

## Development

### Backend (Go)
```bash
go run cmd/server/main.go
```

### Frontend (React/Vite)
```bash
cd web
npm install
npm run dev
```

## License
GNU General Public License v3.0
