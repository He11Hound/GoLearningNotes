package Composite

import (
	"fmt"
)

func MakeSlice() {
	// Слайс — это "обёртка" над массивом:
	// хранит указатель на массив, длину и ёмкость
	// можно изменять длину (через append), гибкий по сравнению с массивом
	//При добавление проверяется текущий размер (cap), если месмто есть, то добавляет, если нет, то создаёт новый слайс с cap в 2 раза больше прошлого копирует все элементы и добавляет прошлое значение

	// Полная запись
	var SliceExample1 []int = []int{1, 2, 3}
	// Пустой слайс
	var SliceExample2 []int // nil, длина 0, ёмкость 0
	// Короткая запись
	SliceExample3 := []int{4, 5, 6}

	fmt.Println("Слайс 1:", SliceExample1)
	fmt.Println("Слайс 2:", SliceExample2)
	fmt.Println("Слайс 3:", SliceExample3)

	// Создание слайса с make
	// make([]Type, length, capacity)
	SliceExample4 := make([]int, 3, 5) // длина 3, ёмкость 5
	fmt.Println("Слайс 4:", SliceExample4)
	fmt.Println("Длина:", len(SliceExample4), "Ёмкость:", cap(SliceExample4))

	// Изменение слайса через append
	SliceExample4 = append(SliceExample4, 10, 20)
	fmt.Println("После append:", SliceExample4)
	fmt.Println("Длина:", len(SliceExample4), "Ёмкость:", cap(SliceExample4))

	// Срез слайса
	subSlice := SliceExample4[1:4] // с 1 по 3 индекс
	fmt.Println("Подслайс:", subSlice)
}

func CompareSlices() {
	// Слайсы нельзя сравнивать через ==
	// Исключение: можно сравнивать с nil
	var a []int
	var b []int
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // true

	// Чтобы сравнить содержимое — нужно использовать цикл или функцию reflect.DeepEqual
	// import "reflect"
	// reflect.DeepEqual(a, b)
}

func FunctionsSlices() {
	// len(slice) - длина слайса
	// cap(slice) - ёмкость (максимум, который можно вместить без перераспределения)
	// append(slice, elems...) - добавление элементов
	// copy(destSlice, sourceSlice) - копирование элементов
	// s[i:j] - создание подслайса
}
