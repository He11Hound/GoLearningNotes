package io

/*
	ИНТЕРФЕЙС CLOSER

	type Closer interface {
		Close() error
	}

	Closer - это интерфейс для закрытия ресурсов.
	Этот интерфейс определяет стандартный способ освобождения
	ресурсов и очистки после завершения работы.

	Основной метод:

	Close() error
		- Закрывает ресурс и освобождает связанные с ним ресурсы
		- Возвращает ошибку если закрытие не удалось
		- Должен быть идемпотентным (безопасно вызывать несколько раз)
		- Должен освобождать все связанные ресурсы

	Поведение:
	- Close должен освобождать все ресурсы
	- Close должен быть безопасным для повторного вызова
	- Close должен возвращать ошибку только при реальных проблемах
	- Close может блокировать до завершения операций
	- Close должен быть потокобезопасным

	Примеры реализации:
	- os.File (файлы)
	- net.Conn (сетевые соединения)
	- io.PipeReader/io.PipeWriter (pipe)
	- bufio.Reader/Writer (буферизованные потоки)
	- compress/gzip.Reader/Writer (сжатие)

	Примеры использования:

	// Закрытие файла
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // автоматическое закрытие

	// Закрытие с проверкой ошибки
	err = file.Close()
	if err != nil {
		log.Printf("Ошибка закрытия файла: %v", err)
	}

	// Закрытие сетевого соединения
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Закрытие pipe
	reader, writer := io.Pipe()
	defer reader.Close()
	defer writer.Close()

	Особенности работы:
	- Close должен быть вызван для каждого открытого ресурса
	- Close может блокировать до завершения операций
	- Close должен быть потокобезопасным
	- Close должен освобождать все связанные ресурсы
	- Close может возвращать ошибки даже при успешном закрытии

	Связь с другими интерфейсами:
	- ReadCloser - для Reader с методом Close
	- WriteCloser - для Writer с методом Close
	- ReadWriteCloser - для Reader и Writer с методом Close
	- io.Closer - базовый интерфейс для всех закрываемых ресурсов

	Лучшие практики:
	- Всегда закрывайте ресурсы с помощью defer
	- Проверяйте ошибки при закрытии
	- Используйте defer для автоматического закрытия
	- Закрывайте ресурсы в порядке, обратном их созданию
	- Обрабатывайте ошибки закрытия отдельно от ошибок использования

	Паттерны использования:

	// Паттерн 1: defer для автоматического закрытия
	func processFile(filename string) error {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close() // закрытие при выходе из функции
		
		// работа с файлом
		return nil
	}

	// Паттерн 2: явное закрытие с проверкой ошибки
	func processFileExplicit(filename string) error {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		
		// работа с файлом
		
		if err := file.Close(); err != nil {
			return fmt.Errorf("ошибка закрытия файла: %w", err)
		}
		return nil
	}

	// Паттерн 3: закрытие нескольких ресурсов
	func processMultipleFiles() error {
		file1, err := os.Open("file1.txt")
		if err != nil {
			return err
		}
		defer file1.Close()
		
		file2, err := os.Open("file2.txt")
		if err != nil {
			return err
		}
		defer file2.Close()
		
		// работа с файлами
		return nil
	}
*/
