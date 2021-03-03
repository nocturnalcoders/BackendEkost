package handler

import (
	"backendEkost/helper"
	"backendEkost/kost"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 1. Tangkap Parameter di handler
// 2. Handler ke service (Diservice yang menentukan apakah repository yang di call)
// 3. Ke Repository : FindAll , FindByUserID
// 4. Ke DB

//GET : api/v1/kosts?user_id=10
//Akan mengambil user id yang ber ID 10
//GET : api/v1/kosts
//Akan mengambil semua data kosts

type kostHandler struct {
	service kost.Service
}

func NewKostHandler(service kost.Service) *kostHandler {
	return &kostHandler{service}
}

//api/v1/kosts
func (h *kostHandler) GetKosts(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	//Kost lebih dari 1 , kumpulan dari object kost
	kosts, err := h.service.GetKosts(userID)
	if err != nil {
		response := helper.APIResponse("Error to Get Kosts", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of Kosts", http.StatusOK, "success", kost.FormatKosts(kosts))
	c.JSON(http.StatusOK, response)

}
