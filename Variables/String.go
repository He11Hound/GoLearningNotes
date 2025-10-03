package Variables

import (
	"fmt"
	"github.com/k0kubun/pp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func MakeString() {
	var ExampleString1 string = "hello world" //Полная запись
	var ExampleString2 string                 // Короткая запись, опустили значение, значение будет ""
	ExampleString3 := "Hello World Again"     //Короткая запись, опустили тип будет string

	fmt.Println("Variable1:", ExampleString1)
	fmt.Println("Variable2:", ExampleString2)
	fmt.Println("Variable3:", ExampleString3)
	fmt.Printf("Тип x: %T\n", ExampleString3) // вернёт тип

	//Также существуеют строки с обратными кавычками
	//Их отличие от обычных, что
	//В них не обрабатываются escape-последовательности.
	//Всё, что написано внутри, выводится как есть.

	TextWithSimpleQuotes := "Hello \n World "
	TextWithReverseQuotes := `Hello \n World `

	fmt.Println("Variable2:", TextWithSimpleQuotes)
	fmt.Println("Variable3:", TextWithReverseQuotes)

	//Также в go есть одиночные ковычки, но с их помощью задаются руны (rune — это псевдоним для int32, то есть число, представляющее Unicode-код символа.)
}

func ComparisonString() {
	String1 := "Hello World"
	String2 := "Hello World Again"

	fmt.Println(String1 > String2)
	fmt.Println(String1 < String2)
	fmt.Println(String1 >= String2)
	fmt.Println(String1 <= String2)
	fmt.Println(String1 != String2)
	fmt.Println(String1 == String2)

	String1 += String2
	fmt.Println(String1)
}

func MainFunctionsWithStrings() {
	SimpleText := "Hello World"
	TextWithSpace := "     Hello World     "
	TextWithComma := "Hello, World"
	NumberText := "42"
	FloatText := "3.14"

	fmt.Println(len(SimpleText))                    //возвращает количество символов (UTF-8 по 2 байта на символ)
	fmt.Println(utf8.RuneCountInString(SimpleText)) // количество символов (правильное для Unicode).

	//Поиск и проверка
	fmt.Println(strings.Contains(SimpleText, "World"))    //проверяет, содержится ли подстрока.
	fmt.Println(strings.ContainsAny(SimpleText, "World")) //проверяет, содержится ли хотя бы один символ из набора
	fmt.Println(strings.HasPrefix(SimpleText, "Hello"))   //начинается ли строка с подстроки.
	fmt.Println(strings.HasSuffix(SimpleText, "Hello"))   //заканчивается ли строка на подстроку.
	fmt.Println(strings.Index(SimpleText, "World"))       //индекс первого вхождения
	fmt.Println(strings.LastIndex(SimpleText, "World"))   //индекс последнего вхождения
	fmt.Println(strings.Count(SimpleText, "Hello"))       //возвращает количество непересекающихся

	//Изменение и обработка
	fmt.Println(strings.Replace(SimpleText, "Hello", "Hi", -1)) //замена подстроки, параметр n указывает количество раз сколько будет замен, при указании -1 будут заменяться все вхождения
	fmt.Println(strings.Repeat(SimpleText, 2))                  // Повторяет строку определённое количество раз
	fmt.Println(strings.ToLower(SimpleText))                    // выведет всё маленькими буквами
	fmt.Println(strings.ToUpper(SimpleText))                    // выведет всё Большими буквами
	fmt.Println(strings.Trim(SimpleText, "H"))                  // обрезка указанных символов по краям. Кажется
	fmt.Println(strings.TrimSpace(TextWithSpace))               //убрать пробелы/табуляции/переводы строк по краям.

	//Разделение и объединение
	pp.Println((strings.Split(TextWithComma, ","))) //разделить строку по разделителю. В итоге будет слайс. Обратный метод - JOIN
	pp.Println((strings.Fields(SimpleText)))        //разделение по пробелам/табам/новым строкам

	//Сравнение
	//== - без учёта регистра
	fmt.Println(strings.EqualFold(SimpleText, strings.ToLower(SimpleText))) // проверка с учётом регистра

	//Преоьразование строки в число

	number, _ := strconv.Atoi(NumberText)         // string → int
	float, _ := strconv.ParseFloat(FloatText, 64) // string → float64
	fmt.Println(number)
	fmt.Println(float)
}

func EffectiveMakeStrings() {
	BasicStringMaker := ""
	for i := 0; i < 5; i++ {
		BasicStringMaker += "Go "
	}
	fmt.Println(BasicStringMaker)

	//В Go строки неизменяемы.
	//Каждая операция конкатенации (s = s + "...") создаёт новую строку в памяти и копирует старое содержимое.
	//Если строк много (например, цикл на тысячи элементов), это очень дорого.
	//strings.Builder хранит промежуточный буфер (слайс байтов) и не пересоздаёт строку каждый раз.
	//strings.Builder нужен, когда строишь строку из многих частей (особенно в цикле).

	var StringBuilder strings.Builder
	for i := 0; i < 5; i++ {
		StringBuilder.WriteString("Go ")
	}
	fmt.Println(StringBuilder.String())

	//Когда использовать
	//При генерации HTML, JSON, SQL и т. п.
	//При обработке больших текстов.
	//Когда в цикле склеиваются сотни/тысячи строк.
	//А для пары строк использовать Builder не нужно — обычная конкатенация быстрее и проще.
}

func CutStringByIndex() {
	//Строка — это байтовый массив, и при срезе можно случайно «разрезать» Unicode-символ:

	SimpleString := "Привет"
	fmt.Println(SimpleString[:3]) // неверно, обрежет по байтам -> "Пр�"

	rune := []rune(SimpleString)
	//rune := []rune{1047, 1088, 1080, 1074, 1077, 1090}
	fmt.Println(string(rune[:3])) // "При"
}
