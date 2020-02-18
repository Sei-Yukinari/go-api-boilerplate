package interfaces

// A SQLHandler belong to the interface layer.
type SQLHandler interface {
	Begin() (Tx, error)
	Query(string, ...interface{}) (Row, error)
	Exec(string, ...interface{}) (Result, error)
}

// A Tx belong to the interface layer.
type Tx interface {
	Commit() error
	Rollback() error
	Exec(string, ...interface{}) (Result, error)
}

// A Result belong to the interface layer.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// A Row belong to the interface layer.
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}
