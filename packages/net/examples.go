package net

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// DialExample демонстрирует использование функции net.Dial
func DialExample() {
	// Установка TCP соединения
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("TCP соединение установлено")
}

// DialTimeoutExample демонстрирует использование функции net.DialTimeout
func DialTimeoutExample() {
	// Установка TCP соединения с таймаутом
	conn, err := net.DialTimeout("tcp", "example.com:80", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("TCP соединение установлено с таймаутом")
}

// ListenExample демонстрирует использование функции net.Listen
func ListenExample() {
	// Создание TCP слушателя
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("TCP слушатель создан на порту 8080")
}

// ListenTCPExample демонстрирует использование функции net.ListenTCP
func ListenTCPExample() {
	// Создание TCP слушателя с конкретным адресом
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("TCP слушатель создан с конкретным адресом")
}

// DialTCPExample демонстрирует использование функции net.DialTCP
func DialTCPExample() {
	// Установка TCP соединения с конкретными адресами
	raddr, err := net.ResolveTCPAddr("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("TCP соединение установлено с конкретными адресами")
}

// ListenUDPExample демонстрирует использование функции net.ListenUDP
func ListenUDPExample() {
	// Создание UDP слушателя
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("UDP слушатель создан на порту 8080")
}

// DialUDPExample демонстрирует использование функции net.DialUDP
func DialUDPExample() {
	// Установка UDP соединения
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("UDP соединение установлено")
}

// ResolveUDPAddrExample демонстрирует использование функции net.ResolveUDPAddr
func ResolveUDPAddrExample() {
	// Разрешение UDP адреса
	addr, err := net.ResolveUDPAddr("udp", "example.com:53")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UDP адрес разрешен: %s\n", addr.String())
}

// ListenUnixExample демонстрирует использование функции net.ListenUnix
func ListenUnixExample() {
	// Создание Unix сокет слушателя
	addr, err := net.ResolveUnixAddr("unix", "/tmp/socket")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Unix сокет слушатель создан")
}

// DialUnixExample демонстрирует использование функции net.DialUnix
func DialUnixExample() {
	// Установка Unix сокет соединения
	raddr, err := net.ResolveUnixAddr("unix", "/tmp/socket")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUnix("unix", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Unix сокет соединение установлено")
}

// ResolveUnixAddrExample демонстрирует использование функции net.ResolveUnixAddr
func ResolveUnixAddrExample() {
	// Разрешение Unix адреса
	addr, err := net.ResolveUnixAddr("unix", "/tmp/socket")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unix адрес разрешен: %s\n", addr.String())
}

// ResolveTCPAddrExample демонстрирует использование функции net.ResolveTCPAddr
func ResolveTCPAddrExample() {
	// Разрешение TCP адреса
	addr, err := net.ResolveTCPAddr("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TCP адрес разрешен: %s\n", addr.String())
}

// ResolveIPAddrExample демонстрирует использование функции net.ResolveIPAddr
func ResolveIPAddrExample() {
	// Разрешение IP адреса
	addr, err := net.ResolveIPAddr("ip", "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP адрес разрешен: %s\n", addr.String())
}

// LookupHostExample демонстрирует использование функции net.LookupHost
func LookupHostExample() {
	// DNS поиск хоста
	hosts, err := net.LookupHost("example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Найдены хосты: %v\n", hosts)
}

// LookupIPExample демонстрирует использование функции net.LookupIP
func LookupIPExample() {
	// DNS поиск IP адресов
	ips, err := net.LookupIP("example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Найдены IP адреса: %v\n", ips)
}

// LookupPortExample демонстрирует использование функции net.LookupPort
func LookupPortExample() {
	// Поиск порта для сервиса
	port, err := net.LookupPort("tcp", "http")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Порт для HTTP: %d\n", port)
}

// LookupCNAMEExample демонстрирует использование функции net.LookupCNAME
func LookupCNAMEExample() {
	// Поиск канонического имени хоста
	cname, err := net.LookupCNAME("www.example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Каноническое имя: %s\n", cname)
}

// LookupSRVExample демонстрирует использование функции net.LookupSRV
func LookupSRVExample() {
	// Поиск SRV записей DNS
	cname, srvs, err := net.LookupSRV("_http", "_tcp", "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Каноническое имя: %s\n", cname)
	fmt.Printf("SRV записи: %v\n", srvs)
}

// LookupMXExample демонстрирует использование функции net.LookupMX
func LookupMXExample() {
	// Поиск MX записей DNS
	mxs, err := net.LookupMX("example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MX записи: %v\n", mxs)
}

// LookupTXTExample демонстрирует использование функции net.LookupTXT
func LookupTXTExample() {
	// Поиск TXT записей DNS
	txts, err := net.LookupTXT("example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TXT записи: %v\n", txts)
}

// InterfacesExample демонстрирует использование функции net.Interfaces
func InterfacesExample() {
	// Получение списка сетевых интерфейсов
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сетевые интерфейсы:")
	for _, iface := range interfaces {
		fmt.Printf("  %s: %s\n", iface.Name, iface.Flags)
	}
}

// InterfaceAddrsExample демонстрирует использование функции net.InterfaceAddrs
func InterfaceAddrsExample() {
	// Получение списка адресов интерфейсов
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Адреса интерфейсов:")
	for _, addr := range addrs {
		fmt.Printf("  %s\n", addr.String())
	}
}

// InterfaceByIndexExample демонстрирует использование функции net.InterfaceByIndex
func InterfaceByIndexExample() {
	// Получение интерфейса по индексу
	iface, err := net.InterfaceByIndex(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Интерфейс по индексу 1: %s\n", iface.Name)
}

// InterfaceByNameExample демонстрирует использование функции net.InterfaceByName
func InterfaceByNameExample() {
	// Получение интерфейса по имени
	iface, err := net.InterfaceByName("lo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Интерфейс lo: %s\n", iface.Flags)
}

// ParseIPExample демонстрирует использование функции net.ParseIP
func ParseIPExample() {
	// Парсинг IP адреса
	ip := net.ParseIP("192.168.1.1")
	if ip == nil {
		log.Fatal("Невалидный IP адрес")
	}

	fmt.Printf("IP адрес: %s\n", ip.String())
}

// IPv4Example демонстрирует использование функции net.IPv4
func IPv4Example() {
	// Создание IPv4 адреса
	ip := net.IPv4(192, 168, 1, 1)
	fmt.Printf("IPv4 адрес: %s\n", ip.String())
}

// ParseCIDRExample демонстрирует использование функции net.ParseCIDR
func ParseCIDRExample() {
	// Парсинг CIDR нотации
	ip, network, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP адрес: %s\n", ip.String())
	fmt.Printf("Сеть: %s\n", network.String())
}

// SplitHostPortExample демонстрирует использование функции net.SplitHostPort
func SplitHostPortExample() {
	// Разделение адреса на хост и порт
	host, port, err := net.SplitHostPort("example.com:80")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Хост: %s, Порт: %s\n", host, port)
}

// JoinHostPortExample демонстрирует использование функции net.JoinHostPort
func JoinHostPortExample() {
	// Объединение хоста и порта
	addr := net.JoinHostPort("example.com", "80")
	fmt.Printf("Адрес: %s\n", addr)
}

// SetDeadlineExample демонстрирует использование функции net.SetDeadline
func SetDeadlineExample() {
	// Установка дедлайна для соединения
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Установка дедлайна на 5 секунд
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Дедлайн установлен на 5 секунд")
}

// SetReadDeadlineExample демонстрирует использование функции net.SetReadDeadline
func SetReadDeadlineExample() {
	// Установка дедлайна для чтения
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Установка дедлайна для чтения на 5 секунд
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Дедлайн для чтения установлен на 5 секунд")
}

// SetWriteDeadlineExample демонстрирует использование функции net.SetWriteDeadline
func SetWriteDeadlineExample() {
	// Установка дедлайна для записи
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Установка дедлайна для записи на 5 секунд
	err = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Дедлайн для записи установлен на 5 секунд")
}

// TCPClientExample демонстрирует создание TCP клиента
func TCPClientExample() {
	// Создание TCP клиента
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Отправка HTTP запроса
	request := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatal(err)
	}

	// Чтение ответа
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Получен ответ: %s\n", string(response[:n]))
}

// TCPServerExample демонстрирует создание TCP сервера
func TCPServerExample() {
	// Создание TCP сервера
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("TCP сервер запущен на порту 8080")

	// Принятие соединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Обработка соединения в горутине
		go func(conn net.Conn) {
			defer conn.Close()
			conn.Write([]byte("Привет от TCP сервера!\n"))
		}(conn)
	}
}

// UDPClientExample демонстрирует создание UDP клиента
func UDPClientExample() {
	// Создание UDP клиента
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Отправка данных
	message := "Привет от UDP клиента!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP сообщение отправлено")
}

// UDPServerExample демонстрирует создание UDP сервера
func UDPServerExample() {
	// Создание UDP сервера
	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("UDP сервер запущен на порту 8080")

	// Обработка пакетов
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Получено от %s: %s\n", addr.String(), string(buffer[:n]))
	}
}

