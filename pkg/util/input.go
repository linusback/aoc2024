package util

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type (
	RowFunc       func(row []byte, nr int) error
	ExecFunc      func(*bufio.Reader, RowFunc, ...RowFunc) error
	MultiRowFunc  func(row [][]byte, nr int) error
	MultiExecFunc func(*bufio.Reader, int, MultiRowFunc, ...MultiRowFunc) error
)

type Indexer[P interface{ ~uint16 | ~uint32 | ~uint64 }] interface {
	New(y, x int) P
}

type PositionMap[T any, P interface{ ~uint16 | ~uint32 | ~uint64 }] struct {
	Map    []T
	MaxPos P
}

func ToMapOfPositionsByte[P interface{ ~uint16 | ~uint32 | ~uint64 }](filename string, pos Indexer[P]) (posMap *PositionMap[byte, P], err error) {
	return ToMapOfPositions(filename, pos, func(b byte) byte {
		return b
	})
}

func ToMapOfPositions[P interface{ ~uint16 | ~uint32 | ~uint64 }, T any](filename string, pos Indexer[P], transform func(byte) T) (posMap *PositionMap[T, P], err error) {
	var (
		y, x int
		row  []byte
		b    byte
	)
	posMap = new(PositionMap[T, P])
	data := make([][]byte, 0, 255)
	err = DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			x = len(row) - 1
		}
		y = nr
		data = append(data, row)
		return nil
	})
	if err != nil {
		return nil, err
	}
	posMap.MaxPos = pos.New(y, x)
	posMap.Map = make([]T, posMap.MaxPos+1)
	for y, row = range data {
		for x, b = range row {
			posMap.Map[pos.New(y, x)] = transform(b)
		}
	}
	return posMap, nil
}

func DoFile(filename string, execFunc ExecFunc, rowFunc RowFunc, extra ...RowFunc) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			err = errors.Join(err, err2)
		}
	}()
	return execFunc(bufio.NewReader(f), rowFunc, extra...)
}

func DoEachRowFile(filename string, rowFunc RowFunc, extra ...RowFunc) (err error) {
	return DoFile(filename, DoEachRowBuff, rowFunc, extra...)
}

func DoEachRowBuff(r *bufio.Reader, f RowFunc, extra ...RowFunc) (err error) {
	var (
		row  []byte
		done bool
	)
	for i := 0; !done; i++ {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
			err = nil
		}
		if len(row) > 0 {
			// removing delim
			row = row[:len(row)-1]
		}
		if len(row) == 0 {
			if len(extra) > 0 {
				f = extra[0]
				extra = extra[:len(extra)-1]
				continue
			}
			return
		}

		err = f(row, i)
		if err != nil {
			return
		}
	}
	return
}

func DoFileN(filename string, n int, execFunc MultiExecFunc, rowFunc MultiRowFunc, extra ...MultiRowFunc) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			err = errors.Join(err, err2)
		}
	}()
	return execFunc(bufio.NewReader(f), n, rowFunc, extra...)
}

func DoEachRowFileN(filename string, n int, rowFunc MultiRowFunc, extra ...MultiRowFunc) (err error) {
	return DoFileN(filename, n, DoEachRowBuffN, rowFunc, extra...)
}

func DoEachRowBuffN(r *bufio.Reader, n int, f MultiRowFunc, extra ...MultiRowFunc) (err error) {
	var (
		done bool
		row  []byte
		rows = make([][]byte, 0, n)
		i    int
	)
	for i = 0; !done; i++ {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
			err = nil
		}
		if len(row) > 0 {
			// removing delim
			row = row[:len(row)-1]
		}
		if len(row) > 0 {
			rows = append(rows, row)
		} else {
			if len(extra) > 0 {
				f = extra[0]
				extra = extra[:len(extra)-1]
				continue
			}
			return
		}
		if len(rows) == n {
			err = f(rows, i)
			if err != nil {
				return
			}
			rows = rows[:0]
		}
	}
	if len(rows) > 0 {
		err = f(rows, i)
		if err != nil {
			return
		}
		rows = rows[:0]
	}

	return
}
