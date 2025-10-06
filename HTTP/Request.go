package HTTP

import "net/http"

//Краткое описание основных методов и параметров структуры request

func ExampleRequest(request *http.Request) {
	//request.Method - возвращает метод
	//request.Header.Get() - получение заголовков запроса
	//request.RemoteAddr - получение IP
	//request.URL.Path - возвращает путь
	//request.URL.Query().Get("parma1") - получение гет параметра
	//request.Cookie("cookie1") - печенье OREO

	//request.ParseForm()
	//request.FormValue('param1') - получение post параметра
}

/*
При работе с методами лучше всего использовать предопределённые константы пакета net/http
	MethodGet     = "GET"
    MethodHead    = "HEAD"
    MethodPost    = "POST"
    MethodPut     = "PUT"
    MethodPatch   = "PATCH"
    MethodDelete  = "DELETE"
    MethodConnect = "CONNECT"
    MethodOptions = "OPTIONS"
    MethodTrace   = "TRACE"
*/
