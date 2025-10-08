package net_rpc

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args структура для аргументов
type Args struct {
	A, B int
}

// Quotient структура для результата деления
type Quotient struct {
	Quo, Rem int
}

// Arith структура для арифметических операций
type Arith struct{}

// Multiply метод умножения
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide метод деления
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// NewServerExample демонстрирует использование функции rpc.NewServer
func NewServerExample() {
	// Создание нового RPC сервера
	server := rpc.NewServer()

	// Регистрация сервиса
	arith := new(Arith)
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Регистрация HTTP обработчиков
	server.HandleHTTP()

	// Запуск сервера
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ListenAndServeExample демонстрирует использование функции rpc.ListenAndServe
func ListenAndServeExample() {
	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Запуск RPC сервера
	err := rpc.ListenAndServe("tcp", ":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeExample демонстрирует использование функции rpc.Serve
func ServeExample() {
	// Создание слушателя
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Запуск RPC сервера
	err = rpc.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeCodecExample демонстрирует использование функции rpc.ServeCodec
func ServeCodecExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание Gob кодека
	codec := rpc.NewGobServerCodec(conn)

	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Запуск RPC сервера с кодеком
	err = rpc.ServeCodec(codec)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeConnExample демонстрирует использование функции rpc.ServeConn
func ServeConnExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Запуск RPC сервера на соединении
	err = rpc.ServeConn(conn)
	if err != nil {
		log.Fatal(err)
	}
}

// DialExample демонстрирует использование функции rpc.Dial
func DialExample() {
	// Установка соединения с RPC сервером
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// DialHTTPExample демонстрирует использование функции rpc.DialHTTP
func DialHTTPExample() {
	// Установка HTTP соединения с RPC сервером
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// DialHTTPPathExample демонстрирует использование функции rpc.DialHTTPPath
func DialHTTPPathExample() {
	// Установка HTTP соединения с RPC сервером по пути
	client, err := rpc.DialHTTPPath("tcp", "localhost:8080", "/rpc")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// NewClientExample демонстрирует использование функции rpc.NewClient
func NewClientExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание RPC клиента
	client := rpc.NewClient(conn)
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// NewClientWithCodecExample демонстрирует использование функции rpc.NewClientWithCodec
func NewClientWithCodecExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание Gob кодека
	codec := rpc.NewGobClientCodec(conn)

	// Создание RPC клиента с кодеком
	client := rpc.NewClientWithCodec(codec)
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// NewGobServerCodecExample демонстрирует использование функции rpc.NewGobServerCodec
func NewGobServerCodecExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание Gob кодека для сервера
	codec := rpc.NewGobServerCodec(conn)

	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Запуск RPC сервера с Gob кодеком
	err = rpc.ServeCodec(codec)
	if err != nil {
		log.Fatal(err)
	}
}

// NewGobClientCodecExample демонстрирует использование функции rpc.NewGobClientCodec
func NewGobClientCodecExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание Gob кодека для клиента
	codec := rpc.NewGobClientCodec(conn)

	// Создание RPC клиента с Gob кодеком
	client := rpc.NewClientWithCodec(codec)
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// NewJSONServerCodecExample демонстрирует использование функции rpc.NewJSONServerCodec
func NewJSONServerCodecExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание JSON кодека для сервера
	codec := rpc.NewJSONServerCodec(conn)

	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Запуск RPC сервера с JSON кодеком
	err = rpc.ServeCodec(codec)
	if err != nil {
		log.Fatal(err)
	}
}

// NewJSONClientCodecExample демонстрирует использование функции rpc.NewJSONClientCodec
func NewJSONClientCodecExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание JSON кодека для клиента
	codec := rpc.NewJSONClientCodec(conn)

	// Создание RPC клиента с JSON кодеком
	client := rpc.NewClientWithCodec(codec)
	defer client.Close()

	// Вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// HandleHTTPExample демонстрирует использование функции rpc.HandleHTTP
func HandleHTTPExample() {
	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Регистрация HTTP обработчиков
	rpc.HandleHTTP()

	// Запуск HTTP сервера
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleHTTPPathExample демонстрирует использование функции rpc.HandleHTTPPath
func HandleHTTPPathExample() {
	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Регистрация HTTP обработчиков по пути
	rpc.HandleHTTPPath("/rpc")

	// Запуск HTTP сервера
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ServeHTTPExample демонстрирует использование функции rpc.ServeHTTP
func ServeHTTPExample() {
	// Создание сервиса
	arith := new(Arith)
	rpc.Register(arith)

	// Создание HTTP обработчика
	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeHTTP(w, r)
	})

	// Запуск HTTP сервера
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// IsErrorExample демонстрирует использование функции rpc.IsError
func IsErrorExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов несуществующего метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.NonExistent", args, &reply)
	if err != nil {
		// Проверка типа ошибки
		if rpc.IsError(err) {
			fmt.Println("This is an RPC error")
		} else {
			fmt.Println("This is not an RPC error")
		}
	}
}

// IsErrorCodeExample демонстрирует использование функции rpc.IsErrorCode
func IsErrorCodeExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов несуществующего метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.NonExistent", args, &reply)
	if err != nil {
		// Проверка кода ошибки
		if rpc.IsErrorCode(err, 404) {
			fmt.Println("Method not found error")
		} else {
			fmt.Println("Other error")
		}
	}
}

// RegisterExample демонстрирует использование функции rpc.Register
func RegisterExample() {
	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service registered successfully")
}

// RegisterNameExample демонстрирует использование функции rpc.RegisterName
func RegisterNameExample() {
	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса с именем
	err := rpc.RegisterName("Arithmetic", arith)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service registered with name 'Arithmetic'")
}

// UnregisterExample демонстрирует использование функции rpc.Unregister
func UnregisterExample() {
	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Отмена регистрации сервиса
	err = rpc.Unregister("Arith")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service unregistered successfully")
}

// CallExample демонстрирует использование метода client.Call
func CallExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Синхронный вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// GoExample демонстрирует использование метода client.Go
func GoExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Асинхронный вызов метода
	args := &Args{A: 7, B: 8}
	var reply int
	call := client.Go("Arith.Multiply", args, &reply, nil)

	// Ожидание завершения
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatal(replyCall.Error)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// CloseExample демонстрирует использование метода client.Close
func CloseExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Использование клиента
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	// Закрытие клиента
	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Client closed successfully")
}

// ServerRegisterExample демонстрирует использование метода server.Register
func ServerRegisterExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service registered on server")
}

// ServerRegisterNameExample демонстрирует использование метода server.RegisterName
func ServerRegisterNameExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса с именем
	err := server.RegisterName("Arithmetic", arith)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service registered with name 'Arithmetic' on server")
}

