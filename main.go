package random_korean_word

import "crypto/rand"

func Get() string {
	b := [2]byte{}
	rand.Read(b[:])
	m := (int(b[1])<<8 | int(b[0])) % len(data)
	i := 0
	for k := range data {
		if i == m {
			return k
		}
		i++
	}
	return "장미"
}
