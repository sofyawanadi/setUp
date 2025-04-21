# setUp

## Project Overview

This project is a backend service built with Go, designed to handle user authentication and management. It includes the following features:

### Features
- **User Authentication**:
    - Login with JWT-based authentication.
    - User registration with secure password hashing.
    - Forgot password functionality with email-based password reset.

- **User Management**:
    - Full CRUD operations for user data.

- **Security**:
    - JSON Web Tokens (JWT) for secure authentication.
    - Password hashing using industry-standard algorithms.

- **Email Integration**:
    - Send email for password reset functionality.

### Initial Setup

To set up the project, follow these steps:

1. **Install Dependencies**:
     - Use `go mod` to manage dependencies.
     - Install required libraries such as:
         - `gorm` for ORM.
         - `jwt-go` for JWT handling.
         - `gomail` or similar for email sending.

2. **Database Configuration**:
     - Set up a relational database (e.g., PostgreSQL, MySQL).
     - Configure database connection in the project.

3. **Environment Variables**:
     - Use `.env` file to manage sensitive configurations like:
         - Database credentials.
         - JWT secret key.
         - Email server settings.

4. **Folder Structure**:
     - Implement **Clean Architecture** to organize the project:
         - `domain/` for business logic.
         - `usecase/` for application logic.
         - `repository/` for data access.
         - `delivery/` for HTTP handlers.
         - `config/` for configuration management.

5. **SOLID Principles**:
     - Ensure the code adheres to SOLID principles for maintainability and scalability:
         - Single Responsibility Principle.
         - Open/Closed Principle.
         - Liskov Substitution Principle.
         - Interface Segregation Principle.
         - Dependency Inversion Principle.

6. **Testing**:
     - Write unit tests for core functionalities.
     - Use mocking for external dependencies.

### Additional Features
- **Role-Based Access Control (RBAC)**:
    - Implement roles (e.g., Admin, User) for access control.
- **API Documentation**:
    - Use tools like Swagger for API documentation.
- **Rate Limiting**:
    - Prevent abuse by limiting the number of requests per user.

### Getting Started

1. Clone the repository:
     ```bash
     git clone <repository-url>
     cd setUp
     ```

2. Install dependencies:
     ```bash
     go mod tidy
     ```

3. Run the application:
     ```bash
     go run main.go
     ```

4. Access the API at `http://localhost:<port>`.

### Contribution

Feel free to contribute to this project by submitting issues or pull requests.

