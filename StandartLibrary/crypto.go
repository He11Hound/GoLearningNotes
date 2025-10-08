package StandartLibrary

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

/*
Пакет crypto

Современные веб-приложения немыслимы без надёжной системы безопасности.
В Go для этого есть расширенная стандартная библиотека golang.org/x/crypto, которая предоставляет мощные инструменты для защиты данных.

В этом уроке вы изучите три ключевых пакета:
	crypto/acme/autocert для автоматического получения HTTPS-сертификатов,
	crypto/bcrypt для безопасного хеширования паролей,
	crypto/scrypt для генерации криптографических ключей.

Базовое использование

В простейшем случае достаточно использовать функцию NewListener. Она создаёт слушателя, который автоматически обрабатывает HTTPS-соединения и управляет сертификатами:
*/

//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello HTTPS!"))
//	})
//	// autocert.NewListener возвращает net.Listener,
//	// который может работать с соединениями tls.Conn и сертификатами Let's Encrypt
//	http.Serve(autocert.NewListener("mysite.ru"), mux)
//
//}

//Пакет crypto/bcrypt

//Адаптивность позволяет устанавливать нужную скорость работы хеш-функции через параметр cost (стоимость).
//Чем выше значение cost, тем дольше выполняется хеширование, и тем труднее реализовать атаку перебором.
//Значения cost лучше подбирать от 4 до 31(макс)
//В продакшене стоит экспериментально подобрать cost так, чтобы хеширование занимало 100–500 миллисекунд на целевом оборудовании.

//Важная особенность bcrypt — использование так называемой соли (salt). Соль — это случайные данные, которые добавляются к паролю перед хешированием.
//Это защищает от атак с применением радужных таблиц (rainbow tables) — предвычисленных таблиц хешей для популярных паролей.
//То есть хэш нам всегда будет возвращаться разный

func ExampleHashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(passwordBytes), nil
}

func CheckHashOfPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) // вернёт nil если всё ок
}

//Пакет crypto/scrypt

//Scrypt разработан как memory-hard функция — она требует значительного объёма оперативной памяти для вычислений. Это усложняет создание специализированного оборудования (ASIC) для атак брутфорсом, поскольку память стоит дорого.
//С рекомендованными настройками scrypt использует более сложные хеш-функции, чем bcrypt. Он работает дольше, сильнее загружает процессор, требует больше RAM, но обеспечивает лучшую криптозащиту.
//Алгоритм scrypt принимает несколько параметров:
//N — параметр стоимости CPU или памяти, должен быть степенью двойки;
//r — размер блока обычно 8;
//p — параметр параллелизации, 1;
//keyLen — желаемая длина выходного ключа в байтах.

func ScryptExample() {
	// Генерируем криптографически стойкую соль
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	password := "very secret password"

	// Параметры сложности: N=32768, r=8, p=1
	// Длина ключа: 32 байта (256 бит)
	key, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Для читаемости перекодируем результаты в Base64
	fmt.Printf("Salt: %s\n", base64.StdEncoding.EncodeToString(salt))
	fmt.Printf("Key:  %s\n", base64.StdEncoding.EncodeToString(key))

	// Проверим, что тот же пароль даёт тот же ключ
	key2, _ := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	fmt.Printf("Same: %t\n", string(key) == string(key2))
}

// При использовании scrypt для аутентификации важно сохранять не только результат, но и все параметры, использованные для его получения. Обычно создают структуру вида:
type ScryptHash struct {
	N      int    `json:"n"`
	R      int    `json:"r"`
	P      int    `json:"p"`
	Salt   []byte `json:"salt"`
	Hash   []byte `json:"hash"`
	KeyLen int    `json:"keylen"`
}
