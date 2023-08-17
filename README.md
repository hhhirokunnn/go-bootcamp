# overview

Enablement Bootcamp for Gopherizing 〜業務で使えるGoを学ぼう〜

KADAI

impl split command and its test code
to see split command specification, use `man split`

usable options are only
- `-l`
- `-n`
- `-b`

# usage

```sh
$ go build -o sample.exe main.go
$ ./sample.exe -l 3 sample.txt output-file-prefix
```
