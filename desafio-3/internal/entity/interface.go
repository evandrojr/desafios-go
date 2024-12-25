package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
	Retrieve(orderId Order) (Order, error)
	List() ([]Order, error)
}
