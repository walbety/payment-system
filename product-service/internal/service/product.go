package service

import "github.com/walbety/payment-system/product-service/internal/canonical"

type ProductService struct{}

func (s *ProductService) ListProductByCategory(category int32) ([]canonical.Product, error) {

	return make([]canonical.Product, 1), nil

}