// UnixSocketExample демонстрирует работу с Unix сокетами
func UnixSocketExample() {
	// Создание Unix сокет сервера
	listener, err := net.Listen("unix", "/tmp/socket")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Unix сокет сервер запущен")

	// Принятие соединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Обработка соединения
		go func(conn net.Conn) {
			defer conn.Close()
			conn.Write([]byte("Привет от Unix сокет сервера!\n"))
		}(conn)
	}
}

// ContextExample демонстрирует использование контекста с сетью
func ContextExample() {
	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Установка соединения с контекстом
	conn, err := net.DialContext(ctx, "tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Соединение установлено с контекстом")
}

// BufferedIOExample демонстрирует использование буферизованного I/O
func BufferedIOExample() {
	// Создание соединения
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Создание буферизованного читателя
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Отправка данных
	writer.WriteString("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n")
	writer.Flush()

	// Чтение данных
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Получена строка: %s\n", line)
}

// MultipleConnectionsExample демонстрирует работу с множественными соединениями
func MultipleConnectionsExample() {
	// Создание множественных соединений
	hosts := []string{"example.com:80", "google.com:80", "github.com:80"}

	for _, host := range hosts {
		go func(host string) {
			conn, err := net.Dial("tcp", host)
			if err != nil {
				log.Printf("Ошибка соединения с %s: %v\n", host, err)
				return
			}
			defer conn.Close()

			fmt.Printf("Соединение с %s установлено\n", host)
		}(host)
	}

	// Ожидание завершения
	time.Sleep(2 * time.Second)
}

