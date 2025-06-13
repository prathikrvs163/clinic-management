# 🏥 Clinic Management System

A full-stack web application built with **Golang**, **PostgreSQL**, and **vanilla HTML/CSS/JavaScript** to manage clinic operations. It supports two user roles — **Receptionist** and **Doctor** — with secure login and full CRUD functionality for managing patient records. Deployed on **Render** (backend + frontend + PostgreSQL).

🌐 **[Live Demo](https://clinic-management-3fus.onrender.com/)**

---

## 🚀 Features

- ✅ User login with **JWT authentication**
- 👩‍⚕️ Role-based access:
  - **Receptionist**: Add, view, delete patients
  - **Doctor**: View, update patient information
- 🗃️ PostgreSQL database integration
- 🧪 Unit-tested login system
- 🌐 RESTful API with JSON responses
- 🔐 Passwords are securely hashed
- 📄 Full frontend built with HTML/CSS/JS
- 📦 Deployed backend + static frontend on Render

---

## 🧱 Tech Stack

| Layer | Technology |
|-------|-----------|
| Frontend | HTML, CSS, JavaScript |
| Backend | Golang, Gorilla Mux |
| Database | PostgreSQL (Render) |
| Authentication | JWT |
| Hosting | Render.com |

---

## 📁 Project Structure

```
clinic-management/
│
├── config/         # DB connection and environment loading
├── controllers/    # API handlers for auth & patient
├── middleware/     # CORS and JWT auth middleware
├── models/         # Patient and User models
├── routes/         # API route mapping
├── public/         # Frontend (index.html)
├── utils/          # JWT helper functions
├── .env            # Local environment file (not pushed to Git)
├── main.go         # App entry point (server + routes)
├── go.mod          # Go dependencies
└── go.sum          # Dependency checksums
```

---

## 🛠️ Setup Instructions

### 🔧 Local Development

1. **Clone the Repository**
   ```bash
   git clone https://github.com/prathikrvs163/clinic-management.git
   cd clinic-management
   ```

2. **Create `.env` file**
   ```env
   DB_HOST=127.0.0.1
   DB_PORT=5433
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=clinic
   ```

3. **Start PostgreSQL Locally**
   - Use pgAdmin or Docker to run PostgreSQL
   - Create a database named `clinic`

4. **Run the Application**
   ```bash
   go run main.go
   ```

5. Visit `http://localhost:8000` to use the app

### 🚀 Deployment (on Render)

1. **Create PostgreSQL Instance**
   - From Render dashboard → Databases → PostgreSQL
   - Copy connection values (host, port, user, password, db name)

2. **Create New Web Service**
   - From Render dashboard → New Web Service
   - Connect your GitHub repo
   - Use build command: `go build -o main .`
   - Use start command: `./main`

3. **Add Environment Variables in Render**
   ```env
   DB_HOST=your_render_db_host
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   ```

4. **Live URL**
   - Your app will be live at: `https://<your-service-name>.onrender.com`

---

## 🔐 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/register` | Register a user (doctor/receptionist) |
| POST | `/login` | Login and receive JWT |
| GET | `/patients` | View all patients (auth required) |
| POST | `/patients` | Add a new patient (Receptionist only) |
| PUT | `/patients/:id` | Update patient (Doctor only) |
| DELETE | `/patients/:id` | Delete patient (Receptionist only) |

---

## 🧪 Sample Credentials

You can register new users using Postman or curl:

**Receptionist**
```json
{
  "email": "reception@clinic.com",
  "password": "pass123",
  "role": "receptionist"
}
```

**Doctor**
```json
{
  "email": "doctor@clinic.com",
  "password": "pass123",
  "role": "doctor"
}
```

---

## 👨‍💻 Author

**Sriram Prathik**
- [LinkedIn](https://www.linkedin.com/in/sriram-prathik-rachakonda-26654a143/)
- [GitHub](https://github.com/prathikrvs163)

---

## 📄 API Documentation

🔗 [Postman Collection](https://prathik-8229284.postman.co/workspace/Prathik's-Workspace~44924073-f5d1-4886-83b8-27b6eb0cd47b/collection/45415263-8ee02dbb-ad5d-45d5-b15c-b548428df011?action=share&creator=45415263)


