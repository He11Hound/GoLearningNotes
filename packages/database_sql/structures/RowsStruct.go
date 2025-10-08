package database_sql

/*
	СТРУКТУРА ROWS

	type Rows struct {
		dc          *driverConn
		releaseConn func(error)
		rowsi       driver.Rows
		closed      bool
		lastcols    []driver.Value
		lasterr     error
		closeStmt   *Stmt
	}

	Rows представляет результат выполнения SQL-запроса.
	Используется для итерации по множественным строкам результата.

	Основные поля:

	dc *driverConn
		- Соединение с базой данных
		- Используется для выполнения операций
		- Управляется пулом соединений

	releaseConn func(error)
		- Функция освобождения соединения
		- Вызывается при закрытии Rows
		- Возвращает соединение в пул

	rowsi driver.Rows
		- Интерфейс Rows драйвера
		- Предоставляется драйвером базы данных
		- Используется для доступа к результатам

	closed bool
		- Флаг закрытия Rows
		- Предотвращает выполнение операций после закрытия
		- Устанавливается при Close()

	lastcols []driver.Value
		- Кэш последних прочитанных колонок
		- Используется для оптимизации
		- Содержит значения колонок

	lasterr error
		- Последняя ошибка итерации
		- Используется для проверки ошибок
		- Проверяется через Err()

	closeStmt *Stmt
		- Связанный подготовленный запрос
		- Используется для автоматического закрытия
		- Может быть nil для обычных запросов

	Основные методы:

	func (rs *Rows) Next() bool
		- Переходит к следующей строке результата
		- Возвращает false если строк больше нет
		- Должен вызываться перед Scan()
		- Автоматически закрывает Rows при завершении

	func (rs *Rows) Scan(dest ...interface{}) error
		- Сканирует текущую строку результата
		- Используется в цикле с Next()
		- dest: указатели на переменные
		- Автоматически преобразует типы данных

	func (rs *Rows) Close() error
		- Закрывает Rows и освобождает ресурсы
		- Должен вызываться после завершения работы
		- Автоматически вызывается при достижении конца
		- Освобождает соединение

	func (rs *Rows) Columns() ([]string, error)
		- Возвращает имена колонок результата
		- Полезно для динамической работы с данными
		- Возвращает ошибку если Rows закрыт
		- Используется для получения метаданных

	func (rs *Rows) ColumnTypes() ([]*ColumnType, error)
		- Возвращает информацию о типах колонок
		- Содержит детальную информацию о типах данных
		- Полезно для динамической обработки данных
		- Возвращает ошибку если Rows закрыт

	func (rs *Rows) Err() error
		- Возвращает последнюю ошибку итерации
		- Проверяется после завершения цикла Next()
		- Может содержать ошибки чтения данных
		- Должна проверяться для корректной обработки

	Примеры использования:

	// Базовое использование
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Получение информации о колонках
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Колонки: %v\n", columns)

	// Динамическое сканирование
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal(err)
		}
		for i, col := range columns {
			fmt.Printf("%s: %v\n", col, values[i])
		}
	}

	Особенности работы:
	- Rows должен быть закрыт с помощью Close()
	- Next() автоматически закрывает Rows при завершении
	- Scan() должен вызываться после Next()
	- Err() должна проверяться после завершения цикла
	- Соединение блокируется на время работы с Rows

	Связь с другими структурами:
	- Rows создается из DB через Query()
	- Rows создается из Tx через Query()
	- Rows создается из Stmt через Query()
	- Rows использует driverConn для доступа к данным
	- Rows использует driver.Rows интерфейс драйвера

	Лучшие практики:
	- Всегда используйте defer rows.Close()
	- Проверяйте ошибки после каждой операции
	- Используйте Err() для проверки ошибок итерации
	- Используйте Columns() для динамической работы
	- Избегайте длительного удержания Rows
	- Используйте подходящие типы для Scan()

	Паттерны использования:

	// Паттерн 1: Базовое сканирование
	func scanUsers(db *sql.DB) error {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				return err
			}
			fmt.Printf("User: %d - %s\n", id, name)
		}

		return rows.Err()
	}

	// Паттерн 2: Динамическое сканирование
	func scanDynamic(db *sql.DB, query string) error {
		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return err
		}

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		for rows.Next() {
			if err := rows.Scan(valuePtrs...); err != nil {
				return err
			}
			for i, col := range columns {
				fmt.Printf("%s: %v\n", col, values[i])
			}
		}

		return rows.Err()
	}

	// Паттерн 3: Сканирование в структуру
	type User struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	func scanToStruct(db *sql.DB) ([]User, error) {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name); err != nil {
				return nil, err
			}
			users = append(users, user)
		}

		return users, rows.Err()
	}
*/
