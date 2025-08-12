# SnippetBox

SnippetBox is a web application for creating, sharing, and managing code snippets. It provides a secure and easy-to-use platform for developers to store and share their favorite code snippets.

## Features

- **User Authentication:** Secure user signup, login, and logout functionality.
- **Snippet Management:** Create, view, and manage your code snippets.
- **Password Security:** Passwords are securely hashed using the bcrypt algorithm.
- **CSRF Protection:** All forms are protected against Cross-Site Request Forgery attacks.
- **Session Management:** Secure session management with server-side session storage in MySQL.
- **HTTPS/TLS Support:** The application runs on HTTPS for secure communication.
- **Centralized Error Handling:** A robust error handling mechanism is in place.
- **Dependency Injection:** The application uses dependency injection for better testability and maintainability.
- **Self-Contained Deployment:** Go's `embed` package is used to create a self-contained binary with all static assets and HTML templates.
- **Comprehensive Testing:** The application has a comprehensive test suite, including unit and integration tests.

## Technologies Used

- **Backend:** Go
- **Database:** MySQL
- **Frontend:** HTML, CSS, JavaScript
- **Libraries:**
  - `github.com/alexedwards/scs/v2`: Session management
  - `github.com/justinas/nosurf`: CSRF protection
  - `github.com/go-playground/form/v4`: Form decoding and validation
  - `github.com/go-sql-driver/mysql`: MySQL driver
  - `github.com/justinas/alice`: Middleware chaining
  - `golang.org/x/crypto`: Cryptography primitives (used for bcrypt)

## Getting Started

### Prerequisites

- Go (version 1.24.5 or higher)
- MySQL

### Installation

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/print-agent/snippet-box.git
    cd snippet-box
    ```

2.  **Set up the database:**
    - Create a MySQL database named `snippetbox`.
    - Create a user with the necessary privileges to access the database.
    - The application uses the following tables: `snippets`, `users`, and `sessions`. You can find the schema in the `internal/models/testdata/setup.sql` file.

3.  **Configure the application:**
    The application is configured using command-line flags:
    - `-addr`: The network address to run the server on (default: `:4000`).
    - `-dns`: The MySQL data source name (DSN). The default is `web:pass@/snippetbox?parseTime=true`. You should change this to match your MySQL configuration.

4.  **Generate TLS certificates:**
    The application requires TLS certificates to run on HTTPS. You can generate self-signed certificates using the following commands:
    ```sh
    mkdir tls
    cd tls
    go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
    ```
    This will create `cert.pem` and `key.pem` files in the `tls` directory.

### Running the Application

To run the application, use the following command:

```sh
go run ./cmd/web -dns="your_mysql_dsn"
```

Replace `"your_mysql_dsn"` with your MySQL data source name. For example:

```sh
go run ./cmd/web -dns="user:password@/snippetbox?parseTime=true"
```

The server will start on `https://localhost:4000`.

## Testing

To run the test suite, use the following command:

```sh
go test ./...
```

## Project Structure

```
.
├── cmd/web/                # Main application
├── internal/               # Internal packages
│   ├── assert/             # Test assertion helpers
│   ├── models/             # Database models and mocks
│   └── validator/          # Form validation helpers
├── tls/                    # TLS certificates
└── ui/                     # User interface (HTML, CSS, JS)
    ├── efs.go              # Embedded file system
    ├── html/               # HTML templates
    └── static/             # Static assets (CSS, JS, images)
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

