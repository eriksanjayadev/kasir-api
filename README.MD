# Kasir API

Kasir API adalah RESTful API sederhana untuk aplikasi kasir, dibangun menggunakan **Go (Golang)** dengan pendekatan **Clean Architecture**. API ini menyediakan manajemen **Product** dan **Category**, termasuk relasi antara keduanya.

---

## 🚀 Tech Stack

* Go (net/http)
* PostgreSQL / Supabase (database)
* Clean Architecture (Handler → Service → Repository)
* Docker (opsional, untuk deployment)

---

## 📦 Base URL

```
http://localhost:8080/api
```

---

## 📂 Struktur Project (Ringkas)

```
kasir-api/
├── handlers/      # HTTP handlers (request & response)
├── services/      # Business logic
├── repositories/  # Database access
├── models/        # Data models & DTO
├── main.go
└── go.mod
```

---

## 🗂️ Category Endpoints

### 1️⃣ Get All Categories

**GET** `/api/categories`

**Response 200**

```json
[
  {
    "id": 1,
    "name": "Food",
    "description": "Food category"
  }
]
```

---

### 2️⃣ Create Category

**POST** `/api/categories`

**Request Body**

```json
{
  "name": "Drink",
  "description": "Beverages"
}
```

**Response 201**

```json
{
  "id": 2,
  "name": "Drink",
  "description": "Beverages"
}
```

---

### 3️⃣ Get Category By ID

**GET** `/api/categories/{id}`

```json
{
  "id": 5,
  "name": "Drink",
  "description": "Kategori Drink"
}
```

---

### 4️⃣ Update Category

**PUT** `/api/categories/{id}`

**Request Body**

```json
{
  "name": "Updated Category",
  "description": "Updated description"
}
```

---

### 5️⃣ Delete Category

**DELETE** `/api/categories/{id}`

**Response 200**

```json
{
  "message": "Category deleted successfully"
}
```

---

## 📦 Product Endpoints

### 1️⃣ Get All Products

**GET** `/api/produk`

**Response 200**

```json
[
  {
    "id": 1,
    "name": "Nano",
    "price": 1200,
    "stock": 45,
    "category_id": 1,
    "category_name": "Meal"
  }
]
```

---

### 2️⃣ Create Product

**POST** `/api/produk`

**Request Body**

```json
{
  "name": "Keyboard",
  "price": 150000,
  "stock": 10,
  "category_id": 1
}
```

**Response 201**

```json
{
  "id": 2,
  "name": "Keyboard",
  "price": 150000,
  "stock": 10,
  "category_id": 1
}
```

---

### 3️⃣ Get Product By ID

**GET** `/api/produk/{id}`
```json
{
  "id": 4,
  "name": "Bakso",
  "price": 12000,
  "stock": 11,
  "category_id": 1,
  "category_name": "Meal"
}
```
---

### 4️⃣ Update Product

**PUT** `/api/produk/{id}`

**Request Body**

```json
{
  "name": "Updated Product",
  "price": 200000,
  "stock": 5,
  "category_id": 2
}
```

---

### 5️⃣ Delete Product

**DELETE** `/api/produk/{id}`

**Response 200**

```json
{
  "message": "Product deleted successfully"
}
```
