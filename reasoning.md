# Reasoning and Design Decisions

## Overview

This project is a RESTful API built in Go to manage users with their name and date of birth (DOB).  
The API dynamically calculates and returns a user’s age instead of storing it in the database.

The main goal was to follow the given requirements strictly while keeping the code clean, readable, and easy to explain.

---

## Architecture

The application follows a layered architecture:

- **Handler layer**  
  Responsible for handling HTTP requests and responses, parsing input, and returning proper status codes.

- **Service layer**  
  Contains business logic such as calculating the user’s age and coordinating operations between layers.

- **Repository layer**  
  Handles all database interactions using SQLC-generated queries.

- **Routes**  
  Maps API endpoints to their respective handlers.

This separation of concerns improves maintainability and makes the application easier to understand and extend.

---

## Database Design

A single `users` table is used with the following fields:
- `id` (SERIAL PRIMARY KEY)
- `name` (TEXT NOT NULL)
- `dob` (DATE NOT NULL)

The `dob` field is stored as a DATE.  
The user’s age is calculated dynamically when fetching data rather than storing it in the database.

---

## Age Calculation

Age is calculated in the service layer using Go’s `time` package.  
By calculating age at runtime, the API always returns an accurate value and avoids data inconsistency that would occur if age were stored.

---

## SQLC Usage

SQLC is used to generate a type-safe database access layer from raw SQL queries.  
This approach keeps SQL explicit while providing compile-time safety and reducing runtime errors.

PostgreSQL is accessed using the pgx driver as required.

---

## Validation

Request input is validated using `go-playground/validator`.  
This ensures required fields are present and correctly formatted before processing any request.

---

## Logging

Uber Zap is used for structured logging.  
Key application lifecycle events such as server startup and database connection are logged to improve observability.

---

## Error Handling

The API returns appropriate HTTP status codes:
- 400 for invalid input
- 404 for missing resources
- 500 for internal server errors

This ensures clear and predictable API behavior.

---

## Conclusion

This project demonstrates a clean Go backend implementation with proper separation of concerns, dynamic data computation, and strict adherence to the provided requirements.  
The focus was on clarity, correctness, and maintainability rather than over-engineering.
