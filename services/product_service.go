package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]models.ProductResponse, error) {
	products, err := s.repo.GetAll(name)
	if err != nil {
		return nil, err
	}

	result := make([]models.ProductResponse, 0, len(products))
	for _, p := range products {
		result = append(result, mapToProductResponse(p))
	}

	return result, nil
}

func (s *ProductService) GetById(id int) (*models.ProductResponse, error) {
	p, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	result := mapToProductResponse(*p)
	return &result, nil
}

func (s *ProductService) Create(req models.ProductCreateRequest) (*models.ProductResponse, error) {
	product := models.Product{
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	if err := s.repo.Create(&product); err != nil {
		return nil, err
	}

	p, err := s.repo.GetById(product.ID)
	if err != nil {
		return nil, err
	}

	result := mapToProductResponse(*p)
	return &result, nil
}

func (s *ProductService) Update(id int, req models.ProductUpdateRequest) error {
	product := models.Product{
		ID:         id,
		Name:       req.Name,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	return s.repo.Update(&product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}

func mapToProductResponse(p models.ProductWithCategory) models.ProductResponse {
	return models.ProductResponse{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
		Category: models.CategoryResponse{
			ID:   p.CategoryID,
			Name: p.CategoryName,
		},
	}
}
