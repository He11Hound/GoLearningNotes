package log

import (
	"fmt"
	"log"
	"os"
)

// PrintExample демонстрирует использование функции log.Print
func PrintExample() {
	// Простое логирование
	log.Print("Простое сообщение лога")
	log.Print("Сообщение с", " несколькими", " аргументами")
}

// PrintfExample демонстрирует использование функции log.Printf
func PrintfExample() {
	// Форматированное логирование
	name := "Пользователь"
	age := 25
	log.Printf("Пользователь %s, возраст %d лет", name, age)
	log.Printf("Ошибка: %v", fmt.Errorf("что-то пошло не так"))
}

// PrintlnExample демонстрирует использование функции log.Println
func PrintlnExample() {
	// Логирование с переносом строки
	log.Println("Сообщение с переносом строки")
	log.Println("Несколько", "аргументов", "с", "пробелами")
}

// FatalExample демонстрирует использование функции log.Fatal
func FatalExample() {
	// Критическая ошибка - завершение программы
	fmt.Println("Это сообщение будет выведено")
	log.Fatal("Критическая ошибка - программа завершается")
	fmt.Println("Это сообщение НЕ будет выведено")
}

// FatalfExample демонстрирует использование функции log.Fatalf
func FatalfExample() {
	// Критическая ошибка с форматированием
	err := fmt.Errorf("ошибка подключения к базе данных")
	log.Fatalf("Критическая ошибка: %v", err)
}

// FatallnExample демонстрирует использование функции log.Fatalln
func FatallnExample() {
	// Критическая ошибка с переносом строки
	log.Fatalln("Критическая ошибка", "с", "несколькими", "аргументами")
}

// PanicExample демонстрирует использование функции log.Panic
func PanicExample() {
	// Паника - невосстановимая ошибка
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Перехвачена паника: %v\n", r)
		}
	}()
	
	log.Panic("Невосстановимая ошибка - вызываем панику")
}

// PanicfExample демонстрирует использование функции log.Panicf
func PanicfExample() {
	// Паника с форматированием
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Перехвачена паника: %v\n", r)
		}
	}()
	
	err := fmt.Errorf("ошибка валидации")
	log.Panicf("Паника: %v", err)
}

// PaniclnExample демонстрирует использование функции log.Panicln
func PaniclnExample() {
	// Паника с переносом строки
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Перехвачена паника: %v\n", r)
		}
	}()
	
	log.Panicln("Паника", "с", "несколькими", "аргументами")
}

// SetFlagsExample демонстрирует использование функции log.SetFlags
func SetFlagsExample() {
	// Установка флагов форматирования
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Сообщение с датой, временем и файлом")
	
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Print("Сообщение с микросекундами")
	
	log.SetFlags(log.LstdFlags)
	log.Print("Стандартное форматирование")
}

// SetPrefixExample демонстрирует использование функции log.SetPrefix
func SetPrefixExample() {
	// Установка префикса
	log.SetPrefix("APP: ")
	log.Print("Сообщение с префиксом")
	
	log.SetPrefix("ERROR: ")
	log.Print("Сообщение об ошибке")
	
	log.SetPrefix("")
	log.Print("Сообщение без префикса")
}

