package Multithreading

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var number atomic.Int64

// атомики - это отдельные, числа для работы с которыми внутри процессора предусмотрено отдельное место, благодаря им под капотом операция получения изменения записи переменной в памяти происходит в одну, что позволяет сразу несколькими горутинам работать с переменной. Но работа с атомиком тратит больше процессорного времени
// sync.Mutex Мьютекс — это механизм, который позволяет выполнить критические участки кода только одной горутиной
// sync.RWMutex - особый вид мьютекса, который позволяет одновременно выполняться либо произвольному количеству операций чтения, либо одной операции записи
func RushStateWithAtomicExample() {
	wg := sync.WaitGroup{}

	wg.Add(10)

	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)
	go increaseAtomic(&wg)

	wg.Wait()

	fmt.Println(number.Load())
}

func increaseAtomic(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		number.Add(1)
	}
	wg.Done()
}

func SimpleMutexExample() {
	var m sync.Mutex
	cache := map[int]int{}

	// горутины, которые изменяют мапу
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m.Lock() //Блокируем для записи
				cache[rand.Intn(5)] = rand.Intn(100)
				m.Unlock() //Разблокируем после записи
				time.Sleep(time.Second / 20)
			}
		}()
	}

	// горутины, которые читают мапу
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m.Lock() //Блокируем для чтения
				fmt.Printf("%#v\n", cache)
				m.Unlock() //Разблокируем после чтения
				time.Sleep(time.Second / 100)
			}
		}()
	}

	time.Sleep(1 * time.Second)
}

func SimpleRWMutexExample() {
	// меняем тип мьютекса
	var m sync.RWMutex
	cache := map[int]int{}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				// здесь остаются блокировки на запись, ничего отличного от примера выше, лучше даже вынести функцию
				m.Lock() //Блокируем для записи
				cache[rand.Intn(5)] = rand.Intn(100)
				m.Unlock() //Разблокируем после записи
				time.Sleep(time.Second / 20)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				// при чтении используем Rlock() и RUnlock()
				m.RLock() //Блокируем для чтения
				fmt.Printf("%#v\n", cache)
				m.RUnlock() //Разблокируем после чтения
				time.Sleep(time.Second / 100)
			}
		}()
	}

	time.Sleep(1 * time.Second)
}
