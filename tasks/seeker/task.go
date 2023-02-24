package seeker

import (
	"io"
	"math"
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/
func SeekTillHalfOfString(strReader *strings.Reader) string {
	halfSize := float64(strReader.Size()) / 2.0
	readSize := int64(math.Ceil(float64(halfSize)))
	strReader.Seek(int64(halfSize), io.SeekStart)
	readerBuf := make([]byte, readSize)
	strReader.Read(readerBuf)
	return string(readerBuf)
}