// SetOutputExample демонстрирует использование функции log.SetOutput
func SetOutputExample() {
	// Перенаправление вывода в файл
	file, err := os.Create("log_output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// Сохраняем оригинальный вывод
	originalOutput := os.Stderr
	
	// Перенаправляем в файл
	log.SetOutput(file)
	log.Print("Это сообщение будет записано в файл")
	
	// Возвращаем оригинальный вывод
	log.SetOutput(originalOutput)
	log.Print("Это сообщение будет выведено в консоль")
}

// NewExample демонстрирует использование функции log.New
func NewExample() {
	// Создание кастомного логгера
	file, err := os.Create("custom_log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// Создание логгера с кастомными настройками
	customLogger := log.New(file, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)
	customLogger.Print("Сообщение от кастомного логгера")
	customLogger.Printf("Форматированное сообщение: %s", "значение")
	
	// Создание логгера для консоли
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	consoleLogger.Print("Сообщение в консоль")
}

// FlagsExample демонстрирует использование функции log.Flags
func FlagsExample() {
	// Получение текущих флагов
	currentFlags := log.Flags()
	fmt.Printf("Текущие флаги: %d\n", currentFlags)
	
	// Изменение флагов
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	newFlags := log.Flags()
	fmt.Printf("Новые флаги: %d\n", newFlags)
}

// PrefixExample демонстрирует использование функции log.Prefix
func PrefixExample() {
	// Получение текущего префикса
	currentPrefix := log.Prefix()
	fmt.Printf("Текущий префикс: '%s'\n", currentPrefix)
	
	// Изменение префикса
	log.SetPrefix("NEW: ")
	newPrefix := log.Prefix()
	fmt.Printf("Новый префикс: '%s'\n", newPrefix)
}

// OutputExample демонстрирует использование функции log.Output
func OutputExample() {
	// Запись сообщения с указанием глубины стека
	err := log.Output(1, "Сообщение через Output с глубиной 1")
	if err != nil {
		fmt.Printf("Ошибка записи: %v\n", err)
	}
	
	// Запись с другой глубиной
	err = log.Output(2, "Сообщение через Output с глубиной 2")
	if err != nil {
		fmt.Printf("Ошибка записи: %v\n", err)
	}
}

// BasicLoggingExample демонстрирует базовое логирование
func BasicLoggingExample() {
	// Простые сообщения
	log.Print("Начало работы программы")
	log.Printf("Загружено %d конфигураций", 5)
	log.Println("Программа готова к работе")
}

// ErrorLoggingExample демонстрирует логирование ошибок
func ErrorLoggingExample() {
	// Логирование различных типов ошибок
	log.Print("Предупреждение: низкий уровень памяти")
	log.Printf("Ошибка подключения: %v", fmt.Errorf("timeout"))
	log.Println("Информация: пользователь вошел в систему")
}

// FileLoggingExample демонстрирует логирование в файл
func FileLoggingExample() {
	// Создание файла для логов
	file, err := os.Create("application.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// Настройка логгера для файла
	fileLogger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	
	// Запись в файл
	fileLogger.Print("Запуск приложения")
	fileLogger.Printf("Версия: %s", "1.0.0")
	fileLogger.Print("Приложение готово")
}

// MultiLoggerExample демонстрирует использование нескольких логгеров
func MultiLoggerExample() {
	// Создание файла для ошибок
	errorFile, err := os.Create("errors.log")
	if err != nil {
		log.Fatal(err)
	}
	defer errorFile.Close()
	
	// Создание файла для информации
	infoFile, err := os.Create("info.log")
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()
	
	// Логгер для ошибок
	errorLogger := log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime)
	
	// Логгер для информации
	infoLogger := log.New(infoFile, "INFO: ", log.Ldate|log.Ltime)
	
	// Использование разных логгеров
	infoLogger.Print("Пользователь вошел в систему")
	errorLogger.Print("Ошибка валидации данных")
	infoLogger.Print("Данные успешно сохранены")
}

// ConditionalLoggingExample демонстрирует условное логирование
func ConditionalLoggingExample() {
	debug := true
	verbose := false
	
	// Условное логирование
	if debug {
		log.Print("Отладочная информация")
		log.Printf("Переменная debug = %v", debug)
	}
	
	if verbose {
		log.Print("Подробная информация")
	} else {
		log.Print("Базовая информация")
	}
}

// StructuredLoggingExample демонстрирует структурированное логирование
func StructuredLoggingExample() {
	// Структурированные сообщения
	userID := 123
	action := "login"
	timestamp := "2023-01-01T12:00:00Z"
	
	log.Printf("UserID=%d Action=%s Timestamp=%s", userID, action, timestamp)
	log.Printf("Event=user_action UserID=%d Action=%s Status=success", userID, action)
	log.Printf("Metric=response_time Value=%dms Endpoint=/api/users", 150)
}
