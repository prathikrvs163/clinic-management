# 🏥 Clinic Management System

A full-stack Clinic Management System built with **Golang** (backend) and **vanilla HTML/CSS/JS** (frontend), supporting role-based access for **Receptionists** and **Doctors**. The app is deployed on **Render**, with **PostgreSQL** as the database and **JWT** for secure authentication.

## 🔗 Live Demo

🌐 [View Application on Render](https://clinic-management-3fus.onrender.com)

---

## 📌 Features

### 👤 Authentication
- JWT-based login for Receptionists and Doctors
- Single login page for both roles

### 🩺 Receptionist Portal
- Register new patients
- View all patients
- Update or delete patient records

### 👨‍⚕️ Doctor Portal
- View all registered patients
- Update patient condition or details

### 🌐 Frontend
- Responsive UI built with HTML, CSS, and JavaScript
- Role-based actions (only receptionists can add/delete)

### 🛠 Backend
- RESTful APIs built with Go (Golang)
- Gorilla Mux router
- PostgreSQL integration with secure `.env` configuration
- CORS enabled
- API documented with Postman

---

## 🧱 Tech Stack

| Layer      | Technology           |
|------------|----------------------|
| Frontend   | HTML, CSS, JS        |
| Backend    | Golang, Gorilla Mux  |
| Database   | PostgreSQL (Render)  |
| Auth       | JWT                  |
| Hosting    | Render.com           |

---

## 🧪 Running Locally

### 1. Clone the repository

```bash
git clone https://github.com/prathikrvs163/clinic-management.git
cd clinic-management
```

2. Create a .env file

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=clinic
```

3. Run the application
   
```bash
go run main.go
```

🚀 Deployment (Render)

PostgreSQL instance created via Render

.env variables added in Render Dashboard → Environment

Go service deployed with go build -o main . && ./main

Static frontend served from /public

📂 Folder Structure

```bash
clinic-management/
│
├── config/         # DB connection and environment loading
├── controllers/    # Business logic and API handlers
├── middleware/     # CORS and auth middleware
├── models/         # Data models
├── routes/         # API route mappings
├── public/         # Frontend (index.html, styles, JS)
├── utils/          # Helper functions
├── .env            # Environment variables (excluded in Git)
├── main.go         # Entry point

```

📄 API Documentation

🔗 Postman Collection – (Add link if exported)

👨‍💻 Author
Sri Ram Prathik
LinkedIn • GitHub
