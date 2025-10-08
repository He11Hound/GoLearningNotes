package structures

/*
	Структура Process

	type Process struct {
		Pid    int           // идентификатор процесса
		handle uintptr       // дескриптор процесса (Windows) или 0 (Unix)
		isdone int32         // флаг завершения процесса
		sigMu  sync.RWMutex  // мьютекс для синхронизации сигналов
	}

	Process представляет внешний процесс, запущенный через os.StartProcess.
	Предоставляет методы для управления процессом: ожидание завершения,
	отправка сигналов, получение информации о состоянии.

	Основные методы:

	func (p *Process) Kill() error
		- Принудительно завершает процесс
		- На Unix отправляет SIGKILL
		- На Windows использует TerminateProcess
		- Нельзя отменить или обработать процессом

	func (p *Process) Release() error
		- Освобождает ресурсы, связанные с процессом
		- Должен вызываться после завершения работы с процессом
		- Не завершает сам процесс, только освобождает ресурсы

	func (p *Process) Signal(sig Signal) error
		- Отправляет сигнал процессу
		- На Unix использует kill() системный вызов
		- На Windows поддерживает ограниченный набор сигналов
		- Может использоваться для graceful shutdown (SIGTERM)

	func (p *Process) Wait() (*ProcessState, error)
		- Ожидает завершения процесса
		- Блокирующий вызов до завершения процесса
		- Возвращает ProcessState с информацией о завершении
		- Можно вызвать только один раз для каждого процесса

	Примеры использования:

	// Запуск процесса
	cmd := exec.Command("ls", "-la")
	process, err := os.StartProcess("", []string{"ls", "-la"}, &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Получение PID процесса
	fmt.Printf("PID процесса: %d\n", process.Pid)

	// Ожидание завершения процесса
	state, err := process.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Процесс завершился с кодом: %d\n", state.ExitCode())
	fmt.Printf("Успешно завершился: %v\n", state.Success())

	// Отправка сигнала процессу
	err = process.Signal(os.Interrupt)
	if err != nil {
		fmt.Printf("Ошибка отправки сигнала: %v\n", err)
	}

	// Принудительное завершение
	err = process.Kill()
	if err != nil {
		fmt.Printf("Ошибка завершения процесса: %v\n", err)
	}

	// Освобождение ресурсов
	err = process.Release()
	if err != nil {
		fmt.Printf("Ошибка освобождения ресурсов: %v\n", err)
	}

	// Запуск процесса с настройками
	attr := &os.ProcAttr{
		Dir:   "/tmp",                    // рабочая директория
		Env:   []string{"PATH=/bin"},     // переменные окружения
		Files: []*os.File{                // файловые дескрипторы
			os.Stdin,   // 0 - стандартный ввод
			os.Stdout,  // 1 - стандартный вывод
			os.Stderr,  // 2 - стандартный вывод ошибок
		},
	}

	process, err = os.StartProcess("/bin/ls", []string{"ls", "-la"}, attr)

	// Мониторинг процесса
	go func() {
		state, err := process.Wait()
		if err != nil {
			fmt.Printf("Ошибка ожидания процесса: %v\n", err)
			return
		}

		if state.Success() {
			fmt.Println("Процесс завершился успешно")
		} else {
			fmt.Printf("Процесс завершился с ошибкой: %v\n", state)
		}
	}()

	// Graceful shutdown
	err = process.Signal(os.Interrupt)
	if err != nil {
		// Если graceful shutdown не удался, принудительно завершаем
		process.Kill()
	}

	Особенности работы:
	- Process не является потокобезопасным для одновременного использования
	- Wait() можно вызвать только один раз для каждого процесса
	- Release() должен вызываться для освобождения ресурсов
	- На разных ОС поведение может отличаться

	Связь с другими структурами:
	- ProcessState возвращается методом Wait()
	- Signal используется для отправки сигналов
	- ProcAttr используется при создании процесса через StartProcess
	- exec.Cmd предоставляет более высокоуровневый API для работы с процессами
*/
