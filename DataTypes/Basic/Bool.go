package DataTypes

import "fmt"

func MakeBool() {
	var ExampleBool1 bool = true // Полная запись - переменная занимает 1 байт
	var ExampleBool2 bool        // Короткая запись, опустили значение, будет false (zero value)
	ExampleBool3 := false        // Короткая запись, тип автоматически bool

	fmt.Println("Variable1:", ExampleBool1)
	fmt.Println("Variable2:", ExampleBool2)
	fmt.Println("Variable3:", ExampleBool3)
	fmt.Printf("Тип x: %T\n", ExampleBool3) // вернёт тип

	// В Go тип bool может принимать только два значения: true или false
	// Zero value (значение по умолчанию) для bool = false
}

func ComparisonBool() {
	Bool1 := true
	Bool2 := false

	// Прямое сравнение
	fmt.Println(Bool1 == Bool2) // false
	fmt.Println(Bool1 != Bool2) // true

	// Булевы выражения
	fmt.Println(3 > 2)  // true
	fmt.Println(10 < 5) // false

	// Логические операции
	fmt.Println(Bool1 && Bool2) // Логическое И (true && false → false)
	fmt.Println(Bool1 || Bool2) // Логическое ИЛИ (true || false → true)
	fmt.Println(!Bool1)         // Логическое отрицание (не true → false)
}
