package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//note :
//pembuatan object baru dari type repository struct dan nilai db dari &repository mengambil value dari db *gorm.DB

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Note :
// 1. kenapa dilakukannya Interface ?
// Karena nanti untuk object / struct lain atau bagian lain mengacu ke
// repository ga kemana"

// 2. Interface adalah sebuah kontrak
// 3. r kecil pada repository tidak bersifat public tidak dapat di akses oleh package lain / pun dari luar
// 4. R besar sebaliknya dengan r kecil
