package kost

type KostFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	LiverCount       int    `json:"liver_count"`
	SpaceCount       int    `json:"space_count"`
}

//Single Pembuatan Kost
func FormatKost(kost Kost) KostFormatter {
	kostFormatter := KostFormatter{}
	kostFormatter.ID = kost.ID
	kostFormatter.UserID = kost.UserID
	kostFormatter.Name = kost.Name
	kostFormatter.ShortDescription = kost.ShortDescription
	kostFormatter.LiverCount = kost.LiverCount
	kostFormatter.SpaceCount = kost.SpaceCount
	kostFormatter.ImageURL = ""

	//Mengecek Image Url ada atau tidak
	if len(kost.KostImages) > 0 {
		kostFormatter.ImageURL = kost.KostImages[0].FileName
	}

	return kostFormatter

}

//Multi Pembuatan Kost
func FormatKosts(kosts []Kost) []KostFormatter {
	// if len(kosts) == 0 {
	// 	return []KostFormatter{}
	// }
	// var kostsFormatter []KostFormatter

	kostsFormatter := []KostFormatter{}

	//Setiap 1 perulangan kita dapet single object kost,
	//Setelah didapat kita ubah menjadi struct kost formatter menggunakan fungsi format kost
	//Jika sudah didapatkan kita masukan kedalam slice kost , dengan memakai kosts formatter (append)
	for _, kost := range kosts {
		kostFormatter := FormatKost(kost)
		kostsFormatter = append(kostsFormatter, kostFormatter)
	}

	return kostsFormatter
}
