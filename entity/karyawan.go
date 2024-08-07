package entity

//utk mysql
import "database/sql"

type Karyawan struct {
	id           int
	nama         string
	jabatan      string
	married      bool
	pasangan     sql.NullString
	tanggallahir sql.NullTime
	gaji         sql.NullInt64
}
