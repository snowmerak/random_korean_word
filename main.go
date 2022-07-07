package random_korean_word

import (
	"crypto/rand"
	"encoding/binary"
)

func Get() string {
	b := [8]byte{}
	rand.Read(b[:])
	m := binary.BigEndian.Uint64(b[:])
	return data[m%uint64(len(data))]
}

func GetLength(length int) string {
	result := ""
	for i := 1; i <= length; i++ {
		b := [8]byte{}
		rand.Read(b[:])
		m := binary.BigEndian.Uint64(b[:])
		if i != 1 {
			result = result + " " + data[m%uint64(len(data))]
		} else {
			result = data[m%uint64(len(data))]
		}
	}
	return result
}
