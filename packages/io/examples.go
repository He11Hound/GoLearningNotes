package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// CopyExample демонстрирует использование функции io.Copy
func CopyExample() {
	// Создание источника данных
	source := strings.NewReader("Hello, World! This is a test string.")
	
	// Создание буфера для записи
	var buffer bytes.Buffer
	
	// Копирование данных
	written, err := io.Copy(&buffer, source)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Скопировано %d байтов\n", written)
	fmt.Printf("Содержимое буфера: %s\n", buffer.String())
}

// CopyBufferExample демонстрирует использование функции io.CopyBuffer
func CopyBufferExample() {
	// Создание источника данных
	source := strings.NewReader("Large data string for buffered copying")
	
	// Создание буфера для записи
	var buffer bytes.Buffer
	
	// Создание буфера для копирования
	copyBuffer := make([]byte, 8) // буфер размером 8 байтов
	
	// Копирование с использованием буфера
	written, err := io.CopyBuffer(&buffer, source, copyBuffer)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Скопировано %d байтов с буфером\n", written)
	fmt.Printf("Содержимое: %s\n", buffer.String())
}

// CopyNExample демонстрирует использование функции io.CopyN
func CopyNExample() {
	// Создание источника данных
	source := strings.NewReader("This is a long string with many characters")
	
	// Создание буфера для записи
	var buffer bytes.Buffer
	
	// Копирование только первых 10 байтов
	written, err := io.CopyN(&buffer, source, 10)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Скопировано %d байтов (ограничено 10)\n", written)
	fmt.Printf("Содержимое: %s\n", buffer.String())
}

// ReadAllExample демонстрирует использование функции io.ReadAll
func ReadAllExample() {
	// Создание источника данных
	source := strings.NewReader("Complete file content")
	
	// Чтение всех данных
	data, err := io.ReadAll(source)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Прочитано %d байтов\n", len(data))
	fmt.Printf("Содержимое: %s\n", string(data))
}

// ReadAtLeastExample демонстрирует использование функции io.ReadAtLeast
func ReadAtLeastExample() {
	// Создание источника данных
	source := strings.NewReader("Short data")
	
	// Создание буфера
	buffer := make([]byte, 20)
	
	// Чтение минимум 5 байтов
	n, err := io.ReadAtLeast(source, buffer, 5)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Прочитано %d байтов (минимум 5)\n", n)
	fmt.Printf("Содержимое: %s\n", string(buffer[:n]))
}

// ReadFullExample демонстрирует использование функции io.ReadFull
func ReadFullExample() {
	// Создание источника данных
	source := strings.NewReader("Full data")
	
	// Создание буфера размером 9 байтов
	buffer := make([]byte, 9)
	
	// Чтение ровно 9 байтов
	n, err := io.ReadFull(source, buffer)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Прочитано %d байтов (полностью)\n", n)
	fmt.Printf("Содержимое: %s\n", string(buffer))
}

// WriteStringExample демонстрирует использование функции io.WriteString
func WriteStringExample() {
	// Создание буфера для записи
	var buffer bytes.Buffer
	
	// Запись строки
	n, err := io.WriteString(&buffer, "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Записано %d байтов\n", n)
	fmt.Printf("Содержимое буфера: %s\n", buffer.String())
}

// PipeExample демонстрирует использование функции io.Pipe
func PipeExample() {
	// Создание pipe
	reader, writer := io.Pipe()
	
	// Горутина для записи
	go func() {
		defer writer.Close()
		writer.Write([]byte("Data from pipe"))
	}()
	
	// Чтение из pipe
	buffer := make([]byte, 20)
	n, err := reader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Прочитано из pipe: %s\n", string(buffer[:n]))
}

// LimitReaderExample демонстрирует использование функции io.LimitReader
func LimitReaderExample() {
	// Создание источника данных
	source := strings.NewReader("This is a very long string with many characters")
	
	// Ограничение чтения до 10 байтов
	limitedReader := io.LimitReader(source, 10)
	
	// Чтение данных
	data, err := io.ReadAll(limitedReader)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Прочитано (ограничено 10 байтов): %s\n", string(data))
}

// TeeReaderExample демонстрирует использование функции io.TeeReader
func TeeReaderExample() {
	// Создание источника данных
	source := strings.NewReader("Original data")
	
	// Создание буфера для дублирования
	var duplicateBuffer bytes.Buffer
	
	// Создание TeeReader
	teeReader := io.TeeReader(source, &duplicateBuffer)
	
	// Чтение данных
	data, err := io.ReadAll(teeReader)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Оригинальные данные: %s\n", string(data))
	fmt.Printf("Дублированные данные: %s\n", duplicateBuffer.String())
}

// MultiReaderExample демонстрирует использование функции io.MultiReader
func MultiReaderExample() {
	// Создание нескольких источников
	reader1 := strings.NewReader("First part ")
	reader2 := strings.NewReader("Second part ")
	reader3 := strings.NewReader("Third part")
	
	// Объединение Reader
	multiReader := io.MultiReader(reader1, reader2, reader3)
	
	// Чтение из объединенного Reader
	data, err := io.ReadAll(multiReader)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Объединенные данные: %s\n", string(data))
}

