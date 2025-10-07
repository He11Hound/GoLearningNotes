package HTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Alice Johnson", Email: "alice@example.com"},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)

		return
	}

	// Генерация ID для нового пользователя (просто увеличение на 1)
	newUser.ID = fmt.Sprint(len(users) + 1)

	// Добавление нового пользователя в список
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/users/"):] // например, "2"

	for i, u := range users {
		if u.ID == id {
			var updatedUser User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				http.Error(w, "Invalid user data", http.StatusBadRequest)
				return
			}

			// сохраняем прежний ID, чтобы не потерять его
			updatedUser.ID = id
			users[i] = updatedUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Получение ID пользователя из URL
	id := r.URL.Path[len("/api/v1/users/"):]

	// Удаление пользователя по ID
	for i, u := range users {
		if u.ID == id {
			// Удаление элемента из среза
			users = append(users[:i], users[i+1:]...)

			w.WriteHeader(http.StatusNoContent)

			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func ProductionServer() {
	mux := http.NewServeMux()

	// Создаем эндпоинт для получения и создания пользователей.
	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUsers(w, r)
		case http.MethodPost:
			CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Создаем эндпоинт для изменения и удаления пользователя. ID извлекается из пути URL.
	mux.HandleFunc("/api/v1/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			UpdateUser(w, r)
		case http.MethodDelete:
			DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8383", mux)
}

/*
func TestCreateUser(t *testing.T) {
	// Создание тестового сервера
	server := httptest.NewServer(http.HandlerFunc(CreateUser))
	defer server.Close()

	// Тестовый JSON для создания пользователя
	newUser := User{Name: "New User", Email: "new@example.com"}
	newUserJSON, _ := json.Marshal(newUser)

	// Создание POST-запроса
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(newUserJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправка POST-запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверка статус кода
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	// Проверка, что новый пользователь был создан
	var createdUser User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == "" {
		t.Error("Expected a valid ID for the created user")
	}

	// Дополнительные проверки, например, можно проверить, что созданный пользователь соответствует ожиданиям
}

func TestUpdateUser(t *testing.T) {
	// Создание тестового сервера
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			UpdateUser(w, r)
		}
	}))
	defer server.Close()

	// Тестовый JSON для обновления пользователя
	updatedUser := User{ID: "1", Name: "Updated User", Email: "updated@example.com"}
	updatedUserJSON, _ := json.Marshal(updatedUser)

	// Создание PUT-запроса
	req, err := http.NewRequest(http.MethodPut, server.URL+"/api/users/1", bytes.NewBuffer(updatedUserJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправка PUT-запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверка статус кода
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Проверка, что пользователь был обновлен
	var updatedUserResponse User
	err = json.NewDecoder(resp.Body).Decode(&updatedUserResponse)
	if err != nil {
		t.Fatal(err)
	}

	if updatedUserResponse.ID != "1" || updatedUserResponse.Name != "Updated User" ||
		updatedUserResponse.Email != "updated@example.com" {
		t.Error("User update response does not match expectations")
	}
}

func TestDeleteUser(t *testing.T) {
	// Создание тестового сервера
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			DeleteUser(w, r)
		}
	}))
	defer server.Close()

	// Создание DELETE-запроса
	req, err := http.NewRequest("DELETE", server.URL+"/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Отправка DELETE-запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверка статус кода
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", resp.StatusCode)
	}

	// Проверка, что пользователь с ID "1" был удален
	// Для более сложных проверок, можно предварительно создать пользователя с ID "1" и потом попытаться получить его, чтобы убедиться, что его больше нет
}
*/
