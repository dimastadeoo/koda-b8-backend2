# Endpoint Backend Login, getalldata, dan Register


Aplikasi sederhana **User Management** menggunakan **Golang**, **Gin Framework**, **PostgreSQL**, dan **HTML/CSS/JavaScript**.

Project ini menerapkan arsitektur berlapis (**Repository Pattern**) sehingga setiap layer memiliki tanggung jawab masing-masing.

## Tech Stack

### Backend

- Golang
- Gin Framework
- PostgreSQL
- pgxpool
- bcrypt

### Frontend

- HTML5
- CSS3
- JavaScript (Vanilla JS)
- Fetch API

---

# Arsitektur Project

Project menggunakan pola Repository Pattern.

```text
Models
   │
   ▼
Repository
   │
   ▼
Service
   │
   ▼
Handler
   │
   ▼
Container (Dependency Injection)
```

Penyimpanan data menggunakan **PostgreSQL**.

---

# Fitur

Backend

- Register User
- Login User
- Get All Users
- Get User By ID
- Create User
- Update User
- Delete User

Frontend

- Login
- Register
- Dashboard User
- Tambah User
- Edit User
- Delete User
- Logout

---

# Tampilan Aplikasi

<table>
    <tr>
        <td>Register data</td>
        <td>Tampilkan Get All Data Users</td>
        <td>Login</td>
    </tr>
    <tr>
        <td><img src="img/Screenshot_2026-07-20_15-02-03.png" alt="Register"></td>
        <td><img src="img/Screenshot_2026-07-20_15-02-27.png" alt="All Data"></td>
        <td><img src="img/Screenshot_2026-07-20_15-02-36.png" alt="Login"></td>
    </tr>
</table>


# Pengujian Backend

Pengujian endpoint menggunakan **VS Code REST Client**.
```bash
### Register
POST /auth/register HTTP/1.1
Host: localhost:8080
Content-Type: application/x-www-form-urlencoded

fullname=dimas&email=dimas1@mail.com&password=123


### Get All Data dengan Authorization
GET /users HTTP/1.1
Host: localhost:8080
Authorization: hello

### Get Data Profile
GET /users/12 HTTP/1.1
Host: localhost:8080
Authorization: hello

### Tambah Data user
POST /users HTTP/1.1
Host: localhost:8080
Authorization: hello
Content-Type: application/x-www-form-urlencoded

fullname=dimas&email=dimas@mail.com&password=123

### Delete Data User
DELETE /users/17 HTTP/1.1
Host: localhost:8080
Authorization: hello

### Update data user
PATCH /users/2 HTTP/1.1
Host: localhost:8080
Authorization: hello
Content-Type: application/x-www-form-urlencoded

fullname=dimas1&email=dimas4@mail.com

### Login
POST /auth/login HTTP/1.1
Host: localhost:8080
Content-Type: application/x-www-form-urlencoded

email=dimas2@mail.com&password=123

```

## ERD Table

```mermaid
erDiagram
    users {
        BIGINT id PK
        VARCHAR fullname
        VARCHAR email UK
        VARCHAR password
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }
```

---

# Struktur Project

```text
.
├── internal/
│   ├── di/
│   ├── handler/
│   ├── middlewares/
│   ├── models/
│   ├── repo/
│   ├── services/
│   └── lib/
│
├── frontend/
│   ├── css/
│   ├── js/
│   ├── index.html
│   ├── login.html
│   ├── register.html
│   └── users.html
│
├── migrations/
├── .env
├── go.mod
├── go.sum
├── main.go
├── main.go
├── Makefile
└── README.md
```

---

# Menjalankan Project

## Clone Repository

```bash
git clone https://github.com/username/repository.git

cd repository
```

---

## Install Dependency

```bash
go mod tidy
```

---

## Konfigurasi Environment

Buat file **.env**

```env
DATABASE_URL=postgres://username:password@host:port/database?sslmode=disable
PORT=Port Backend
PORT_FRONTEND= Port Frontend

PGUSER=user postgres
PGPASSWORD=password postgres
PGHOST=host postgres
PGPORT=port postgres
PGDATABASE=nama database
```

---

## Jalankan Backend

```bash
go run main.go
```

## Jalankan Frontend

Karena frontend menggunakan HTML biasa, cukup jalankan menggunakan **Live Server** di Visual Studio Code.

# Alur Aplikasi

```text
Home
 │
 ├── Register
 │      │
 │      ▼
 │   Login
 │
 ▼
Login
 │
 ▼
Dashboard User
 │
 ├── Get All
 ├── Tambah User
 ├── Edit User
 ├── Delete User
 └── Logout
```

---

# Catatan

- Password disimpan menggunakan **bcrypt hash**.
- Endpoint `/users` dilindungi menggunakan middleware Authorization.
- Frontend menggunakan **Fetch API** dengan `application/x-www-form-urlencoded`.
- Token Authorization disimpan di **localStorage** setelah login berhasil.

---

# Author

**Dimas Tadeo**