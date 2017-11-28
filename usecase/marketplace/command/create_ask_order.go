package command

// CreateOrder is a command to create an ask order.
type CreateAskOrder struct {
	// Order ID, UUIDv4
	ID string
	// Supplier's Ethereum ID
	SupplierID string
	// Order price
	Price int64
	// Slot a slot
	Slot Slot
}

// CommandID implements Command interface.
func (c CreateAskOrder) CommandID() string {
	return "CreateOrder"
}