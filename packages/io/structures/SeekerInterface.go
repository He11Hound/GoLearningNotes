package io

/*
	ИНТЕРФЕЙС SEEKER

	type Seeker interface {
		Seek(offset int64, whence int) (int64, error)
	}

	Seeker - это интерфейс для изменения позиции в потоке данных.
	Этот интерфейс позволяет навигировать по потоку данных,
	изменяя текущую позицию для чтения или записи.

	Основной метод:

	Seek(offset int64, whence int) (int64, error)
		- Изменяет позицию в потоке данных
		- offset: смещение в байтах
		- whence: точка отсчета (SeekStart, SeekCurrent, SeekEnd)
		- Возвращает новую позицию и ошибку
		- Позиция отсчитывается от начала потока

	Параметры:
	- offset: смещение в байтах (может быть отрицательным)
	- whence: точка отсчета:
		* SeekStart (0) - от начала потока
		* SeekCurrent (1) - от текущей позиции
		* SeekEnd (2) - от конца потока

	Поведение:
	- Seek должен возвращать новую позицию в потоке
	- Seek должен обрабатывать отрицательные смещения
	- Seek должен возвращать ошибку для недопустимых позиций
	- Seek может расширить поток при записи за конец
	- Seek должен быть потокобезопасным

	Примеры реализации:
	- os.File (файлы)
	- bytes.Reader (байты в памяти)
	- strings.Reader (строки)
	- compress/gzip.Reader (сжатые данные)

	Примеры использования:

	// Позиционирование в файле
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Позиционирование в начало
	pos, err := file.Seek(0, SeekStart)
	fmt.Printf("Позиция в начале: %d\n", pos)

	// Позиционирование в конец
	pos, err = file.Seek(0, SeekEnd)
	fmt.Printf("Позиция в конце: %d\n", pos)

	// Позиционирование на 100 байтов от начала
	pos, err = file.Seek(100, SeekStart)
	fmt.Printf("Позиция на 100 байтов: %d\n", pos)

	// Позиционирование на 50 байтов назад от текущей позиции
	pos, err = file.Seek(-50, SeekCurrent)
	fmt.Printf("Позиция на 50 байтов назад: %d\n", pos)

	// Позиционирование на 10 байтов от конца
	pos, err = file.Seek(-10, SeekEnd)
	fmt.Printf("Позиция на 10 байтов от конца: %d\n", pos)

	Особенности работы:
	- Seek работает с байтовыми позициями
	- Seek может расширить файл при записи за конец
	- Seek должен быть потокобезопасным
	- Seek может быть медленным для некоторых типов потоков
	- Seek не поддерживается всеми типами потоков

	Связь с другими интерфейсами:
	- ReadSeeker - для Reader с возможностью Seek
	- WriteSeeker - для Writer с возможностью Seek
	- ReadWriteSeeker - для Reader и Writer с возможностью Seek
	- ReaderAt - для чтения с определенной позиции
	- WriterAt - для записи в определенную позицию

	Лучшие практики:
	- Всегда проверяйте ошибки после Seek
	- Используйте именованные константы вместо чисел
	- Учитывайте, что не все потоки поддерживают Seek
	- Проверяйте поддержку Seek перед использованием
	- Используйте Seek для эффективного доступа к данным

	Паттерны использования:

	// Паттерн 1: чтение с определенной позиции
	func readFromPosition(file *os.File, offset int64) ([]byte, error) {
		_, err := file.Seek(offset, SeekStart)
		if err != nil {
			return nil, err
		}
		
		buffer := make([]byte, 1024)
		n, err := file.Read(buffer)
		if err != nil && err != EOF {
			return nil, err
		}
		
		return buffer[:n], nil
	}

	// Паттерн 2: запись в конец файла
	func appendToFile(file *os.File, data []byte) error {
		_, err := file.Seek(0, SeekEnd)
		if err != nil {
			return err
		}
		
		_, err = file.Write(data)
		return err
	}

	// Паттерн 3: проверка поддержки Seek
	func canSeek(reader io.Reader) bool {
		_, ok := reader.(io.Seeker)
		return ok
	}

	// Паттерн 4: безопасный Seek
	func safeSeek(seeker io.Seeker, offset int64, whence int) (int64, error) {
		pos, err := seeker.Seek(offset, whence)
		if err != nil {
			return 0, fmt.Errorf("ошибка Seek: %w", err)
		}
		return pos, nil
	}
*/
