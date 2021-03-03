package kost

type Service interface {
	GetKosts(userID int) ([]Kost, error)
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
