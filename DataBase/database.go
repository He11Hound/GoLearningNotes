package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

//Изучение пакета databse/sql
/*
Для подключения к бд мы обязательно должны инициализировать (init функция нужного пакета) драйвер бд которой используем, для этого просто в import передаём через _ название пакета с драйвером
Само соединение делается через функцию sql.Open("driverName", "connectionString")
	driverName - название драйвера, указан в init функции
	connectionString - строка с параметрами подключения

	Возвращает нам ссылку на структуру пакета sql.DB и ошибку
*/

func main() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}

	defer db.Close()

	orders, err := SimpleSelectQuery(db)

	if err != nil {
		log.Fatal(err)
	}

	if len(orders) > 0 {
		for _, order := range orders {
			fmt.Println(order)
		}
	}

}

type Order struct {
	ID          int
	UserID      int
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ConnectToDB() (*sql.DB, error) {
	connStr := "host=postgres user=go_user password=go_password dbname=go_learning_db sslmode=disable"
	return sql.Open("postgres", connStr)
}

func SimpleSelectQuery(db *sql.DB) ([]Order, error) {
	var orders []Order

	rows, err := db.Query("SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders")

	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return orders, errors.New("Ошибка при выполнении запроса")
	}

	defer rows.Close()

	for rows.Next() {
		var o Order
		err := rows.Scan(&o.ID, &o.UserID, &o.TotalAmount, &o.Status, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			log.Println("Ошибка при чтении строки:", err)
			continue
		}
		orders = append(orders, o) // добавляем в слайс
	}

	if err = rows.Err(); err != nil {
		log.Println("Ошибка при итерации по строкам:", err)
	}

	return orders, nil
}

func SimpleInsertQuery(db *sql.DB) {

}

func SimpleDeleteQuery(db *sql.DB) {

}
