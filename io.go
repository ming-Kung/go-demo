package main

import (
	"errors"
	"fmt"
	"os"
)

type MyFile struct {
	FileName string
	ByteNum  int64
	File     *os.File
}

func (f *MyFile) Create() {
	file, err := os.Create(f.FileName)
	if err != nil {
		fmt.Println("创建文件失败，原因：", err)
		return
	}
	f.File = file
}

func (f *MyFile) Open() {
	file, err := os.Open(f.FileName)
	if err != nil {
		fmt.Println("打开文件失败，原因：", err)
		return
	}
	f.File = file
}

func (f *MyFile) Read() string {
	data := make([]byte, f.ByteNum)
	if f.File == nil {
		fmt.Println("读取文件失败，文件句柄为nil")
		return ""
	}
	_, err := f.File.Read(data)
	if err != nil {
		fmt.Println("读取文件失败，原因：", err)
		return ""
	}
	return string(data)
}

// 只能一次性全部读取，适合读取中小型文件
func (f *MyFile) ReadFile() string {
	data, err := os.ReadFile(f.FileName)
	if err != nil {
		fmt.Println("读取文件失败，原因：", err)
		return ""
	}
	return string(data)
}

func (f *MyFile) Write(str string) error {
	if f.File == nil {
		fmt.Println("写入文件失败，文件句柄为nil")
		return errors.New("文件句柄为nil")
	}
	_, err := f.File.Write([]byte(str))
	if err != nil {
		fmt.Println("写入文件失败，原因：", err)
	}
	return err
}

// 只能一次性全部写入，适合中小型文件
func (f *MyFile) WriteFile(str string) error {
	return os.WriteFile(f.FileName, []byte(str), 0644)
}

func main() {
	f := MyFile{
		FileName: "/Users/gm/www/go/src/write.txt",
		ByteNum:  1024,
	}
	defer f.File.Close()
	f.Create()
	f.Write("我是gm，我是男的\n")
	f.Write("今年才27岁")
	f.Open()
	fmt.Println("文件内容：", f.Read())
	fmt.Println("文件内容：", f.ReadFile())

	f.WriteFile("我是gm")
	fmt.Println("文件内容：", f.ReadFile())
}
