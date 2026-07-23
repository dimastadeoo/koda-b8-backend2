# Endpoint Backend Login, getalldata, dan Register


Aplikasi sederhana **User Management** menggunakan **Golang**, **Gin Framework**, **PostgreSQL**, dan **HTML/CSS/JavaScript**.

Project ini menerapkan arsitektur berlapis (**Repository Pattern**) sehingga setiap layer memiliki tanggung jawab masing-masing.

## API Documentation

http://localhost:8080/swagger/index.html

## Tech Stack

### Backend

- Golang
- Gin Gonic
- PostgreSQL
- pgxpool
- JWT
- Swagger Go
- golang-migrate

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

- User Authentication
    - Register
    - Login
    - JWT Authentication

- User Management
    - Create User
    - Get All User
    - Get User By ID
    - Update User
    - Delete User

- Upload Profile Picture
    - Upload Image
    - Validation Image Extension
    - Validation File Size
    - Delete Old Image

- Documentation
    - Swagger UI

- Database
    - PostgreSQL
    - Migration

- Frontend
    - Login
    - Register
    - User CRUD

---

# Tampilan Aplikasi

<table>
    <tr>
        <td>Register data</td>
        <td>Tampilkan Get All Data Users</td>
        <td>Login</td>
    </tr>
    <tr>
        <td><img src="img/Screenshot_2026-07-21_21-59-47.png" alt="Register"></td>
        <td><img src="img/Screenshot_2026-07-21_22-00-03.png" alt="All Data"></td>
        <td><img src="img/Screenshot_2026-07-21_21-59-42.png" alt="Login"></td>
    </tr>
</table>

<table>
    <tr>
        <td>Tambah data</td>
        <td>Update Data</td>
        <td>Delete Data</td>
    </tr>
    <tr>
        <td><img src="img/Screenshot_2026-07-21_22-00-13.png" alt="Add Data"></td>
        <td><img src="img/Screenshot_2026-07-21_22-00-21.png" alt="Update"></td>
        <td><img src="img/Screenshot_2026-07-21_22-00-27.png" alt="Delete"></td>
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
Authorization: Bearer JWT Token

### Get Data Profile
GET /users/12 HTTP/1.1
Host: localhost:8080
Authorization: Bearer JWT Token

### Tambah Data user
POST /users HTTP/1.1
Host: localhost:8080
Authorization: Bearer JWT Token
Content-Type: application/x-www-form-urlencoded

fullname=dimas&email=dimas@mail.com&password=123

### Delete Data User
DELETE /users/17 HTTP/1.1
Host: localhost:8080
Authorization: Bearer JWT Token

### Update data user
PATCH /users/2 HTTP/1.1
Host: localhost:8080
Authorization: Bearer JWT Token
Content-Type: application/x-www-form-urlencoded

fullname=dimas1&email=dimas4@mail.com

PATCH /users/2/picture HTTP/1.1
Host: localhost:8080
Authorization: Bearer JWT Token
Content-Type: multipart/form-data; boundary=webf

--webf
Content-Disposition: form-data; name="picture"; filename="images.jpeg"
Content-Type: image/jpg

< /home/dimastadeo/Downloads/test/images.jpeg
--webf--

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
        VARCHAR picture
        BIGINT created_at FK
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
├── docs/
├── uploads/
├── .env
├── go.mod
├── go.sum
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

JWT_KEY=key jwt

PGUSER=user postgres
PGPASSWORD=password postgres
PGHOST=host postgres
PGPORT=port postgres
PGDATABASE=nama database
```

---

## Migration

```bash
make migrate-up
```

```bash
make migrate-down
```

```bash
make migrate-create name=create_users
```

```bash
make migrate-force version=1
```

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
 ├── Edit User & upload profile image
 ├── Delete User
 └── Logout
```

---

## Security

- JWT Authentication
- Protected Routes
- Password Hashing (bcrypt)
- Image Extension Validation
- File Size Validation

# Catatan

- Password disimpan menggunakan **bcrypt hash**.
- Endpoint `/users` dilindungi menggunakan middleware Authorization dengan JWT Token, setiap token hanya aktif 15 menit.
- Frontend menggunakan **Fetch API** dengan `application/x-www-form-urlencoded`.
- Token Authorization disimpan di **localStorage** setelah login berhasil.

---

# Author

**Dimas Tadeo**