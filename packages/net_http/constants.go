package net_http

/*
	КОНСТАНТЫ ПАКЕТА NET/HTTP

	Пакет net/http предоставляет константы для работы с HTTP протоколом,
	статус кодами, методами и заголовками. Эти константы используются
	для стандартизации HTTP операций.

	=== HTTP МЕТОДЫ ===

	"GET" - получение ресурса
		- Используется для получения данных
		- Идемпотентный метод
		- Не должен изменять состояние сервера
		- Может кэшироваться

	"POST" - создание ресурса
		- Используется для отправки данных
		- Не идемпотентный метод
		- Может изменять состояние сервера
		- Не кэшируется

	"PUT" - обновление ресурса
		- Используется для полного обновления ресурса
		- Идемпотентный метод
		- Может изменять состояние сервера
		- Не кэшируется

	"DELETE" - удаление ресурса
		- Используется для удаления ресурса
		- Идемпотентный метод
		- Может изменять состояние сервера
		- Не кэшируется

	"HEAD" - получение заголовков
		- Используется для получения только заголовков
		- Идемпотентный метод
		- Не возвращает тело ответа
		- Может кэшироваться

	"OPTIONS" - получение опций
		- Используется для получения доступных методов
		- Идемпотентный метод
		- Не изменяет состояние сервера
		- Может кэшироваться

	"PATCH" - частичное обновление
		- Используется для частичного обновления ресурса
		- Не идемпотентный метод
		- Может изменять состояние сервера
		- Не кэшируется

	"TRACE" - трассировка
		- Используется для диагностики
		- Идемпотентный метод
		- Возвращает полученный запрос
		- Не кэшируется

	"CONNECT" - подключение
		- Используется для туннелирования
		- Не идемпотентный метод
		- Создает туннель к серверу
		- Не кэшируется

	=== HTTP СТАТУС КОДЫ ===

	// 1xx - Информационные
	StatusContinue           = 100 // Продолжить
	StatusSwitchingProtocols = 101 // Переключение протоколов
	StatusProcessing         = 102 // Обработка
	StatusEarlyHints         = 103 // Ранние подсказки

	// 2xx - Успешные
	StatusOK                   = 200 // OK
	StatusCreated              = 201 // Создано
	StatusAccepted             = 202 // Принято
	StatusNonAuthoritativeInfo = 203 // Неавторитетная информация
	StatusNoContent            = 204 // Нет содержимого
	StatusResetContent         = 205 // Сброс содержимого
	StatusPartialContent       = 206 // Частичное содержимое
	StatusMultiStatus          = 207 // Множественный статус
	StatusAlreadyReported      = 208 // Уже сообщено
	StatusIMUsed               = 226 // Использовано IM

	// 3xx - Перенаправления
	StatusMultipleChoices   = 300 // Множественный выбор
	StatusMovedPermanently  = 301 // Перемещено навсегда
	StatusFound             = 302 // Найдено
	StatusSeeOther          = 303 // См. другое
	StatusNotModified       = 304 // Не изменено
	StatusUseProxy          = 305 // Использовать прокси
	StatusTemporaryRedirect = 307 // Временное перенаправление
	StatusPermanentRedirect = 308 // Постоянное перенаправление

	// 4xx - Ошибки клиента
	StatusBadRequest                  = 400 // Плохой запрос
	StatusUnauthorized                = 401 // Неавторизован
	StatusPaymentRequired             = 402 // Требуется оплата
	StatusForbidden                   = 403 // Запрещено
	StatusNotFound                    = 404 // Не найдено
	StatusMethodNotAllowed            = 405 // Метод не разрешен
	StatusNotAcceptable               = 406 // Неприемлемо
	StatusProxyAuthRequired           = 407 // Требуется аутентификация прокси
	StatusRequestTimeout              = 408 // Таймаут запроса
	StatusConflict                    = 409 // Конфликт
	StatusGone                        = 410 // Удалено
	StatusLengthRequired              = 411 // Требуется длина
	StatusPreconditionFailed          = 412 // Предусловие не выполнено
	StatusRequestEntityTooLarge       = 413 // Сущность запроса слишком большая
	StatusRequestURITooLong          = 414 // URI запроса слишком длинный
	StatusUnsupportedMediaType       = 415 // Неподдерживаемый тип медиа
	StatusRequestedRangeNotSatisfiable = 416 // Запрошенный диапазон не выполним
	StatusExpectationFailed           = 417 // Ожидание не выполнено
	StatusTeapot                      = 418 // Я чайник
	StatusMisdirectedRequest          = 421 // Неправильно направленный запрос
	StatusUnprocessableEntity         = 422 // Необрабатываемая сущность
	StatusLocked                      = 423 // Заблокировано
	StatusFailedDependency            = 424 // Неудачная зависимость
	StatusTooEarly                    = 425 // Слишком рано
	StatusUpgradeRequired             = 426 // Требуется обновление
	StatusPreconditionRequired        = 428 // Требуется предусловие
	StatusTooManyRequests              = 429 // Слишком много запросов
	StatusRequestHeaderFieldsTooLarge = 431 // Поля заголовка запроса слишком большие
	StatusUnavailableForLegalReasons  = 451 // Недоступно по юридическим причинам

	// 5xx - Ошибки сервера
	StatusInternalServerError           = 500 // Внутренняя ошибка сервера
	StatusNotImplemented               = 501 // Не реализовано
	StatusBadGateway                    = 502 // Плохой шлюз
	StatusServiceUnavailable           = 503 // Сервис недоступен
	StatusGatewayTimeout                = 504 // Таймаут шлюза
	StatusHTTPVersionNotSupported       = 505 // Версия HTTP не поддерживается
	StatusVariantAlsoNegotiates         = 506 // Вариант также согласовывается
	StatusInsufficientStorage           = 507 // Недостаточно места
	StatusLoopDetected                  = 508 // Обнаружена петля
	StatusNotExtended                   = 510 // Не расширено
	StatusNetworkAuthenticationRequired = 511 // Требуется сетевая аутентификация

	=== HTTP ЗАГОЛОВКИ ===

	// Общие заголовки
	"Accept"                    // Типы содержимого, которые принимает клиент
	"Accept-Charset"            // Кодировки, которые принимает клиент
	"Accept-Encoding"          // Методы сжатия, которые принимает клиент
	"Accept-Language"          // Языки, которые принимает клиент
	"Authorization"            // Учетные данные для аутентификации
	"Cache-Control"            // Директивы кэширования
	"Connection"               // Управление соединением
	"Content-Encoding"         // Метод сжатия содержимого
	"Content-Language"         // Язык содержимого
	"Content-Length"           // Длина содержимого в байтах
	"Content-Type"             // Тип содержимого
	"Date"                     // Дата и время сообщения
	"Expect"                   // Ожидания клиента
	"Expires"                  // Время истечения содержимого
	"From"                     // Email адрес пользователя
	"Host"                     // Имя хоста и порт
	"If-Match"                 // Условный запрос с ETag
	"If-Modified-Since"        // Условный запрос с датой
	"If-None-Match"            // Условный запрос с ETag
	"If-Range"                 // Условный запрос с диапазоном
	"If-Unmodified-Since"      // Условный запрос с датой
	"Last-Modified"            // Дата последней модификации
	"Location"                 // URL для перенаправления
	"Max-Forwards"             // Максимальное количество пересылок
	"Pragma"                   // Директивы реализации
	"Proxy-Authorization"      // Учетные данные для прокси
	"Range"                    // Диапазон байтов для запроса
	"Referer"                  // URL источника запроса
	"Retry-After"              // Время повторной попытки
	"Server"                   // Информация о сервере
	"TE"                       // Методы кодирования передачи
	"Trailer"                  // Заголовки в конце сообщения
	"Transfer-Encoding"        // Методы кодирования передачи
	"Upgrade"                  // Протоколы для обновления
	"User-Agent"               // Информация о клиенте
	"Vary"                     // Заголовки для вариации
	"Via"                      // Промежуточные прокси
	"Warning"                  // Предупреждения
	"WWW-Authenticate"         // Схемы аутентификации

	// Заголовки для cookies
	"Cookie"                   // Cookies от клиента
	"Set-Cookie"               // Cookies для установки

	// Заголовки для CORS
	"Access-Control-Allow-Origin"      // Разрешенные источники
	"Access-Control-Allow-Methods"     // Разрешенные методы
	"Access-Control-Allow-Headers"     // Разрешенные заголовки
	"Access-Control-Allow-Credentials" // Разрешение учетных данных
	"Access-Control-Expose-Headers"    // Заголовки для экспозиции
	"Access-Control-Max-Age"           // Время кэширования preflight
	"Access-Control-Request-Method"    // Запрашиваемый метод
	"Access-Control-Request-Headers"   // Запрашиваемые заголовки

	// Заголовки для безопасности
	"X-Content-Type-Options"  // Защита от MIME sniffing
	"X-Frame-Options"          // Защита от clickjacking
	"X-XSS-Protection"        // Защита от XSS
	"Strict-Transport-Security" // Принуждение HTTPS
	"Content-Security-Policy" // Политика безопасности содержимого

	=== ТИПЫ СОДЕРЖИМОГО ===

	"text/plain"                    // Простой текст
	"text/html"                     // HTML документ
	"text/css"                      // CSS стили
	"text/javascript"               // JavaScript код
	"text/csv"                      // CSV данные
	"text/xml"                      // XML документ
	"text/markdown"                 // Markdown документ

	"application/json"              // JSON данные
	"application/xml"                // XML данные
	"application/pdf"               // PDF документ
	"application/zip"               // ZIP архив
	"application/octet-stream"      // Бинарные данные
	"application/x-www-form-urlencoded" // Form данные
	"application/javascript"       // JavaScript код
	"application/ld+json"           // JSON-LD данные

	"image/jpeg"                    // JPEG изображение
	"image/png"                     // PNG изображение
	"image/gif"                     // GIF изображение
	"image/svg+xml"                 // SVG изображение
	"image/webp"                    // WebP изображение

	"audio/mpeg"                    // MP3 аудио
	"audio/wav"                     // WAV аудио
	"audio/ogg"                     // OGG аудио

	"video/mp4"                     // MP4 видео
	"video/webm"                    // WebM видео
	"video/ogg"                     // OGG видео

	"multipart/form-data"           // Multipart form данные
	"multipart/mixed"               // Multipart mixed данные
	"multipart/alternative"         // Multipart alternative данные

	=== ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ ===

	// HTTP методы
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	req, _ := http.NewRequest("POST", "https://api.example.com", body)
	req, _ := http.NewRequest("PUT", "https://api.example.com/resource", body)
	req, _ := http.NewRequest("DELETE", "https://api.example.com/resource", nil)

	// Статус коды
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusNotFound)
	w.WriteHeader(http.StatusInternalServerError)

	// Заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	req.Header.Set("User-Agent", "MyApp/1.0")

	// Типы содержимого
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Type", "image/jpeg")

	// Cookies
	cookie := &http.Cookie{
		Name:  "session",
		Value: "abc123",
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	// CORS заголовки
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Заголовки безопасности
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000")

	=== ОСОБЕННОСТИ РАБОТЫ ===

	1. HTTP методы
		- GET и HEAD идемпотентны
		- POST и PUT могут изменять состояние
		- DELETE идемпотентен
		- OPTIONS используется для CORS

	2. Статус коды
		- 1xx: информационные
		- 2xx: успешные
		- 3xx: перенаправления
		- 4xx: ошибки клиента
		- 5xx: ошибки сервера

	3. Заголовки
		- Регистр не важен
		- Могут содержать множественные значения
		- Некоторые заголовки имеют специальное значение

	4. Типы содержимого
		- Определяют формат данных
		- Используются для парсинга
		- Влияют на поведение браузера

	=== РЕКОМЕНДАЦИИ ПО ИСПОЛЬЗОВАНИЮ ===

	1. Выбор методов
		- Используйте GET для получения данных
		- Используйте POST для создания
		- Используйте PUT для обновления
		- Используйте DELETE для удаления

	2. Статус коды
		- Используйте правильные статус коды
		- 200 для успешных операций
		- 404 для не найденных ресурсов
		- 500 для ошибок сервера

	3. Заголовки
		- Устанавливайте правильные Content-Type
		- Используйте CORS для веб-приложений
		- Добавляйте заголовки безопасности

	4. Типы содержимого
		- Используйте application/json для API
		- Используйте text/html для веб-страниц
		- Используйте image/* для изображений
*/
