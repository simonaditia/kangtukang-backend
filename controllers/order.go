package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/models"
)

func Order(c *gin.Context) {
	var user models.User
	var order models.Orders
	var id_customer string = c.Query("id_customer")
	const ROLE = "tukang"
	// err := models.DB.Table("users").
	// 	Where("nama LIKE ? AND kategori LIKE ? AND role = ?", "%"+nama+"%", "%"+kategori+"%", ROLE).Find(&users).Error
	if err := models.DB.Where("id = ? AND role = ?", c.Param("id"), ROLE).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pesanan := models.Orders{IDCustomer: id_customer, IDTukang: c.Param("id"), DetailPerbaikan: order.DetailPerbaikan, WaktuPerbaikan: order.WaktuPerbaikan, Status: order.Status, Alamat: order.Alamat}
	models.DB.Create(&pesanan)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   pesanan,
	})
}

func StatusOrderCustomerMenunggu(c *gin.Context) {
	var orders []models.Orders
	// var order models.Orders
	var users []models.User
	var id_customer string = c.Query("id_customer")
	var STATUS = "Menunggu Konfirmasi"
	if err := models.DB.Where("id_customer = ? AND status = ?", id_customer, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// for i, order := range orders {
	// 	id_tukangs := make([]string, len(order.IDTukang))
	// 	for j, id_tukang := range order.IDTukang {
	// 		id_tukangs[j] = id_tukang.
	// 	}
	// }

	// var id_tukang string
	var nama_tukang = ""
	for _, order := range orders {
		// fmt.Println(order.IDTukang)
		// id_tukang = append(id_tukang, order.IDTukang)
		if err := models.DB.Where("id = ?", order.IDTukang).Find(&users).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Record not found!",
			})
			return
		}
		for _, user := range users {
			nama_tukang = user.Nama
			fmt.Println(nama_tukang)
			c.JSON(http.StatusOK, gin.H{
				"status":      "success",
				"data":        orders,
				"nama tukang": nama_tukang,
			})
		}
	}
	// fmt.Println("diluar for: ", id_tukang)

}

func StatusOrderCustomerBerlangsung(c *gin.Context) {
	var orders []models.Orders
	var id_customer string = c.Query("id_customer")
	var STATUS = "Sedang Berlangsung"
	if err := models.DB.Where("id_customer = ? AND status = ?", id_customer, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}

func StatusOrderCustomerSelesai(c *gin.Context) {
	var orders []models.Orders
	var id_customer string = c.Query("id_customer")
	var STATUS_SELESAI = "Selesai"
	var STATUS_BATAL = "Dibatalkan"
	if err := models.DB.Where("id_customer = ? AND (status = ? OR status = ?)", id_customer, STATUS_SELESAI, STATUS_BATAL).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}
