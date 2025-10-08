package structures

/*
	Интерфейс DirEntry

	type DirEntry interface {
		Name() string
		IsDir() bool
		Type() FileMode
		Info() (FileInfo, error)
	}

	DirEntry представляет запись в директории и является интерфейсом.
	Предоставляет информацию о файле или директории без необходимости
	получать полную информацию о файле (что может быть дорогостоящей операцией).

	Основные методы:

	func Name() string
		- Возвращает имя файла или директории
		- Базовое имя без пути
		- Например: "file.txt", "subdir"

	func IsDir() bool
		- Возвращает true если запись является директорией
		- Быстрая проверка типа без полного чтения метаданных файла
		- Более эффективно чем вызов Info().IsDir()

	func Type() FileMode
		- Возвращает тип файла (без прав доступа)
		- Содержит только биты типа файла (ModeDir, ModeSymlink и т.д.)
		- Не включает права доступа (ModePerm)
		- Быстрее чем Info().Mode() для определения типа

	func Info() (FileInfo, error)
		- Возвращает полную информацию о файле
		- Может быть дорогостоящей операцией
		- Используется когда нужна детальная информация (размер, время модификации и т.д.)

	Примеры использования:

	// Чтение директории
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// Итерация по записям
	for _, entry := range entries {
		name := entry.Name()

		// Быстрая проверка типа
		if entry.IsDir() {
			fmt.Printf("%s/ (директория)\n", name)
		} else {
			fmt.Printf("%s (файл)\n", name)
		}

		// Проверка типа через FileMode
		fileType := entry.Type()
		if fileType&os.ModeSymlink != 0 {
			fmt.Printf("%s -> символическая ссылка\n", name)
		}

		// Получение полной информации (только при необходимости)
		if needFullInfo {
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("Ошибка получения информации о %s: %v\n", name, err)
				continue
			}

			fmt.Printf("Размер: %d байт\n", info.Size())
			fmt.Printf("Время модификации: %v\n", info.ModTime())
		}
	}

	// Фильтрация файлов по типу
	var files []string
	var dirs []string

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		} else {
			files = append(files, entry.Name())
		}
	}

	// Проверка специальных типов файлов
	for _, entry := range entries {
		fileType := entry.Type()

		switch {
		case fileType&os.ModeDir != 0:
			fmt.Printf("%s - директория\n", entry.Name())
		case fileType&os.ModeSymlink != 0:
			fmt.Printf("%s - символическая ссылка\n", entry.Name())
		case fileType&os.ModeSocket != 0:
			fmt.Printf("%s - сокет\n", entry.Name())
		case fileType&os.ModeNamedPipe != 0:
			fmt.Printf("%s - именованный канал\n", entry.Name())
		case fileType&os.ModeDevice != 0:
			fmt.Printf("%s - устройство\n", entry.Name())
		default:
			fmt.Printf("%s - обычный файл\n", entry.Name())
		}
	}

	Преимущества DirEntry:
	- Эффективность: не требует системных вызовов для получения базовой информации
	- Быстрое определение типа файла через IsDir() и Type()
	- Ленивое получение полной информации через Info()
	- Оптимизировано для итерации по большим директориям

	Сравнение с FileInfo:
	- DirEntry: быстрый доступ к базовой информации
	- FileInfo: полная информация, но требует дополнительных системных вызовов
	- DirEntry.Info() возвращает FileInfo при необходимости

	Реализации DirEntry:
	- os.unixDirent (внутренняя структура для Unix-систем)
	- fs.DirEntry (из пакета io/fs)
	- Кастомные структуры, реализующие интерфейс

	Связь с другими структурами:
	- File.ReadDir() возвращает []DirEntry
	- os.ReadDir() возвращает []DirEntry
	- DirEntry.Info() возвращает FileInfo
	- DirEntry.Type() возвращает FileMode
*/
