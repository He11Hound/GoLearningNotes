package database_sql

/*
	СТРУКТУРА ROW

	type Row struct {
		err error
		rows *Rows
	}

	Row представляет одну строку результата SQL-запроса.
	Используется для получения единственного результата из запроса.

	Основные поля:

	err error
		- Ошибка выполнения запроса
		- Содержит ошибку если запрос не выполнен
		- Используется для проверки ошибок

	rows *Rows
		- Указатель на объект Rows
		- Используется для доступа к результатам
		- Может быть nil если произошла ошибка

	Основные методы:

	func (r *Row) Scan(dest ...interface{}) error
		- Сканирует строку результата в переменные
		- dest: указатели на переменные для сохранения данных
		- Возвращает ошибку если строка не найдена
		- Автоматически преобразует типы данных
		- Возвращает sql.ErrNoRows если строка не найдена

	Примеры использования:

	// Базовое использование
	var id int
	var name string
	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", 1).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Пользователь не найден")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Пользователь: ID=%d, Name=%s\n", id, name)
	}

	// Использование с проверкой ошибок
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Всего пользователей: %d\n", count)

	// Сканирование в несколько переменных
	var id int
	var name string
	var email string
	var createdAt time.Time
	err = db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", 1).Scan(&id, &name, &email, &createdAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Пользователь: %d - %s (%s) создан %v\n", id, name, email, createdAt)

	// Использование с nullable полями
	var id int
	var name string
	var email sql.NullString
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", 1).Scan(&id, &name, &email)
	if err != nil {
		log.Fatal(err)
	}
	if email.Valid {
		fmt.Printf("Email: %s\n", email.String)
	} else {
		fmt.Println("Email не указан")
	}

	Особенности работы:
	- Row автоматически закрывает Rows после Scan()
	- Scan() должен вызываться только один раз
	- Если строка не найдена, возвращается sql.ErrNoRows
	- Row не требует явного закрытия
	- Соединение автоматически освобождается

	Связь с другими структурами:
	- Row создается из DB через QueryRow()
	- Row создается из Tx через QueryRow()
	- Row создается из Stmt через QueryRow()
	- Row использует Rows для доступа к данным
	- Row автоматически управляет жизненным циклом Rows

	Лучшие практики:
	- Всегда проверяйте ошибки после Scan()
	- Обрабатывайте sql.ErrNoRows отдельно
	- Используйте подходящие типы для Scan()
	- Используйте sql.NullString для nullable полей
	- Проверяйте количество аргументов Scan()

	Паттерны использования:

	// Паттерн 1: Проверка существования
	func userExists(db *sql.DB, id int) (bool, error) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}

	// Паттерн 2: Получение одной записи
	func getUserByID(db *sql.DB, id int) (*User, error) {
		var user User
		err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // пользователь не найден
			}
			return nil, err
		}
		return &user, nil
	}

	// Паттерн 3: Получение агрегатных данных
	func getUserStats(db *sql.DB) (*UserStats, error) {
		var stats UserStats
		err := db.QueryRow(`
			SELECT 
				COUNT(*) as total_users,
				COUNT(CASE WHEN created_at > ? THEN 1 END) as new_users
			FROM users
		`, time.Now().AddDate(0, -1, 0)).Scan(&stats.TotalUsers, &stats.NewUsers)
		if err != nil {
			return nil, err
		}
		return &stats, nil
	}

	// Паттерн 4: Работа с nullable полями
	func getUserWithOptionalEmail(db *sql.DB, id int) (*User, error) {
		var user User
		var email sql.NullString
		err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if email.Valid {
			user.Email = email.String
		}
		return &user, nil
	}

	// Паттерн 5: Получение конфигурации
	func getConfigValue(db *sql.DB, key string) (string, error) {
		var value string
		err := db.QueryRow("SELECT value FROM config WHERE key = ?", key).Scan(&value)
		if err != nil {
			if err == sql.ErrNoRows {
				return "", fmt.Errorf("конфигурация %s не найдена", key)
			}
			return "", err
		}
		return value, nil
	}

	// Паттерн 6: Проверка уникальности
	func isEmailUnique(db *sql.DB, email string) (bool, error) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
		if err != nil {
			return false, err
		}
		return count == 0, nil
	}

	// Паттерн 7: Получение последней записи
	func getLastUser(db *sql.DB) (*User, error) {
		var user User
		err := db.QueryRow("SELECT id, name, email FROM users ORDER BY id DESC LIMIT 1").Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return &user, nil
	}

	// Паттерн 8: Получение случайной записи
	func getRandomUser(db *sql.DB) (*User, error) {
		var user User
		err := db.QueryRow("SELECT id, name, email FROM users ORDER BY RANDOM() LIMIT 1").Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return &user, nil
	}

	// Паттерн 9: Получение с условием
	func getActiveUserByEmail(db *sql.DB, email string) (*User, error) {
		var user User
		err := db.QueryRow("SELECT id, name, email FROM users WHERE email = ? AND status = 'active'", email).Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return &user, nil
	}

	// Паттерн 10: Получение с JOIN
	func getUserWithProfile(db *sql.DB, userID int) (*UserWithProfile, error) {
		var user UserWithProfile
		err := db.QueryRow(`
			SELECT u.id, u.name, u.email, p.bio, p.avatar
			FROM users u
			LEFT JOIN profiles p ON u.id = p.user_id
			WHERE u.id = ?
		`, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Bio, &user.Avatar)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return &user, nil
	}
*/
