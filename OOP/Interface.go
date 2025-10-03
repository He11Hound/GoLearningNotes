package OOP

import "fmt"

//интерфейс — это набор методов, которые должен реализовать тип.
//Если тип реализует все методы интерфейса → он автоматически считается реализацией этого интерфейса без implements и т.д.
//пустой интерфейс реализует любой тип данных

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeItSpeak(dog) // Woof!
	makeItSpeak(cat) // Meow!
}

// Интерфейсы могут включать другие интерфейсы:
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

//Интерфейс может передаваться в функцию как возвращаемое значение
//И интерпетируется это как:
//👉 "Я ожидаю тип, у которого есть такие методы."

type Greeter interface {
	Greet() string
}

type PersonExample struct {
	Name string
}

func (p PersonExample) Greet() string {
	return "Hello, my name is " + p.Name
}

func NewGreeter(name string) Greeter {
	return PersonExample{Name: name}
}
