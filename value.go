package infinite

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const maxNameLen = 255

// encodeValue encodes the given value to names.
//
// The value is first split into chunks. These chunks are encoded using base64.
// They are then enumerated. These will be returned as the names, to be saved as
// files to store data.
func encodeValue(bb []byte) ([]string, error) {
	var (
		nn        []name
		buf       bytes.Buffer
		maxBufLen int
	)

	next := func() {
		if buf.Len() > 0 {
			bb := make([]byte, buf.Len())
			copy(bb, buf.Bytes())

			n := name{
				index: len(nn),
				chunk: bb,
			}

			nn = append(nn, n)
			buf = bytes.Buffer{}
		}
		maxBufLen = (maxNameLen / 4 * 3) - digits(len(nn)) - 1
	}

	next()
	for _, b := range bb {
		if buf.Len() >= maxBufLen {
			next()
		}

		if err := buf.WriteByte(b); err != nil {
			return nil, err
		}
	}
	next()

	ss := make([]string, len(nn))

	for i, n := range nn {
		ss[i] = n.String()
	}

	return ss, nil
}

// decodeValue decodes the given names to a value.
//
// Will fail when chunks cannot be constructed in order or when an invalid value
// is found. The chunk and index are extracted from the name, and then sorted in
// order. The chunk will be decoded using base64 before being concatenated as
// the value.
func decodeValue(ss []string) ([]byte, error) {
	nn := make([]name, len(ss))

	for i, s := range ss {
		n, err := newName(s)
		if err != nil {
			return nil, err
		}

		nn[i] = n
	}

	sort.Sort(names(nn))

	var (
		bb   []byte
		prev = -1
	)

	for _, n := range nn {
		if n.index-1 != prev {
			return nil, ErrInvalidValue
		}
		prev = n.index

		bb = append(bb, n.chunk...)
	}

	return bb, nil
}

// name represents an enumerated chunk.
//
// name is an intermediate representation of encoding or decoding a value.
type name struct {
	index int
	chunk []byte
}

// newName creates a new name.
//
// Will fail when the wrong amount of components are found, or when the second
// component does not contain an integer.
func newName(s string) (name, error) {
	comps := strings.Split(s, ".")
	if len(comps) != 2 {
		return name{}, ErrInvalidValue
	}

	index, err := strconv.Atoi(comps[1])
	if err != nil {
		return name{}, ErrInvalidValue
	}

	chunk, err := base64.URLEncoding.DecodeString(comps[0])
	if err != nil {
		return name{}, ErrInvalidValue
	}

	return name{
		index: index,
		chunk: chunk,
	}, nil
}

func (n name) String() string {
	chunk := base64.URLEncoding.EncodeToString(n.chunk)
	return fmt.Sprintf("%s.%d", chunk, n.index)
}

// names represents a sortable slice of names.
type names []name

func (n names) Len() int {
	return len(n)
}

func (n names) Less(i int, j int) bool {
	return n[i].index < n[j].index
}

func (n names) Swap(i int, j int) {
	n[i], n[j] = n[j], n[i]
}

// digits returns the number of digits in the given number.
//
// If a negative number is given, it will first be converted to a positive
// one before its digits are checked.
func digits(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}

	var c int
	for n > 0 {
		n /= 10
		c++
	}
	return c
}
