package storage

import (
	"github.com/sonm-io/marketplace/infra/storage/inmemory"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// OrderStorage stores and retrieves Orders (Read side).
type OrderStorage struct {
	e Engine
}

// NewOrderStorage creates an new instance of OrderStorage.
func NewOrderStorage(e Engine) *OrderStorage {
	return &OrderStorage{
		e: e,
	}
}

// Adds the given Order to the storage.
func (s *OrderStorage) Add(o *ds.Order) error {
	return s.e.Add(o, o.ID)
}

// Remove removes an Order with the given ID from OrderStorage.
// If no orders found, an error is returned.
func (s *OrderStorage) Remove(ID string) error {
	return s.e.Remove(ID)
}

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderStorage) ByID(ID string) (ds.Order, error) {

	el, err := s.e.Get(ID)
	if err != nil {
		return ds.Order{}, err
	}
	order := el.(*ds.Order)
	return *order, nil
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) ([]ds.Order, error) {

	b := inmemory.NewBuilder()
	b.WithLimit(limit)
	b.WithSpec(spec)

	elements, err := s.e.Match(b.Build())
	if err != nil {
		return nil, err
	}

	var orders []ds.Order
	for _, el := range elements {
		order := el.(*ds.Order)
		orders = append(orders, *order)
	}

	return orders, nil
}