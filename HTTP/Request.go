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
	//request.URL.Query() - Получает сразу всe параметры, записывает в строковую мапу

	//request.PathValue("pathValue1") - получаем значение из подстановочной переменной {} в самом пути хэндлера пишем путь вида /api/v1/test/{pathValue1}

	//request.ParseForm()
	//request.FormValue("param1") - возвращает первое значение параметра с указанным именем. При этом может быть возвращён как POST, так и GET-параметр, но поиск начинается с POST.
	//request.PostFormValue("param1") - отличается от предыдущего тем, что ищет только среди параметров POST-запроса.
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
