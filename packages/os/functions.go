package os

/*
	ФУНКЦИИ ПАКЕТА OS

	Пакет os предоставляет множество функций для работы с операционной системой.
	Эти функции можно разделить на несколько категорий по функциональности.

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ФАЙЛАМИ ===

	func Open(name string) (*File, error)
		- Открывает файл для чтения
		- Возвращает *File и error
		- Файл должен существовать
		- Эквивалентно OpenFile(name, O_RDONLY, 0)

	func Create(name string) (*File, error)
		- Создает новый файл или обрезает существующий до нуля
		- Открывает файл для записи
		- Права доступа: 0666 (rw-rw-rw-)
		- Эквивалентно OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)

	func OpenFile(name string, flag int, perm FileMode) (*File, error)
		- Универсальная функция для открытия файлов
		- flag: флаги открытия (O_RDONLY, O_WRONLY, O_RDWR, O_CREAT, O_TRUNC, O_APPEND)
		- perm: права доступа для новых файлов
		- Наиболее гибкий способ открытия файлов

	func Stat(name string) (FileInfo, error)
		- Получает информацию о файле по имени
		- Следует символическим ссылкам
		- Возвращает FileInfo с метаданными файла

	func Lstat(name string) (FileInfo, error)
		- Получает информацию о файле, не следуя символическим ссылкам
		- Возвращает информацию о самой ссылке, а не о файле, на который она указывает
		- Полезно для работы с символическими ссылками

	func ReadFile(name string) ([]byte, error)
		- Читает весь файл в память
		- Возвращает содержимое файла как []byte
		- Удобно для небольших файлов
		- Автоматически закрывает файл

	func WriteFile(name string, data []byte, perm FileMode) error
		- Записывает данные в файл
		- Создает файл если не существует
		- Обрезает файл до нуля перед записью
		- Автоматически закрывает файл

	func ReadDir(name string) ([]DirEntry, error)
		- Читает содержимое директории
		- Возвращает []DirEntry для эффективного чтения
		- Современная альтернатива Readdir

	func Readdir(name string) ([]FileInfo, error)
		- Читает содержимое директории (устаревший метод)
		- Возвращает []FileInfo с полной информацией
		- Менее эффективен чем ReadDir

	func Readdirnames(name string) ([]string, error)
		- Возвращает только имена файлов в директории
		- Самый быстрый способ получить список имен
		- Не предоставляет информацию о типах файлов

	Примеры использования функций для работы с файлами:

	// Открытие файла для чтения
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создание нового файла
	file, err = os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Открытие файла с флагами
	file, err = os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Получение информации о файле
	info, err := os.Stat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Размер файла: %d байт\n", info.Size())

	// Чтение всего файла
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Содержимое: %s\n", string(data))

	// Запись данных в файл
	err = os.WriteFile("output.txt", []byte("Привет, мир!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Чтение директории
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Printf("%s\n", entry.Name())
	}

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ДИРЕКТОРИЯМИ ===

	func Mkdir(name string, perm FileMode) error
		- Создает директорию
		- perm: права доступа для новой директории
		- Не создает родительские директории
		- Возвращает ошибку если директория уже существует

	func MkdirAll(path string, perm FileMode) error
		- Создает директорию и все необходимые родительские директории
		- Не возвращает ошибку если директория уже существует
		- Полезно для создания сложных путей

	func Remove(name string) error
		- Удаляет файл или пустую директорию
		- Не удаляет непустые директории
		- Возвращает ошибку если файл не существует

	func RemoveAll(path string) error
		- Удаляет файл или директорию со всем содержимым
		- Опасная операция - удаляет все рекурсивно
		- Не возвращает ошибку если файл не существует

	func Rename(oldpath, newpath string) error
		- Переименовывает или перемещает файл/директорию
		- Работает в пределах одной файловой системы
		- Атомарная операция

	func Truncate(name string, size int64) error
		- Изменяет размер файла
		- Если size меньше текущего размера - обрезает файл
		- Если size больше - расширяет файл нулями

	Примеры использования функций для работы с директориями:

	// Создание директории
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Создание сложной структуры директорий
	err = os.MkdirAll("path/to/deep/directory", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Удаление файла
	err = os.Remove("temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Удаление директории со всем содержимым
	err = os.RemoveAll("old_directory")
	if err != nil {
		log.Fatal(err)
	}

	// Переименование файла
	err = os.Rename("old_name.txt", "new_name.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Изменение размера файла
	err = os.Truncate("large_file.txt", 1024) // обрезать до 1KB
	if err != nil {
		log.Fatal(err)
	}

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ПРОЦЕССАМИ ===

	func Getpid() int
		- Возвращает идентификатор текущего процесса (PID)
		- Уникальный номер процесса в системе

	func Getppid() int
		- Возвращает идентификатор родительского процесса
		- PID процесса, который создал текущий процесс

	func Getuid() int
		- Возвращает user ID текущего пользователя
		- Работает только на Unix-системах
		- На Windows возвращает -1

	func Getgid() int
		- Возвращает group ID текущего пользователя
		- Работает только на Unix-системах
		- На Windows возвращает -1

	func Geteuid() int
		- Возвращает effective user ID
		- Работает только на Unix-системах
		- На Windows возвращает -1

	func Getegid() int
		- Возвращает effective group ID
		- Работает только на Unix-системах
		- На Windows возвращает -1

	func Exit(code int)
		- Завершает программу с указанным кодом выхода
		- Немедленно завершает выполнение
		- Не выполняет отложенные функции defer

	func Getwd() (dir string, err error)
		- Возвращает текущую рабочую директорию
		- Абсолютный путь к текущей директории

	func Chdir(dir string) error
		- Изменяет текущую рабочую директорию
		- Влияет на все последующие операции с относительными путями

	func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
		- Запускает новый процесс
		- name: имя программы или путь к исполняемому файлу
		- argv: аргументы командной строки
		- attr: атрибуты процесса (рабочая директория, переменные окружения, файловые дескрипторы)

	Примеры использования функций для работы с процессами:

	// Получение информации о процессе
	fmt.Printf("PID текущего процесса: %d\n", os.Getpid())
	fmt.Printf("PID родительского процесса: %d\n", os.Getppid())

	// Получение информации о пользователе (Unix)
	if uid := os.Getuid(); uid != -1 {
		fmt.Printf("User ID: %d\n", uid)
		fmt.Printf("Group ID: %d\n", os.Getgid())
		fmt.Printf("Effective User ID: %d\n", os.Geteuid())
		fmt.Printf("Effective Group ID: %d\n", os.Getegid())
	}

	// Получение текущей директории
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Текущая директория: %s\n", wd)

	// Изменение директории
	err = os.Chdir("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	// Запуск процесса
	process, err := os.StartProcess("/bin/ls", []string{"ls", "-la"}, &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ожидание завершения процесса
	state, err := process.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Процесс завершился с кодом: %d\n", state.ExitCode())

	// Завершение программы
	os.Exit(0)

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ПЕРЕМЕННЫМИ ОКРУЖЕНИЯ ===

	func Getenv(key string) string
		- Возвращает значение переменной окружения
		- Возвращает пустую строку если переменная не установлена
		- Не различает отсутствие переменной и пустое значение

	func Setenv(key, value string) error
		- Устанавливает значение переменной окружения
		- Изменения применяются только к текущему процессу
		- Наследуются дочерними процессами

	func Unsetenv(key string) error
		- Удаляет переменную окружения
		- Переменная перестает существовать для текущего процесса
		- Не влияет на родительский процесс

	func Environ() []string
		- Возвращает все переменные окружения
		- Формат: "KEY=value"
		- Полезно для копирования окружения

	func LookupEnv(key string) (string, bool)
		- Возвращает значение переменной окружения и флаг существования
		- Позволяет различить отсутствие переменной и пустое значение
		- Более надежный способ проверки переменных

	func ExpandEnv(s string) string
		- Заменяет переменные окружения в строке
		- Формат: $VAR или ${VAR}
		- Полезно для обработки шаблонов

	func Expand(s string, mapping func(string) string) string
		- Заменяет переменные в строке используя пользовательскую функцию
		- Более гибкий способ замены переменных
		- Позволяет кастомную логику обработки

	Примеры использования функций для работы с переменными окружения:

	// Получение переменной окружения
	path := os.Getenv("PATH")
	fmt.Printf("PATH: %s\n", path)

	// Установка переменной окружения
	err := os.Setenv("MY_VAR", "my_value")
	if err != nil {
		log.Fatal(err)
	}

	// Проверка существования переменной
	if value, exists := os.LookupEnv("MY_VAR"); exists {
		fmt.Printf("MY_VAR существует: %s\n", value)
	} else {
		fmt.Println("MY_VAR не существует")
	}

	// Удаление переменной окружения
	err = os.Unsetenv("MY_VAR")
	if err != nil {
		log.Fatal(err)
	}

	// Получение всех переменных окружения
	env := os.Environ()
	for _, e := range env {
		fmt.Printf("%s\n", e)
	}

	// Замена переменных в строке
	config := "Database host: $DB_HOST, Port: ${DB_PORT}"
	expanded := os.ExpandEnv(config)
	fmt.Printf("Конфигурация: %s\n", expanded)

	// Кастомная замена переменных
	customExpanded := os.Expand(config, func(key string) string {
		switch key {
		case "DB_HOST":
			return "localhost"
		case "DB_PORT":
			return "5432"
		default:
			return ""
		}
	})
	fmt.Printf("Кастомная конфигурация: %s\n", customExpanded)

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ПРАВАМИ ДОСТУПА ===

	func Chmod(name string, mode FileMode) error
		- Изменяет права доступа к файлу или директории
		- mode: новые права доступа (например, 0644)
		- Работает на всех платформах

	func Chown(name string, uid, gid int) error
		- Изменяет владельца файла или директории
		- uid: user ID нового владельца
		- gid: group ID новой группы
		- Работает только на Unix-системах

	func Lchown(name string, uid, gid int) error
		- Изменяет владельца символической ссылки (не файла, на который она указывает)
		- Работает только на Unix-системах
		- Полезно для работы с символическими ссылками

	func Getgroups() ([]int, error)
		- Возвращает список групп, к которым принадлежит текущий пользователь
		- Работает только на Unix-системах
		- На Windows возвращает ошибку

	Примеры использования функций для работы с правами доступа:

	// Изменение прав доступа
	err := os.Chmod("file.txt", 0644) // rw-r--r--
	if err != nil {
		log.Fatal(err)
	}

	// Изменение владельца файла (Unix)
	err = os.Chown("file.txt", 1000, 1000) // uid=1000, gid=1000
	if err != nil {
		log.Fatal(err)
	}

	// Изменение владельца символической ссылки (Unix)
	err = os.Lchown("symlink", 1000, 1000)
	if err != nil {
		log.Fatal(err)
	}

	// Получение групп пользователя (Unix)
	groups, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Группы пользователя: %v\n", groups)

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С СИМВОЛИЧЕСКИМИ ССЫЛКАМИ ===

	func Readlink(name string) (string, error)
		- Читает содержимое символической ссылки
		- Возвращает путь, на который указывает ссылка
		- Работает только на Unix-системах

	func Symlink(oldname, newname string) error
		- Создает символическую ссылку
		- oldname: путь к файлу, на который указывает ссылка
		- newname: имя символической ссылки
		- Работает только на Unix-системах

	func Link(oldname, newname string) error
		- Создает жесткую ссылку
		- oldname: существующий файл
		- newname: имя новой ссылки
		- Работает только на Unix-системах

	Примеры использования функций для работы с символическими ссылками:

	// Создание символической ссылки (Unix)
	err := os.Symlink("/path/to/original", "link")
	if err != nil {
		log.Fatal(err)
	}

	// Чтение символической ссылки (Unix)
	target, err := os.Readlink("link")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ссылка указывает на: %s\n", target)

	// Создание жесткой ссылки (Unix)
	err = os.Link("original.txt", "hardlink.txt")
	if err != nil {
		log.Fatal(err)
	}

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ВРЕМЕННЫМИ ФАЙЛАМИ ===

	func TempDir() string
		- Возвращает директорию для временных файлов
		- Обычно /tmp на Unix, %TEMP% на Windows
		- Может быть изменена через переменную окружения TMPDIR

	func CreateTemp(dir, pattern string) (*File, error)
		- Создает временный файл
		- dir: директория для создания файла ("" = TempDir())
		- pattern: шаблон имени файла (например, "prefix*")
		- Возвращает открытый файл

	func MkdirTemp(dir, pattern string) (string, error)
		- Создает временную директорию
		- dir: родительская директория ("" = TempDir())
		- pattern: шаблон имени директории
		- Возвращает путь к созданной директории

	Примеры использования функций для работы с временными файлами:

	// Получение директории для временных файлов
	tempDir := os.TempDir()
	fmt.Printf("Временная директория: %s\n", tempDir)

	// Создание временного файла
	tempFile, err := os.CreateTemp("", "prefix_*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name()) // очистка
	defer tempFile.Close()

	fmt.Printf("Временный файл: %s\n", tempFile.Name())

	// Запись в временный файл
	_, err = tempFile.WriteString("Временные данные")
	if err != nil {
		log.Fatal(err)
	}

	// Создание временной директории
	tempDirPath, err := os.MkdirTemp("", "tempdir_*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDirPath) // очистка

	fmt.Printf("Временная директория: %s\n", tempDirPath)

	=== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ДЛЯ ОБРАБОТКИ ОШИБОК ===

	func IsExist(err error) bool
		- Проверяет, является ли ошибка "файл уже существует"
		- Полезно для обработки ошибок создания файлов
		- Возвращает true если файл уже существует

	func IsNotExist(err error) bool
		- Проверяет, является ли ошибка "файл не существует"
		- Полезно для проверки существования файлов
		- Возвращает true если файл не найден

	func IsPermission(err error) bool
		- Проверяет, является ли ошибка связанной с правами доступа
		- Полезно для обработки ошибок доступа
		- Возвращает true если недостаточно прав

	func IsTimeout(err error) bool
		- Проверяет, является ли ошибка связанной с таймаутом
		- Полезно для обработки сетевых операций
		- Возвращает true если операция превысила время ожидания

	func IsPathSeparator(c uint8) bool
		- Проверяет, является ли символ разделителем пути
		- Возвращает true для '/' на Unix и '\' на Windows
		- Полезно для работы с путями

	func IsPathError(err error) bool
		- Проверяет, является ли ошибка *PathError
		- PathError содержит информацию о пути и операции
		- Возвращает true если это ошибка пути

	Примеры использования вспомогательных функций:

	// Проверка существования файла
	_, err := os.Stat("file.txt")
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Файл существует")
	}

	// Создание файла с проверкой существования
	file, err := os.Create("newfile.txt")
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Файл уже существует")
		} else if os.IsPermission(err) {
			fmt.Println("Недостаточно прав для создания файла")
		} else {
			log.Fatal(err)
		}
	} else {
		defer file.Close()
		fmt.Println("Файл создан успешно")
	}

	// Проверка таймаута
	file, err = os.OpenFile("network_file.txt", os.O_RDONLY, 0)
	if err != nil {
		if os.IsTimeout(err) {
			fmt.Println("Операция превысила время ожидания")
		} else {
			log.Fatal(err)
		}
	}

	// Проверка типа ошибки
	if os.IsPathError(err) {
		fmt.Println("Это ошибка пути")
	}

	// Проверка разделителя пути
	if os.IsPathSeparator('/') {
		fmt.Println("Это разделитель пути")
	}

	=== КОНСТАНТЫ ПАКЕТА OS ===

	Флаги для OpenFile:
	- O_RDONLY: открыть только для чтения
	- O_WRONLY: открыть только для записи
	- O_RDWR: открыть для чтения и записи
	- O_APPEND: добавлять в конец файла
	- O_CREATE: создать файл если не существует
	- O_EXCL: использовать с O_CREATE, файл не должен существовать
	- O_SYNC: синхронная запись
	- O_TRUNC: обрезать файл до нуля при открытии

	Права доступа (FileMode):
	- 0400: чтение для владельца
	- 0200: запись для владельца
	- 0100: выполнение для владельца
	- 0040: чтение для группы
	- 0020: запись для группы
	- 0010: выполнение для группы
	- 0004: чтение для остальных
	- 0002: запись для остальных
	- 0001: выполнение для остальных

	Специальные типы файлов (FileMode):
	- ModeDir: директория
	- ModeAppend: только добавление
	- ModeExclusive: эксклюзивное использование
	- ModeTemporary: временный файл
	- ModeSymlink: символическая ссылка
	- ModeDevice: устройство
	- ModeNamedPipe: именованный канал
	- ModeSocket: сокет
	- ModeSetuid: setuid бит
	- ModeSetgid: setgid бит
	- ModeCharDevice: символьное устройство
	- ModeSticky: sticky бит
	- ModeIrregular: нерегулярный файл

	Сигналы (Signal):
	- Interrupt: SIGINT (прерывание)
	- Kill: SIGKILL (принудительное завершение)
	- SIGTERM: завершение процесса
	- SIGQUIT: выход с дампом памяти
	- SIGUSR1, SIGUSR2: пользовательские сигналы
*/
