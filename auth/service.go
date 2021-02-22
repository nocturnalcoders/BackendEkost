package auth

//service generate token
import "github.com/dgrijalva/jwt-go"

//tentukan dahulu servicenya mau ngapain
//1. bagaimana cara membuat token (Generate)
//2. Melakukan validasi token

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

//Secret key ->syarat -> sebuah token dianggap valid jika dia dibuat dengan secret key yang sama dengan yang terdaftar di server
var SECRET_KEY = []byte("EKOST_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	//Variabel dama JWT.IO Payload disebut juga dengan claim
	//Melakukan passing dengan User ID sebagai key, dan valuenya adalah user id yang dipassing melalui func generate token
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//Membuat Token
	//menggunakan HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
