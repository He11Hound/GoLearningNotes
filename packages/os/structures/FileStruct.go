package structures

/*
	Структура File

	type File struct {
		*file // указатель на внутреннюю структуру
	}

	type file struct {
		fd      int         // файловый дескриптор ОС
		name    string      // имя файла
		dirinfo *dirInfo    // данные о директории (если это директория)
		nepipe  int32       // количество ошибок "broken pipe"
	}

	File представляет открытый файл или директорию в файловой системе.
	Предоставляет методы для чтения, записи, управления файлом и получения информации.

	Основные методы:

	func (f *File) Chdir() error
		- Изменяет текущую рабочую директорию на директорию, представленную файлом
		- Работает только если File представляет директорию
		- Эквивалентно os.Chdir() но для конкретного файла

	func (f *File) Chmod(mode FileMode) error
		- Изменяет права доступа к файлу
		- Принимает FileMode с новыми правами доступа
		- Эквивалентно os.Chmod() но для конкретного файла

	func (f *File) Chown(uid, gid int) error
		- Изменяет владельца файла (user ID и group ID)
		- Работает только на Unix-системах
		- Требует соответствующих прав доступа

	func (f *File) Close() error
		- Закрывает файл и освобождает ресурсы
		- Должен вызываться для каждого открытого файла
		- Рекомендуется использовать defer file.Close()

	func (f *File) Fd() uintptr
		- Возвращает файловый дескриптор ОС
		- Используется для низкоуровневых операций
		- Не рекомендуется для обычного использования

	func (f *File) Name() string
		- Возвращает имя файла, переданное при открытии
		- Может отличаться от реального пути к файлу
		- Полезно для логирования и отладки

	func (f *File) ReadAt(b []byte, off int64) (n int, err error)
		- Читает данные из файла с конкретного смещения
		- Не изменяет текущее положение указателя чтения
		- Полезно для параллельного чтения файла из разных горутин

	func (f *File) ReadDir(n int) ([]DirEntry, error)
		- Читает записи директории (современный метод)
		- Возвращает []DirEntry для эффективного чтения
		- n = -1 читает все записи, n > 0 читает n записей

	func (f *File) ReadFrom(r io.Reader) (n int64, err error)
		- Читает данные из io.Reader и записывает их в файл
		- Реализует интерфейс io.WriterTo
		- Полезно для копирования данных между потоками

	func (f *File) Readdir(n int) ([]FileInfo, error)
		- Читает записи директории (устаревший метод)
		- Возвращает []FileInfo с полной информацией
		- Менее эффективен чем ReadDir()

	func (f *File) Readdirnames(n int) (names []string, err error)
		- Возвращает только имена файлов/директорий
		- Самый быстрый способ получить список имен
		- Не предоставляет информацию о типе файлов

	func (f *File) Seek(offset int64, whence int) (ret int64, err error)
		- Перемещает указатель чтения/записи в файле
		- whence: 0=начало, 1=текущая позиция, 2=конец
		- Возвращает новую позицию в файле

	func (f *File) SetDeadline(t time.Time) error
		- Устанавливает общий таймаут для операций чтения/записи
		- Применяется ко всем операциям ввода-вывода
		- Полезно для предотвращения зависания

	func (f *File) SetReadDeadline(t time.Time) error
		- Устанавливает таймаут только для операций чтения
		- Более точный контроль над таймаутами
		- Не влияет на операции записи

	func (f *File) SetWriteDeadline(t time.Time) error
		- Устанавливает таймаут только для операций записи
		- Более точный контроль над таймаутами
		- Не влияет на операции чтения

	func (f *File) Stat() (FileInfo, error)
		- Получает информацию о файле
		- Возвращает FileInfo с метаданными файла
		- Эквивалентно os.Stat() но для конкретного файла

	func (f *File) Sync() error
		- Принудительно сбрасывает данные на диск
		- Гарантирует, что данные записаны на физическое устройство
		- Важно для критически важных данных

	func (f *File) SyscallConn() (syscall.RawConn, error)
		- Получает низкоуровневый доступ к файловому дескриптору
		- Используется для системных вызовов
		- Требует глубокого понимания системного программирования

	func (f *File) Truncate(size int64) error
		- Обрезает или расширяет файл до указанного размера
		- Если size меньше текущего размера - обрезает
		- Если size больше - расширяет нулями

	func (f *File) Write(b []byte) (n int, err error)
		- Записывает слайс байт в файл
		- Записывает с текущей позиции указателя
		- Возвращает количество записанных байт

	func (f *File) WriteAt(b []byte, off int64) (n int, err error)
		- Записывает данные в файл с конкретного смещения
		- Не изменяет текущее положение указателя записи
		- Полезно для параллельной записи в файл

	func (f *File) WriteString(s string) (n int, err error)
		- Записывает строку в файл
		- Удобный метод для записи текстовых данных
		- Эквивалентно Write([]byte(s))

	func (f *File) WriteTo(w io.Writer) (n int64, err error)
		- Копирует содержимое файла в io.Writer
		- Реализует интерфейс io.ReaderTo
		- Полезно для копирования файлов

	Примеры использования:

	// Открытие файла для чтения
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Чтение всего файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Содержимое файла: %s\n", string(data))

	// Открытие файла для записи
	file, err = os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Запись данных
	_, err = file.WriteString("Привет, мир!")
	if err != nil {
		log.Fatal(err)
	}

	// Получение информации о файле
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Размер файла: %d байт\n", info.Size())

	// Чтение директории
	dir, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("%s/ (директория)\n", entry.Name())
		} else {
			fmt.Printf("%s (файл)\n", entry.Name())
		}
	}

	// Работа с позицией в файле
	file, err = os.OpenFile("data.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Переход к концу файла
	pos, err := file.Seek(0, 2) // 2 = конец файла
	if err != nil {
		log.Fatal(err)
	}

	// Добавление данных в конец
	_, err = file.WriteString("\nНовая строка")
	if err != nil {
		log.Fatal(err)
	}

	// Параллельное чтение файла
	file, err = os.Open("large_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err = file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Чтение первой половины файла
	firstHalf := make([]byte, info.Size()/2)
	_, err = file.ReadAt(firstHalf, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Чтение второй половины файла
	secondHalf := make([]byte, info.Size()-info.Size()/2)
	_, err = file.ReadAt(secondHalf, info.Size()/2)
	if err != nil {
		log.Fatal(err)
	}

	// Установка таймаутов
	file, err = os.OpenFile("network_file.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Таймаут 5 секунд для всех операций
	err = file.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	// Изменение прав доступа
	err = file.Chmod(0644) // rw-r--r--
	if err != nil {
		log.Fatal(err)
	}

	// Обрезание файла
	err = file.Truncate(1024) // обрезать до 1KB
	if err != nil {
		log.Fatal(err)
	}

	// Принудительная запись на диск
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	Особенности работы:
	- File не является потокобезопасным
	- Close() должен вызываться для каждого открытого файла
	- ReadAt/WriteAt не изменяют позицию указателя
	- Seek работает с байтовыми смещениями
	- Таймауты применяются ко всем операциям ввода-вывода

	Связь с другими структурами:
	- File.Stat() возвращает FileInfo
	- File.ReadDir() возвращает []DirEntry
	- File.Chmod() принимает FileMode
	- File.SyscallConn() возвращает syscall.RawConn
	- Реализует интерфейсы io.Reader, io.Writer, io.ReaderAt, io.WriterAt
*/
