package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if !ValidateArgs(args) {
		DescribeUsage()
		return // noop
	}

	if len(args) == 3 {
		option, file := args[0], args[1]
		fmt.Println(option)
		fmt.Println(file)
	} else if len(args) == 1 {
		// only len(args) is 1
		CopyFile(args[0])
	} else {
		fmt.Errorf("unexpected")
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
	if len(args) > 3 {
		fmt.Printf("too much args")
		return false
	}
	if len(args) == 3 {
		op, num, file := args[0], args[1], args[2]
		if op == "-l" || op == "-n" || op == "-b" {
			// noop
		} else {
			fmt.Printf("unsupported options")
			return false
		}
		_, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("second args should be num")
			return false
		}
		if !ExistFile(file) {
			fmt.Printf("file not found")
			return false
		}
		return true
	} else if len(args) == 2 {
		return false
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
