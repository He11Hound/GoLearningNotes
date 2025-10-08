package structures

/*
	Структура SyscallError

	type SyscallError struct {
		Syscall string
		Err     error
	}

	SyscallError представляет ошибку системного вызова.
	Содержит имя системного вызова, который вызвал ошибку,
	и саму ошибку. Используется для предоставления контекста
	о том, какой именно системный вызов привел к ошибке.

	Поля структуры:

	Syscall string
		- Имя системного вызова, который вызвал ошибку
		- Примеры: "open", "read", "write", "stat", "chmod"
		- Помогает понять, какая операция привела к ошибке

	Err error
		- Оригинальная ошибка системного вызова
		- Обычно это syscall.Errno или другая системная ошибка
		- Содержит детальную информацию об ошибке

	Основные методы:

	func (e *SyscallError) Error() string
		- Реализует интерфейс error
		- Возвращает строковое представление ошибки
		- Формат: "syscall_name: error_description"
		- Пример: "open: no such file or directory"

	func (e *SyscallError) Unwrap() error
		- Возвращает оригинальную ошибку
		- Позволяет использовать errors.Is() и errors.As()
		- Поддерживает цепочку ошибок

	func (e *SyscallError) Timeout() bool
		- Возвращает true если ошибка связана с таймаутом
		- Проверяет, является ли ошибка временной
		- Используется для определения возможности повтора операции

	func (e *SyscallError) Temporary() bool
		- Возвращает true если ошибка является временной
		- Временные ошибки могут исчезнуть при повторной попытке
		- Примеры: EAGAIN, EINTR, EWOULDBLOCK

	Примеры использования:

	// Создание SyscallError
	syscallErr := &os.SyscallError{
		Syscall: "open",
		Err:     syscall.ENOENT, // "no such file or directory"
	}

	fmt.Println(syscallErr.Error()) // "open: no such file or directory"

	// Проверка типа ошибки
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		var sysErr *os.SyscallError
		if errors.As(err, &sysErr) {
			fmt.Printf("Системный вызов %s завершился с ошибкой: %v\n",
				sysErr.Syscall, sysErr.Err)

			// Проверка конкретной ошибки
			if errors.Is(sysErr.Err, syscall.ENOENT) {
				fmt.Println("Файл не найден")
			}
		}
	}

	// Обработка временных ошибок
	for i := 0; i < 3; i++ {
		file, err := os.Open("file.txt")
		if err != nil {
			var sysErr *os.SyscallError
			if errors.As(err, &sysErr) {
				if sysErr.Temporary() {
					fmt.Printf("Временная ошибка в %s, повторяем попытку...\n",
						sysErr.Syscall)
					time.Sleep(time.Second)
					continue
				}
			}
			log.Fatal(err)
		}
		file.Close()
		break
	}

	// Обработка таймаутов
	conn, err := net.DialTimeout("tcp", "example.com:80", 5*time.Second)
	if err != nil {
		var sysErr *os.SyscallError
		if errors.As(err, &sysErr) {
			if sysErr.Timeout() {
				fmt.Println("Таймаут соединения")
			}
		}
	}

	// Логирование системных ошибок
	func logSyscallError(err error) {
		var sysErr *os.SyscallError
		if errors.As(err, &sysErr) {
			log.Printf("Ошибка системного вызова %s: %v",
				sysErr.Syscall, sysErr.Err)

			// Дополнительная информация в зависимости от типа ошибки
			switch {
			case errors.Is(sysErr.Err, syscall.ENOENT):
				log.Println("Ресурс не найден")
			case errors.Is(sysErr.Err, syscall.EACCES):
				log.Println("Недостаточно прав доступа")
			case errors.Is(sysErr.Err, syscall.ENOSPC):
				log.Println("Недостаточно места на диске")
			case errors.Is(sysErr.Err, syscall.EINTR):
				log.Println("Системный вызов был прерван")
			}
		}
	}

	// Создание кастомной SyscallError
	func newSyscallError(syscall string, err error) *os.SyscallError {
		return &os.SyscallError{
			Syscall: syscall,
			Err:     err,
		}
	}

	// Использование в функциях
	func safeOpen(filename string) (*os.File, error) {
		file, err := os.Open(filename)
		if err != nil {
			var sysErr *os.SyscallError
			if errors.As(err, &sysErr) {
				// Добавляем контекст к ошибке
				return nil, fmt.Errorf("не удалось открыть файл %s: %w",
					filename, sysErr)
			}
			return nil, err
		}
		return file, nil
	}

	// Проверка конкретных системных вызовов
	func checkFileOperation(filename string) error {
		file, err := os.Open(filename)
		if err != nil {
			var sysErr *os.SyscallError
			if errors.As(err, &sysErr) {
				switch sysErr.Syscall {
				case "open":
					return fmt.Errorf("ошибка открытия файла: %v", sysErr.Err)
				case "stat":
					return fmt.Errorf("ошибка получения информации о файле: %v", sysErr.Err)
				default:
					return fmt.Errorf("неизвестная ошибка системного вызова %s: %v",
						sysErr.Syscall, sysErr.Err)
				}
			}
			return err
		}
		file.Close()
		return nil
	}

	Общие системные ошибки:
	- ENOENT: "no such file or directory" - файл не найден
	- EACCES: "permission denied" - недостаточно прав
	- ENOSPC: "no space left on device" - нет места на диске
	- EINTR: "interrupted system call" - прерванный системный вызов
	- EAGAIN: "resource temporarily unavailable" - ресурс временно недоступен
	- EWOULDBLOCK: "operation would block" - операция заблокирована

	Особенности работы:
	- SyscallError автоматически создается пакетом os при системных ошибках
	- Поддерживает цепочку ошибок через Unwrap()
	- Методы Timeout() и Temporary() помогают классифицировать ошибки
	- Имя системного вызова помогает в отладке

	Связь с другими структурами:
	- Содержит syscall.Errno в поле Err
	- Используется в методах File для обертывания системных ошибок
	- Поддерживает интерфейс error и временные ошибки
	- Интегрируется с пакетом errors для обработки ошибок
*/
