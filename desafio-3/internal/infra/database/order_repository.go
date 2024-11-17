package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) Retrieve(entityInput entity.Order) (entity.Order, error) {

	query := "Select id, price, tax, final_price from orders where id = ?"
	order := entity.Order{}

	err := r.Db.QueryRow(query, entityInput.ID).Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice) // Busca a linha e mapeia os dados
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum resultado encontrado.")
		} else {
			log.Fatalf("Erro ao executar a query: %v", err)
		}
		return entity.Order{}, err
	}
	return order, nil

}
