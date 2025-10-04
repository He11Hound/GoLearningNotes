package Multithreading

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup - инструмент который позволяет дожидаться момента, когда горутины закончат свою работу

func postman(text string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Я почтальон я отнёс газету", text, i, "раз")
		time.Sleep(1 * time.Second)
	}

	wg.Done() //Вычитает из счётчика значение
}

func WgExample() {
	wg := sync.WaitGroup{} // создание

	wg.Add(1)                //Перед каждым запуском горутины обновляем счётчик, для рижема ожидания
	go postman("maxim", &wg) //передаём указатель
	wg.Add(1)
	go postman("test", &wg)

	wg.Wait() //останавливаемся на данном этапе, ждём пока значение счётчика не станет 0

	fmt.Println("end")
}