// ErrorHandlingExample демонстрирует обработку ошибок сети
func ErrorHandlingExample() {
	// Попытка соединения с несуществующим хостом
	conn, err := net.Dial("tcp", "nonexistent.example:80")
	if err != nil {
		// Проверка типа ошибки
		if netErr, ok := err.(net.Error); ok {
			if netErr.Timeout() {
				fmt.Println("Ошибка таймаута")
			} else if netErr.Temporary() {
				fmt.Println("Временная ошибка")
			} else {
				fmt.Printf("Сетевая ошибка: %v\n", err)
			}
		} else {
			fmt.Printf("Другая ошибка: %v\n", err)
		}
		return
	}
	defer conn.Close()

	fmt.Println("Соединение установлено")
}

// NetworkInterfaceExample демонстрирует работу с сетевыми интерфейсами
func NetworkInterfaceExample() {
	// Получение всех интерфейсов
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range interfaces {
		fmt.Printf("Интерфейс: %s\n", iface.Name)
		fmt.Printf("  Флаги: %s\n", iface.Flags)
		fmt.Printf("  MTU: %d\n", iface.MTU)

		// Получение адресов интерфейса
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			fmt.Printf("  Адрес: %s\n", addr.String())
		}
		fmt.Println()
	}
}

// IPOperationsExample демонстрирует операции с IP адресами
func IPOperationsExample() {
	// Парсинг IP адреса
	ip := net.ParseIP("192.168.1.1")
	if ip == nil {
		log.Fatal("Невалидный IP адрес")
	}

	// Проверка типа IP
	if ip.To4() != nil {
		fmt.Println("Это IPv4 адрес")
	} else if ip.To16() != nil {
		fmt.Println("Это IPv6 адрес")
	}

	// Проверка приватности
	if ip.IsPrivate() {
		fmt.Println("Это приватный IP адрес")
	}

	// Проверка loopback
	if ip.IsLoopback() {
		fmt.Println("Это loopback IP адрес")
	}

	// Проверка multicast
	if ip.IsMulticast() {
		fmt.Println("Это multicast IP адрес")
	}

	fmt.Printf("IP адрес: %s\n", ip.String())
}

