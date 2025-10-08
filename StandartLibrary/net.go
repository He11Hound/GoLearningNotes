package StandartLibrary

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
)

//Пакет net/url предназначен для парсинга и манипулирования URL-адресами

func parseUrl() {
	rawURL := "https://example.com:8080/path/to/resource?param1=value1&param2=value2#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("Ошибка парсинга: %v\n", err)
		return
	}

	fmt.Printf("Схема: %s\n", parsedURL.Scheme)         // https
	fmt.Printf("Хост: %s\n", parsedURL.Host)            // example.com:8080
	fmt.Printf("Имя хоста: %s\n", parsedURL.Hostname()) // example.com
	fmt.Printf("Порт: %s\n", parsedURL.Port())          // 8080
	fmt.Printf("Путь: %s\n", parsedURL.Path)            // /path/to/resource
	fmt.Printf("Фрагмент: %s\n", parsedURL.Fragment)    // fragment
}

//Пакет net/html

//Пакет net/html предназначен для парсинга HTML-документов.
//Он строит дерево DOM, которое можно обходить и анализировать для извлечения нужной информации.
//Это особенно полезно для веб-скрейпинга и анализа HTML-контента.

/*
html.Node — это структура, которая представляет узел дерева HTML.

Каждый узел может быть:
	ElementNode (тег, например <h1> или <div>),
	TextNode (текст внутри тега),
	CommentNode и т.д.
FirstChild и NextSibling позволяют обходить дерево рекурсивно.
*/

// findH1 рекурсивно ищет первый h1-тэг в HTML-узле
func findH1(n *html.Node) *html.Node {
	// Проверяем, является ли текущий узел HTML-элементом и равен ли он "h1"
	if n.Type == html.ElementNode && n.Data == "h1" {
		return n
	}
	// Идем по всем дочерним узлам текущего узла
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Рекурсивно ищем <h1> в каждом дочернем узле
		if result := findH1(c); result != nil {
			return result
		}
	}
	return nil
}

func GetSiteTitle() {
	url := "https://chr.rbc.ru" // URL сайта

	// Отправляем GET-запрос к указанному URL
	// net/http выполняет весь процесс TCP-соединения, TLS (если HTTPS),
	// отправку HTTP-запроса и получение ответа
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// Обязательно закрываем тело ответа после использования, чтобы
	// освободить сетевые ресурсы и соединение
	defer resp.Body.Close()

	// html.Parse читает поток HTML из resp.Body
	// и строит дерево узлов (DOM-подобная структура)
	// Каждый узел — это тег, текст или комментарий
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	h1 := findH1(doc)
	if h1 != nil && h1.FirstChild != nil {
		// h1.FirstChild.Data содержит текст внутри тега <h1>
		fmt.Println("H1:", h1.FirstChild.Data)
	} else {
		fmt.Println("H1 тег не найден")
	}
}
