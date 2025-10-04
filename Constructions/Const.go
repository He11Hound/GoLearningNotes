package Constructions

import "math"

const testConst = 52
const flag uint8 = 128

//Именованные константы
//Ключевое слово const определяет именованную константу, при этом константе можно присвоить результат некоторого выражения. В одном объявлении const можно определить несколько констант.

//Группировка

//const pi = 3.14159
//const doublePi = pi * 2
//const version = "1.0.0"

// эквивалентно:

const (
	pi       = math.Pi
	doublePi = pi * 2
	version  = "1.0.0"
)
