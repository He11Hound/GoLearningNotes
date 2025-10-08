package database_sql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// OpenExample демонстрирует использование функции sql.Open
func OpenExample() {
	// Подключение к базе данных
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Соединение с базой данных установлено")
}

// OpenDBExample демонстрирует использование функции sql.OpenDB
func OpenDBExample() {
	// Создание коннектора (пример для MySQL)
	// В реальном коде нужно использовать конкретный драйвер
	fmt.Println("OpenDB используется с конкретными драйверами")
	fmt.Println("Пример: sql.OpenDB(mysql.NewConnector(config))")
}

// CloseExample демонстрирует использование функции db.Close
func CloseExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// Закрытие соединения
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Соединение с базой данных закрыто")
}

// PingExample демонстрирует использование функции db.Ping
func PingExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Соединение с базой данных активно")
}

// StatsExample демонстрирует использование функции db.Stats
func StatsExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Получение статистики
	stats := db.Stats()
	fmt.Printf("Максимум открытых соединений: %d\n", stats.MaxOpenConns)
	fmt.Printf("Открытые соединения: %d\n", stats.OpenConns)
	fmt.Printf("Неактивные соединения: %d\n", stats.IdleConns)
	fmt.Printf("Ожидающие соединения: %d\n", stats.WaitCount)
}

// ExecExample демонстрирует использование функции db.Exec
func ExecExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Вставка данных
	result, err := db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
	if err != nil {
		log.Fatal(err)
	}

	// Получение количества затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Вставлено строк: %d\n", rowsAffected)
}

// QueryExample демонстрирует использование функции db.Query
func QueryExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie')")

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Итерация по результатам
	fmt.Println("Пользователи:")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Проверка ошибок итерации
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// QueryRowExample демонстрирует использование функции db.QueryRow
func QueryRowExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice')")

	// Выполнение запроса с одним результатом
	var id int
	var name string
	err = db.QueryRow("SELECT id, name FROM users WHERE name = ?", "Alice").Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Пользователь не найден")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Найден пользователь: ID=%d, Name=%s\n", id, name)
	}
}

// PrepareExample демонстрирует использование функции db.Prepare
func PrepareExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

	// Подготовка запроса
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Выполнение подготовленного запроса несколько раз
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		_, err = stmt.Exec(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Вставлен пользователь: %s\n", name)
	}
}

// BeginExample демонстрирует использование функции db.Begin
func BeginExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("CREATE TABLE orders (id INTEGER PRIMARY KEY, user_id INTEGER, amount REAL)")

	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение операций в транзакции
	_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "John")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec("INSERT INTO orders (user_id, amount) VALUES (?, ?)", 1, 100.50)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Подтверждение транзакции
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Транзакция успешно выполнена")
}

// BeginTxExample демонстрирует использование функции db.BeginTx
func BeginTxExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Начало транзакции с контекстом и опциями
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение операции в транзакции
	_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Подтверждение транзакции
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Транзакция с контекстом успешно выполнена")
}

// SetConnMaxLifetimeExample демонстрирует использование функции db.SetConnMaxLifetime
func SetConnMaxLifetimeExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Установка максимального времени жизни соединения
	db.SetConnMaxLifetime(time.Hour)

	fmt.Println("Максимальное время жизни соединения установлено на 1 час")
}

// SetMaxOpenConnsExample демонстрирует использование функции db.SetMaxOpenConns
func SetMaxOpenConnsExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Установка максимального количества открытых соединений
	db.SetMaxOpenConns(25)

	fmt.Println("Максимальное количество открытых соединений установлено на 25")
}

// SetMaxIdleConnsExample демонстрирует использование функции db.SetMaxIdleConns
func SetMaxIdleConnsExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Установка максимального количества неактивных соединений
	db.SetMaxIdleConns(5)

	fmt.Println("Максимальное количество неактивных соединений установлено на 5")
}

// LastInsertIdExample демонстрирует использование функции result.LastInsertId
func LastInsertIdExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы с автоинкрементом
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")

	// Вставка данных
	result, err := db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
	if err != nil {
		log.Fatal(err)
	}

	// Получение ID последней вставленной записи
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID последней вставленной записи: %d\n", lastID)
}

// RowsAffectedExample демонстрирует использование функции result.RowsAffected
func RowsAffectedExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie')")

	// Обновление данных
	result, err := db.Exec("UPDATE users SET name = ? WHERE name = ?", "Alice Updated", "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// Получение количества затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Обновлено строк: %d\n", rowsAffected)
}

