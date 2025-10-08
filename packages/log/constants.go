package log

/*
	КОНСТАНТЫ ПАКЕТА LOG

	Пакет log предоставляет константы для настройки форматирования логов.
	Эти константы используются с функцией SetFlags() для управления
	форматом вывода сообщений логгера.

	=== ФЛАГИ ФОРМАТИРОВАНИЯ ===

	Ldate = 1 << iota
		- Включает отображение даты
		- Формат: 2009/01/23
		- Показывает дату создания лога

	Ltime
		- Включает отображение времени
		- Формат: 01:23:23
		- Показывает время создания лога

	Lmicroseconds
		- Включает отображение микросекунд
		- Формат: 01:23:23.123123
		- Более точное время для отладки

	Llongfile
		- Включает полный путь к файлу и номер строки
		- Формат: /a/b/c/d.go:23
		- Полезно для отладки

	Lshortfile
		- Включает только имя файла и номер строки
		- Формат: d.go:23
		- Более компактный формат

	LUTC
		- Использует UTC время вместо локального
		- Все временные метки в UTC
		- Полезно для серверных приложений

	LstdFlags = Ldate | Ltime
		- Стандартные флаги по умолчанию
		- Включает дату и время
		- Наиболее часто используемая комбинация

	=== ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ ===

	// Только дата
	log.SetFlags(log.Ldate)
	log.Print("Сообщение с датой")

	// Дата и время
	log.SetFlags(log.Ldate | log.Ltime)
	log.Print("Сообщение с датой и временем")

	// Полная информация
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Print("Подробное сообщение")

	// Стандартные флаги
	log.SetFlags(log.LstdFlags)
	log.Print("Стандартное сообщение")

	// UTC время
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	log.Print("Сообщение с UTC временем")

	=== КОМБИНАЦИИ ФЛАГОВ ===

	// Минимальная информация
	log.SetFlags(0)
	log.Print("Только сообщение")

	// Базовая информация
	log.SetFlags(log.Ldate | log.Ltime)
	log.Print("Дата и время")

	// Расширенная информация
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Print("Дата, время и микросекунды")

	// Отладочная информация
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Дата, время и файл")

	// Полная отладочная информация
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.Print("Вся доступная информация")

	=== РЕКОМЕНДАЦИИ ПО ИСПОЛЬЗОВАНИЮ ===

	1. Для продакшена: log.Ldate | log.Ltime
	2. Для разработки: log.Ldate | log.Ltime | log.Lshortfile
	3. Для отладки: log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	4. Для серверов: log.Ldate | log.Ltime | log.LUTC
	5. Для минимального вывода: 0 (только сообщение)
*/
