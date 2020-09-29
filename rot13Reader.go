package main

/*
Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
For more info: https://tour.golang.org/methods/23
*/
import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(buff []byte) (int, error) {
	transientBuff := make([]byte, len(buff))
	var n int
	var err error
	n, err = rd.r.Read(transientBuff)
	if err != nil {
		return 0, err
	}
	for i, v := range transientBuff[:n] {
		if unicode.IsLetter(rune(v)) {
			//Only apply substitution if the character is an alphabet
			newVal := v + 13
			if string(v) == strings.ToUpper(string(v)) {
				//If upper case character
				if newVal > uint8('Z') {
					//If we go beyond 'Z'
					buff[i] = uint8('A') + (newVal % uint8('Z')) - 1
				} else {
					buff[i] = newVal
				}
			} else {
				//Lower case character
				if newVal > uint8('z') {
					//If we go beyond 'z'
					buff[i] = uint8('a') + (newVal % uint8('z')) - 1
				} else {
					buff[i] = newVal
				}
			}
		} else {
			buff[i] = v
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
