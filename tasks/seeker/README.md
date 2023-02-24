## IO Seeker Task

SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.

```go
func SeekTillHalfOfString(strReader *strings.Reader, n int) []string {
	panic("Not implemented yet...")
}
```

- Task is to implement the `SeekTillHalfOfString` method based on the above requirements

Example Inputs:
```
input: HELLO-WORLD
output: -WORLD
```
```
input: HELLOWORLD
output: WORLD
```
```
input: GOLANG-UNITED
output: -UNITED
```