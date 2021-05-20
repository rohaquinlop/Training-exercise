package graph

import "github.com/rohaquinlop/Trainin-exercise/backend/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	buyers       []*model.Buyer
	products     []*model.Product
	transactions []*model.Transaction
}
