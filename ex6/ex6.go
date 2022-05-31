package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Fields    int
	Delim     string
	Separated bool
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		flag.IntVar(&instance.Fields, "f", 0, "fields")
		flag.StringVar(&instance.Delim, "d", "\t", "delim")
		flag.BoolVar(&instance.Separated, "s", false, "separated")
		flag.Parse()
	})
	return instance
}

func main() {
	cfg := NewConfig()

	if cfg.Fields == 0 {
		log.Fatalln("-f must be > 0")
	}

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	res := Cut(text, cfg)
	fmt.Println(res)
}

func Cut(str string, cfg Config) string {
	if cfg.Separated {
		if !strings.Contains(str, cfg.Delim) {
			return ""
		}
	}
	spl := strings.Split(str, cfg.Delim)
	if cfg.Fields <= len(spl) {
		return spl[cfg.Fields-1] + "\n"
	}
	return ""
}
