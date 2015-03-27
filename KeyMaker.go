// KeyMaker project KeyMaker.go
package KeyMaker

import (
	"bytes"
	"crypto/md5"
	"strings"
)

func KeysToObjectId(keys ...string) ([16]byte, error) {
	var buffer bytes.Buffer
	var err error
	var emptyResult [16]byte

	// Concatenate strings
	for i := range keys {
		_, err := buffer.WriteString(strings.ToLower(keys[i]))
		if err != nil {
			return emptyResult, err
		}
	}

	// Create MD5 Sum
	output := md5.Sum(buffer.Bytes())

	return output, err
}
