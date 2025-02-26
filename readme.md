# Intitek Project - Setup & API Documentation

## 📌 Table of Contents

- [Clone Repository](#1-clone-repository)
- [Backend Setup](#2-menjalankan-project-backend)
- [Frontend Setup](#3-menjalankan-project-frontend)
- [API Documentation](#4-api-documentation)
  - [Authentication](#authentication)
  - [Product Management](#product-management)

---

## 1️⃣ Clone Repository

```sh
git clone https://github.com/anzalass/test-intitek
cd test-intitek
```

---

## 2️⃣ Menjalankan Project Backend

### 🔹 a. Masuk ke directory backend

```sh
cd intitek-be
```

### 🔹 b. Konfigurasi Environment

- Ubah nama file `.env.example` menjadi `.env`
- Sesuaikan konfigurasi dalam file `.env` sesuai kebutuhan

### 🔹 c. Install Dependencies

```sh
go mod tidy
```

### 🔹 d. Jalankan Backend

```sh
go run main.go
```

---

## 3️⃣ Menjalankan Project Frontend

### 🔹 a. Masuk ke directory frontend

```sh
cd ../intitek-fe
```

### 🔹 b. Konfigurasi Environment

- Ubah nama file `.env.example` menjadi `.env`
- Sesuaikan konfigurasi dalam file `.env` sesuai kebutuhan

### 🔹 c. Install Dependencies

```sh
npm install
```

### 🔹 d. Jalankan Frontend

```sh
npm run dev
```

---

## 4️⃣ API Documentation

### 🔑 Authentication

#### 🔹 Login

**Endpoint:** `POST /login`

**Request Body:**

```json
{
  "username": "admin12345",
  "password": "admin12345"
}
```

**Response:**

```json
{
  "success": true,
  "user": {
    "ID": 4,
    "Username": "anzalas100",
    "Password": "anzalas100"
  }
}
```

#### 🔹 Register

**Endpoint:** `POST /register`

**Request Body:**

```json
{
  "username": "admin12345",
  "password": "admin12345"
}
```

**Response:**

```json
{
  "data": {
    "ID": 2,
    "Username": "anzalas111",
    "Password": "anzalas111"
  },
  "success": true
}
```

---

### 📦 Product Management

#### 🔹 Create Product

**Endpoint:** `POST /product`

**Request Body:**

```json
{
  "sku": "abc123",
  "name": "Samsung",
  "quantity": 10,
  "location": "Tangerang",
  "status": "Out of Stock"
}
```

**Response:**

```json
{
  "data": {
    "name": "Samsung",
    "sku": "abc123",
    "quantity": 10,
    "location": "Tangerang",
    "status": "Out of Stock"
  },
  "success": true
}
```

#### 🔹 Update Product

**Endpoint:** `PUT /product/:sku`

**Request Body:**

```json
{
  "sku": "abc123",
  "name": "Samsung",
  "quantity": 10,
  "location": "Tangerang",
  "status": "Out of Stock"
}
```

**Response:**

```json
{
  "data": {
    "name": "Samsung",
    "sku": "abc124",
    "quantity": 10,
    "location": "Tangerang",
    "status": "Out of Stock"
  },
  "success": true
}
```

#### 🔹 Get Product by SKU

**Endpoint:** `GET /product/:sku`

**Response:**

```json
{
  "data": {
    "name": "Buku Tulis",
    "sku": "abc123a",
    "quantity": 122,
    "location": "Pasar Kemis",
    "status": "Available"
  },
  "success": true
}
```

#### 🔹 Get All Products (with Filtering & Pagination)

**Endpoint:** `GET /products?page=1&pageSize=10&low_stock=&status=`

**Response:**

```json
{
  "data": [
    {
      "name": "Samsung",
      "sku": "abc124",
      "quantity": 10,
      "location": "Tangerang",
      "status": "Out of Stock"
    },
    {
      "name": "Buku",
      "sku": "que1",
      "quantity": 100,
      "location": "Jakarta",
      "status": "Out of Stock"
    }
  ],
  "page": 1,
  "pageSize": 10,
  "success": true,
  "total": 3
}
```

#### 🔹 Delete Product

**Endpoint:** `DELETE /product/:sku`

**Response:**

```json
{
  "data": {
    "name": "Samsung",
    "sku": "abc12345",
    "quantity": 10,
    "location": "Tangerang",
    "status": "Out of Stock"
  },
  "success": true
}
```

---

## 🎯 Notes

- Pastikan backend berjalan di `http://localhost:8000`
- Pastikan frontend berjalan di `http://localhost:5173`
- Simpan token autentikasi di `localStorage` untuk mengakses halaman yang membutuhkan login

🚀 **Selamat menggunakan aplikasi!**
