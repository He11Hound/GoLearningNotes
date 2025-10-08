package io

/*
	ФУНКЦИИ ПАКЕТА IO

	Пакет io предоставляет базовые интерфейсы для операций ввода-вывода
	и их реализации. Это один из самых важных пакетов в Go, так как он
	определяет основные интерфейсы для работы с потоками данных.

	=== ОСНОВНЫЕ ИНТЕРФЕЙСЫ ===

	Reader interface
		- Базовый интерфейс для чтения данных
		- Метод: Read([]byte) (int, error)
		- Используется для чтения из файлов, сетевых соединений, буферов

	Writer interface
		- Базовый интерфейс для записи данных
		- Метод: Write([]byte) (int, error)
		- Используется для записи в файлы, сетевые соединения, буферы

	Closer interface
		- Интерфейс для закрытия ресурсов
		- Метод: Close() error
		- Используется для освобождения ресурсов

	Seeker interface
		- Интерфейс для изменения позиции в потоке
		- Метод: Seek(offset int64, whence int) (int64, error)
		- Используется для навигации по файлам

	=== КОМБИНИРОВАННЫЕ ИНТЕРФЕЙСЫ ===

	ReadWriter interface
		- Комбинация Reader и Writer
		- Позволяет читать и записывать данные

	ReadCloser interface
		- Комбинация Reader и Closer
		- Позволяет читать и закрывать ресурс

	WriteCloser interface
		- Комбинация Writer и Closer
		- Позволяет записывать и закрывать ресурс

	ReadWriteCloser interface
		- Комбинация Reader, Writer и Closer
		- Полнофункциональный интерфейс для работы с ресурсами

	ReadSeeker interface
		- Комбинация Reader и Seeker
		- Позволяет читать и изменять позицию

	WriteSeeker interface
		- Комбинация Writer и Seeker
		- Позволяет записывать и изменять позицию

	ReadWriteSeeker interface
		- Комбинация Reader, Writer и Seeker
		- Полнофункциональный интерфейс с навигацией

	=== ФУНКЦИИ ДЛЯ КОПИРОВАНИЯ ДАННЫХ ===

	func Copy(dst Writer, src Reader) (written int64, err error)
		- Копирует данные из src в dst
		- Возвращает количество скопированных байтов
		- Автоматически обрабатывает буферизацию

	func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
		- Копирует данные с использованием предоставленного буфера
		- Позволяет контролировать размер буфера
		- Более эффективен для больших объемов данных

	func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
		- Копирует ровно n байтов из src в dst
		- Останавливается после копирования n байтов
		- Полезно для ограничения объема данных

	=== ФУНКЦИИ ДЛЯ ЧТЕНИЯ ДАННЫХ ===

	func ReadAll(r Reader) ([]byte, error)
		- Читает все данные из Reader до EOF
		- Возвращает содержимое как []byte
		- Полезно для небольших файлов

	func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
		- Читает минимум min байтов в буфер
		- Возвращает количество прочитанных байтов
		- Возвращает ошибку если не удалось прочитать минимум

	func ReadFull(r Reader, buf []byte) (n int, err error)
		- Читает ровно len(buf) байтов в буфер
		- Эквивалентно ReadAtLeast(r, buf, len(buf))
		- Возвращает ошибку если не удалось прочитать все

	=== ФУНКЦИИ ДЛЯ ЗАПИСИ ДАННЫХ ===

	func WriteString(w Writer, s string) (n int, err error)
		- Записывает строку в Writer
		- Эквивалентно w.Write([]byte(s))
		- Удобная функция для записи строк

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С ПОТОКАМИ ===

	func Pipe() (*PipeReader, *PipeWriter)
		- Создает синхронную пару pipe
		- PipeReader и PipeWriter связаны
		- Полезно для связи между горутинами

	func LimitReader(r Reader, n int64) Reader
		- Ограничивает количество читаемых байтов
		- Возвращает Reader который читает максимум n байтов
		- Полезно для ограничения размера данных

	func TeeReader(r Reader, w Writer) Reader
		- Создает Reader который дублирует данные в Writer
		- Читает из r и одновременно записывает в w
		- Полезно для логирования или мониторинга

	func MultiReader(readers ...Reader) Reader
		- Объединяет несколько Reader в один
		- Читает последовательно из всех Reader
		- Полезно для объединения потоков

	func MultiWriter(writers ...Writer) Writer
		- Создает Writer который записывает в несколько Writer
		- Записывает данные во все Writer одновременно
		- Полезно для дублирования вывода

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С БУФЕРАМИ ===

	func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
		- Создает Reader для чтения части данных
		- Читает n байтов начиная с позиции off
		- Полезно для чтения фрагментов файлов

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С СТРОКАМИ ===

	func NewReader(s string) *strings.Reader
		- Создает Reader из строки
		- Позволяет читать строку как поток данных
		- Полезно для тестирования

	func NewWriter() *strings.Builder
		- Создает Writer для построения строк
		- Позволяет записывать данные в строку
		- Полезно для построения строк

	=== ФУНКЦИИ ДЛЯ РАБОТЫ С БАЙТАМИ ===

	func NewReader(b []byte) *bytes.Reader
		- Создает Reader из среза байтов
		- Позволяет читать байты как поток данных
		- Полезно для работы с данными в памяти

	func NewWriter() *bytes.Buffer
		- Создает Writer для записи в буфер
		- Позволяет записывать данные в буфер
		- Полезно для накопления данных

	=== СПЕЦИАЛЬНЫЕ ФУНКЦИИ ===

	func Discard() Writer
		- Возвращает Writer который отбрасывает все данные
		- Полезно для игнорирования вывода
		- Эквивалентно /dev/null в Unix

	func NopCloser(r Reader) ReadCloser
		- Оборачивает Reader в ReadCloser с пустой функцией Close
		- Полезно когда нужно добавить метод Close
		- Не выполняет никаких действий при закрытии



*/
