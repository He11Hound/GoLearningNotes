package DataTypes

import (
	"fmt"
	"github.com/k0kubun/pp"
	"math"
	"strconv"
)

func MakeFloat() {
	var ExampleFloat1 float64 = 32.2 // Полная запись
	var ExampleFloat2 float64        // Короткая запись, опустили значение, значение будет 0
	ExampleFloat3 := 52.2            //Короткая запись, опустили тип будет зависеть от битности системы пользователя, для 32 - float32, 64 - float64

	fmt.Println("Variable1:", ExampleFloat1)
	fmt.Println("Variable2:", ExampleFloat2)
	fmt.Println("Variable3:", ExampleFloat3)
	fmt.Printf("Тип x: %T\n", ExampleFloat3) // вернёт тип

	//Для чисел с плавающей точкой в Go есть несколько типов
	//
	//| Тип       | Размер  | Пример диапазона | Точность             |
	//| --------- | ------- | ---------------- | -------------------- |
	//| `float32` | 4 байта | ~±3.4e38         | ~6–7 значащих цифр   |
	//| `float64` | 8 байт  | ~±1.7e308        | ~15–16 значащих цифр |
	//
	//| Тип          | Размер  | Компоненты                |
	//| ------------ | ------- | ------------------------- |
	//| `complex64`  | 8 байт  | 2 × float32 (real + imag) |
	//| `complex128` | 16 байт | 2 × float64 (real + imag) |
}

func ComparisonFloat() {
	Float1 := 322.2
	Float2 := 228.3

	fmt.Println(Float1 > Float2)
	fmt.Println(Float1 < Float2)
	fmt.Println(Float1 >= Float2)
	fmt.Println(Float1 <= Float2)
	fmt.Println(Float1 != Float2)
	fmt.Println(Float1 == Float2) // тут нужно быть аккуратнее

	//a := 0.1 + 0.2
	//b := 0.3
	//fmt.Println(a == b) // false!

	Float1 += Float2
	fmt.Println(Float1)
}

func MainFunctionsWithFloat() {
	FloatVal1 := 3.7
	FloatVal2 := -2.4

	//Округление и модуль
	fmt.Println("Abs:", math.Abs(FloatVal2))     // берёт модуль числа
	fmt.Println("Ceil:", math.Ceil(FloatVal1))   // округляет в большую сторону
	fmt.Println("Floor:", math.Floor(FloatVal1)) // округляет в меньшую
	fmt.Println("Trunc:", math.Trunc(FloatVal1)) // Отбрасывает дробную часть, тип переменной не меняется
	fmt.Println("Round:", math.Round(FloatVal1)) // Округляет в ближайшую сторону

	//В go нет функции для округления до сотых или тысячных, если есть задача, то нужно будет писать что то своё

	//Степени и корни
	fmt.Println("Pow:", math.Pow(2, 3)) //возвращает первое число в корне второго
	fmt.Println("Sqrt:", math.Sqrt(16)) // Корень из числа
	fmt.Println("Cbrt:", math.Cbrt(8))  // Кубический корень из числа

	//Мин макс
	fmt.Println("Min:", math.Min(FloatVal1, FloatVal2)) // Возвращает наименьшее значение
	fmt.Println("Max:", math.Max(FloatVal1, FloatVal2)) // Возвращает наибольшее значение

	//Константы
	fmt.Println("Число π:", math.Pi)     // Число π
	fmt.Println("Число Эйлера:", math.E) // Число Эйлера

	//Преобразования

	StringFormInt := strconv.FormatFloat(FloatVal1, 'f', 2, 64) // float → string
	pp.Println(StringFormInt)

	Integer := int(FloatVal1)     // float64 -> int
	Integer32 := int32(FloatVal1) // float64 -> int32
	Integer64 := int64(FloatVal1) // float64 -> int64

	pp.Println(Integer)
	pp.Println(Integer32)
	pp.Println(Integer64)
}
