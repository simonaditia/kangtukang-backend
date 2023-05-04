package controllers

import (
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
	// var user models.User
	var ordersResponse []models.OrderResponse
	var id_customer string = c.Query("id_customer")
	var STATUS = "Menunggu Konfirmasi"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_customer = ? AND status = ?", id_customer, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:              order.ID,
			CustomerID:      order.IDCustomer,
			TukangID:        order.IDTukang,
			Status:          order.Status,
			DetailPerbaikan: order.DetailPerbaikan,
			WaktuPerbaikan:  order.WaktuPerbaikan,
			Alamat:          order.Alamat,
			CustomerName:    order.CustomerName,
			TukangName:      order.TukangName,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func StatusOrderCustomerBerlangsung(c *gin.Context) {
	var orders []models.Orders
	var ordersResponse []models.OrderResponse
	var id_customer string = c.Query("id_customer")
	var STATUS = "Sedang Berlangsung"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_customer = ? AND status = ?", id_customer, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:              order.ID,
			CustomerID:      order.IDCustomer,
			TukangID:        order.IDTukang,
			Status:          order.Status,
			DetailPerbaikan: order.DetailPerbaikan,
			WaktuPerbaikan:  order.WaktuPerbaikan,
			Alamat:          order.Alamat,
			CustomerName:    order.CustomerName,
			TukangName:      order.TukangName,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func StatusOrderCustomerSelesai(c *gin.Context) {
	var orders []models.Orders
	var ordersResponse []models.OrderResponse
	var id_customer string = c.Query("id_customer")
	var STATUS_SELESAI = "Selesai"
	var STATUS_BATAL = "Dibatalkan"
	var STATUS_TOLAK = "Ditolak"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_customer = ? AND (status = ? OR status = ? OR status = ?)", id_customer, STATUS_SELESAI, STATUS_BATAL, STATUS_TOLAK).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:              order.ID,
			CustomerID:      order.IDCustomer,
			TukangID:        order.IDTukang,
			Status:          order.Status,
			DetailPerbaikan: order.DetailPerbaikan,
			WaktuPerbaikan:  order.WaktuPerbaikan,
			Alamat:          order.Alamat,
			CustomerName:    order.CustomerName,
			TukangName:      order.TukangName,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func StatusOrderTukangMenunggu(c *gin.Context) {
	var orders []models.Orders
	var ordersResponse []models.OrderResponse
	var id_tukang string = c.Query("id_tukang")
	var STATUS = "Menunggu Konfirmasi"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_tukang = ? AND status = ?", id_tukang, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:              order.ID,
			CustomerID:      order.IDCustomer,
			TukangID:        order.IDTukang,
			Status:          order.Status,
			DetailPerbaikan: order.DetailPerbaikan,
			WaktuPerbaikan:  order.WaktuPerbaikan,
			Alamat:          order.Alamat,
			CustomerName:    order.CustomerName,
			TukangName:      order.TukangName,
		}
		orderResponse.CustomerName = order.CustomerName

		ordersResponse = append(ordersResponse, orderResponse)

		// var user1 models.User
		// models.DB.Where("id = ?", order.IDTukang).Find(&user1)
		// orderResponse.TukangName = user1.Nama
		// ordersResponse = append(ordersResponse, orderResponse)

		// var user models.User
		// models.DB.Where("id = ?", order.IDCustomer).Find(&user)
		// orderResponse.CustomerName = user.Nama
		// ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func StatusOrderTukangBerlangsung(c *gin.Context) {
	var orders []models.Orders
	var ordersResponse []models.OrderResponse
	var id_tukang string = c.Query("id_tukang")
	var STATUS = "Sedang Berlangsung"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_tukang = ? AND status = ?", id_tukang, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:              order.ID,
			CustomerID:      order.IDCustomer,
			TukangID:        order.IDTukang,
			Status:          order.Status,
			DetailPerbaikan: order.DetailPerbaikan,
			WaktuPerbaikan:  order.WaktuPerbaikan,
			Alamat:          order.Alamat,
			CustomerName:    order.CustomerName,
			TukangName:      order.TukangName,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func StatusOrderTukangSelesai(c *gin.Context) {
	var orders []models.Orders
	var ordersResponse []models.OrderResponse
	var id_tukang string = c.Query("id_tukang")
	var STATUS_SELESAI = "Selesai"
	var STATUS_BATAL = "Dibatalkan"
	var STATUS_TOLAK = "Ditolak"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_tukang = ? AND (status = ? OR status = ? OR status = ?)", id_tukang, STATUS_SELESAI, STATUS_BATAL, STATUS_TOLAK).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:              order.ID,
			CustomerID:      order.IDCustomer,
			TukangID:        order.IDTukang,
			Status:          order.Status,
			DetailPerbaikan: order.DetailPerbaikan,
			WaktuPerbaikan:  order.WaktuPerbaikan,
			Alamat:          order.Alamat,
			CustomerName:    order.CustomerName,
			TukangName:      order.TukangName,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}
