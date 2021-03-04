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

func (h *kostHandler) GetKost(c *gin.Context) {
	// bentuknya api/v1/kosts/id -> id bisa 1 / 2 / 3 brapapun
	//Handler -> Mapping id yang di URL ke Struct Input utk dimasukan ke Service , Call formatter
	//Service -> Inputan Struct untuk menangkap ID di URL , pakai shouldbindJSOn
	//Repository  -> untuk get kost by id

	var input kost.GetKostDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to Get Detail of Kost", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// kost, err := h.service.GetKostByID(input)
	// if err != nil {
	// 	response := helper.APIResponse("Failed to Get Detail of Kost", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	kostDetail, err := h.service.GetKostByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to Get Detail of Kost", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Kost Detail", http.StatusOK, "success", kost.FormatKostDetail(kostDetail))
	c.JSON(http.StatusOK, response)
}
