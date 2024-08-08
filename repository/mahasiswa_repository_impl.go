package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type mahasiswaRepositoryImpl struct {
	DB *sql.DB
}

func NewMahasiswaRepository(db *sql.DB) MahasiswaRepository {
	return &mahasiswaRepositoryImpl{DB: db}
}

// FindAll implements MahasiswaRepository.
func (m *mahasiswaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Mahasiswa, error) {
	sqlquery := "select id,nama, jurusan, tgllahir from Mahasiswa"
	baris, err := m.DB.QueryContext(ctx, sqlquery)

	if err != nil {
		return nil, err
	}
	lstmhs := []entity.Mahasiswa{}
	for baris.Next() {
		var nama, jurusan string
		var id int
		var tgllahir time.Time

		baris.Scan(&id, &nama, &jurusan, &tgllahir)

		lstmhs = append(lstmhs, entity.Mahasiswa{
			Id:       id,
			Nama:     nama,
			Jurusan:  jurusan,
			TglLahir: tgllahir,
		})
	}
	defer baris.Close()
	return lstmhs, nil
}

// FindById implements MahasiswaRepository.
func (m *mahasiswaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Mahasiswa, error) {
	sqlquery := "select top 1 id,nama, jurusan, tgllahir from Mahasiswa where id=@id"
	baris, err := m.DB.QueryContext(ctx, sqlquery, sql.Named("id", id))
	objentity := entity.Mahasiswa{}
	if err != nil {
		return objentity, err
	}
	defer baris.Close()
	if baris.Next() {
		var nama, jurusan string
		var id int
		var tgllahir time.Time

		baris.Scan(&id, &nama, &jurusan, &tgllahir)

		objentity = entity.Mahasiswa{
			Id:       id,
			Nama:     nama,
			Jurusan:  jurusan,
			TglLahir: tgllahir,
		}
		return objentity, nil
	} else {
		return objentity, errors.New("Id " + strconv.Itoa(int(id)) + " tidak ada")
	}

}

// FindByJurusan implements MahasiswaRepository.
func (m *mahasiswaRepositoryImpl) FindByJurusan(ctx context.Context, jurusan string) ([]entity.Mahasiswa, error) {
	sqlquery := "select top 1 id,nama, jurusan, tgllahir from Mahasiswa where jurusan=@jurusan"
	baris, err := m.DB.QueryContext(ctx, sqlquery, sql.Named("jurusan", jurusan))
	if err != nil {
		return nil, err
	}

	lstmhs := []entity.Mahasiswa{}
	for baris.Next() {
		var nama, jurusan string
		var id int
		var tgllahir time.Time

		baris.Scan(&id, &nama, &jurusan, &tgllahir)

		lstmhs = append(lstmhs, entity.Mahasiswa{
			Id:       id,
			Nama:     nama,
			Jurusan:  jurusan,
			TglLahir: tgllahir,
		})
	}
	defer baris.Close()
	return lstmhs, nil
}

// Insert implements MahasiswaRepository.
func (m *mahasiswaRepositoryImpl) Insert(ctx context.Context, objentity entity.Mahasiswa) (entity.Mahasiswa, error) {
	sqlcmd := "insert into Mahasiswa(nama, jurusan, tgllahir) OUTPUT INSERTED.id VALUES(@nama, @jurusan, @tgllahir)"
	lastid := 0
	_, err := m.DB.QueryContext(ctx, sqlcmd, objentity.Nama, objentity.Jurusan, objentity.TglLahir).Scan(&lastid)
	objentity.Id = lastid
	if err != nil {
		return objentity, err
	}
	objentity.Id = lastid
	return objentity, nil
}
