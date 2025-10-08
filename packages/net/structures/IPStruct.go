package net

/*
	Тип IP

	type IP []byte

	IP представляет IP адрес. Это срез байтов, который может содержать
	как IPv4, так и IPv6 адреса. IP адреса используются для идентификации
	сетевых узлов и маршрутизации трафика.

	Основные методы:

	func (ip IP) String() string
		- Возвращает строковое представление IP адреса
		- IPv4 адреса возвращаются в формате "192.168.1.1"
		- IPv6 адреса возвращаются в формате "2001:db8::1"
		- Используется для отображения и логирования

	func (ip IP) To4() IP
		- Преобразует IP в IPv4 адрес
		- Возвращает IPv4 адрес если это возможно
		- Возвращает nil если IP не является IPv4
		- Используется для проверки типа IP

	func (ip IP) To16() IP
		- Преобразует IP в IPv6 адрес
		- Возвращает IPv6 адрес если это возможно
		- Возвращает nil если IP не является IPv6
		- Используется для проверки типа IP

	func (ip IP) IsLoopback() bool
		- Проверяет является ли IP loopback адресом
		- IPv4: 127.0.0.0/8
		- IPv6: ::1
		- Возвращает true если это loopback адрес

	func (ip IP) IsPrivate() bool
		- Проверяет является ли IP приватным адресом
		- IPv4: 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16
		- IPv6: fc00::/7
		- Возвращает true если это приватный адрес

	func (ip IP) IsMulticast() bool
		- Проверяет является ли IP multicast адресом
		- IPv4: 224.0.0.0/4
		- IPv6: ff00::/8
		- Возвращает true если это multicast адрес

	func (ip IP) IsUnspecified() bool
		- Проверяет является ли IP неспецифицированным адресом
		- IPv4: 0.0.0.0
		- IPv6: ::
		- Возвращает true если это неспецифицированный адрес

	func (ip IP) IsGlobalUnicast() bool
		- Проверяет является ли IP глобальным unicast адресом
		- Возвращает true если это глобальный unicast адрес
		- Исключает loopback, multicast, private адреса

	func (ip IP) IsLinkLocalUnicast() bool
		- Проверяет является ли IP link-local unicast адресом
		- IPv4: 169.254.0.0/16
		- IPv6: fe80::/10
		- Возвращает true если это link-local адрес

	func (ip IP) IsLinkLocalMulticast() bool
		- Проверяет является ли IP link-local multicast адресом
		- IPv4: 224.0.0.0/24
		- IPv6: ff02::/16
		- Возвращает true если это link-local multicast адрес

	func (ip IP) IsInterfaceLocalMulticast() bool
		- Проверяет является ли IP interface-local multicast адресом
		- IPv6: ff01::/16
		- Возвращает true если это interface-local multicast адрес

	func (ip IP) IsValid() bool
		- Проверяет является ли IP валидным
		- Возвращает true если IP имеет корректную длину
		- IPv4: 4 байта, IPv6: 16 байт
		- Используется для валидации IP адресов

	func (ip IP) Equal(other IP) bool
		- Сравнивает два IP адреса
		- Возвращает true если адреса одинаковые
		- Учитывает длину и содержимое
		- Используется для сравнения адресов

	func (ip IP) MarshalText() ([]byte, error)
		- Сериализует IP в текстовый формат
		- Возвращает байты для JSON/XML сериализации
		- Используется для сохранения IP в файлы

	func (ip *IP) UnmarshalText(text []byte) error
		- Десериализует IP из текстового формата
		- Парсит IP из строки
		- Используется для загрузки IP из файлов

	Особенности работы:
	- IP является срезом байтов
	- IPv4 адреса имеют длину 4 байта
	- IPv6 адреса имеют длину 16 байт
	- IP адреса могут быть созданы различными способами
	- IP адреса используются для сетевых операций

	Примеры использования:

	// Создание IP адресов
	ipv4 := net.ParseIP("192.168.1.1")
	ipv6 := net.ParseIP("2001:db8::1")
	ipv4FromBytes := net.IPv4(192, 168, 1, 1)

	// Проверка типа IP
	if ipv4.To4() != nil {
		fmt.Println("Это IPv4 адрес")
	}
	if ipv6.To16() != nil {
		fmt.Println("Это IPv6 адрес")
	}

	// Проверка свойств IP
	if ipv4.IsPrivate() {
		fmt.Println("Это приватный IP адрес")
	}
	if ipv4.IsLoopback() {
		fmt.Println("Это loopback IP адрес")
	}
	if ipv4.IsMulticast() {
		fmt.Println("Это multicast IP адрес")
	}

	// Сравнение IP адресов
	ip1 := net.ParseIP("192.168.1.1")
	ip2 := net.ParseIP("192.168.1.1")
	if ip1.Equal(ip2) {
		fmt.Println("IP адреса одинаковые")
	}

	// Строковое представление
	fmt.Printf("IPv4: %s\n", ipv4.String())
	fmt.Printf("IPv6: %s\n", ipv6.String())

	// Валидация IP
	if ipv4.IsValid() {
		fmt.Println("IP адрес валидный")
	}

	// Проверка глобальности
	if ipv4.IsGlobalUnicast() {
		fmt.Println("Это глобальный unicast адрес")
	}

	// Проверка link-local
	if ipv4.IsLinkLocalUnicast() {
		fmt.Println("Это link-local unicast адрес")
	}

	// Проверка multicast
	if ipv4.IsMulticast() {
		fmt.Println("Это multicast адрес")
	}

	// Проверка interface-local multicast (только IPv6)
	if ipv6.IsInterfaceLocalMulticast() {
		fmt.Println("Это interface-local multicast адрес")
	}

	// Проверка link-local multicast
	if ipv4.IsLinkLocalMulticast() {
		fmt.Println("Это link-local multicast адрес")
	}

	// Проверка неспецифицированного адреса
	unspecified := net.ParseIP("0.0.0.0")
	if unspecified.IsUnspecified() {
		fmt.Println("Это неспецифицированный адрес")
	}

	// Создание IP из байтов
	ipv4Bytes := []byte{192, 168, 1, 1}
	ipv4FromBytes = net.IP(ipv4Bytes)

	ipv6Bytes := []byte{0x20, 0x01, 0x0d, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	ipv6FromBytes := net.IP(ipv6Bytes)

	// Преобразование в байты
	ipv4Bytes = []byte(ipv4)
	ipv6Bytes = []byte(ipv6)

	// Использование в сетевых операциях
	tcpAddr := &net.TCPAddr{IP: ipv4, Port: 8080}
	udpAddr := &net.UDPAddr{IP: ipv4, Port: 8080}

	// Создание соединений
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Создание слушателей
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Логирование IP адресов
	log.Printf("Соединение с IP: %s\n", ipv4.String())

	// Отладка IP адресов
	fmt.Printf("IP: %s\n", ipv4.String())
	fmt.Printf("Тип: %T\n", ipv4)
	fmt.Printf("Длина: %d\n", len(ipv4))
	fmt.Printf("Байты: %v\n", []byte(ipv4))

	// Сериализация IP
	text, err := ipv4.MarshalText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Сериализованный IP: %s\n", string(text))

	// Десериализация IP
	var newIP net.IP
	err = newIP.UnmarshalText(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Десериализованный IP: %s\n", newIP.String())

	=== СПЕЦИАЛЬНЫЕ IP АДРЕСА ===

	// Loopback адреса
	loopback4 := net.ParseIP("127.0.0.1")
	loopback6 := net.ParseIP("::1")

	// Приватные адреса
	private4_1 := net.ParseIP("10.0.0.1")      // 10.0.0.0/8
	private4_2 := net.ParseIP("172.16.0.1")     // 172.16.0.0/12
	private4_3 := net.ParseIP("192.168.1.1")   // 192.168.0.0/16
	private6 := net.ParseIP("fc00::1")          // fc00::/7

	// Multicast адреса
	multicast4 := net.ParseIP("224.0.0.1")      // 224.0.0.0/4
	multicast6 := net.ParseIP("ff00::1")        // ff00::/8

	// Link-local адреса
	linkLocal4 := net.ParseIP("169.254.1.1")    // 169.254.0.0/16
	linkLocal6 := net.ParseIP("fe80::1")        // fe80::/10

	// Неспецифицированные адреса
	unspecified4 := net.ParseIP("0.0.0.0")
	unspecified6 := net.ParseIP("::")

	=== ОБРАБОТКА ОШИБОК ===

	1. Ошибки парсинга:
		- Невалидные IP адреса
		- Неправильный формат
		- Неподдерживаемые версии IP

	2. Ошибки преобразования:
		- Невозможность преобразования типа
		- Неподдерживаемые операции
		- Ошибки сериализации

	3. Ошибки валидации:
		- Невалидная длина IP
		- Неправильные байты
		- Неподдерживаемые форматы

	=== ПРОИЗВОДИТЕЛЬНОСТЬ ===

	1. Кэширование IP:
		- Кэшируйте часто используемые IP
		- Избегайте повторного парсинга
		- Используйте пулы IP адресов

	2. Оптимизация операций:
		- Используйте To4() и To16() для проверки типа
		- Избегайте лишних преобразований
		- Минимизируйте копирование

	3. Память:
		- IP адреса занимают мало памяти
		- IPv4: 4 байта, IPv6: 16 байт
		- Избегайте создания лишних копий

	=== БЕЗОПАСНОСТЬ ===

	1. Валидация IP:
		- Проверяйте все входящие IP
		- Валидируйте формат адресов
		- Проверяйте диапазоны адресов

	2. Фильтрация IP:
		- Блокируйте нежелательные IP
		- Используйте whitelist/blacklist
		- Проверяйте приватность адресов

	3. Логирование:
		- Логируйте все IP адреса
		- Отслеживайте подозрительные IP
		- Анализируйте паттерны доступа

	=== МОНИТОРИНГ ===

	1. Статистика IP:
		- Считайте количество соединений по IP
		- Мониторьте популярные IP
		- Отслеживайте изменения IP

	2. Анализ трафика:
		- Анализируйте источники трафика
		- Отслеживайте географическое распределение
		- Выявляйте аномалии

	3. Производительность:
		- Измеряйте время парсинга IP
		- Мониторьте использование памяти
		- Оптимизируйте операции с IP
*/
