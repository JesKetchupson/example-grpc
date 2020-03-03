package db

//DB is Interface for Database connection
type DB interface {
	//TODO: sep openR openW
	Open() (DB, error)
	Insert(b []byte) error
	Delete() error
	List() (b [][]byte, err error)
	Close() error
}
