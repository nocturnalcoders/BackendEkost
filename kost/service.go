package kost

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetKosts(userID int) ([]Kost, error)
	GetKostByID(input GetKostDetailInput) (Kost, error)
	CreateKost(input CreateKostInput) (Kost, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetKosts(userID int) ([]Kost, error) {
	//mengapa tidak memakai json karena parameter yang dikirimkan oleh user akan lsg mendapatkan integer
	//nanti di cek apakah userID ada atau tidak
	//hanya mengambil userID yang bersangkutan
	//jika kosong ,kita akan menampilkan data kosts
	if userID != 0 {
		kosts, err := s.repository.FindByUserID(userID)
		if err != nil {
			return kosts, err
		}

		return kosts, nil
	}

	kosts, err := s.repository.FindAll()
	if err != nil {
		return kosts, err
	}

	return kosts, nil

}

func (s *service) GetKostByID(input GetKostDetailInput) (Kost, error) {
	kost, err := s.repository.FindByID(input.ID)
	if err != nil {
		return kost, err
	}

	return kost, nil
}

func (s *service) CreateKost(input CreateKostInput) (Kost, error) {
	//Melakukan mapping dari inputan user ke createkostinput
	//Kemudian dari CreateCampaignInput menjadi object Kost
	kost := Kost{}
	kost.Name = input.Name
	kost.ShortDescription = input.ShortDescription
	kost.Description = input.Description
	kost.Perks = input.Perks
	kost.LiverCount = input.LiverCount
	kost.UserID = input.User.ID
	//Nama kost + id user ke brp -> nama-kost-10
	//Membuat variabel untuk membuat gabungan antara nama kost dan id

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	kost.Slug = slug.Make(slugCandidate)

	//Proses pembuatan Slug
	//Memanggil Repository

	newKost, err := s.repository.Save(kost)
	if err != nil {
		return newKost, err
	}

	return newKost, nil
}
