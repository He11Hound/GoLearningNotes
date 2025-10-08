package os

/*
	Структура File

	type File struct {
		*file // указатель на внутреннюю структуру
	}

	type file struct {
		fd      int         // файловый дескриптор ОС
		name    string      // имя файла
		dirinfo *dirInfo    // данные о директории (если это директория)
		nepipe  int32       // количество ошибок "broken pipe"
	}

	Методы:

	func (f *File) Chdir() error - изменяет дирректорию
	func (f *File) Chmod(mode FileMode) error - изменяет права на файл
	func (f *File) Chown(uid, gid int) error - изменяет владельца файла
	func (f *File) Close() error - закрытие файла
	func (f *File) Fd() uintptr - возвращает дескриптор файла
	func (f *File) Name() string - возвращает имя файла
	func (f *File) ReadAt(b []byte, off int64) (n int, err error) - чтения данных из файла с конкретного смещения, не изменяя текущее положение указателя чтения файла. Скорее всего для работ с горутинами, при чтении одного файла
	func (f *File) ReadDir(n int) ([]DirEntry, error) - Читает записи директории
	func (f *File) ReadFrom(r io.Reader) (n int64, err error) - Читает данные из io.Reader и записывает их в файл
	func (f *File) Readdir(n int) ([]FileInfo, error) - Читает записи директории
	func (f *File) Readdirnames(n int) (names []string, err error) - Возвращает только имена файлов/директорий
	func (f *File) Seek(offset int64, whence int) (ret int64, err error) - Перемещает указатель чтения/записи в файле
	func (f *File) SetDeadline(t time.Time) error - Устанавливает общий таймаут для операций чтения/записи
	func (f *File) SetReadDeadline(t time.Time) error - Таймаут только для чтения
	func (f *File) SetWriteDeadline(t time.Time) error - Таймаут только для записи
	func (f *File) Stat() (FileInfo, error) - Получает информацию о файле
	func (f *File) Sync() error - Принудительно сбрасывает данные на диск
	func (f *File) SyscallConn() (syscall.RawConn, error) - Получает низкоуровневый доступ к файловому дескриптору
	func (f *File) Truncate(size int64) error - Обрезает или расширяет файл до указанного размера
	func (f *File) Write(b []byte) (n int, err error) - Записывает слайс байт в файл
	func (f *File) WriteAt(b []byte, off int64) (n int, err error) - Записывает данные в файл с конкретного смещения
	func (f *File) WriteString(s string) (n int, err error) - Записывает строку в файл
	func (f *File) WriteTo(w io.Writer) (n int64, err error) - Копирует содержимое файла в io.Writer
*/
