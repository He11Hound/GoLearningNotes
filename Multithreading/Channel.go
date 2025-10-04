package Multithreading

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Что такое каналы?
//При работе с горутинами и какими-либо переменнами невозможно обойтись без Каналов
//Канал - пункт передачи значений, каналы служат для синхронизации между разными потоками управления

//Виды каналов
//Буферизированный - имеет буфер вместимых значений, если в буфере есть свободное место не блокирует выполнение горутины
//Небуферизированный - не имеет буфер, блокирует горутину до момента, пока из канала не будет прочитано значение

func MakeChannel() {
	var channel chan int = make(chan int) //создание небуферизрованного канала, полная запись
	defer close(channel)                  //закрытие канала
	//bufferizedChannel := make(chan int, 10) //создание буферизрованного канала,
	workers := 3
	bank := 0
	days := 7

	getMoney := func(channel chan int, days int) {
		for i := 0; i < days; i++ {
			channel <- rand.Intn(10) //запись в канал
			time.Sleep(time.Second)
		}
	}

	for i := 0; i < workers; i++ {
		go getMoney(channel, days)
	}

	for i := 0; i < workers; i++ {
		for k := 0; k < days; k++ {
			bank += <-channel // чтение из канала
			fmt.Println(bank)
		}
	}

	time.Sleep(10 * time.Second)
}

//Нужно быть аккуратным при работе с горутинами и канналами
//Так как очень легко можно получить ошибку, которая будет аффектить на систему

//deadlock (взаимная блокировка) - можно получить когда мы пытаемся записать что либо в канал, в котором нет места
//Утечка каннала - когда мы пытаемся что-то прочитать из канала в который уже ничего не пишется, ошибки не будет, но будет жраться cpu на планироащике рантайма го, так как он будет всё время держать канал в режиме ожидания

//Select {} - конструкция, напоминающия switch, предназначена для работы с нескольким каналами
//Конструкция предназначена для чтения нескольких канналов, default отрабатывает, когда ничего не отрабатывает

func SelectChannelExample() {
	intChan := make(chan int)
	stringChan := make(chan string)

	go func() {
		i := 1
		for {
			intChan <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i := 1
		for {
			stringChan <- "hi" + strconv.Itoa(i)
			i++
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case str := <-stringChan:
			fmt.Println(str)
		case intVal := <-intChan:
			fmt.Println(intVal)
		}
	}
}
