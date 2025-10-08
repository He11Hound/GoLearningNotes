package net_http

import (
	"io"
	"net/http"
	"time"
)

/*
	Структура Response

	type Response struct {
		Status     string // e.g. "200 OK"
		StatusCode int    // e.g. 200
		Proto      string // e.g. "HTTP/1.0"
		ProtoMajor int    // e.g. 1
		ProtoMinor int    // e.g. 0
		Header Header
		Body io.ReadCloser
		ContentLength int64
		TransferEncoding []string
		Close bool
		Uncompressed bool
		Trailer Header
		Request *Request
		TLS *tls.ConnectionState
	}

	Response представляет HTTP ответ. Это основная структура для
	представления исходящих HTTP ответов на сервере и входящих
	HTTP ответов на клиенте.

	Основные поля:

	Status string
		- Статус ответа ("200 OK", "404 Not Found")
		- Содержит код статуса и сообщение
		- Используется для отображения статуса

	StatusCode int
		- Код статуса ответа (200, 404, 500)
		- Числовой код статуса
		- Используется для программной обработки

	Proto string
		- Версия HTTP протокола ("HTTP/1.1", "HTTP/2.0")
		- Определяет возможности протокола
		- Влияет на обработку ответа

	ProtoMajor int
		- Мажорная версия HTTP протокола (1, 2)
		- Используется для определения версии протокола
		- Влияет на обработку ответа

	ProtoMinor int
		- Минорная версия HTTP протокола (0, 1)
		- Используется для определения версии протокола
		- Влияет на обработку ответа

	Header Header
		- HTTP заголовки ответа
		- Содержит метаданные ответа
		- Используется для передачи дополнительной информации

	Body io.ReadCloser
		- Тело ответа
		- Содержит данные ответа
		- Может быть nil для ответов без тела

	ContentLength int64
		- Длина содержимого ответа
		- Устанавливается автоматически
		- Используется для валидации ответа

	TransferEncoding []string
		- Методы кодирования передачи
		- Содержит информацию о кодировании
		- Используется для обработки тела ответа

	Close bool
		- Флаг закрытия соединения
		- Определяет нужно ли закрывать соединение
		- Используется для управления соединением

	Uncompressed bool
		- Флаг сжатия
		- Определяет сжато ли содержимое
		- Используется для обработки сжатых данных

	Trailer Header
		- Заголовки в конце сообщения
		- Содержит дополнительные заголовки
		- Используется для передачи метаданных

	Request *Request
		- Запрос, который вызвал ответ
		- Содержит информацию о запросе
		- Используется для связи запроса и ответа

	TLS *tls.ConnectionState
		- Состояние TLS соединения
		- Содержит информацию о SSL/TLS
		- Используется для HTTPS ответов

	Основные методы:

	func (r *Response) Cookies() []*Cookie
		- Возвращает cookies из ответа
		- Возвращает список cookies
		- Используется для обработки cookies

	func (r *Response) Location() (*url.URL, error)
		- Возвращает URL редиректа
		- Возвращает URL из заголовка Location
		- Используется для обработки редиректов

	func (r *Response) Write(w io.Writer) error
		- Записывает ответ в Writer
		- w: Writer для записи
		- Возвращает ошибку если запись не удалась
		- Используется для отладки

	func (r *Response) WriteProxy(w io.Writer) error
		- Записывает ответ в Writer для прокси
		- w: Writer для записи
		- Возвращает ошибку если запись не удалась
		- Используется для прокси

	func (r *Response) Dump() ([]byte, error)
		- Возвращает дамп ответа
		- Возвращает байты ответа
		- Используется для отладки

	func (r *Response) DumpResponse() ([]byte, error)
		- Возвращает дамп ответа
		- Возвращает байты ответа
		- Используется для отладки

	func (r *Response) Close() error
		- Закрывает тело ответа
		- Освобождает ресурсы
		- Используется для очистки

	func (r *Response) Read(p []byte) (n int, err error)
		- Читает данные из тела ответа
		- p: буфер для чтения
		- n: количество прочитанных байт
		- err: ошибка если чтение не удалось
		- Реализует io.Reader интерфейс

	func (r *Response) ReadAll() ([]byte, error)
		- Читает все данные из тела ответа
		- Возвращает все байты ответа
		- Используется для получения полного ответа

	func (r *Response) UnmarshalJSON(v interface{}) error
		- Парсит JSON из тела ответа
		- v: структура для парсинга
		- Возвращает ошибку если парсинг не удался
		- Используется для обработки JSON

	func (r *Response) UnmarshalXML(v interface{}) error
		- Парсит XML из тела ответа
		- v: структура для парсинга
		- Возвращает ошибку если парсинг не удался
		- Используется для обработки XML

	Особенности работы:
	- Response является основной структурой для HTTP ответов
	- Все поля доступны для чтения и записи
	- Методы предоставляют удобный доступ к данным
	- Тело ответа должно быть закрыто после использования
	- Cookies извлекаются автоматически

	Примеры использования:

	// Создание ответа
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("Hello, World!")),
	}

	// Настройка заголовков
	resp.Header.Set("Content-Type", "text/plain")
	resp.Header.Set("Content-Length", "13")

	// Получение cookies
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie: %s=%s\n", cookie.Name, cookie.Value)
	}

	// Получение URL редиректа
	location, err := resp.Location()
	if err == nil {
		fmt.Printf("Redirect to: %s\n", location.String())
	}

	// Запись ответа
	err = resp.Write(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	// Запись для прокси
	err = resp.WriteProxy(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	// Дамп ответа
	dump, err := resp.Dump()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response dump: %s\n", string(dump))

	// Закрытие ответа
	err = resp.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Чтение данных
	buffer := make([]byte, 1024)
	n, err := resp.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))

	// Чтение всех данных
	body, err := resp.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Full body: %s\n", string(body))

	// Парсинг JSON
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var user User
	err = resp.UnmarshalJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %+v\n", user)

	// Парсинг XML
	type Data struct {
		Value string `xml:"value"`
	}
	var data Data
	err = resp.UnmarshalXML(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", data)

	// Обработка статус кодов
	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("Success")
	case http.StatusNotFound:
		fmt.Println("Not found")
	case http.StatusInternalServerError:
		fmt.Println("Server error")
	default:
		fmt.Printf("Status: %d\n", resp.StatusCode)
	}

	// Проверка заголовков
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		fmt.Println("JSON response")
	}

	// Обработка сжатия
	if resp.Uncompressed {
		fmt.Println("Response is compressed")
	}

	// Обработка редиректов
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location, err := resp.Location()
		if err == nil {
			fmt.Printf("Redirect to: %s\n", location.String())
		}
	}

	=== ОБРАБОТКА ОШИБОК ===

	1. Ошибки чтения:
		- io.EOF: достигнут конец потока
		- net.Error: сетевая ошибка
		- Ошибки декодирования

	2. Ошибки парсинга:
		- Невалидный JSON
		- Невалидный XML
		- Ошибки кодирования

	3. Ошибки закрытия:
		- Ошибки закрытия тела
		- Ошибки освобождения ресурсов

	=== ПРОИЗВОДИТЕЛЬНОСТЬ ===

	1. Чтение данных:
		- Используйте буферизованное чтение
		- Ограничивайте размер буферов
		- Используйте потоковое чтение

	2. Парсинг:
		- Парсите только необходимые данные
		- Используйте streaming парсеры
		- Ограничивайте размер данных

	3. Память:
		- Закрывайте тело ответа
		- Освобождайте ресурсы
		- Используйте пулы буферов

	=== БЕЗОПАСНОСТЬ ===

	1. Валидация данных:
		- Проверяйте все входящие данные
		- Ограничивайте размер ответов
		- Валидируйте заголовки

	2. Обработка ошибок:
		- Всегда обрабатывайте ошибки
		- Логируйте ошибки для отладки
		- Не передавайте ошибки пользователю

	3. Ресурсы:
		- Ограничивайте использование ресурсов
		- Мониторьте количество соединений
		- Используйте таймауты

	=== МОНИТОРИНГ ===

	1. Логирование:
		- Логируйте все ответы
		- Отслеживайте ошибки
		- Анализируйте паттерны

	2. Метрики:
		- Считайте количество ответов
		- Измеряйте время обработки
		- Мониторьте использование ресурсов

	3. Отладка:
		- Используйте Dump для отладки
		- Логируйте заголовки
		- Анализируйте ошибки
*/
