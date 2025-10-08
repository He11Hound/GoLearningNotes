package os

import (
	"fmt"
	"log"
	"os"
)

// OpenExample демонстрирует использование функции os.Open
func OpenExample() {
	// Открытие файла для чтения
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Файл успешно открыт для чтения")
}

// CreateExample демонстрирует использование функции os.Create
func CreateExample() {
	// Создание нового файла
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Запись данных
	_, err = file.WriteString("Привет, мир!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл создан и данные записаны")
}

// OpenFileExample демонстрирует использование функции os.OpenFile
func OpenFileExample() {
	// Открытие файла с флагами
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Файл открыт с флагами O_CREATE|O_WRONLY|O_APPEND")
}

// StatExample демонстрирует использование функции os.Stat
func StatExample() {
	// Получение информации о файле
	info, err := os.Stat("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Размер файла: %d байт\n", info.Size())
	fmt.Printf("Время модификации: %v\n", info.ModTime())
	fmt.Printf("Права доступа: %v\n", info.Mode())
}

// LstatExample демонстрирует использование функции os.Lstat
func LstatExample() {
	// Получение информации о файле без следования символическим ссылкам
	info, err := os.Lstat("symlink_file")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Информация о символической ссылке:\n")
	fmt.Printf("Имя: %s\n", info.Name())
	fmt.Printf("Размер: %d байт\n", info.Size())
	fmt.Printf("Режим: %v\n", info.Mode())
	fmt.Printf("Это символическая ссылка: %v\n", info.Mode()&os.ModeSymlink != 0)
}

// ReadFileExample демонстрирует использование функции os.ReadFile
func ReadFileExample() {
	// Чтение всего файла
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Содержимое файла: %s\n", string(data))
}

// WriteFileExample демонстрирует использование функции os.WriteFile
func WriteFileExample() {
	// Запись данных в файл
	err := os.WriteFile("output.txt", []byte("Привет, мир!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные записаны в файл")
}

// ReadDirExample демонстрирует использование функции os.ReadDir
func ReadDirExample() {
	// Чтение директории
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Содержимое текущей директории:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("%s/ (директория)\n", entry.Name())
		} else {
			fmt.Printf("%s (файл)\n", entry.Name())
		}
	}
}

// MkdirExample демонстрирует использование функции os.Mkdir
func MkdirExample() {
	// Создание директории
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Директория создана")
}

