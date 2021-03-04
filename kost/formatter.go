package kost

import "strings"

type KostFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	LiverCount       int    `json:"liver_count"`
	SpaceCount       int    `json:"space_count"`
	Slug             string `json:"slug"`
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
	kostFormatter.Slug = kost.Slug
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

type KostDetailFormatter struct {
	ID               int                  `json:"id"`
	Name             string               `json:"name"`
	ShortDescription string               `json:"short_description"`
	Description      string               `json:"description"`
	ImageURL         string               `json:"images_url"`
	LiverCount       int                  `json:"liver_count"`
	SpaceCount       int                  `json:"space_count"`
	UserID           int                  `json:"user_id"`
	Slug             string               `json:"slug"`
	Perks            []string             `json:"perks"`
	User             KostUserFormatter    `json:"user"`
	Images           []KostImageFormatter `json:"images"`
}

type KostUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"images_url"`
}

type KostImageFormatter struct {
	ImageURL string `json:"images_url"`
	IsPrmary bool   `json:"is_primary"`
}

func FormatKostDetail(kost Kost) KostDetailFormatter {
	kostDetailFormatter := KostDetailFormatter{}
	kostDetailFormatter.ID = kost.ID
	kostDetailFormatter.Name = kost.Name
	kostDetailFormatter.ShortDescription = kost.ShortDescription
	kostDetailFormatter.Description = kost.Description
	kostDetailFormatter.LiverCount = kost.LiverCount
	kostDetailFormatter.SpaceCount = kost.SpaceCount
	kostDetailFormatter.UserID = kost.UserID
	kostDetailFormatter.Slug = kost.Slug
	kostDetailFormatter.ImageURL = ""

	//Mengecek Image Url ada atau tidak
	if len(kost.KostImages) > 0 {
		kostDetailFormatter.ImageURL = kost.KostImages[0].FileName
	}

	//Function Split untuk memecah string menurut apa yang dikehendaki misal , . or apapun
	var perks []string

	for _, perk := range strings.Split(kost.Perks, ",") {
		perks = append(perks, perk)
	}

	kostDetailFormatter.Perks = perks

	user := kost.User

	kostUserFormatter := KostUserFormatter{}
	kostUserFormatter.Name = user.Name
	kostUserFormatter.ImageURL = user.AvatarFileName
	kostDetailFormatter.User = kostUserFormatter

	//Loop 1 per 1 images
	//Melakukan append kost images formatter yang satuan kedalam slice image
	images := []KostImageFormatter{}
	for _, image := range kost.KostImages {
		kostImageFormatter := KostImageFormatter{}
		kostImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		kostImageFormatter.IsPrmary = isPrimary

		images = append(images, kostImageFormatter)
	}

	kostDetailFormatter.Images = images

	return kostDetailFormatter
}
