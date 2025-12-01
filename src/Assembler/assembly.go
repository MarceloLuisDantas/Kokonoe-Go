package assembler

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func load_file(file_name string) string {
	data, err := os.ReadFile(file_name)
	check(err)
	return string(data)
}

func Assembler() {
	fmt.Println("Assembler: ")

	if len(os.Args) < 2 {
		fmt.Println("Arquivo não dado")
		os.Exit(0)
	}

	file_name := os.Args[1]
	data := load_file(file_name)

	tk := newTokenizer(data)
	tk.Tokenize()
}
