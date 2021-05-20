package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rohaquinlop/Trainin-exercise/backend/graph/generated"
	"github.com/rohaquinlop/Trainin-exercise/backend/graph/model"
)

func (r *mutationResolver) CreateBuyer(ctx context.Context, input model.NewBuyer) (*model.Buyer, error) {
	buyer := &model.Buyer{
		ID:   input.ID,
		Name: input.Name,
		Age:  input.Age,
	}
	r.buyers = append(r.buyers, buyer)
	return buyer, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	product := &model.Product{
		ID:    input.ID,
		Name:  input.Name,
		Price: input.Price,
	}
	r.products = append(r.products, product)
	return product, nil
}

func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.NewTransaction) (*model.Transaction, error) {
	transaction := &model.Transaction{
		ID:         input.ID,
		BuyerID:    input.BuyerID,
		IP:         input.IP,
		Device:     input.Device,
		ProductIds: input.ProductIds,
	}
	r.transactions = append(r.transactions, transaction)
	return transaction, nil
}

func (r *queryResolver) Buyers(ctx context.Context) ([]*model.Buyer, error) {
	return r.buyers, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return r.products, nil
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	return r.transactions, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
