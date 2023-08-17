package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	if !ValidateArgs(args) {
		DescribeUsage()
		return // noop
	}

	if len(args) == 4 {
		option, strNum, file, name := args[0], args[1], args[2], args[3]
		num, _ := strconv.ParseUint(strNum, 10, 64)
		if option == "-l" {
			SplitByLine(file, num, name)
		} else if option == "-n" {
			SplitByNum(file, num, name)
		} else if option == "-b" {
			SplitByByte(file, num, name)
		} else {
			fmt.Errorf("Argument error: unsupported option %v", option)
		}
	} else if len(args) == 1 {
		// only len(args) is 1
		CopyFile(args[0])
	} else {
		fmt.Errorf("no case to come")
	}

	return
}

func DescribeUsage() {
	fmt.Printf(`
	usage
	go run main.go [OPTION]... [FILE [PREFIX]] 
	
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
	if len(args) > 4 {
		fmt.Printf("too much args")
		return false
	}
	if len(args) == 4 {
		op, num, file, _ := args[0], args[1], args[2], args[3]
		if op == "-l" || op == "-n" || op == "-b" {
			// noop
		} else {
			fmt.Printf("unsupported options")
			return false
		}
		_, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			fmt.Printf("second args should be num")
			return false
		}
		if !ExistFile(file) {
			fmt.Printf("file not found")
			return false
		}
		return true
	} else if len(args) == 1 {
		file := args[0]
		if !ExistFile(file) {
			fmt.Printf("file not found")
			return false
		}

		return true
	}

	return false
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

func SplitByLine(f string, n uint64, p string) error {
	data, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	var tmp []string
	for i, l := range lines {
		tmp = append(tmp, l)
		if uint64(len(tmp)) == n {
			newfile, err := os.Create(p + strconv.Itoa(i))
			if err != nil {
				return fmt.Errorf("file create error")
			}
			for _, t := range tmp {
				_, err := newfile.WriteString(t + "\n")
				if err != nil {
					newfile.Close()
					return fmt.Errorf("file write error")
				}
			}
			newfile.Close()
			tmp = []string{}
		}
	}
	if len(tmp) > 0 {
		newfile, err := os.Create(p + strconv.Itoa(len(lines)))
		if err != nil {
			return fmt.Errorf("file create error")
		}
		for _, t := range tmp {
			_, err := newfile.WriteString(t + "\n")
			if err != nil {
				newfile.Close()
				return fmt.Errorf("file write error")
			}
		}
		newfile.Close()
	}

	return nil
}

func SplitByNum(f string, n uint64, p string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := int(uint64(fileinfo.Size()) / n)
	var bf []byte
	for i := 0; i < int(n); i++ {
		if i < int(n)-1 {
			bf = make([]byte, fileSize)
		} else {
			bf = make([]byte, fileSize+int(uint64(fileinfo.Size())%n))
		}
		_, err := file.Read(bf)
		if err != nil {
			return err
		}

		newfile, err := os.Create(p + strconv.Itoa(i))
		if err != nil {
			return err
		}
		_, err = newfile.Write(bf)
		newfile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func SplitByByte(f string, n uint64, p string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return err
	}

	filenum := int(uint64(fileinfo.Size()) / n)
	if uint64(fileinfo.Size())%n != 0 {
		filenum++
	}

	bf := make([]byte, n)
	for i := 0; i < filenum; i++ {
		_, err := file.Read(bf)
		if err != nil {
			return err
		}

		newfile, err := os.Create(p + strconv.Itoa(i))
		if err != nil {
			return err
		}
		_, err = newfile.Write(bf)
		if err != nil {
			newfile.Close()
			return err
		}
		newfile.Close()
	}
	return nil
}
