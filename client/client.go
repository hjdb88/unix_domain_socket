package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var file = "/Users/hjdb88/Desktop/dream/unix_domain_socket/uds.sock" // 用于 unix domain socket 的文件

func main() {
	conn, err := net.Dial("unix", file) // 发起请求
	if err != nil {
		log.Fatal(err) // 如果发生错误，直接退出程序，因为请求失败所以不需要 close
	}
	defer conn.Close() // 习惯性的写上

	input := bufio.NewScanner(os.Stdin) // 创建一个读取输入的处理器
	reader := bufio.NewReader(conn)     // 创建一个读取网络的处理器
	for {
		fmt.Print("请输入需要发送的数据: ")       // 打印提示
		input.Scan()                    // 读取终端输入
		data := input.Text()            // 提取输入内容
		conn.Write([]byte(data + "\n")) // 将输入的内容发送出去，需要将 string 转 byte 加 \n  作为读取的分割符

		msg, err := reader.ReadString('\n') // 读取对端的数据
		if err != nil {
			log.Println(err)
		}
		fmt.Println(msg) // 打印接收的消息
	}
}
