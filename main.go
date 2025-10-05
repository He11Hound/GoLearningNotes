package main

import (
	"learning/Multithreading"
)

//go:generate go run tools/genconstants.go "Version" "1.2.3" "Name" "myapp" "Title" "My Application"

//Команда запускает файл genconstants.go и передаёт параметры

func main() {
	Multithreading.SelectChannelExample()
}
