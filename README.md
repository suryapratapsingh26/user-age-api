# User Age API

A simple RESTful API built in Go to manage users with their name and date of birth (DOB).  
The API calculates and returns the user’s age dynamically when fetching user details.

---

## Features

- Create, update, delete users
- Fetch user details with dynamically calculated age
- List all users
- PostgreSQL database integration
- Input validation and structured logging

---

## Tech Stack

- Go (Golang)
- Fiber
- PostgreSQL
- SQLC
- go-playground/validator
- Uber Zap

---

## Project Structure

/cmd/server/main.go  
/config/  
/db/  
/internal/  

---

## API Endpoints

POST /users  
GET /users/:id  
PUT /users/:id  
DELETE /users/:id  
GET /users  

---

## Database

The `users` table stores:
- id
- name
- date of birth (dob)

Age is not stored in the database and is calculated dynamically.

---

## Setup & Run

1. Install Go and PostgreSQL  
2. Configure PostgreSQL and set the DATABASE_URL environment variable  
3. Run the database migration from db/migrations  
4. Start the server  

The server runs on port 3000.

---


