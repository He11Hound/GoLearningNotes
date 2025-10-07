package Testing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Пакет тестирования встроен в среду выполнения Go и предлагает всё необходимое для написания сложного тестового кода.
Тестовый код размещается в отдельном файле, но в том же пакете, что и основной исходный код.
По конвенции в нём должны находиться файлы, которые заканчиваются на _test.go

handlers.go - handlers_test.go

Пакет net/http/httptest в стандартной библиотеке Go предоставляет инструменты для тестирования HTTP-обработчиков и клиентов. Он используется для создания виртуальных HTTP-серверов и клиентов для тестирования HTTP-решений.
*/

func ExampleServerTest() {
	handeler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	server := httptest.NewServer(handeler)

	defer server.Close()
	//Важно закрывать виртуальный сервер после завершения тестов или другого использования. Это освободит ресурсы и порт, занятый сервером.

	req, err := http.NewRequest("GET", server.URL, nil)
	//httptest.NewRequest — это функция, предназначенная для создания тестовых HTTP-запросов.
	//Отличие httptest.NewRequest от  http.Request состоит в том, что первый в случае ошибки вызывает функцию panic(). Запрос, созданный с помощью пакета net/http/httptest нельзя использовать совместно с виртуальным сервером.

	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func TestMyHandler(t *testing.T) {
	// Создание тестового запроса
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Set("Content-Type", "application/json")

	// Вызов вашего обработчика
	w := httptest.NewRecorder()
	//MyHandler(w, req)

	// Проверка ответа
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
