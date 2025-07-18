# Go Authenticator Project

This project is a **backend authenticator service built in Go**, designed using **Clean Architecture** to keep the code **organized and easy to understand**. The resource: [clean architecture](https://threedots.tech/post/introducing-clean-architecture/)

---

##  What this project does

--> Allows users to **register** with email and password (securely hashed with bcrypt).  
--> Allows users to **log in** and receive a **JWT token stored in an HTTP-only cookie** for authentication.  
--> Allows logged-in users to **fetch their profile via a protected route (`/getprofile`)** using JWT validation.  
--> Uses **MongoDB** to store user data.  
--> Uses **Clean Architecture** to keep concerns separated for clarity and scalability.

---

##  Implemented via Clean Architecture 

### 1. The Core (Entities)

**Directory:** `/internal/models/`  
- **`user.go`**: Defines the `User` struct (`ID`, `Email`, `Password`). Pure data, no dependency on database or HTTP.

---

### 2. Business Logic (Application Layer)

**Directory:** `/internal/services/`  
- **`user_service.go`**: Contains the application’s logic:
  - `RegisterUser`: hashes password, saves user.
  - `LoginUser`: verifies credentials, generates JWT.
  - `GetUserProfile`: retrieves user profile.

---

### 3. Interface (Adapter Layer)

**Directory:** `/internal/middleware/`  
- **`auth_middleware.go`**: Validates JWT from cookies for protected routes.

**Directory:** `/internal/handlers/`  
- **`register.go`**: Handles `POST /register`, decodes JSON, calls service to register the user.  
- **`login.go`**: Handles `POST /login`, decodes JSON, calls service to log in, sets JWT cookie.  
- **`profile_handler.go`**: Handles `GET /getprofile`, retrieves user profile if authenticated.

---

### 4. Infrastructure (Outer Layer)

**Directory:** `/internal/repositories/`  
- **`mongo.go`**: Manages MongoDB connection.  
- **`user_repository.go`**: Functions to interact with the `users` collection:
  - `CreateUser`: inserts a new user.
  - `GetUserByEmail`: retrieves a user by email.

**Directory:** `/pkg/utils/`  
- **`hash_utils.go`**: Provides password hashing and verification using `bcrypt`.  
- **`jwt_utils.go`**: Provides JWT generation and validation functions.

---

### 5. Entry Point

**Directory:** `/cmd/httpserver/`  
- **`main.go`**: The application entry point:
  - Connects to MongoDB.
  - Sets up HTTP routes (`/register`, `/login`, `/getprofile`).
  - Applies `AuthMiddleware` to protect routes.
  - Starts the HTTP server on `:8080`.

---


