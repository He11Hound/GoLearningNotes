package HTTP

import (
	"fmt"
	"net/http"
)

/*
Реализация HTTP-сервера на Go включает в себя сам сервер, который слушает порт и принимает запросы, поступающие от HTTP-клиентов, и одну или несколько функций-обработчиков, которые обрабатывают эти запросы.
Функции-обработчики называются хендлерами (handlers).
*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//обязательные 2 параметра, так как гошка сама вызывает под капотом функцию переданную в хэндлере с этими параметрами
	fmt.Println("Hello, world!") //печатает в консоль при запросе
}

type HandlerStruct struct{}

// метод структуры для регистрации обработчика
func (handlerStruct HandlerStruct) ServeHTTP(write http.ResponseWriter, _ *http.Request) {

	/*
		Функции для работы с заголовками ответа
		write.Header().Add("Content-Type", "text/html; charset=utf-8")
		write.Header().Set("Content-Type", "json")
		write.Header().Del("Content-Type")

		Передача статуса ответа - вызывать СТРОГО после работы с заголовками
		write.WriteHeader(http.StatusOK)

		Если что-то явно не так возвращаем
		http.Error(write, err.Error(), http.StatusInternalServerError)
	*/

	write.Write([]byte("BYTE BYTE BYTE"))

	//fmt.Fprintln(write, "Hello, world!") //поток вывода на страницу
}

func ExampleSimpleHttpServer() {
	var hs HandlerStruct

	http.HandleFunc("/", helloHandler)
	//регистрация отбработчика через функцию
	//первый параметр - путь
	//второй - функция обработчик

	http.Handle("/test", hs)
	//регистрация отбработчика через структуру

	err := http.ListenAndServe(":8383", nil)
	//Запускает http-server,
	//первый параметр - порт, который слушается
	if err != nil {
		panic(err)
	}
}

/*
	--Handler через создание структур--


	Тип http.Handler — это интерфейсный тип с единственной функцией ServeHTTP(...)
	Она будет вызвана для обработки любого HTTP-запроса.

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

	Обработчику нужно передать:
	http.ResponseWriter — интерфейс потоковой записи, куда обработчик может писать ответные данные для клиента;
	http.Request представляет собой структуру данных, которая содержит всю информацию о входящем HTTP-запросе.


	type ResponseWriter interface {
		Header() Header
		Write([]byte) (int, error)
		WriteHeader(statusCode int)
	}

	Здесь методы Header и WriteHеader используются для работы с заголовками,
	а метод Write записывает тело ответа:

	Второй параметр — типа *Request — это указатель на структуру, которая содержит информацию о заголовках HTTP-запроса и данные, отправленные клиентом.
*/

/*
Создание и настройка HTTP-сервера
В Go можно создать HTTP-сервер и настроить его. Для этого в пакете net/http есть тип Server. Вот определение его структуры:

type Server struct {
    Addr string
    Handler Handler // handler to invoke, http.DefaultServeMux if nil
    DisableGeneralOptionsHandler bool
    TLSConfig *tls.Config
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout time.Duration
    IdleTimeout time.Duration
    MaxHeaderBytes int
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
    BaseContext func(net.Listener) context.Context
    ConnContext func(ctx context.Context, c net.Conn) context.Context
    HTTP2 *HTTP2Config
    Protocols *Protocols
}

Запуск своего сервера
func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/home", homeHandler)
    mux.HandleFunc("/about", aboutHandler)

    s := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  time.Duration(5 * time.Second),
        WriteTimeout: time.Duration(10 * time.Second),
    }

    log.Fatal(s.ListenAndServe())
}


*/
