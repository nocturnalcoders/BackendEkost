package kost

//tidak memakai json , karena dikirim lewat body biasanya POST
//Ada beberapa mengirim JSON
//Dengan bentuk body, Query Param, Menyatu dengan URL -> Memakai URI
type GetKostDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
