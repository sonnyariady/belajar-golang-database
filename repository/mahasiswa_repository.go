package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type MahasiswaRepository interface {
	Insert(ctx context.Context, objentity entity.Mahasiswa) (entity.Mahasiswa, error)
	FindById(ctx context.Context, id int32) (entity.Mahasiswa, error)
	FindAll(ctx context.Context) ([]entity.Mahasiswa, error)
	FindByJurusan(ctx context.Context, jurusan string) ([]entity.Mahasiswa, error)
}
