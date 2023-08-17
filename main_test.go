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
		{[]string{"-l", "dummy_file.txt"}, false},
		{[]string{"-invalid", "dummy_file.txt"}, false},
		{[]string{"-l", "-b", "dummy_file.txt"}, false},
		{[]string{"test_file.txt"}, true},
		{[]string{"-l", "test_file.txt"}, true},
		{[]string{"-n", "test_file.txt"}, true},
		{[]string{"-b", "test_file.txt"}, true},
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
