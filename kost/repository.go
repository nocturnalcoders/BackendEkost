package kost

import "gorm.io/gorm"

//Campaign -> Kost

type Repository interface {
	FindAll([]Kost, error)
	FindByUserID(UserID int) ([]Kost, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Kost, error) {
	//Mengambil Semua nilai yang ada di DB
	//Slice of = untuk mengambil banyaknya Data
	var kosts []Kost

	err := r.db.Preload("KostImages", "kost_images.is_primary = 1").Find(&kosts).Error
	if err != nil {
		return kosts, err
	}

	return kosts, nil
}

func (r *repository) FindByUserID(userID int) ([]Kost, error) {
	var kosts []Kost
	//preload akan load sebuah relasi kost_images

	err := r.db.Where("user_id = ?", userID).Preload("KostImages", "kost_images.is_primary = 1").Find(&kosts).Error
	// err := r.db.Where("user_id = ?", userID).Preload("KostImages", "kost_images.is_primary = 1").Find(&kosts).Error
	//"KostImages" -> nama field
	//"kost_images.is_primary" -> nama tabelnya
	//"kost_images.is_primary = 1" -> melakukan filter bahwa kost images saat kita melakukan load kost yang dibuat user id , kita skalian mau ambil datanya images , tapi yang diambil hanya is_primary 1
	if err != nil {
		return kosts, err
	}

	return kosts, nil

}
