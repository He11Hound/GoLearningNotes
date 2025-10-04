package Multithreading

import (
	"context"
	"fmt"
	"time"
)

//Контекст - инстурумент который позволяет группировать и завершать работу нескольких горутин
//Отмена контекста - отменяет работу всех горутин, которые к нему привязаны
//В го самый главный контекст - background, все остальные создаются от него

func ContextExample() {
	contextFromParent, cancelFromParent := context.WithCancel(context.Background()) // контекст от backround контекста
	contextFromChild, cancelFromChild := context.WithCancel(contextFromParent)      // контекст который завязали от родительского (тот что выше название может быть любое, может также быть завязана на дочерний) контеста

	go example2(contextFromChild)
	go example1(contextFromParent)

	time.Sleep(10 * time.Second)

	cancelFromParent() //отмена родительского контекста
	//cancelFromChild() //отмена дочернего контекста

	time.Sleep(10 * time.Second)

	//cancelFromParent() //отмена родительского контекста
	cancelFromChild() //отмена дочернего контекста

	time.Sleep(1 * time.Second)
}

// Функция для родительского контекста
func example1(ctx context.Context) {
	for { // бесконечное отслеживание
		select { //Данная констркукция позволяет проверять состояние контекста
		case <-ctx.Done(): // выполняется когда контекст заблокирован, отменён
			fmt.Println("end first")
			return // завершаем горутину
		default: //выполняется всегда пока жив канал
			fmt.Println("output first") //какая-то логика
		}

		time.Sleep(1 * time.Second)
	}
}

// Функция для дочернего
func example2(ctx context.Context) {
	for { // бесконечное отслеживание
		select { //Данная констркукция позволяет проверять состояние контекста
		case <-ctx.Done(): // выполняется когда контекст заблокирован, отменён
			fmt.Println("end second")
			return // завершаем горутину
		default: //выполняется всегда пока жив канал
			fmt.Println("output second") //какая-то логика
		}

		time.Sleep(3 * time.Second)
	}
}
