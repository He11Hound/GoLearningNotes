package net_http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ListenAndServeExample демонстрирует использование функции http.ListenAndServe
func ListenAndServeExample() {
	// Создание HTTP сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ListenAndServeTLSExample демонстрирует использование функции http.ListenAndServeTLS
func ListenAndServeTLSExample() {
	// Создание HTTPS сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTPS World!")
	})

	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeExample демонстрирует использование функции http.Serve
func ServeExample() {
	// Создание слушателя
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание HTTP сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from custom listener!")
	})

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeTLSExample демонстрирует использование функции http.ServeTLS
func ServeTLSExample() {
	// Создание слушателя
	listener, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}

	// Создание HTTPS сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from custom HTTPS listener!")
	})

	err = http.ServeTLS(listener, nil, "cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
}

// GetExample демонстрирует использование функции http.Get
func GetExample() {
	// Выполнение GET запроса
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Body: %s\n", string(body))
}

// PostExample демонстрирует использование функции http.Post
func PostExample() {
	// Подготовка данных
	data := `{"name": "John", "age": 30}`
	body := strings.NewReader(data)

	// Выполнение POST запроса
	resp, err := http.Post("https://httpbin.org/post", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(responseBody))
}

// PostFormExample демонстрирует использование функции http.PostForm
func PostFormExample() {
	// Подготовка form данных
	data := url.Values{}
	data.Set("name", "John")
	data.Set("age", "30")

	// Выполнение POST запроса с form данными
	resp, err := http.PostForm("https://httpbin.org/post", data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}

// HeadExample демонстрирует использование функции http.Head
func HeadExample() {
	// Выполнение HEAD запроса
	resp, err := http.Head("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("Content-Length: %s\n", resp.Header.Get("Content-Length"))
}

// DoExample демонстрирует использование функции http.Do
func DoExample() {
	// Создание запроса
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Добавление заголовков
	req.Header.Set("User-Agent", "MyApp/1.0")
	req.Header.Set("Accept", "application/json")

	// Выполнение запроса
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}

// NewRequestExample демонстрирует использование функции http.NewRequest
func NewRequestExample() {
	// Создание нового запроса
	req, err := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader(`{"key": "value"}`))
	if err != nil {
		log.Fatal(err)
	}

	// Настройка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token123")

	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("URL: %s\n", req.URL.String())
	fmt.Printf("Headers: %v\n", req.Header)
}

// NewRequestWithContextExample демонстрирует использование функции http.NewRequestWithContext
func NewRequestWithContextExample() {
	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Создание запроса с контекстом
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
}

// ReadResponseExample демонстрирует использование функции http.ReadResponse
func ReadResponseExample() {
	// Создание запроса
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", string(body))
}

// ServeFileExample демонстрирует использование функции http.ServeFile
func ServeFileExample() {
	// Создание обработчика для отдачи файлов
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeContentExample демонстрирует использование функции http.ServeContent
func ServeContentExample() {
	// Создание обработчика для отдачи содержимого
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		content := strings.NewReader("Hello, World!")
		http.ServeContent(w, r, "hello.txt", time.Now(), content)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// FileServerExample демонстрирует использование функции http.FileServer
func FileServerExample() {
	// Создание файлового сервера
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ProxyURLExample демонстрирует использование функции http.ProxyURL
func ProxyURLExample() {
	// Создание прокси функции
	proxyURL, _ := url.Parse("http://proxy.example.com:8080")
	proxyFunc := http.ProxyURL(proxyURL)

	// Создание клиента с прокси
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxyFunc,
		},
	}

	// Выполнение запроса через прокси
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
}

// ProxyFromEnvironmentExample демонстрирует использование функции http.ProxyFromEnvironment
func ProxyFromEnvironmentExample() {
	// Создание клиента с прокси из переменных окружения
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	// Выполнение запроса
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
}

// SetCookieExample демонстрирует использование функции http.SetCookie
func SetCookieExample() {
	// Создание обработчика с установкой cookie
	http.HandleFunc("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:  "session",
			Value: "abc123",
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Cookie установлен")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// CanonicalHeaderKeyExample демонстрирует использование функции http.CanonicalHeaderKey
func CanonicalHeaderKeyExample() {
	// Преобразование заголовка в канонический формат
	header := "content-type"
	canonical := http.CanonicalHeaderKey(header)
	fmt.Printf("Исходный: %s, Канонический: %s\n", header, canonical)
}

// ParseHTTPVersionExample демонстрирует использование функции http.ParseHTTPVersion
func ParseHTTPVersionExample() {
	// Парсинг версии HTTP
	major, minor, ok := http.ParseHTTPVersion("HTTP/1.1")
	if ok {
		fmt.Printf("HTTP версия: %d.%d\n", major, minor)
	}
}

// ErrorExample демонстрирует использование функции http.Error
func ErrorExample() {
	// Создание обработчика с ошибкой
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// NotFoundExample демонстрирует использование функции http.NotFound
func NotFoundExample() {
	// Создание обработчика с 404 ошибкой
	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// RedirectExample демонстрирует использование функции http.Redirect
func RedirectExample() {
	// Создание обработчика с редиректом
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-location", http.StatusMovedPermanently)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// RequestWithContextExample демонстрирует использование функции http.RequestWithContext
func RequestWithContextExample() {
	// Создание контекста
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Создание запроса
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Добавление контекста к запросу
	req = req.WithContext(ctx)

	// Выполнение запроса
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
}

// CustomClientExample демонстрирует создание кастомного HTTP клиента
func CustomClientExample() {
	// Создание кастомного клиента
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			DisableCompression:  true,
		},
	}

	// Выполнение запроса
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
}

// CustomServerExample демонстрирует создание кастомного HTTP сервера
func CustomServerExample() {
	// Создание кастомного сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from custom server!")
		}),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запуск сервера
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// MiddlewareExample демонстрирует создание middleware
func MiddlewareExample() {
	// Middleware для логирования
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		})
	}

	// Middleware для CORS
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			next.ServeHTTP(w, r)
		})
	}

	// Создание обработчика с middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello with middleware!")
	})

	// Применение middleware
	finalHandler := loggingMiddleware(corsMiddleware(handler))

	http.Handle("/", finalHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// JSONAPIExample демонстрирует создание JSON API
func JSONAPIExample() {
	// Структура для данных
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Обработчик для получения пользователей
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, Name: "John", Age: 30},
			{ID: 2, Name: "Jane", Age: 25},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// Обработчик для создания пользователя
	http.HandleFunc("/api/users/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// FileUploadExample демонстрирует загрузку файлов
func FileUploadExample() {
	// Обработчик для загрузки файлов
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Парсинг multipart form
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Получение файла
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error getting file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded file: %s, Size: %d\n", handler.Filename, handler.Size)
		fmt.Fprintf(w, "File uploaded successfully")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// WebSocketExample демонстрирует работу с WebSocket
func WebSocketExample() {
	// Обработчик для WebSocket
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Здесь должна быть логика WebSocket
		// Для примера просто возвращаем ответ
		fmt.Fprintf(w, "WebSocket endpoint")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// RateLimitingExample демонстрирует ограничение скорости запросов
func RateLimitingExample() {
	// Простое ограничение скорости
	rateLimiter := make(chan struct{}, 10) // Максимум 10 одновременных запросов

	// Middleware для ограничения скорости
	rateLimitMiddleware := func(next http.Handler) http.Handler {
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

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request processed")
	})

	http.Handle("/", rateLimitMiddleware(handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// AuthenticationExample демонстрирует аутентификацию
func AuthenticationExample() {
	// Middleware для аутентификации
	authMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token != "Bearer valid-token" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Authenticated request")
	})

	http.Handle("/protected", authMiddleware(handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// HTTPSRedirectExample демонстрирует принудительное перенаправление на HTTPS
func HTTPSRedirectExample() {
	// Middleware для принудительного HTTPS
	httpsRedirectMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Forwarded-Proto") != "https" {
				httpsURL := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, httpsURL, http.StatusMovedPermanently)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Secure connection")
	})

	http.Handle("/", httpsRedirectMiddleware(handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
