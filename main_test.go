package main

import (
	"os"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	f, _ := os.Create("test_file.txt")
	f.Sync()
	defer os.Remove("test_file.txt")

	tests := []struct {
		args []string
		exp  bool
	}{
		{[]string{}, false},
		{[]string{"dummy_file.txt"}, false},
		{[]string{"-l", "test_file.txt"}, false},
		{[]string{"-l", "-b", "test_file.txt"}, false},
		{[]string{"-l", "-b", "1", "test_file.txt", "hoge"}, false},
		{[]string{"-l", "1", "test_file.txt"}, false},
		{[]string{"-invalid", "1", "test_file.txt", "hoge"}, false},
		{[]string{"test_file.txt"}, true},
		{[]string{"-l", "1", "test_file.txt", "hoge"}, true},
		{[]string{"-n", "1", "test_file.txt", "hoge"}, true},
		{[]string{"-b", "1", "test_file.txt", "hoge"}, true},
	}

	for _, tt := range tests {
		got := ValidateArgs(tt.args)
		if got != tt.exp {
			t.Errorf("given: %v / got: %v / expected: %v", tt.args, got, tt.exp)
		}
	}
}

func TestExistFile(t *testing.T) {
	f, _ := os.Create("test_file.txt")
	f.Sync()
	defer os.Remove("test_file.txt")

	tests := []struct {
		arg string
		exp bool
	}{
		{"dummy_file.txt", false},
		{"test_file.txt", true},
	}

	for _, tt := range tests {
		got := ExistFile(tt.arg)
		if got != tt.exp {
			t.Errorf("given: %v / got: %v / expected: %v", tt.arg, got, tt.exp)
		}
	}
}

func TestCopyFile(t *testing.T) {
	f, _ := os.Create("test_file.txt")
	f.Sync()
	defer os.Remove("test_file.txt")

	case1 := CopyFile("")
	if case1 == nil {
		t.Errorf("given: %v / should return error", "")
	}

	case2 := CopyFile("test_file.txt")
	if case2 != nil {
		t.Errorf("given: %v / should not return error: %v", "test_file.txt", case2)
	}
	defer os.Remove("new_test_file.txt")
}

func TestSplitByLine(t *testing.T) {
	// TODO
}

func TestSplitByNum(t *testing.T) {
	// TODO
}

func TestSplitByByte(t *testing.T) {
	// TODO
}
