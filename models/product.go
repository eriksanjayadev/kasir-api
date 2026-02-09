package models

// =====================
// DATABASE MODELS
// =====================

// Product adalah entity database (repo only)
type Product struct {
	ID         int
	Name       string
	Price      int
	Stock      int
	CategoryID int
}

// ProductWithCategory digunakan untuk hasil JOIN (repo → service)
type ProductWithCategory struct {
	ID           int
	Name         string
	Price        int
	Stock        int
	CategoryID   int
	CategoryName string
}

// =====================
// REQUEST MODELS
// =====================

// ProductCreateRequest digunakan untuk POST /products
type ProductCreateRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

// ProductUpdateRequest digunakan untuk PUT /products/{id}
type ProductUpdateRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

// =====================
// RESPONSE MODELS
// =====================

// ProductResponse adalah response API dengan category nested
type ProductResponse struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Price    int              `json:"price"`
	Stock    int              `json:"stock"`
	Category CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
