package net_http

/*
	Структура Request

	type Request struct {
		Method string
		URL    *url.URL
		Proto      string // "HTTP/1.0"
		ProtoMajor int    // 1
		ProtoMinor int    // 0
		Header Header
		Body io.ReadCloser
		GetBody func() (io.ReadCloser, error)
		ContentLength int64
		TransferEncoding []string
		Close bool
		Host string
		Form url.Values
		PostForm url.Values
		MultipartForm *multipart.Form
		Trailer Header
		RemoteAddr string
		RequestURI string
		TLS *tls.ConnectionState
		Cancel <-chan struct{}
		Response *Response
		ctx context.Context
	}

	Request представляет HTTP запрос. Это основная структура для
	представления входящих HTTP запросов на сервере и исходящих
	HTTP запросов на клиенте.

	Основные поля:

	Method string
		- HTTP метод запроса ("GET", "POST", "PUT", "DELETE")
		- Определяет тип операции с ресурсом
		- Стандартные методы: GET, POST, PUT, DELETE, HEAD, OPTIONS

	URL *url.URL
		- URL запроса
		- Содержит схему, хост, путь, параметры запроса
		- Используется для определения целевого ресурса

	Proto string
		- Версия HTTP протокола ("HTTP/1.1", "HTTP/2.0")
		- Определяет возможности протокола
		- Влияет на обработку запроса

	ProtoMajor int
		- Мажорная версия HTTP протокола (1, 2)
		- Используется для определения версии протокола
		- Влияет на обработку запроса

	ProtoMinor int
		- Минорная версия HTTP протокола (0, 1)
		- Используется для определения версии протокола
		- Влияет на обработку запроса

	Header Header
		- HTTP заголовки запроса
		- Содержит метаданные запроса
		- Используется для передачи дополнительной информации

	Body io.ReadCloser
		- Тело запроса
		- Содержит данные запроса
		- Может быть nil для запросов без тела

	GetBody func() (io.ReadCloser, error)
		- Функция для получения тела запроса
		- Используется для повторного чтения тела
		- Позволяет переиспользовать запрос

	ContentLength int64
		- Длина содержимого запроса
		- Устанавливается автоматически
		- Используется для валидации запроса

	TransferEncoding []string
		- Методы кодирования передачи
		- Содержит информацию о кодировании
		- Используется для обработки тела запроса

	Close bool
		- Флаг закрытия соединения
		- Определяет нужно ли закрывать соединение
		- Используется для управления соединением

	Host string
		- Имя хоста запроса
		- Содержит имя хоста из заголовка Host
		- Используется для виртуальных хостов

	Form url.Values
		- Параметры формы из URL
		- Содержит параметры запроса
		- Используется для обработки GET параметров

	PostForm url.Values
		- Параметры формы из тела запроса
		- Содержит параметры POST запроса
		- Используется для обработки POST данных

	MultipartForm *multipart.Form
		- Multipart форма
		- Содержит данные multipart запроса
		- Используется для загрузки файлов

	Trailer Header
		- Заголовки в конце сообщения
		- Содержит дополнительные заголовки
		- Используется для передачи метаданных

	RemoteAddr string
		- Адрес клиента
		- Содержит IP адрес клиента
		- Используется для логирования и безопасности

	RequestURI string
		- URI запроса
		- Содержит полный URI запроса
		- Используется для обработки запроса

	TLS *tls.ConnectionState
		- Состояние TLS соединения
		- Содержит информацию о SSL/TLS
		- Используется для HTTPS запросов

	Cancel <-chan struct{}
		- Канал для отмены запроса
		- Используется для отмены запроса
		- Устарел, используйте context

	Response *Response
		- Ответ на запрос
		- Содержит ответ сервера
		- Используется для клиентских запросов

	ctx context.Context
		- Контекст запроса
		- Содержит контекст для отмены
		- Используется для управления жизненным циклом

	Основные методы:

	func (r *Request) Context() context.Context
		- Возвращает контекст запроса
		- Используется для отмены запроса
		- Позволяет управлять жизненным циклом

	func (r *Request) WithContext(ctx context.Context) *Request
		- Создает новый запрос с контекстом
		- ctx: новый контекст
		- Возвращает новый запрос
		- Используется для обновления контекста

	func (r *Request) Clone(ctx context.Context) *Request
		- Клонирует запрос
		- ctx: контекст для нового запроса
		- Возвращает копию запроса
		- Используется для создания копий

	func (r *Request) ParseForm() error
		- Парсит форму из URL и тела запроса
		- Заполняет поля Form и PostForm
		- Возвращает ошибку если парсинг не удался
		- Используется для обработки форм

	func (r *Request) ParseMultipartForm(maxMemory int64) error
		- Парсит multipart форму
		- maxMemory: максимальный размер памяти
		- Заполняет поле MultipartForm
		- Возвращает ошибку если парсинг не удался
		- Используется для загрузки файлов

	func (r *Request) FormValue(key string) string
		- Возвращает значение параметра формы
		- key: ключ параметра
		- Возвращает первое значение параметра
		- Используется для получения параметров

	func (r *Request) PostFormValue(key string) string
		- Возвращает значение POST параметра
		- key: ключ параметра
		- Возвращает первое значение параметра
		- Используется для получения POST данных

	func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		- Возвращает файл из multipart формы
		- key: ключ файла
		- Возвращает файл и его заголовок
		- Используется для загрузки файлов

	func (r *Request) BasicAuth() (username, password string, ok bool)
		- Возвращает учетные данные Basic Auth
		- Возвращает имя пользователя и пароль
		- ok: успешность извлечения
		- Используется для аутентификации

	func (r *Request) SetBasicAuth(username, password string)
		- Устанавливает учетные данные Basic Auth
		- username: имя пользователя
		- password: пароль
		- Используется для аутентификации

	func (r *Request) UserAgent() string
		- Возвращает User-Agent заголовок
		- Возвращает строку User-Agent
		- Используется для определения клиента

	func (r *Request) Referer() string
		- Возвращает Referer заголовок
		- Возвращает URL источника
		- Используется для отслеживания переходов

	func (r *Request) Write(w io.Writer) error
		- Записывает запрос в Writer
		- w: Writer для записи
		- Возвращает ошибку если запись не удалась
		- Используется для отладки

	func (r *Request) WriteProxy(w io.Writer) error
		- Записывает запрос в Writer для прокси
		- w: Writer для записи
		- Возвращает ошибку если запись не удалась
		- Используется для прокси

	Особенности работы:
	- Request является основной структурой для HTTP запросов
	- Все поля доступны для чтения и записи
	- Методы предоставляют удобный доступ к данным
	- Контекст используется для управления жизненным циклом
	- Формы парсятся автоматически при необходимости

	Примеры использования:

	// Создание запроса
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Настройка заголовков
	req.Header.Set("User-Agent", "MyApp/1.0")
	req.Header.Set("Accept", "application/json")

	// Получение контекста
	ctx := req.Context()

	// Создание запроса с контекстом
	req = req.WithContext(context.Background())

	// Клонирование запроса
	newReq := req.Clone(context.Background())

	// Парсинг формы
	err = req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Получение параметров формы
	value := req.FormValue("key")
	postValue := req.PostFormValue("key")

	// Парсинг multipart формы
	err = req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		log.Fatal(err)
	}

	// Получение файла
	file, handler, err := req.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Basic Auth
	username, password, ok := req.BasicAuth()
	if ok {
		fmt.Printf("User: %s, Password: %s\n", username, password)
	}

	// Установка Basic Auth
	req.SetBasicAuth("user", "password")

	// Получение User-Agent
	userAgent := req.UserAgent()

	// Получение Referer
	referer := req.Referer()

	// Запись запроса
	err = req.Write(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	// Запись для прокси
	err = req.WriteProxy(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	=== ОБРАБОТКА ОШИБОК ===

	1. Ошибки создания запроса:
		- Невалидный URL
		- Неподдерживаемый метод
		- Ошибки парсинга

	2. Ошибки парсинга формы:
		- Невалидные данные формы
		- Превышение лимитов
		- Ошибки кодирования

	3. Ошибки multipart:
		- Невалидные multipart данные
		- Превышение лимитов памяти
		- Ошибки загрузки файлов

	=== ПРОИЗВОДИТЕЛЬНОСТЬ ===

	1. Парсинг форм:
		- Парсите формы только при необходимости
		- Используйте FormValue для быстрого доступа
		- Ограничивайте размер multipart форм

	2. Контекст:
		- Используйте контекст для отмены запросов
		- Устанавливайте таймауты
		- Отменяйте запросы при необходимости

	3. Память:
		- Ограничивайте размер тела запроса
		- Используйте потоковое чтение
		- Освобождайте ресурсы

	=== БЕЗОПАСНОСТЬ ===

	1. Валидация данных:
		- Проверяйте все входящие данные
		- Ограничивайте размер запросов
		- Валидируйте параметры

	2. Аутентификация:
		- Проверяйте учетные данные
		- Используйте безопасные методы
		- Логируйте попытки доступа

	3. Загрузка файлов:
		- Ограничивайте размер файлов
		- Проверяйте типы файлов
		- Сканируйте на вирусы

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
		- Используйте Write для отладки
		- Логируйте заголовки
		- Анализируйте ошибки
*/
