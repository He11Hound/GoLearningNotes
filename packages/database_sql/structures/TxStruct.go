package database_sql

/*
	СТРУКТУРА TX

	type Tx struct {
		db *DB
		mu sync.RWMutex
		closed bool
		dc *driverConn
		txi driver.Tx
		releaseConn func(error)
		done bool
		keepConnOnRollback bool
		stmtCache map[string]*Stmt
		closedmu sync.RWMutex
	}

	Tx представляет транзакцию базы данных.
	Все операции в рамках транзакции выполняются атомарно.

	Основные поля:

	db *DB
		- Указатель на объект DB
		- Используется для доступа к пулу соединений
		- Связывает транзакцию с базой данных

	mu sync.RWMutex
		- Мьютекс для обеспечения потокобезопасности
		- Защищает доступ к транзакции
		- Позволяет множественное чтение

	closed bool
		- Флаг закрытия транзакции
		- Предотвращает выполнение операций после закрытия
		- Устанавливается при Commit() или Rollback()

	dc *driverConn
		- Соединение с базой данных
		- Используется для выполнения операций
		- Управляется пулом соединений

	txi driver.Tx
		- Интерфейс транзакции драйвера
		- Предоставляется драйвером базы данных
		- Используется для выполнения операций

	releaseConn func(error)
		- Функция освобождения соединения
		- Вызывается при завершении транзакции
		- Возвращает соединение в пул

	done bool
		- Флаг завершения транзакции
		- Устанавливается при Commit() или Rollback()
		- Предотвращает повторное завершение

	keepConnOnRollback bool
		- Флаг сохранения соединения при откате
		- Используется для оптимизации
		- Зависит от настроек драйвера

	stmtCache map[string]*Stmt
		- Кэш подготовленных запросов
		- Используется для переиспользования Stmt
		- Ключ - текст SQL запроса

	closedmu sync.RWMutex
		- Мьютекс для управления закрытием
		- Защищает операции закрытия
		- Используется для синхронизации

	Основные методы:

	func (tx *Tx) Commit() error
		- Подтверждает транзакцию
		- Сохраняет все изменения в базе данных
		- Освобождает ресурсы транзакции
		- Должен вызываться для завершения транзакции

	func (tx *Tx) Rollback() error
		- Отменяет транзакцию
		- Откатывает все изменения
		- Освобождает ресурсы транзакции
		- Должен вызываться при ошибках

	func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
		- Выполняет запрос в рамках транзакции
		- Аналогично db.Exec но в транзакции
		- Все изменения будут откачены при Rollback()
		- args: параметры для запроса

	func (tx *Tx) Query(query string, args ...interface{}) (*Rows, error)
		- Выполняет запрос в рамках транзакции
		- Аналогично db.Query но в транзакции
		- Результаты видны только в рамках транзакции
		- args: параметры для запроса

	func (tx *Tx) QueryRow(query string, args ...interface{}) *Row
		- Выполняет запрос в рамках транзакции
		- Аналогично db.QueryRow но в транзакции
		- Результат виден только в рамках транзакции
		- args: параметры для запроса

	func (tx *Tx) Prepare(query string) (*Stmt, error)
		- Подготавливает запрос в рамках транзакции
		- Аналогично db.Prepare но в транзакции
		- Stmt привязан к транзакции
		- Автоматически закрывается при завершении транзакции

	func (tx *Tx) Stmt(stmt *Stmt) *Stmt
		- Привязывает существующий Stmt к транзакции
		- Позволяет использовать подготовленные запросы в транзакции
		- Возвращает новый Stmt привязанный к транзакции
		- Полезно для переиспользования подготовленных запросов

	Примеры использования:

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

	_, err = tx.Exec("UPDATE users SET status = ? WHERE name = ?", "active", "Alice")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Подтверждение транзакции
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// Использование с контекстом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Операции в транзакции
	_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Bob")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	Особенности работы:
	- Tx потокобезопасен благодаря мьютексу
	- Все операции в транзакции выполняются атомарно
	- Транзакция должна быть завершена Commit() или Rollback()
	- Соединение блокируется на время транзакции
	- Подготовленные запросы кэшируются в транзакции

	Связь с другими структурами:
	- Tx создается из DB через Begin() или BeginTx()
	- Tx создает и управляет объектами Stmt
	- Tx возвращает объекты Rows и Row
	- Tx использует driverConn для выполнения операций
	- Tx использует driver.Tx интерфейс драйвера

	Лучшие практики:
	- Всегда завершайте транзакцию Commit() или Rollback()
	- Используйте defer tx.Rollback() для автоматического отката
	- Проверяйте ошибки после каждой операции
	- Используйте контексты для отмены длительных транзакций
	- Избегайте длительных транзакций
	- Используйте подходящий уровень изоляции

	Паттерны использования:

	// Паттерн 1: Автоматический откат
	func transferMoney(db *sql.DB, from, to string, amount float64) error {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback() // автоматический откат при ошибке

		// Операции в транзакции
		_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE name = ?", amount, from)
		if err != nil {
			return err
		}

		_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE name = ?", amount, to)
		if err != nil {
			return err
		}

		return tx.Commit()
	}

	// Паттерн 2: Транзакция с контекстом
	func processWithTimeout(db *sql.DB, ctx context.Context) error {
		tx, err := db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})
		if err != nil {
			return err
		}
		defer tx.Rollback()

		// Операции в транзакции
		_, err = tx.Exec("INSERT INTO logs (message) VALUES (?)", "Processing")
		if err != nil {
			return err
		}

		return tx.Commit()
	}
*/
