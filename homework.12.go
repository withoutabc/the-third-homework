package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建文件
	file1, err := os.Create("plan.txt")
	if err != nil {
		fmt.Println("创建文件失败")
	}
	fmt.Printf("成功创建：plan.txt\n")
	//写入文字
	_, err = file1.WriteString("I’m not afraid of difficulties and insist on learning programming")
	if err != nil {
		return
	}
	fmt.Println("已写入")
	//打开文件
	file2, err := os.Open("plan.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	fmt.Println("打开文件")
	bytes := make([]byte, 200)
	_, err = file2.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("读取：%s\n", bytes)
	// 程序运行结束关闭文件
	defer func() {
		err := file2.Close()
		if err != nil {
			return
		}
		fmt.Println("文件关闭")
	}()

}
