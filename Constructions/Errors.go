package Constructions

import (
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
)

//Ошибки частое явление, обработка ошибок - зачастую обязанность разработчика

type User struct {
	Name    string
	Balance int
}

func (user *User) Pay(usd int) error {
	if user.Balance-usd < 0 {
		return errors.New("Оплата не была произведена, так как недостаточно средств")
	}

	user.Balance -= usd

	return nil
}

func UserBalanceErrorExample() {
	user := &User{
		Name:    "John",
		Balance: 0,
	}

	errorMsg := user.Pay(10)

	if errorMsg != nil {
		pp.Println(errorMsg.Error())
	} else {
		pp.Println("Операция успешно проведена")
	}

	pp.Println(user)
}

func PanicExample() {
	defer func() {
		Examplepanic := recover()
		if Examplepanic != nil {
			fmt.Println("Была паника:", Examplepanic)
		}
	}()
	a := 0
	b := 1 / a

	fmt.Println(b)
}
