package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

/*
Изучение пакета databse/sql

Для подключения к бд мы обязательно должны инициализировать (init функция нужного пакета) драйвер бд которой используем, для этого просто в import передаём через _ название пакета с драйвером

Само соединение делается через функцию sql.Open("driverName", "connectionString")
	driverName - название драйвера, указан в init функции
	connectionString - строка с параметрами подключения

	НЕ устанавливает реальное соединение с БД! Она только:
		Регистрирует драйвер (если не зарегистрирован)
		Создает структуру sql.DB с настройками по умолчанию
		Возвращает "ленивое" соединение - реальное подключение произойдет при первом запросе

sql.DB - это структура, которая представляет пул соединений с базой данных

db.Close() закрывает все соединения в пуле:

db.Query() возвращает *sql.Rows, error - структуру для итерации по результатам: Используется когда нужно получить несколько записей

db.QueryRow() возвращает один Row, без ошибки, к которому сразу применям Scan, и уже тут если ответ всё таки был пустым, то формируем ошибку

db.Exec() для Insert запросов, возвращает объект типа sql.Result - позволяет узнать, сколько строк было добавлено, удалено или изменено, и получить последний добавленный идентификатор при вставке новых строк.

sql.Result.LastInsertId() - возвращает последний добавочный ID, Postgres - не поддерживается
sql.Result.RowsAffected() - возвращает количество вставленных записей, Postgres - не поддерживается

rows.Close() обязательно нужно вызывать, потому что:
	Освобождает соединение обратно в пул
	Закрывает курсор в базе данных
	Предотвращает утечки памяти
	Освобождает ресурсы драйвера

rows.Next() перемещает курсор к следующей строке:
rows.Scan() копирует значения текущей строки в переменные:


Работа с NULL-значениями
В SQL есть специальное значение NULL, которое может быть записано в поле любого типа. Оно означает, что данное поле не содержит значения. Например, NULL нужен, чтобы отличить пустую строку в текстовом поле от отсутствия значения (значение не определено).  Чтобы отслеживать такие ситуации, пакет database/sql содержит типы, начинающиеся с Null: sql.NullString, sql.NullInt64, sql.NullFloat64, sql.NullByte и так далее.
*/

/*
Параметры запроса

Так как мы не пишем все запросы руками, а используем переменные, то появляется вопрос как вставлять переменные в запрос
Есть 3 способа
	Позиционный:
		В SQL-запросе там, куда нужно подставить параметр, указывается знак вопроса ? - rows, err := db.Query("SELECT product FROM products WHERE price > ?", 500)
	Нумерованный:
		Вместо ? используется знак $ после которого идёт число, Индексация начинается с 1
		rows, err := db.Query("SELECT product FROM products WHERE price > $1", 500)
	Именованные. В запросе указывается имя параметра, перед именем идёт знак : !!!!!!!!!!!! 🔍 POSTGRES - НЕ ПОДДЕРЖИВАЕТ  🔍!!!!!!!!!
		В качестве значения передаётся объект типа sql.NamedArg. Чтобы его получить, нужно вызвать функцию Named() из пакета database/sql. В качестве аргументов этой функции нужно передать имя параметра и значение:
		rows, err := db.Query("SELECT product FROM products WHERE price > :price", sql.Named("price", 500))
*/

func main() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}

	defer db.Close()

	//SimpleDeleteQuery(db)
	//SimpleUpdateQuery(db)
	//SimpleInsertQuery(db)

	//order := SimpleSelectWithOneRoew(db)
	//fmt.Println(order)

	//orders, err := SimpleSelectQuery(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if len(orders) > 0 {
	//	for _, order := range orders {
	//		fmt.Println(order)
	//	}
	//}

}

type Order struct { //Тестовая струтура заказов
	ID          int
	UserID      int
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ConnectToDB() (*sql.DB, error) {
	connStr := "host=postgres user=go_user password=go_password dbname=go_learning_db sslmode=disable"
	return sql.Open("postgres", connStr) //возвращает структуру для пула Соединений с БД
}

func SimpleSelectQuery(db *sql.DB) ([]Order, error) { //Работа со всеми полями
	var orders []Order

	rows, err := db.Query("SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders") //Выполняем запрос, получаем структуру для итерации по результатам

	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return orders, errors.New("Ошибка при выполнении запроса")
	}

	defer rows.Close() //возвращает соединение обратно в пул

	for rows.Next() { //проходится по результатам
		var o Order
		err := rows.Scan(&o.ID, &o.UserID, &o.TotalAmount, &o.Status, &o.CreatedAt, &o.UpdatedAt) //заполняет переданную структуру, согласно очерёдности данных я так полагаю
		if err != nil {
			log.Println("Ошибка при чтении строки:", err)
			continue
		}
		orders = append(orders, o) // добавляем в слайс
	}

	if err = rows.Err(); err != nil { //Проверка на ошибки
		log.Println("Ошибка при итерации по строкам:", err)
	}

	return orders, nil
}

func SimpleSelectWithOneRow(db *sql.DB) Order {
	var order Order

	row := db.QueryRow("SELECT id, user_id, total_amount, status, created_at, updated_at FROM orders LIMIT 1")

	err := row.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		log.Println("Ошибка при чтении строки:", err)
	}

	return order
}

func SimpleInsertQuery(db *sql.DB) {
	now := time.Now()

	res, err := db.Exec(
		`INSERT INTO orders (user_id, total_amount, status, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		2,
		1200.25,
		"pending",
		now,
		now,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
	return
}

func SimpleUpdateQuery(db *sql.DB) {
	res, err := db.Exec(
		`UPDATE orders SET total_amount = $1 WHERE id = $2`,
		1555.55,
		4,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	return
}

func SimpleDeleteQuery(db *sql.DB) {
	_, err := db.Exec(
		`DELETE FROM orders WHERE id = $1`,
		11,
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}
