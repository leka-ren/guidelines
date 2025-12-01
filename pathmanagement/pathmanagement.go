package pathmanagement

import (
	"fmt"
	"path/filepath"
)

func PathAbs() {
	// Метод Abs возвращает абсолютный путь до папки откуда была вызвана команда, не включая ее в итоговую строку
	// если в аргументах вызова метода Abs указать какую то строку, то произойдет конкатенация полученного пути и переданной строки
	// Допустим мы на рабочем столе, пример: /Users/username/Desktop/someString
	extantionedAbsolutePath, _ := filepath.Abs("someString")
	fmt.Println(extantionedAbsolutePath)

	// аналогично с пустой строкой: /Users/username/Desktop
	emptyAbsolutePath, _ := filepath.Abs("")
	fmt.Println(emptyAbsolutePath)

	// А так же если передать "путь", который начинается со слэша(так начинается абсолютный путь), то никаких изменений не произойдет
	absolutePath, _ := filepath.Abs("/somePartOfPath/anotherOnePart/andAnotherOne")

	// В результате будет тот же "путь": /somePartOfPath/anotherOnePart/andAnotherOne
	fmt.Println(absolutePath)
}

func LastPathName() {
	// Base используется для получения последнего элемента из полученного пути(или прописанного в ручную через слэш)
	fileName := filepath.Base("/Users/username/Desktop")

	// Таким образом, из пути /Users/username/Desktop, мы получим Desktop
	fmt.Println(fileName)
}

func CutDirPath() {
	// Метод dir используют для получения пути, исключая последний элемент
	pathWithoutLast := filepath.Dir("/Users/username/Desktop")

	// Таким образом, из пути /Users/username/Desktop, мы получим /Users/username
	fmt.Println(pathWithoutLast)
}

func ExtCheck() {
	// Ext нужен для определения расширения файла и путь до файла включая сам файл, либо просто имя файла с указанным форматом
	extName := filepath.Ext("Users/username/Desktop/somefile.go")

	// Если передать строку вида Users/username/Desktop/somefile.go или просто имя файла somefile.go, то в выоде будет ".go"
	fmt.Println(extName)
}

func NewPath() {
	// Join используется для создания нового кастомного пути
	// можно передать неограниченное количество аргументов строкового типа и все они будут конкатенированы в одну строку через слэш
	newPath := filepath.Join("somestring1", "somestring2", "somestring3", "somestring4")

	// в результате будет образован "путь" somestring1/somestring2/somestring3/somestring4
	fmt.Println(newPath)
}
