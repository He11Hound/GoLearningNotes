package OOP

import (
	"fmt"
	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Age     int
	IsAlive bool
}

//Структура - логически объединённый набор данных, под какую-то сущность

func MakeNewUser() {
	var ExampleUser1 User = User{
		"Sanya",
		24,
		true,
	}

	var Exampleuser2 User // Exampleuser2 := User{} Когда не объявляем значение, каждое не объявленное задаётся по умолчанию, "", 0, false и т.д.

	ExampleUser3 := User{
		"Danya",
		25,
		false,
	}

	pp.Println(ExampleUser1)
	pp.Println(Exampleuser2)
	pp.Println(ExampleUser3)
}

type ExampleStruct1 struct {
	Name    string
	Age     int
	IsAlive bool
}

type ExampleStruct2 struct {
	name    string
	age     int
	isAlive bool
}

//Если поле структуры начинается с заглавной буквы → оно экспортируемое (public), доступно из других пакетов.
//Если поле структуры начинается со строчной буквы → оно неэкспортируемое (private), доступно только внутри того же пакета.

func StructAccessExample() {
	exampleStruct1 := ExampleStruct1{
		"Test1",
		10,
		true,
	}

	pp.Println("было", exampleStruct1)
	exampleStruct1.Name = "Произвольное имя"
	exampleStruct1.Age = 200
	exampleStruct1.IsAlive = false
	pp.Println("стало", exampleStruct1)

	//Здесь всё отработает правильно но при переносе вне пакет будет ошибка, поэтому на уровне пакета для полей нужно будет реализовывать методы для работы с данными полями
	//exampleStruct2 := ExampleStruct2{
	//	"Test2",
	//	10,
	//	true,
	//}
	//
	//pp.Println("было", exampleStruct2)
	//exampleStruct2.name = "Произвольное имя"
	//exampleStruct2.age = 200
	//exampleStruct2.isAlive = false
	//pp.Println("стало", exampleStruct2)
}

type Person struct {
	string
	int
}

// Анонимные поля
// В структуре можно хранить поле без имени — просто тип. Тогда имя поля = имя типа:
// Обычно такое используется для встраивания структур.
func ExampleAnon() {
	p := Person{"Alex", 30}
	fmt.Println(p.string) // "Alex"
	fmt.Println(p.int)    // 30
}

// Встраивание структур (псевдо-наследование)
// В Go нет наследования, но можно "встроить" одну структуру в другую. Это даёт композицию и доступ к полям/методам вложенной структуры:
type Address struct {
	City   string
	Street string
}

type Citizien struct {
	Name    string
	Age     int
	Address // встраивание
}

func ExampleEmbed() {
	u := Citizien{"Sanya", 24, Address{"Moscow", "Lenina"}}
	fmt.Println(u.City)   // "Moscow"
	fmt.Println(u.Street) // "Lenina"
}

//К структурам можно привязывать методы:
//Метод — это функция с привязкой к конкретному типу (обычно к структуре).
//Важно: методы можно делать с value receiver (копия структуры) или с pointer receiver (работа с оригиналом):
//Если метод объявлен через value receiver — в метод передаётся копия структуры.
//Если используем pointer receiver, метод получает доступ к "оригиналу".

func (u User) BirthdayValue() { //копия объекта
	u.Age++
}

func (u *User) BirthdayPointer() { //ссылка на объект - изменит передаваемое значение
	u.Age++
}

// Конструкторы
// В Go нет явных конструкторов, их пишут как обычные функции:
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age, IsAlive: true}
}

//Структуры можно сравнивать ==, если все их поля сравнимы.
//Если хотя бы одно поле — срез, карта, функция → сравнение будет недопустимо.

/*
--todo: Отдельно разобрать устройство тэгов в go --
*/

//Теги структур
//Можно добавлять метаданные для сериализации:

type UserSerialize struct {
	Name string `json:"name" db:"user_name" validate:"required"`
	Age  int    `json:"age" validate:"min=18"`
}

//json:"name" → указывает, как поле будет называться при сериализации в JSON.
//db:"user_name" → имя колонки в базе данных.
//validate:"required" → правило для валидации (например, в пакете go-playground/validator).
