package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if !ValidateArgs(args) {
		DescribeUsage()
		// noop
		return
	}

	if len(args) == 2 {
		option, file := args[0], args[1]
		fmt.Println(option)
		fmt.Println(file)
	} else {
		// only len(args) is 1
		CopyFile(args[0])
	}

	return
}

func DescribeUsage() {
	fmt.Printf(`
	usage
	go run main.go [OPTION]... [FILE]
	
	DESCRIOTION
	  -l
		split by line number
	  -n
		split per specific number
	  -b
		split by byte size
			`)
}

func ValidateArgs(args []string) bool {
	if len(args) < 1 {
		fmt.Printf("no args")
		return false
	}
	if len(args) > 2 {
		fmt.Printf("too much args")
		return false
	}
	if len(args) == 2 {
		op, file := args[0], args[1]
		if op == "-l" || op == "-n" || op == "-b" {
			// noop
		} else {
			fmt.Printf("unsupported options")
			return false
		}
		if !ExistFile(file) {
			fmt.Printf("file not found")
			return false
		}
		return true
	}

	file := args[0]
	if !ExistFile(file) {
		fmt.Printf("file not found")
		return false
	}

	return true
}

func ExistFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return true
}

func CopyFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("file open error")
	}
	defer f.Close()

	cf, err := os.Create("new_" + file)
	if err != nil {
		return fmt.Errorf("file create error")
	}
	defer cf.Close()

	_, err = io.Copy(cf, f)
	if err != nil {
		return fmt.Errorf("write file error")
	}

	return cf.Sync()
}
