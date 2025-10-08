package net_http

import (
	"net/http"
)

/*
	Интерфейс Handler

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

	Handler представляет интерфейс для обработки HTTP запросов.
	Это основной интерфейс для создания HTTP обработчиков в Go.
	Реализуется различными типами для разных целей.

	Основные методы:

	func ServeHTTP(ResponseWriter, *Request)
		- Обрабатывает HTTP запрос
		- w: ResponseWriter для записи ответа
		- r: Request с параметрами запроса
		- Должен обработать запрос и отправить ответ
		- Используется для создания обработчиков

	Типы Handler:

	http.HandlerFunc - функция обработчик
		- Преобразует функцию в Handler
		- Используется для простых обработчиков
		- Наиболее распространенный тип

	http.ServeMux - мультиплексор
		- Маршрутизирует запросы по URL
		- Используется для создания роутеров
		- Поддерживает паттерны URL

	http.FileServer - файловый сервер
		- Обслуживает статические файлы
		- Используется для статических ресурсов
		- Поддерживает HTTP кэширование

	http.TimeoutHandler - обработчик с таймаутом
		- Добавляет таймаут к обработчику
		- Используется для ограничения времени
		- Предотвращает зависание

	http.RedirectHandler - обработчик редиректа
		- Выполняет редирект на другой URL
		- Используется для перенаправлений
		- Поддерживает различные коды редиректа

	http.StripPrefix - обработчик с удалением префикса
		- Удаляет префикс из URL
		- Используется для статических файлов
		- Позволяет обслуживать файлы из подпапки

	http.CanonicalHost - обработчик канонического хоста
		- Перенаправляет на канонический хост
		- Используется для SEO
		- Предотвращает дублирование контента

	http.SingleHost - обработчик одного хоста
		- Ограничивает доступ к одному хосту
		- Используется для безопасности
		- Предотвращает атаки на другие хосты

	Особенности работы:
	- Handler является интерфейсом
	- ServeHTTP должен обработать запрос полностью
	- Обработчик может быть вложенным
	- Middleware реализуется как Handler
	- Обработчики могут быть объединены

	Примеры использования:

	// Простой обработчик
	func simpleHandler(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}

	// Использование HandlerFunc
	handler := http.HandlerFunc(simpleHandler)
	http.Handle("/", handler)

	// Обработчик с параметрами
	func paramHandler(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	}

	// Обработчик с методами
	func methodHandler(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("GET request"))
		case "POST":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("POST request"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}

	// Обработчик с JSON
	func jsonHandler(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "Hello",
			"status":  "success",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}

	// Обработчик с ошибками
	func errorHandler(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/error" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}

	// Обработчик с редиректом
	func redirectHandler(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-location", http.StatusMovedPermanently)
	}

	// Обработчик с cookies
	func cookieHandler(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:  "session",
			Value: "abc123",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cookie set"))
	}

	// Обработчик с аутентификацией
	func authHandler(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "password" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authenticated"))
	}

	// Обработчик с CORS
	func corsHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("CORS enabled"))
	}

	// Обработчик с логированием
	func loggingHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		})
	}

	// Обработчик с восстановлением
	func recoveryHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Panic recovered: %v", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с ограничением скорости
	func rateLimitHandler(next http.Handler) http.Handler {
		rateLimiter := make(chan struct{}, 10)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case rateLimiter <- struct{}{}:
				defer func() { <-rateLimiter }()
				next.ServeHTTP(w, r)
			default:
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
			}
		})
	}

	// Обработчик с сжатием
	func compressionHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				w.Header().Set("Content-Encoding", "gzip")
				gz := gzip.NewWriter(w)
				defer gz.Close()
				w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}
			}
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с кэшированием
	func cacheHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "public, max-age=3600")
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с безопасностью
	func securityHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с метриками
	func metricsHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)
			
			// Отправка метрик
			fmt.Printf("Request: %s %s, Duration: %v\n", r.Method, r.URL.Path, duration)
		})
	}

	// Обработчик с контекстом
	func contextHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "user", "admin")
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с валидацией
	func validationHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > 1024*1024 { // 1MB
				http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с версионированием
	func versionHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			version := r.Header.Get("API-Version")
			if version == "" {
				version = "v1"
			}
			w.Header().Set("API-Version", version)
			next.ServeHTTP(w, r)
		})
	}

	// Обработчик с локализацией
	func localizationHandler(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lang := r.Header.Get("Accept-Language")
			if lang == "" {
				lang = "en"
			}
			w.Header().Set("Content-Language", lang)
			next.ServeHTTP(w, r)
		})
	}

	=== ОБРАБОТКА ОШИБОК ===

	1. Ошибки обработки:
		- Паника в обработчике
		- Ошибки записи ответа
		- Таймауты обработки

	2. Ошибки валидации:
		- Невалидные параметры
		- Неподдерживаемые методы
		- Ошибки аутентификации

	3. Ошибки ресурсов:
		- Недоступные ресурсы
		- Ошибки базы данных
		- Ошибки внешних сервисов

	=== ПРОИЗВОДИТЕЛЬНОСТЬ ===

	1. Middleware:
		- Используйте middleware для общей логики
		- Минимизируйте количество middleware
		- Оптимизируйте порядок middleware

	2. Обработка:
		- Обрабатывайте запросы быстро
		- Используйте пулы ресурсов
		- Кэшируйте часто используемые данные

	3. Память:
		- Ограничивайте использование памяти
		- Используйте потоковую обработку
		- Освобождайте ресурсы

	=== БЕЗОПАСНОСТЬ ===

	1. Валидация:
		- Проверяйте все входящие данные
		- Ограничивайте размер запросов
		- Валидируйте параметры

	2. Аутентификация:
		- Проверяйте учетные данные
		- Используйте безопасные методы
		- Логируйте попытки доступа

	3. Заголовки:
		- Устанавливайте заголовки безопасности
		- Используйте CORS правильно
		- Добавляйте защиту от атак

	=== МОНИТОРИНГ ===

	1. Логирование:
		- Логируйте все запросы
		- Отслеживайте ошибки
		- Анализируйте паттерны

	2. Метрики:
		- Считайте количество запросов
		- Измеряйте время обработки
		- Мониторьте использование ресурсов

	3. Отладка:
		- Используйте middleware для отладки
		- Логируйте параметры запросов
		- Анализируйте ошибки
*/
