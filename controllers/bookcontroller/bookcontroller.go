package bookcontroller

import (
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Book's", books)
}
