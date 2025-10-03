package Variables

import (
	"fmt"
	"strconv"
)

func MakeInteger() {
	var ExampleInteger1 int = 322 //Полная запись
	var ExampleInteger2 int       // Короткая запись, опустили значение, значение будет 0
	ExampleInteger3 := 52         //Короткая запись, опустили тип будет int

	fmt.Println("Variable1:", ExampleInteger1)
	fmt.Println("Variable2:", ExampleInteger2)
	fmt.Println("Variable3:", ExampleInteger3)
	fmt.Printf("Тип x: %T\n", ExampleInteger3) // вернёт тип

	//Для чисел в Go придумана куча типов
	//| Тип      | Размер        | Диапазон (примерно)                                                        | Знаковый?               |
	//| -------- | ------------- | -------------------------------------------------------------------------- | ----------------------- |
	//| `int`    | 4 или 8 байт* | зависит от платформы, обычно -2³¹..2³¹-1 (32 бит) или -2⁶³..2⁶³-1 (64 бит) | да                      |
	//| `int8`   | 1 байт        | -128 … 127                                                                 | да                      |
	//| `int16`  | 2 байта       | -32,768 … 32,767                                                           | да                      |
	//| `int32`  | 4 байта       | -2,147,483,648 … 2,147,483,647                                             | да                      |
	//| `int64`  | 8 байт        | -9,223,372,036,854,775,808 … 9,223,372,036,854,775,807                     | да                      |
	//| `uint`   | 4 или 8 байт* | 0 … 2³²-1 (32 бит) или 0 … 2⁶⁴-1 (64 бит)                                  | нет                     |
	//| `uint8`  | 1 байт        | 0 … 255                                                                    | нет                     |
	//| `uint16` | 2 байта       | 0 … 65,535                                                                 | нет                     |
	//| `uint32` | 4 байта       | 0 … 4,294,967,295                                                          | нет                     |
	//| `uint64` | 8 байт        | 0 … 18,446,744,073,709,551,615                                             | нет                     |
	//| `byte`   | 1 байт        | 0 … 255                                                                    | нет (alias для `uint8`) |
	//| `rune`   | 4 байта       | -2,147,483,648 … 2,147,483,647                                             | да (alias для `int32`)  |
}

func ComparisonInteger() {
	Integer1 := 322
	Integer2 := 228

	fmt.Println(Integer1 > Integer2)
	fmt.Println(Integer1 < Integer2)
	fmt.Println(Integer1 >= Integer2)
	fmt.Println(Integer1 <= Integer2)
	fmt.Println(Integer1 != Integer2)
	fmt.Println(Integer1 == Integer2)
	fmt.Println(Integer1%2 == 0)

	Integer1 += Integer2
	fmt.Println(Integer1)
}

func MainFunctionsWithInteger() {
	//Особо операций при работе с числами типа int нет, кроме арифмитических
	//Преобразования
	Integer1 := 322

	StringFormInt := strconv.Itoa(Integer1) // int → string
	float := float64(Integer1)              // int → float64
	var Integer2 int64 = int64(Integer1)    // int32 -> int64
	fmt.Println(StringFormInt)
	fmt.Println(float)
	fmt.Println(Integer2)
	//В Go все преобразования типов явные, иначе будет ошибка компиляции.
}

func IfVariableMoreThanSize() {
	//Что будет если переполнить переменную
	var number int8 = 127
	fmt.Println(number)
	number++
	fmt.Println(number) // -128 (переполнение)
}
