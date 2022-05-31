package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// go run ex10.go google.com 80
func main() {
  timeout := flag.Int("timeout", 5, "Timeout")
  flag.Parse()
  args := flag.Args()
  // устанавливаем tcp соединение к серверу с timeout
  conn, err := net.DialTimeout("tcp", args[0]+":"+args[1], time.Duration(*timeout)*time.Second)
  if err != nil {
    fmt.Println("dial error:", err)
    return
  }
  defer conn.Close() // закрываем сокет при выходе из функции

  // назначаем буфер для чтения
  reader := bufio.NewScanner(os.Stdin)
  go func() {
    reader := bufio.NewReader(conn)
    for {
      message, err := reader.ReadString('\n')

      if err == io.EOF { // ошибка Ctrl+D
        fmt.Println("Connection is closed EOF")
        os.Exit(0)
      }
      if err != nil {
        fmt.Println(err)
        continue
      }
      fmt.Print(message)
      fmt.Fprintf(os.Stdin, "hello")
    }
  }()

  // Сканирование чтения в сокет
  for reader.Scan() {
    _, err := fmt.Fprintf(conn, reader.Text()+" HTTP/1.0\r\n\r\n")
    if err != nil {
      fmt.Println(err)
      break
    }
  }
  fmt.Println("Exit")
}