// ServerUnregisterExample демонстрирует использование метода server.Unregister
func ServerUnregisterExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Отмена регистрации сервиса
	err = server.Unregister("Arith")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service unregistered from server")
}

// ServerServeCodecExample демонстрирует использование метода server.ServeCodec
func ServerServeCodecExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Создание Gob кодека
	codec := rpc.NewGobServerCodec(conn)

	// Запуск сервера с кодеком
	err = server.ServeCodec(codec)
	if err != nil {
		log.Fatal(err)
	}
}

// ServerServeConnExample демонстрирует использование метода server.ServeConn
func ServerServeConnExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Создание соединения
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера на соединении
	err = server.ServeConn(conn)
	if err != nil {
		log.Fatal(err)
	}
}

// ServerServeHTTPExample демонстрирует использование метода server.ServeHTTP
func ServerServeHTTPExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Создание HTTP обработчика
	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})

	// Запуск HTTP сервера
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ComplexServiceExample демонстрирует сложный сервис
func ComplexServiceExample() {
	// Создание сервера
	server := rpc.NewServer()

	// Создание сервиса
	arith := new(Arith)

	// Регистрация сервиса
	err := server.Register(arith)
	if err != nil {
		log.Fatal(err)
	}

	// Регистрация HTTP обработчиков
	server.HandleHTTP()

	// Запуск сервера
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Ожидание запуска сервера
	time.Sleep(100 * time.Millisecond)

	// Создание клиента
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов метода умножения
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Вызов метода деления
	args = &Args{A: 15, B: 4}
	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
}

// ErrorHandlingExample демонстрирует обработку ошибок
func ErrorHandlingExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Вызов несуществующего метода
	args := &Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.NonExistent", args, &reply)
	if err != nil {
		// Проверка типа ошибки
		if rpc.IsError(err) {
			fmt.Println("RPC error:", err)
		} else {
			fmt.Println("Non-RPC error:", err)
		}
	}

	// Вызов метода с ошибкой
	args = &Args{A: 7, B: 0}
	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		fmt.Println("Division error:", err)
	}
}

// ConcurrentCallsExample демонстрирует параллельные вызовы
func ConcurrentCallsExample() {
	// Создание клиента
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Параллельные вызовы
	for i := 0; i < 5; i++ {
		go func(i int) {
			args := &Args{A: i, B: i + 1}
			var reply int
			err := client.Call("Arith.Multiply", args, &reply)
			if err != nil {
				log.Printf("Error in goroutine %d: %v", i, err)
				return
			}
			fmt.Printf("Goroutine %d: %d*%d=%d\n", i, args.A, args.B, reply)
		}(i)
	}

	// Ожидание завершения
	time.Sleep(1 * time.Second)
}
