package infinite

import "strings"

type value []string

func encodeValue(s string) (value, error) {
	return strings.Split(s, "|"), nil
}

func (v value) decode() (string, error) {
	return strings.Join(v, "|"), nil
}
