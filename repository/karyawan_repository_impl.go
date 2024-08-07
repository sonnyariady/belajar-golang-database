package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type karyawanRepositoryImpl struct {
	DB *sql.DB
}

func NewKaryawanRepository(db *sql.DB) KaryawanRepository {
	return &karyawanRepositoryImpl{DB: db}
}

func (repository *karyawanRepositoryImpl) Insert(ctx context.Context, karyawan entity.Karyawan) (entity.Karyawan, error) {
	sqlcmd := "insert into Karyawan(nama, married, pasangan, jabatan, tanggallahir, gaji) VALUES(?,?,?,?,?,?)"
	hasil, err := repository.DB.ExecContext(ctx, sqlcmd, karyawan.nama, karyawan.married, karyawan.pasangan, karyawan.jabatan, karyawan.tanggallahir, karyawan.gaji)
	if err != nil {
		return karyawan, err
	}
	id, err := hasil.LastInsertId()
	if err != nil {
		return karyawan, err
	}
	karyawan.id = id
	return karyawan, nil
}

func (repository *karyawanRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Karyawan, error) {
	sqlquery := "select id,nama,married, pasangan, jabatan, tanggallahir, gaji where id=? from karyawan limit 1"
	baris, err := repository.DB.QueryContext(ctx, sqlquery, id)
	objentity := entity.Karyawan{}

	if err != nil {
		return objentity, err
	}
	defer baris.Close()
	if baris.Next() {
		var id int
		var gaji sql.NullInt32
		var nama, jabatan string
		var pasangan sql.NullString
		var married []byte
		var isMarried bool = false
		var tanggallahir sql.NullTime

		err := baris.Scan(&id, &nama, &married, &pasangan, &jabatan, &tanggallahir, &gaji)

		if err != nil {
			return objentity, err
		}

		if len(married) > 0 && married[0] == 1 {
			isMarried = true
			objentity.married = isMarried
		}

		if pasangan.Valid {
			objentity.pasangan = pasangan.String
		}

		objentity.jabatan = jabatan
		if tanggallahir.Valid {
			objentity.tanggallahir = tanggallahir.Time
		}

		if gaji.Valid {
			objentity.gaji = gaji.Int32
		}
		return objentity, nil
	} else {
		return objentity, errors.New("Id " + strconv.Itoa(id) + " tidak ada")
	}

}
func (repository *karyawanRepositoryImpl) FindAll(ctx context.Context) ([]entity.Karyawan, error) {
	sqlquery := "select id,nama,married, pasangan, jabatan, tanggallahir, gaji from karyawan"
	baris, err := repository.DB.QueryContext(ctx, sqlquery)
	objentity := entity.Karyawan{}

	if err != nil {
		return nil, err
	}
	listkaryawan := []entity.Karyawan{}
	defer baris.Close()
	for baris.Next() {
		var id int
		var gaji sql.NullInt32
		var nama, jabatan string
		var pasangan sql.NullString
		var married []byte
		var isMarried bool = false
		var tanggallahir sql.NullTime

		err := baris.Scan(&id, &nama, &married, &pasangan, &jabatan, &tanggallahir, &gaji)

		if err != nil {
			return objentity, err
		}

		if len(married) > 0 && married[0] == 1 {
			isMarried = true
			objentity.married = isMarried
		}

		if pasangan.Valid {
			objentity.pasangan = pasangan.String
		}

		objentity.jabatan = jabatan
		if tanggallahir.Valid {
			objentity.tanggallahir = tanggallahir.Time
		}

		if gaji.Valid {
			objentity.gaji = gaji.Int32
		}
		listkaryawan = append(listkaryawan, objentity)
	}
	return listkaryawan, nil
}
