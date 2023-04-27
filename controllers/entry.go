package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/helper"
	"github.com/simonaditia/kangtukang-backend/models"
)

func AddEntry(context *gin.Context) {
	var input models.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input.UserID = user.ID
	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": savedEntry,
	})
}

func GetALlEntries(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": user.Entries,
	})
}
