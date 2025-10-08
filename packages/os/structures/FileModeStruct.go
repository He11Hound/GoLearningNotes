package structures

/*
	Тип FileMode

	type FileMode uint32

	FileMode представляет режим файла - комбинацию прав доступа и типа файла.
	Это битовая маска, где различные биты имеют разное значение.

	Константы для типов файлов:

	const (
		ModeDir        FileMode = 1 << (32 - 1 - iota) // d: директория
		ModeAppend                                      // a: только добавление
		ModeExclusive                                   // l: эксклюзивное использование
		ModeTemporary                                   // T: временный файл
		ModeSymlink                                     // L: символическая ссылка
		ModeDevice                                      // D: устройство
		ModeNamedPipe                                   // p: именованный канал
		ModeSocket                                      // S: сокет
		ModeSetuid                                      // u: setuid
		ModeSetgid                                      // g: setgid
		ModeCharDevice                                  // c: символьное устройство
		ModeSticky                                      // t: sticky
		ModeIrregular                                   // ?: нерегулярный файл
	)

	Константы для прав доступа:

	const (
		ModePerm FileMode = 0777 // права доступа (rwxrwxrwx)
	)

	Основные методы:

	func (m FileMode) String() string
		- Возвращает строковое представление режима файла
		- Формат: "-rwxrwxrwx" (тип файла + права доступа)
		- Примеры:
		  * "drwxr-xr-x" - директория с правами 755
		  * "-rw-r--r--" - обычный файл с правами 644
		  * "lrwxrwxrwx" - символическая ссылка

	func (m FileMode) IsDir() bool
		- Возвращает true если файл является директорией
		- Проверяет бит ModeDir

	func (m FileMode) IsRegular() bool
		- Возвращает true если файл является обычным файлом
		- Проверяет отсутствие специальных битов

	func (m FileMode) Perm() FileMode
		- Возвращает только права доступа (без типа файла)
		- Маскирует все биты кроме ModePerm (0777)

	Примеры использования:

	// Получение режима файла
	info, err := os.Stat("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := info.Mode()

	// Проверка типа файла
	if mode.IsDir() {
		fmt.Println("Это директория")
	}

	if mode.IsRegular() {
		fmt.Println("Это обычный файл")
	}

	// Проверка прав доступа
	perm := mode.Perm()
	if perm&0400 != 0 { // проверка права на чтение для владельца
		fmt.Println("Владелец может читать файл")
	}

	// Установка прав доступа
	err = os.Chmod("file.txt", 0644) // rw-r--r--

	// Создание файла с определенными правами
	file, err := os.OpenFile("newfile.txt", os.O_CREATE|os.O_WRONLY, 0644)

	// Проверка специальных атрибутов
	if mode&os.ModeSymlink != 0 {
		fmt.Println("Это символическая ссылка")
	}

	if mode&os.ModeSetuid != 0 {
		fmt.Println("Файл имеет setuid бит")
	}

	// Работа с правами доступа
	ownerRead := mode&0400 != 0
	ownerWrite := mode&0200 != 0
	ownerExecute := mode&0100 != 0

	groupRead := mode&0040 != 0
	groupWrite := mode&0020 != 0
	groupExecute := mode&0010 != 0

	otherRead := mode&0004 != 0
	otherWrite := mode&0002 != 0
	otherExecute := mode&0001 != 0

	Битовые операции:
	- Используется битовая маска для комбинации различных атрибутов
	- Можно комбинировать права доступа: 0755 = 0400|0200|0100|0040|0004|0001
	- Проверка прав: mode&0400 != 0 для проверки права на чтение владельца
*/
