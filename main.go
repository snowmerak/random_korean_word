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