// PortScanningExample демонстрирует сканирование портов
func PortScanningExample() {
	// Сканирование портов
	host := "127.0.0.1"
	ports := []int{22, 80, 443, 8080}

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			fmt.Printf("Порт %d закрыт\n", port)
		} else {
			fmt.Printf("Порт %d открыт\n", port)
			conn.Close()
		}
	}
}

// DNSResolutionExample демонстрирует разрешение DNS
func DNSResolutionExample() {
	// Разрешение имени хоста
	ips, err := net.LookupIP("example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IP адреса для example.com:")
	for _, ip := range ips {
		fmt.Printf("  %s\n", ip.String())
	}

	// Обратное разрешение
	names, err := net.LookupAddr("93.184.216.34")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Имена для 93.184.216.34:")
	for _, name := range names {
		fmt.Printf("  %s\n", name)
	}
}

// NetworkMonitoringExample демонстрирует мониторинг сети
func NetworkMonitoringExample() {
	// Получение статистики интерфейсов
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range interfaces {
		// Проверка состояния интерфейса
		if iface.Flags&net.FlagUp != 0 {
			fmt.Printf("Интерфейс %s активен\n", iface.Name)
		} else {
			fmt.Printf("Интерфейс %s неактивен\n", iface.Name)
		}

		// Проверка возможностей интерфейса
		if iface.Flags&net.FlagBroadcast != 0 {
			fmt.Printf("  Поддерживает широковещание\n")
		}
		if iface.Flags&net.FlagMulticast != 0 {
			fmt.Printf("  Поддерживает multicast\n")
		}
		if iface.Flags&net.FlagLoopback != 0 {
			fmt.Printf("  Это петлевой интерфейс\n")
		}
	}
}

// AdvancedTCPExample демонстрирует продвинутые возможности TCP
func AdvancedTCPExample() {
	// Создание TCP соединения с настройками
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Установка таймаутов
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	// Получение информации о соединении
	localAddr := conn.LocalAddr()
	remoteAddr := conn.RemoteAddr()

	fmt.Printf("Локальный адрес: %s\n", localAddr.String())
	fmt.Printf("Удаленный адрес: %s\n", remoteAddr.String())

	// Отправка данных
	request := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatal(err)
	}

	// Чтение ответа
	response := make([]byte, 4096)
	n, err := conn.Read(response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Получено %d байт ответа\n", n)
}

// CustomProtocolExample демонстрирует создание пользовательского протокола
func CustomProtocolExample() {
	// Создание TCP сервера для пользовательского протокола
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Пользовательский протокол сервер запущен")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			// Чтение команды
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				return
			}

			command := strings.TrimSpace(string(buffer[:n]))
			fmt.Printf("Получена команда: %s\n", command)

			// Обработка команды
			switch command {
			case "HELLO":
				conn.Write([]byte("Привет!\n"))
			case "TIME":
				conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05\n")))
			case "QUIT":
				conn.Write([]byte("До свидания!\n"))
				return
			default:
				conn.Write([]byte("Неизвестная команда\n"))
			}
		}(conn)
	}
}