// MkdirAllExample демонстрирует использование функции os.MkdirAll
func MkdirAllExample() {
	// Создание сложной структуры директорий
	err := os.MkdirAll("path/to/deep/directory", 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Структура директорий создана")
}

// RemoveExample демонстрирует использование функции os.Remove
func RemoveExample() {
	// Удаление файла
	err := os.Remove("temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл удален")
}

// RemoveAllExample демонстрирует использование функции os.RemoveAll
func RemoveAllExample() {
	// Удаление директории со всем содержимым
	err := os.RemoveAll("old_directory")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Директория и все содержимое удалены")
}

// RenameExample демонстрирует использование функции os.Rename
func RenameExample() {
	// Переименование файла
	err := os.Rename("old_name.txt", "new_name.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл переименован")
}

// TruncateExample демонстрирует использование функции os.Truncate
func TruncateExample() {
	// Изменение размера файла
	err := os.Truncate("large_file.txt", 1024) // обрезать до 1KB
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Размер файла изменен")
}

// GetpidExample демонстрирует использование функции os.Getpid
func GetpidExample() {
	// Получение PID текущего процесса
	pid := os.Getpid()
	fmt.Printf("PID текущего процесса: %d\n", pid)
}

// GetppidExample демонстрирует использование функции os.Getppid
func GetppidExample() {
	// Получение PID родительского процесса
	ppid := os.Getppid()
	fmt.Printf("PID родительского процесса: %d\n", ppid)
}

// GetuidExample демонстрирует использование функции os.Getuid
func GetuidExample() {
	// Получение User ID (Unix)
	if uid := os.Getuid(); uid != -1 {
		fmt.Printf("User ID: %d\n", uid)
	} else {
		fmt.Println("Функция работает только на Unix-системах")
	}
}

// GetgidExample демонстрирует использование функции os.Getgid
func GetgidExample() {
	// Получение Group ID (Unix)
	if gid := os.Getgid(); gid != -1 {
		fmt.Printf("Group ID: %d\n", gid)
	} else {
		fmt.Println("Функция работает только на Unix-системах")
	}
}

// GeteuidExample демонстрирует использование функции os.Geteuid
func GeteuidExample() {
	// Получение Effective User ID (Unix)
	if euid := os.Geteuid(); euid != -1 {
		fmt.Printf("Effective User ID: %d\n", euid)
	} else {
		fmt.Println("Функция работает только на Unix-системах")
	}
}

// GetegidExample демонстрирует использование функции os.Getegid
func GetegidExample() {
	// Получение Effective Group ID (Unix)
	if egid := os.Getegid(); egid != -1 {
		fmt.Printf("Effective Group ID: %d\n", egid)
	} else {
		fmt.Println("Функция работает только на Unix-системах")
	}
}

// ExitExample демонстрирует использование функции os.Exit
func ExitExample() {
	fmt.Println("Программа завершается...")
	os.Exit(0) // Завершение с кодом 0 (успешно)
}

// GetwdExample демонстрирует использование функции os.Getwd
func GetwdExample() {
	// Получение текущей директории
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Текущая директория: %s\n", wd)
}

// ChdirExample демонстрирует использование функции os.Chdir
func ChdirExample() {
	// Изменение директории
	err := os.Chdir("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Директория изменена на /tmp")
}

// StartProcessExample демонстрирует использование функции os.StartProcess
func StartProcessExample() {
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
}

// GetenvExample демонстрирует использование функции os.Getenv
func GetenvExample() {
	// Получение переменной окружения
	path := os.Getenv("PATH")
	fmt.Printf("PATH: %s\n", path)

	home := os.Getenv("HOME")
	fmt.Printf("HOME: %s\n", home)
}

// SetenvExample демонстрирует использование функции os.Setenv
func SetenvExample() {
	// Установка переменной окружения
	err := os.Setenv("MY_VAR", "my_value")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Переменная окружения установлена")
}

// UnsetenvExample демонстрирует использование функции os.Unsetenv
func UnsetenvExample() {
	// Удаление переменной окружения
	err := os.Unsetenv("MY_VAR")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Переменная окружения удалена")
}

// EnvironExample демонстрирует использование функции os.Environ
func EnvironExample() {
	// Получение всех переменных окружения
	env := os.Environ()
	fmt.Printf("Всего переменных окружения: %d\n", len(env))

	// Вывод первых 5 переменных
	for i, e := range env {
		if i >= 5 {
			break
		}
		fmt.Printf("%s\n", e)
	}
}

// LookupEnvExample демонстрирует использование функции os.LookupEnv
func LookupEnvExample() {
	// Проверка существования переменной
	if value, exists := os.LookupEnv("MY_VAR"); exists {
		fmt.Printf("MY_VAR существует: %s\n", value)
	} else {
		fmt.Println("MY_VAR не существует")
	}
}

// ExpandEnvExample демонстрирует использование функции os.ExpandEnv
func ExpandEnvExample() {
	// Замена переменных в строке
	config := "Database host: $DB_HOST, Port: ${DB_PORT}"
	expanded := os.ExpandEnv(config)
	fmt.Printf("Конфигурация: %s\n", expanded)
}

// ExpandExample демонстрирует использование функции os.Expand
func ExpandExample() {
	// Кастомная замена переменных
	config := "Database host: $DB_HOST, Port: ${DB_PORT}"
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
}

// ChmodExample демонстрирует использование функции os.Chmod
func ChmodExample() {
	// Изменение прав доступа
	err := os.Chmod("file.txt", 0644) // rw-r--r--
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Права доступа изменены")
}

// ChownExample демонстрирует использование функции os.Chown
func ChownExample() {
	// Изменение владельца файла (Unix)
	err := os.Chown("file.txt", 1000, 1000) // uid=1000, gid=1000
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Владелец файла изменен")
}

// LchownExample демонстрирует использование функции os.Lchown
func LchownExample() {
	// Изменение владельца символической ссылки (Unix)
	err := os.Lchown("symlink", 1000, 1000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Владелец символической ссылки изменен")
}

// GetgroupsExample демонстрирует использование функции os.Getgroups
func GetgroupsExample() {
	// Получение групп пользователя (Unix)
	groups, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Группы пользователя: %v\n", groups)
}

// ReadlinkExample демонстрирует использование функции os.Readlink
func ReadlinkExample() {
	// Чтение символической ссылки (Unix)
	target, err := os.Readlink("link")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ссылка указывает на: %s\n", target)
}

// SymlinkExample демонстрирует использование функции os.Symlink
func SymlinkExample() {
	// Создание символической ссылки (Unix)
	err := os.Symlink("/path/to/original", "link")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Символическая ссылка создана")
}

// LinkExample демонстрирует использование функции os.Link
func LinkExample() {
	// Создание жесткой ссылки (Unix)
	err := os.Link("original.txt", "hardlink.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Жесткая ссылка создана")
}

// TempDirExample демонстрирует использование функции os.TempDir
func TempDirExample() {
	// Получение директории для временных файлов
	tempDir := os.TempDir()
	fmt.Printf("Временная директория: %s\n", tempDir)
}

// CreateTempExample демонстрирует использование функции os.CreateTemp
func CreateTempExample() {
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
}

// MkdirTempExample демонстрирует использование функции os.MkdirTemp
func MkdirTempExample() {
	// Создание временной директории
	tempDirPath, err := os.MkdirTemp("", "tempdir_*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDirPath) // очистка

	fmt.Printf("Временная директория: %s\n", tempDirPath)
}

// IsExistExample демонстрирует использование функции os.IsExist
func IsExistExample() {
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
}

// IsNotExistExample демонстрирует использование функции os.IsNotExist
func IsNotExistExample() {
	// Проверка существования файла
	_, err := os.Stat("file.txt")
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Файл существует")
	}
}

// IsPermissionExample демонстрирует использование функции os.IsPermission
func IsPermissionExample() {
	// Проверка прав доступа
	_, err := os.OpenFile("/root/restricted_file", os.O_RDONLY, 0)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Недостаточно прав для доступа к файлу")
		} else {
			log.Fatal(err)
		}
	}
}

// IsTimeoutExample демонстрирует использование функции os.IsTimeout
func IsTimeoutExample() {
	// Проверка таймаута
	file, err := os.OpenFile("network_file.txt", os.O_RDONLY, 0)
	if err != nil {
		if os.IsTimeout(err) {
			fmt.Println("Операция превысила время ожидания")
		} else {
			log.Fatal(err)
		}
	} else {
		file.Close()
	}
}

// IsPathSeparatorExample демонстрирует использование функции os.IsPathSeparator
func IsPathSeparatorExample() {
	// Проверка разделителя пути
	if os.IsPathSeparator('/') {
		fmt.Println("Это разделитель пути")
	}

	if os.IsPathSeparator('\\') {
		fmt.Println("Это тоже разделитель пути (Windows)")
	}
}
