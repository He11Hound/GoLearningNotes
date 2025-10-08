package net

/*
	ФУНКЦИИ ПАКЕТА NET

	Пакет net предоставляет переносимый интерфейс для сетевого ввода-вывода,
	включая TCP/IP, UDP, доменные сокеты и другие сетевые протоколы.
	Это один из самых важных пакетов в Go для сетевого программирования.

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С TCP ===

	func Dial(network, address string) (Conn, error)
		- Устанавливает соединение с удаленным адресом
		- network: тип сети ("tcp", "tcp4", "tcp6", "udp", "udp4", "udp6")
		- address: адрес в формате "host:port"
		- Возвращает Conn для обмена данными
		- Блокирующая операция

	func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
		- Устанавливает соединение с таймаутом
		- timeout: максимальное время ожидания
		- Возвращает ошибку если соединение не установлено в течение timeout
		- Полезно для избежания зависания

	func Listen(network, address string) (Listener, error)
		- Создает слушатель на указанном адресе
		- network: тип сети ("tcp", "tcp4", "tcp6")
		- address: адрес для прослушивания ("host:port" или ":port")
		- Возвращает Listener для принятия соединений
		- Используется для создания серверов

	func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
		- Создает TCP слушатель
		- network: "tcp", "tcp4", "tcp6"
		- laddr: локальный адрес для прослушивания
		- Возвращает TCPListener с дополнительными возможностями
		- Более специфичный чем Listen

	func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
		- Устанавливает TCP соединение
		- network: "tcp", "tcp4", "tcp6"
		- laddr: локальный адрес (может быть nil)
		- raddr: удаленный адрес
		- Возвращает TCPConn с дополнительными возможностями

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С UDP ===

	func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
		- Создает UDP слушатель
		- network: "udp", "udp4", "udp6"
		- laddr: локальный адрес для прослушивания
		- Возвращает UDPConn для обмена данными
		- Используется для UDP серверов

	func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
		- Устанавливает UDP соединение
		- network: "udp", "udp4", "udp6"
		- laddr: локальный адрес (может быть nil)
		- raddr: удаленный адрес
		- Возвращает UDPConn для обмена данными

	func ResolveUDPAddr(network, address string) (*UDPAddr, error)
		- Разрешает UDP адрес
		- network: "udp", "udp4", "udp6"
		- address: адрес в формате "host:port"
		- Возвращает UDPAddr структуру
		- Используется перед созданием UDP соединений

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С UNIX СОКЕТАМИ ===

	func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
		- Создает Unix сокет слушатель
		- network: "unix", "unixgram", "unixpacket"
		- laddr: локальный адрес Unix сокета
		- Возвращает UnixListener
		- Используется для межпроцессного взаимодействия

	func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
		- Устанавливает Unix сокет соединение
		- network: "unix", "unixgram", "unixpacket"
		- laddr: локальный адрес (может быть nil)
		- raddr: удаленный адрес
		- Возвращает UnixConn
		- Используется для локального взаимодействия

	func ResolveUnixAddr(network, address string) (*UnixAddr, error)
		- Разрешает Unix адрес
		- network: "unix", "unixgram", "unixpacket"
		- address: путь к файлу сокета
		- Возвращает UnixAddr структуру

	=== ФУНКЦИИ ДЛЯ РАЗРЕШЕНИЯ АДРЕСОВ ===

	func ResolveTCPAddr(network, address string) (*TCPAddr, error)
		- Разрешает TCP адрес
		- network: "tcp", "tcp4", "tcp6"
		- address: адрес в формате "host:port"
		- Возвращает TCPAddr структуру
		- Выполняет DNS разрешение

	func ResolveIPAddr(network, address string) (*IPAddr, error)
		- Разрешает IP адрес
		- network: "ip", "ip4", "ip6"
		- address: IP адрес или имя хоста
		- Возвращает IPAddr структуру
		- Выполняет DNS разрешение

	func LookupHost(host string) ([]string, error)
		- Выполняет DNS поиск хоста
		- host: имя хоста для поиска
		- Возвращает список IP адресов
		- Возвращает как IPv4, так и IPv6 адреса

	func LookupIP(host string) ([]IP, error)
		- Выполняет DNS поиск IP адресов
		- host: имя хоста для поиска
		- Возвращает список IP адресов
		- Более специфичный чем LookupHost

	func LookupPort(network, service string) (int, error)
		- Ищет порт для сервиса
		- network: тип сети ("tcp", "udp")
		- service: имя сервиса ("http", "ssh", "ftp")
		- Возвращает номер порта
		- Использует /etc/services или аналогичные файлы

	func LookupCNAME(cname string) (string, error)
		- Ищет каноническое имя хоста
		- cname: имя хоста для поиска
		- Возвращает каноническое имя
		- Используется для разрешения CNAME записей

	func LookupSRV(service, proto, name string) (string, []*SRV, error)
		- Ищет SRV записи DNS
		- service: имя сервиса ("_http", "_sip")
		- proto: протокол ("_tcp", "_udp")
		- name: доменное имя
		- Возвращает целевой хост и список SRV записей

	func LookupMX(name string) ([]*MX, error)
		- Ищет MX записи DNS
		- name: доменное имя
		- Возвращает список MX записей
		- Используется для поиска почтовых серверов

	func LookupTXT(name string) ([]string, error)
		- Ищет TXT записи DNS
		- name: доменное имя
		- Возвращает список TXT записей
		- Используется для различных целей (SPF, DKIM и т.д.)

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ИНТЕРФЕЙСАМИ ===

	func Interfaces() ([]Interface, error)
		- Возвращает список сетевых интерфейсов
		- Возвращает все доступные интерфейсы
		- Включает информацию о состоянии интерфейсов
		- Полезно для обнаружения сетевых устройств

	func InterfaceAddrs() ([]Addr, error)
		- Возвращает список адресов интерфейсов
		- Возвращает все адреса всех интерфейсов
		- Включает IP адреса и другие типы адресов
		- Полезно для получения локальных адресов

	func InterfaceByIndex(index int) (*Interface, error)
		- Возвращает интерфейс по индексу
		- index: индекс интерфейса
		- Возвращает информацию об интерфейсе
		- Используется для работы с конкретными интерфейсами

	func InterfaceByName(name string) (*Interface, error)
		- Возвращает интерфейс по имени
		- name: имя интерфейса ("eth0", "wlan0", "lo")
		- Возвращает информацию об интерфейсе
		- Используется для работы с конкретными интерфейсами

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С IP АДРЕСАМИ ===

	func ParseIP(s string) IP
		- Парсит IP адрес из строки
		- s: строка с IP адресом ("192.168.1.1", "::1")
		- Возвращает IP структуру
		- Возвращает nil если адрес невалидный

	func IPv4(a, b, c, d byte) IP
		- Создает IPv4 адрес из байтов
		- a, b, c, d: байты IPv4 адреса
		- Возвращает IP структуру
		- Удобно для создания IPv4 адресов

	func ParseCIDR(s string) (IP, *IPNet, error)
		- Парсит CIDR нотацию
		- s: строка в формате CIDR ("192.168.1.0/24")
		- Возвращает IP адрес и сеть
		- Используется для работы с подсетями

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ПОРТАМИ ===

	func SplitHostPort(hostport string) (host, port string, err error)
		- Разделяет адрес на хост и порт
		- hostport: адрес в формате "host:port"
		- Возвращает хост и порт отдельно
		- Обрабатывает IPv6 адреса с квадратными скобками

	func JoinHostPort(host, port string) string
		- Объединяет хост и порт в адрес
		- host: имя хоста или IP адрес
		- port: номер порта
		- Возвращает адрес в формате "host:port"
		- Правильно обрабатывает IPv6 адреса

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ТАЙМАУТАМИ ===

	func SetDeadline(t time.Time) error
		- Устанавливает дедлайн для операций
		- t: время дедлайна
		- Применяется ко всем операциям ввода-вывода
		- Операции завершаются с ошибкой после дедлайна

	func SetReadDeadline(t time.Time) error
		- Устанавливает дедлайн для операций чтения
		- t: время дедлайна
		- Применяется только к операциям чтения
		- Операции чтения завершаются с ошибкой после дедлайна

	func SetWriteDeadline(t time.Time) error
		- Устанавливает дедлайн для операций записи
		- t: время дедлайна
		- Применяется только к операциям записи
		- Операции записи завершаются с ошибкой после дедлайна

	=== ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ ===

	// TCP клиент
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// TCP сервер
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// UDP соединение
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// DNS разрешение
	ips, err := net.LookupIP("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Разделение адреса
	host, port, err := net.SplitHostPort("example.com:80")
	if err != nil {
		log.Fatal(err)
	}
*/
