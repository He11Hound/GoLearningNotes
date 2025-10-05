package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

/*
Тутс описана работа с транзакциями в GOOOOOOOOOOOO

Для начала транзакции нужно вызвать один из методов, которые возвращают переменную типа *sql.Tx:

Begin указывает на начало транзакции;
Commit сохраняет изменения, которые внесли в транзакцию;
Rollback откатывает изменения. Например, Rollback поможет вернуть

func (db *DB) Begin() (*Tx, error)
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

В методе BeginTx передаётся контекст — транзакция будет отменена при отмене контекста. Параметр opts позволяет настраивать транзакцию, например, можно передать указать уровень изоляции.
Метод (tx *Tx) Commit() error сохраняет все изменения, сделанные в рамках транзакции, а (tx *Tx) Rollback() error отменяет все изменения и откатывает транзакцию.
Для выполнения запросов у типа *sql.Tx есть методы tx.Query...() и tx.Exec...(), аналогичные тем, которые мы использовали.
*/

func main() {
	connStr := "host=postgres user=go_user password=go_password dbname=go_learning_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}

	TransactionExample(db)

	defer db.Close()
}

func TransactionExample(db *sql.DB) {
	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Ошибка начала транзакции:", err)
	}

	// В конце фиксируем или откатываем
	defer func() {
		if err != nil {
			tx.Rollback()
			fmt.Println("Транзакция откатена")
		} else {
			tx.Commit()
			fmt.Println("Транзакция успешно выполнена")
		}
	}()

	// Создаем новый заказ для пользователя с id=1
	var orderID int
	err = tx.QueryRow(
		`INSERT INTO orders (user_id, total_amount, status, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		1, 0.0, "pending", time.Now(), time.Now(),
	).Scan(&orderID)
	if err != nil {
		log.Println("Ошибка при создании заказа:", err)
		return
	}

	// Добавляем товары в order_items
	items := []struct {
		ProductID int
		Quantity  int
		Price     float64
	}{
		{1, 2, 999.99}, // 2 iPhone 15
		{3, 3, 29.99},  // 3 Футболки Nike
	}

	total := 0.0
	for _, item := range items {
		_, err = tx.Exec(
			`INSERT INTO order_items (order_id, product_id, quantity, price, created_at)
			 VALUES ($1, $2, $3, $4, $5)`,
			orderID, item.ProductID, item.Quantity, item.Price, time.Now(),
		)
		if err != nil {
			log.Println("Ошибка при добавлении элемента заказа:", err)
			return
		}
		total += float64(item.Quantity) * item.Price
	}

	// Обновляем total_amount в заказе
	_, err = tx.Exec(
		`UPDATE orders SET total_amount=$1, updated_at=$2 WHERE id=$3`,
		total, time.Now(), orderID,
	)
	if err != nil {
		log.Println("Ошибка при обновлении суммы заказа:", err)
		return
	}

	// Если мы дошли сюда, транзакция выполнена успешно
}
