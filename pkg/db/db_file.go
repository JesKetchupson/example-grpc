package db

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

//separator for Split function can be replaced with message lenght
const separator uint64 = 18446744073709551614
const shift = unsafe.Sizeof(separator)

type FileDB struct {
	Path   string
	sourse *os.File
}

func (fdb *FileDB) Open() (DB, error) {
	f, err := os.OpenFile(fdb.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileDB{sourse: f}, nil
}
func (fdb *FileDB) Insert(elem []byte) error {
	l := make([]byte, shift)
	binary.LittleEndian.PutUint64(l, uint64(separator))
	elem = append(elem, l...)

	_, err := fdb.sourse.Write(elem)

	if err != nil {
		fmt.Println(err)
	}

	err = fdb.sourse.Close()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (fdb *FileDB) Delete() error {
	return nil
}

func (fdb *FileDB) Close() error {
	return fdb.sourse.Close()
}

func (fdb *FileDB) List() ([][]byte, error) {
	b, err := ioutil.ReadFile(fdb.sourse.Name())
	if err != nil {
		return nil, err
	}
	l := make([]byte, shift)
	binary.LittleEndian.PutUint64(l, uint64(separator))
	out := bytes.Split(b, l)
	return out, err
}

func (fdb *FileDB) Update(upd []byte, id uint64) error {
	return nil
}