// ScanExample демонстрирует использование функции row.Scan
func ScanExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	db.Exec("INSERT INTO users (name, age) VALUES ('Alice', 25)")

	// Выполнение запроса с одним результатом
	var id int
	var name string
	var age int
	err = db.QueryRow("SELECT id, name, age FROM users WHERE name = ?", "Alice").Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Пользователь: ID=%d, Name=%s, Age=%d\n", id, name, age)
}

// RowsScanExample демонстрирует использование функции rows.Scan
func RowsScanExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	db.Exec("INSERT INTO users (name, age) VALUES ('Alice', 25), ('Bob', 30), ('Charlie', 35)")

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Итерация по результатам
	fmt.Println("Все пользователи:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID=%d, Name=%s, Age=%d\n", id, name, age)
	}

	// Проверка ошибок итерации
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// NextExample демонстрирует использование функции rows.Next
func NextExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie')")

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Итерация по результатам с использованием Next
	fmt.Println("Пользователи:")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID=%d, Name=%s\n", id, name)
	}

	fmt.Println("Итерация завершена")
}

// ColumnsExample демонстрирует использование функции rows.Columns
func ColumnsExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	db.Exec("INSERT INTO users (name, email) VALUES ('Alice', 'alice@example.com')")

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Получение имен колонок
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Колонки: %v\n", columns)
	fmt.Printf("Количество колонок: %d\n", len(columns))
}

// CommitExample демонстрирует использование функции tx.Commit
func CommitExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение операций в транзакции
	_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Подтверждение транзакции
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Транзакция успешно подтверждена")
}

// RollbackExample демонстрирует использование функции tx.Rollback
func RollbackExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение операций в транзакции
	_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Имитация ошибки и откат транзакции
	fmt.Println("Произошла ошибка, откатываем транзакцию")
	err = tx.Rollback()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Транзакция успешно откачена")
}

// StmtExecExample демонстрирует использование функции stmt.Exec
func StmtExecExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

	// Подготовка запроса
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Выполнение подготовленного запроса
	result, err := stmt.Exec("Alice")
	if err != nil {
		log.Fatal(err)
	}

	// Получение количества затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Вставлено строк через подготовленный запрос: %d\n", rowsAffected)
}

// StmtQueryExample демонстрирует использование функции stmt.Query
func StmtQueryExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob')")

	// Подготовка запроса
	stmt, err := db.Prepare("SELECT id, name FROM users WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Выполнение подготовленного запроса
	rows, err := stmt.Query("Alice")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Итерация по результатам
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Найден пользователь: ID=%d, Name=%s\n", id, name)
	}
}

// StmtQueryRowExample демонстрирует использование функции stmt.QueryRow
func StmtQueryRowExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы и вставка данных
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice')")

	// Подготовка запроса
	stmt, err := db.Prepare("SELECT id, name FROM users WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Выполнение подготовленного запроса с одним результатом
	var id int
	var name string
	err = stmt.QueryRow("Alice").Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Найден пользователь: ID=%d, Name=%s\n", id, name)
}

// StmtCloseExample демонстрирует использование функции stmt.Close
func StmtCloseExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")

	// Подготовка запроса
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса
	_, err = stmt.Exec("Alice")
	if err != nil {
		log.Fatal(err)
	}

	// Закрытие подготовленного запроса
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Подготовленный запрос успешно закрыт")
}

// ComplexTransactionExample демонстрирует сложную транзакцию
func ComplexTransactionExample() {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблиц
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, balance REAL)")
	db.Exec("CREATE TABLE transactions (id INTEGER PRIMARY KEY, from_user INTEGER, to_user INTEGER, amount REAL)")

	// Вставка начальных данных
	db.Exec("INSERT INTO users (name, balance) VALUES ('Alice', 1000.0), ('Bob', 500.0)")

	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Перевод денег от Alice к Bob
	amount := 100.0

	// Проверка баланса Alice
	var aliceBalance float64
	err = tx.QueryRow("SELECT balance FROM users WHERE name = ?", "Alice").Scan(&aliceBalance)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if aliceBalance < amount {
		tx.Rollback()
		fmt.Println("Недостаточно средств для перевода")
		return
	}

	// Обновление баланса Alice
	_, err = tx.Exec("UPDATE users SET balance = balance - ? WHERE name = ?", amount, "Alice")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Обновление баланса Bob
	_, err = tx.Exec("UPDATE users SET balance = balance + ? WHERE name = ?", amount, "Bob")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Запись транзакции
	_, err = tx.Exec("INSERT INTO transactions (from_user, to_user, amount) VALUES (?, ?, ?)", 1, 2, amount)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Подтверждение транзакции
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Перевод %f от Alice к Bob успешно выполнен\n", amount)
}