// MultiWriterExample демонстрирует использование функции io.MultiWriter
func MultiWriterExample() {
	// Создание нескольких Writer
	var buffer1 bytes.Buffer
	var buffer2 bytes.Buffer
	
	// Объединение Writer
	multiWriter := io.MultiWriter(&buffer1, &buffer2)
	
	// Запись в объединенный Writer
	n, err := io.WriteString(multiWriter, "Data for both buffers")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Записано %d байтов\n", n)
	fmt.Printf("Буфер 1: %s\n", buffer1.String())
	fmt.Printf("Буфер 2: %s\n", buffer2.String())
}

// NewSectionReaderExample демонстрирует использование функции io.NewSectionReader
func NewSectionReaderExample() {
	// Создание данных
	data := []byte("This is a long string with many characters for section reading")
	
	// Создание ReaderAt
	readerAt := bytes.NewReader(data)
	
	// Создание SectionReader для чтения с позиции 10, 20 байтов
	sectionReader := io.NewSectionReader(readerAt, 10, 20)
	
	// Чтение секции
	sectionData, err := io.ReadAll(sectionReader)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Секция данных: %s\n", string(sectionData))
}

// NewReaderExample демонстрирует использование функции io.NewReader
func NewReaderExample() {
	// Создание Reader из строки
	reader := strings.NewReader("String reader example")
	
	// Чтение данных
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Данные из строки: %s\n", string(data))
}

// NewWriterExample демонстрирует использование функции io.NewWriter
func NewWriterExample() {
	// Создание Writer для строк
	var builder strings.Builder
	
	// Запись данных
	n, err := io.WriteString(&builder, "String builder example")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Записано %d байтов\n", n)
	fmt.Printf("Построенная строка: %s\n", builder.String())
}

// NewBytesReaderExample демонстрирует использование функции io.NewReader для байтов
func NewBytesReaderExample() {
	// Создание Reader из байтов
	data := []byte("Bytes reader example")
	reader := bytes.NewReader(data)
	
	// Чтение данных
	readData, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Данные из байтов: %s\n", string(readData))
}

// NewBytesWriterExample демонстрирует использование функции io.NewWriter для байтов
func NewBytesWriterExample() {
	// Создание Writer для байтов
	var buffer bytes.Buffer
	
	// Запись данных
	n, err := io.WriteString(&buffer, "Bytes buffer example")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Записано %d байтов\n", n)
	fmt.Printf("Содержимое буфера: %s\n", buffer.String())
}

// DiscardExample демонстрирует использование функции io.Discard
func DiscardExample() {
	// Создание источника данных
	source := strings.NewReader("Data to be discarded")
	
	// Запись в Discard (данные отбрасываются)
	written, err := io.Copy(io.Discard, source)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Отброшено %d байтов\n", written)
}

// NopCloserExample демонстрирует использование функции io.NopCloser
func NopCloserExample() {
	// Создание Reader
	reader := strings.NewReader("Data with NopCloser")
	
	// Оборачивание в ReadCloser
	readCloser := io.NopCloser(reader)
	
	// Чтение данных
	data, err := io.ReadAll(readCloser)
	if err != nil {
		log.Fatal(err)
	}
	
	// Закрытие (не выполняет никаких действий)
	err = readCloser.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Данные: %s\n", string(data))
	fmt.Println("ReadCloser успешно закрыт")
}

// SeekExample демонстрирует использование констант Seek
func SeekExample() {
	// Создание временного файла
	file, err := os.CreateTemp("", "seek_example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	
	// Запись данных
	file.WriteString("This is a test file for seeking")
	
	// Позиционирование в начало
	pos, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Позиция в начале: %d\n", pos)
	
	// Позиционирование в конец
	pos, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Позиция в конце: %d\n", pos)
	
	// Позиционирование относительно текущей позиции
	pos, err = file.Seek(-10, io.SeekCurrent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Позиция на 10 байтов назад: %d\n", pos)
}

// FileCopyExample демонстрирует копирование файлов
func FileCopyExample() {
	// Создание исходного файла
	sourceFile, err := os.CreateTemp("", "source")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(sourceFile.Name())
	defer sourceFile.Close()
	
	// Запись данных в исходный файл
	sourceFile.WriteString("This is source file content")
	sourceFile.Seek(0, io.SeekStart)
	
	// Создание целевого файла
	destFile, err := os.CreateTemp("", "destination")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(destFile.Name())
	defer destFile.Close()
	
	// Копирование файла
	written, err := io.Copy(destFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Скопировано %d байтов из файла в файл\n", written)
}

// BufferOperationsExample демонстрирует операции с буферами
func BufferOperationsExample() {
	// Создание буфера
	var buffer bytes.Buffer
	
	// Запись в буфер
	io.WriteString(&buffer, "Hello, ")
	io.WriteString(&buffer, "World!")
	
	// Чтение из буфера
	data, err := io.ReadAll(&buffer)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Данные из буфера: %s\n", string(data))
}

// ErrorHandlingExample демонстрирует обработку ошибок
func ErrorHandlingExample() {
	// Создание Reader который возвращает ошибку
	errorReader := &errorReader{err: fmt.Errorf("simulated error")}
	
	// Попытка чтения
	buffer := make([]byte, 10)
	_, err := errorReader.Read(buffer)
	if err != nil {
		fmt.Printf("Обработана ошибка: %v\n", err)
	}
}

// errorReader - простой Reader который возвращает ошибку
type errorReader struct {
	err error
}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, r.err
}
