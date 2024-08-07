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

func (repository *karyawanRepositoryImpl) Insert(ctx context.Context, karyawan entity.KaryawanTable) (entity.KaryawanTable, error) {
	sqlcmd := "insert into Karyawan(nama, married, pasangan, jabatan, tanggallahir, gaji) VALUES(?,?,?,?,?,?)"
	hasil, err := repository.DB.ExecContext(ctx, sqlcmd, karyawan.Nama, karyawan.Married, karyawan.Pasangan, karyawan.Jabatan, karyawan.Tanggallahir, karyawan.Gaji)
	if err != nil {
		return karyawan, err
	}
	Id, err := hasil.LastInsertId()
	if err != nil {
		return karyawan, err
	}
	karyawan.Id = int64(Id)
	return karyawan, nil
}

func (repository *karyawanRepositoryImpl) FindById(ctx context.Context, id int64) (entity.KaryawanTable, error) {
	sqlquery := "select id,nama,married, pasangan, jabatan, tanggallahir, gaji where id=? from karyawan limit 1"
	baris, err := repository.DB.QueryContext(ctx, sqlquery, id)
	objentity := entity.KaryawanTable{}

	if err != nil {
		return objentity, err
	}
	defer baris.Close()
	if baris.Next() {

		err := baris.Scan(&objentity.Id, &objentity.Nama, &objentity.Married, &objentity.Pasangan, &objentity.Jabatan, &objentity.Tanggallahir, &objentity.Gaji)

		if err != nil {
			return objentity, err
		}

		return objentity, nil
	} else {
		return objentity, errors.New("Id " + strconv.Itoa(int(id)) + " tidak ada")
	}

}
func (repository *karyawanRepositoryImpl) FindAll(ctx context.Context) ([]entity.KaryawanTable, error) {
	sqlquery := "select id,nama,married, pasangan, jabatan, tanggallahir, gaji from karyawan"
	baris, err := repository.DB.QueryContext(ctx, sqlquery)
	objentity := entity.KaryawanTable{}

	if err != nil {
		return nil, err
	}
	listkaryawan := []entity.KaryawanTable{}
	defer baris.Close()
	for baris.Next() {

		baris.Scan(&objentity.Id, &objentity.Nama, &objentity.Married, &objentity.Pasangan, &objentity.Jabatan, &objentity.Tanggallahir, &objentity.Gaji)

		listkaryawan = append(listkaryawan, objentity)
	}
	return listkaryawan, nil
}
