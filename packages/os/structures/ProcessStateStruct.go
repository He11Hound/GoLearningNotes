package structures

/*
	Структура ProcessState

	type ProcessState struct {
		pid    int             // идентификатор процесса
		status syscall.WaitStatus // статус завершения процесса
		rusage *syscall.Rusage   // информация об использовании ресурсов
	}

	ProcessState представляет состояние завершенного процесса.
	Содержит информацию о том, как завершился процесс, код выхода,
	время выполнения и использование ресурсов.

	Основные методы:

	func (p *ProcessState) ExitCode() int
		- Возвращает код выхода процесса
		- 0 обычно означает успешное завершение
		- Не-0 значения указывают на ошибку или особое состояние
		- На Unix: младшие 8 бит статуса
		- На Windows: код выхода из ExitProcess

	func (p *ProcessState) Exited() bool
		- Возвращает true если процесс завершился нормально
		- Процесс завершился через exit() или return из main()
		- Не был завершен сигналом или другим способом

	func (p *ProcessState) Pid() int
		- Возвращает идентификатор процесса (PID)
		- Уникальный номер процесса в системе
		- Используется для идентификации процесса

	func (p *ProcessState) String() string
		- Возвращает строковое представление состояния процесса
		- Формат зависит от платформы
		- Примеры: "exit status 0", "signal: killed"

	func (p *ProcessState) Success() bool
		- Возвращает true если процесс завершился успешно
		- Эквивалентно p.ExitCode() == 0
		- Удобный способ проверки успешности выполнения

	func (p *ProcessState) Sys() interface{}
		- Возвращает системно-зависимую информацию
		- На Unix: syscall.WaitStatus
		- На Windows: syscall.ProcessInformation
		- Содержит низкоуровневые детали завершения

	func (p *ProcessState) SysUsage() interface{}
		- Возвращает информацию об использовании ресурсов
		- На Unix: *syscall.Rusage
		- На Windows: *syscall.ProcessMemoryCounters
		- Содержит данные о CPU времени, памяти и т.д.

	func (p *ProcessState) UserTime() time.Duration
		- Возвращает время CPU, потраченное в пользовательском режиме
		- Доступно только на Unix системах
		- На Windows возвращает 0

	func (p *ProcessState) SystemTime() time.Duration
		- Возвращает время CPU, потраченное в системном режиме
		- Доступно только на Unix системах
		- На Windows возвращает 0

	Примеры использования:

	// Запуск и ожидание процесса
	cmd := exec.Command("ls", "-la")
	err := cmd.Run()

	// Получение состояния процесса
	state := cmd.ProcessState
	if state == nil {
		fmt.Println("Процесс еще не завершился")
		return
	}

	// Проверка успешности выполнения
	if state.Success() {
		fmt.Println("Команда выполнена успешно")
	} else {
		fmt.Printf("Команда завершилась с ошибкой, код: %d\n", state.ExitCode())
	}

	// Получение PID
	fmt.Printf("PID процесса: %d\n", state.Pid())

	// Проверка способа завершения
	if state.Exited() {
		fmt.Println("Процесс завершился нормально")
	} else {
		fmt.Println("Процесс был завершен сигналом или другим способом")
	}

	// Получение времени выполнения (Unix)
	userTime := state.UserTime()
	systemTime := state.SystemTime()
	totalTime := userTime + systemTime

	fmt.Printf("Время выполнения:\n")
	fmt.Printf("  Пользовательское время: %v\n", userTime)
	fmt.Printf("  Системное время: %v\n", systemTime)
	fmt.Printf("  Общее время CPU: %v\n", totalTime)

	// Получение информации об использовании ресурсов (Unix)
	if rusage := state.SysUsage(); rusage != nil {
		if unixRusage, ok := rusage.(*syscall.Rusage); ok {
			fmt.Printf("Максимальное использование памяти: %d KB\n", unixRusage.Maxrss)
			fmt.Printf("Количество страничных ошибок: %d\n", unixRusage.Minflt)
			fmt.Printf("Количество обращений к диску: %d\n", unixRusage.Inblock)
		}
	}

	// Строковое представление
	fmt.Printf("Состояние процесса: %s\n", state.String())

	// Системно-зависимая информация
	sysInfo := state.Sys()
	fmt.Printf("Системная информация: %+v\n", sysInfo)

	// Обработка различных кодов выхода
	switch state.ExitCode() {
	case 0:
		fmt.Println("Успешное завершение")
	case 1:
		fmt.Println("Общая ошибка")
	case 2:
		fmt.Println("Неправильное использование команды")
	case 126:
		fmt.Println("Команда не может быть выполнена")
	case 127:
		fmt.Println("Команда не найдена")
	default:
		fmt.Printf("Неизвестный код выхода: %d\n", state.ExitCode())
	}

	// Мониторинг производительности
	start := time.Now()
	cmd := exec.Command("complex_program")
	err := cmd.Run()
	duration := time.Since(start)

	state := cmd.ProcessState
	if state != nil {
		fmt.Printf("Время выполнения: %v\n", duration)
		fmt.Printf("CPU время: %v\n", state.UserTime()+state.SystemTime())
		fmt.Printf("Эффективность: %.2f%%\n",
			float64(state.UserTime()+state.SystemTime())/float64(duration)*100)
	}

	Особенности работы:
	- ProcessState доступен только после завершения процесса
	- Некоторые методы (UserTime, SystemTime) работают только на Unix
	- SysUsage() может возвращать nil на некоторых платформах
	- Коды выхода интерпретируются по-разному в разных программах

	Связь с другими структурами:
	- Возвращается методом Process.Wait()
	- Связан с exec.Cmd.ProcessState
	- Использует syscall.WaitStatus для хранения статуса
	- Содержит syscall.Rusage для информации о ресурсах
*/
