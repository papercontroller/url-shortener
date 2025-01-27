# URL Shortener

URL Shortener is a web-based application that allows users to generate shortened versions of long URLs for easier sharing and management. It provides features like alias customization, URL deletion, and basic authentication for secure access.

## Features

- **Shorten URLs**: Generate a shortened URL for any valid long URL.
- **Custom Aliases**: Specify a custom alias for your shortened URL.
- **Delete Shortened URLs**: Remove previously created shortened URLs.
- **Basic Authentication**: Protect API endpoints using username and password.

## Technology Stack

- **Backend**: Golang
- **Framework**: [chi](https://github.com/go-chi/chi) for routing
- **Database**: SQLite for storage
- **Logging**: slog for structured logging
- **Validation**: go-playground/validator for input validation

## Endpoints

### 1. Create Shortened URL
**POST** `/url`

#### Request:
```json
{
  "url": "https://www.google.com",
  "alias": "google"
}
```

#### Response:
```json
{
  "status": "OK",
  "alias": "google"
}
```

#### Notes:
- If no alias is provided, a random alias is generated.

### 2. Delete Shortened URL
**DELETE** `/url/{alias}`

#### Request:
- Basic authentication required.
- `alias` is passed as a URL parameter.

#### Response:
```json
{
  "status": "OK"
}
```

### 3. Redirect to Original URL
**GET** `/{alias}`

#### Behavior:
- Redirects to the original URL associated with the alias.
- Returns a `404 Not Found` response if the alias does not exist.

## Authentication

Basic authentication is enabled for protected endpoints like URL deletion. Provide your username and password with each request:

```bash
curl -u <username>:<password> <endpoint>
```

## Configuration

The application uses a configuration file or environment variables for the following:
- `HTTPServer.User`: Username for Basic Auth.
- `HTTPServer.Password`: Password for Basic Auth.

Example configuration:
```yaml
HTTPServer:
  User: "myuser"
  Password: "mypass"
```

## How to Run

### Prerequisites
- Go installed (version 1.18 or higher)

### Steps
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd url-shortener
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The application will be available at `http://localhost:8000`.

## Testing with `curl`

### Create a Shortened URL:
```bash
curl -u myuser:mypass -X POST http://localhost:8000/url \
-d '{"url": "https://www.google.com", "alias": "google"}' \
-H "Content-Type: application/json"
```

### Delete a Shortened URL:
```bash
curl -u myuser:mypass -X DELETE http://localhost:8000/url/google
```

### Redirect:
Open in browser:
```
http://localhost:8000/google
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request to suggest improvements or new features.

---

**Author**: [Your Name]

