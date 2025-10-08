package net_rpc

/*
	ФУНКЦИИ ПАКЕТА NET/RPC

	Пакет net/rpc предоставляет возможность вызывать методы объектов
	через сеть. Это система удаленного вызова процедур (RPC) для Go.
	Позволяет вызывать методы объектов на удаленных серверах.

	=== ФУНКЦИИ ДЛЯ СОЗДАНИЯ СЕРВЕРА ===

	func NewServer() *Server
		- Создает новый RPC сервер
		- Возвращает указатель на Server
		- Используется для создания RPC серверов
		- Позволяет регистрировать методы и объекты

	func ListenAndServe(network, address string, handler Handler) error
		- Запускает RPC сервер на указанном адресе
		- network: тип сети ("tcp", "tcp4", "tcp6")
		- address: адрес для прослушивания (":8080")
		- handler: обработчик RPC запросов
		- Блокирующая функция, запускает сервер

	func Serve(l net.Listener, handler Handler) error
		- Запускает RPC сервер на существующем слушателе
		- l: сетевой слушатель (net.Listener)
		- handler: обработчик RPC запросов
		- Блокирующая функция, запускает сервер
		- Позволяет использовать кастомные слушатели

	func ServeCodec(codec ServerCodec) error
		- Запускает RPC сервер с кастомным кодеком
		- codec: кодек для сериализации/десериализации
		- Блокирующая функция, запускает сервер
		- Используется для кастомных протоколов

	func ServeConn(conn io.ReadWriteCloser) error
		- Запускает RPC сервер на соединении
		- conn: соединение для RPC
		- Блокирующая функция, запускает сервер
		- Используется для одиночных соединений

	=== ФУНКЦИИ ДЛЯ СОЗДАНИЯ КЛИЕНТА ===

	func Dial(network, address string) (*Client, error)
		- Устанавливает соединение с RPC сервером
		- network: тип сети ("tcp", "tcp4", "tcp6")
		- address: адрес сервера ("localhost:8080")
		- Возвращает Client для вызова методов
		- Используется для подключения к серверу

	func DialHTTP(network, address string) (*Client, error)
		- Устанавливает HTTP соединение с RPC сервером
		- network: тип сети ("tcp", "tcp4", "tcp6")
		- address: адрес сервера ("localhost:8080")
		- Возвращает Client для вызова методов
		- Использует HTTP протокол для RPC

	func DialHTTPPath(network, address, path string) (*Client, error)
		- Устанавливает HTTP соединение с RPC сервером по пути
		- network: тип сети ("tcp", "tcp4", "tcp6")
		- address: адрес сервера ("localhost:8080")
		- path: путь к RPC сервису ("/rpc")
		- Возвращает Client для вызова методов
		- Использует HTTP протокол с кастомным путем

	func NewClient(conn io.ReadWriteCloser) *Client
		- Создает новый RPC клиент
		- conn: соединение для RPC
		- Возвращает Client для вызова методов
		- Используется для кастомных соединений

	func NewClientWithCodec(codec ClientCodec) *Client
		- Создает новый RPC клиент с кастомным кодеком
		- codec: кодек для сериализации/десериализации
		- Возвращает Client для вызова методов
		- Используется для кастомных протоколов

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С КОДЕКАМИ ===

	func NewGobServerCodec(conn io.ReadWriteCloser) ServerCodec
		- Создает Gob кодек для сервера
		- conn: соединение для RPC
		- Возвращает ServerCodec для Gob протокола
		- Используется для сериализации Gob

	func NewGobClientCodec(conn io.ReadWriteCloser) ClientCodec
		- Создает Gob кодек для клиента
		- conn: соединение для RPC
		- Возвращает ClientCodec для Gob протокола
		- Используется для сериализации Gob

	func NewJSONServerCodec(conn io.ReadWriteCloser) ServerCodec
		- Создает JSON кодек для сервера
		- conn: соединение для RPC
		- Возвращает ServerCodec для JSON протокола
		- Используется для сериализации JSON

	func NewJSONClientCodec(conn io.ReadWriteCloser) ClientCodec
		- Создает JSON кодек для клиента
		- conn: соединение для RPC
		- Возвращает ClientCodec для JSON протокола
		- Используется для сериализации JSON

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С HTTP ===

	func HandleHTTP()
		- Регистрирует HTTP обработчики для RPC
		- Использует DefaultServer для обработки
		- Регистрирует пути "/_goRPC_" и "/debug/rpc"
		- Используется для HTTP RPC серверов

	func HandleHTTPPath(path string)
		- Регистрирует HTTP обработчики для RPC по пути
		- path: путь к RPC сервису ("/rpc")
		- Использует DefaultServer для обработки
		- Используется для кастомных путей

	func ServeHTTP(w http.ResponseWriter, r *http.Request)
		- Обрабатывает HTTP запросы для RPC
		- w: ResponseWriter для записи ответа
		- r: Request с параметрами запроса
		- Использует DefaultServer для обработки
		- Используется для HTTP RPC серверов

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ОШИБКАМИ ===

	func IsError(err error) bool
		- Проверяет является ли ошибка RPC ошибкой
		- err: ошибка для проверки
		- Возвращает true если это RPC ошибка
		- Используется для проверки типа ошибки

	func IsErrorCode(err error, code int) bool
		- Проверяет является ли ошибка RPC ошибкой с кодом
		- err: ошибка для проверки
		- code: код ошибки для проверки
		- Возвращает true если это RPC ошибка с кодом
		- Используется для проверки конкретных ошибок

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С МЕТОДАМИ ===

	func Register(rcvr interface{}) error
		- Регистрирует объект в DefaultServer
		- rcvr: объект для регистрации
		- Возвращает ошибку если регистрация не удалась
		- Используется для регистрации методов

	func RegisterName(name string, rcvr interface{}) error
		- Регистрирует объект с именем в DefaultServer
		- name: имя для регистрации
		- rcvr: объект для регистрации
		- Возвращает ошибку если регистрация не удалась
		- Используется для регистрации методов с именем

	func Unregister(name string) error
		- Отменяет регистрацию объекта в DefaultServer
		- name: имя объекта для отмены регистрации
		- Возвращает ошибку если отмена не удалась
		- Используется для отмены регистрации методов

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С СЕРВЕРОМ ===

	func (s *Server) Register(rcvr interface{}) error
		- Регистрирует объект в сервере
		- rcvr: объект для регистрации
		- Возвращает ошибку если регистрация не удалась
		- Используется для регистрации методов

	func (s *Server) RegisterName(name string, rcvr interface{}) error
		- Регистрирует объект с именем в сервере
		- name: имя для регистрации
		- rcvr: объект для регистрации
		- Возвращает ошибку если регистрация не удалась
		- Используется для регистрации методов с именем

	func (s *Server) Unregister(name string) error
		- Отменяет регистрацию объекта в сервере
		- name: имя объекта для отмены регистрации
		- Возвращает ошибку если отмена не удалась
		- Используется для отмены регистрации методов

	func (s *Server) ServeCodec(codec ServerCodec) error
		- Запускает сервер с кодеком
		- codec: кодек для сериализации/десериализации
		- Блокирующая функция, запускает сервер
		- Используется для кастомных протоколов

	func (s *Server) ServeConn(conn io.ReadWriteCloser) error
		- Запускает сервер на соединении
		- conn: соединение для RPC
		- Блокирующая функция, запускает сервер
		- Используется для одиночных соединений

	func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)
		- Обрабатывает HTTP запросы
		- w: ResponseWriter для записи ответа
		- r: Request с параметрами запроса
		- Используется для HTTP RPC серверов

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С КЛИЕНТОМ ===

	func (c *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
		- Вызывает метод на сервере
		- serviceMethod: имя сервиса и метода ("Service.Method")
		- args: аргументы для метода
		- reply: указатель на результат
		- Возвращает ошибку если вызов не удался
		- Блокирующий вызов

	func (c *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call
		- Асинхронно вызывает метод на сервере
		- serviceMethod: имя сервиса и метода ("Service.Method")
		- args: аргументы для метода
		- reply: указатель на результат
		- done: канал для уведомления о завершении
		- Возвращает Call для отслеживания
		- Неблокирующий вызов

	func (c *Client) Close() error
		- Закрывает соединение с сервером
		- Освобождает ресурсы клиента
		- Возвращает ошибку если закрытие не удалось
		- Используется для очистки ресурсов

	=== ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ ===

	// RPC сервер
	type Arith struct{}
	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	http.ListenAndServe(":8080", nil)

	// RPC клиент
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	// Асинхронный вызов
	call := client.Go("Arith.Multiply", args, &reply, nil)
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatal(replyCall.Error)
	}

	// Кастомный сервер
	server := rpc.NewServer()
	server.Register(arith)
	server.HandleHTTP()
	http.ListenAndServe(":8080", nil)

	// JSON RPC
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Gob RPC
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()
*/
