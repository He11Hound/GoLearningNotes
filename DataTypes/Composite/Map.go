package Composite

import (
	"fmt"
)

func MakeMap() {
	// Map — это коллекция "ключ → значение"
	// Все ключи одного типа, все значения одного типа
	// Порядок элементов не гарантирован

	// Полная запись
	var MapExample1 map[string]int = map[string]int{
		"apple":  5,
		"banana": 3,
	}

	// Пустая карта
	var MapExample2 map[string]int // nil, нельзя добавлять элементы, иначе panic

	// Короткая запись
	MapExample3 := map[string]int{
		"cat": 2,
		"dog": 7,
	}

	// Создание карты через make
	MapExample4 := make(map[string]int) // пустая карта
	MapExample4["x"] = 10
	MapExample4["y"] = 20

	fmt.Println("Map 1:", MapExample1)
	fmt.Println("Map 2:", MapExample2) // nil
	fmt.Println("Map 3:", MapExample3)
	fmt.Println("Map 4:", MapExample4)
}

func AccessMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	// Доступ к элементу
	fmt.Println(m["a"]) // 1

	// Проверка наличия ключа
	val, ok := m["c"]
	fmt.Println(val, ok) // 0 false, ключа нет

	// Удаление ключа
	delete(m, "b")
	fmt.Println(m)
}

func IterateMap() {
	m := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}

	// Итерация по карте
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func FunctionsMap() {
	// len(map) - количество элементов
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("Количество элементов:", len(m))

	// nil-map нельзя изменять
	//var m2 map[string]int
	// m2["x"] = 10 // panic
}
