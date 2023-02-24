## IO Reader Task

ReaderSplit - contains a code snippet written in Go that
defines a function called ReaderSplit.
The function takes a strings.Reader and an integer n as input,
and splits the contents of the reader into chunks of size n.
The function returns a slice of strings containing the chunks

```go
func ReaderSplit(strReader *strings.Reader, n int) []string {
	panic("Not implemented yet...")
}
```

- Task is to implement the `ReaderSplit` method based on the above requirements

Example Inputs:
```
input: helloworld
output: ["he", "ll", "ow", "or", "ld"]
```
```
input: GolangUnitedSchool
output: ["Golan", "gUnit", "edSch", "ool"]
```
```
input: ABCDEFGHIJKLMNOPQRSTUVWXYZ
output: ["ABCDEF", "GHIJKL", "MNOPQR", "STUVWX", "YZ"]
```