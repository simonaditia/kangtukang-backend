package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/models"
)

type UpdateWaktuInput struct {
	JadwalPerbaikanAwal  string `json:"jadwal_perbaikan_awal"`
	JadwalPerbaikanAkhir string `json:"jadwal_perbaikan_akhir"`
}

type UpdateOrdersAlasan struct {
	Status           string `json:"status"`
	AlasanTolakBatal string `json:"alasan_tolak_batal"`
}

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

	pesanan := models.Orders{IDCustomer: id_customer, IDTukang: c.Param("id"), DetailPerbaikan: order.DetailPerbaikan, JadwalPerbaikanAwal: order.JadwalPerbaikanAwal, JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir, Status: order.Status, Alamat: order.Alamat, LatitudeCustomer: order.LatitudeCustomer, LongitudeCustomer: order.LongitudeCustomer, CustomerName: order.CustomerName, TukangName: order.TukangName, KategoriTukang: order.KategoriTukang, TotalBiaya: order.TotalBiaya}
	models.DB.Create(&pesanan)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   pesanan,
	})
}

func CancelOrderByCustomer(c *gin.Context) {
	id_order := c.Param("id")
	// Check if the order exists
	var order models.Orders
	result := models.DB.First(&order, id_order)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update the order status
	var input UpdateOrdersAlasan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Model(&order).Updates(input)

	// Update the order status
	// result = models.DB.Model(&order).Update("status", "Dibatalkan")
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
	// 	return
	// }

	// Save the changes to the database
	// result = models.DB.Save(&order)
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order changes"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Order status updated successfully",
	})
}

func AcceptOrderByTukang(c *gin.Context) {
	id_order := c.Param("id")
	// Check if the order exists
	var order models.Orders
	result := models.DB.First(&order, id_order)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update the order status
	result = models.DB.Model(&order).Update("status", "Sedang Berlangsung")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	// Save the changes to the database
	result = models.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order changes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Order status updated successfully",
	})
}

func RejectOrderByTukang(c *gin.Context) {
	id_order := c.Param("id")
	// Check if the order exists
	var order models.Orders
	result := models.DB.First(&order, id_order)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update the order status
	var input UpdateOrdersAlasan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Model(&order).Updates(input)
	// result = models.DB.Model(&order).Update("status", "Ditolak")
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
	// 	return
	// }

	// // Save the changes to the database
	// result = models.DB.Save(&order)
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order changes"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Order status updated successfully",
	})
}

func DoneOrderByTukang(c *gin.Context) {
	id_order := c.Param("id")
	// Check if the order exists
	var order models.Orders
	result := models.DB.First(&order, id_order)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update the order status
	result = models.DB.Model(&order).Update("status", "Selesai")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	// Save the changes to the database
	result = models.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order changes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Order status updated successfully",
	})
}

func ReadOrderByTukang(c *gin.Context) {
	var orders []models.Orders
	var ordersResponse []models.OrderResponse
	var id_tukang string = c.Query("id_tukang")
	var STATUS = "Menunggu Konfirmasi"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_tukang = users_2.id").Where("orders.id_tukang = ? AND status = ?", id_tukang, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			TukangID:             order.IDTukang,
			CustomerID:           order.IDCustomer,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			TotalBiaya:           order.TotalBiaya,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func StatusOrderCustomerMenunggu(c *gin.Context) {
	var orders []models.Orders
	// var order models.Orders
	// var user models.User
	var ordersResponse []models.OrderResponse
	var id_customer string = c.Query("id_customer")
	var STATUS = "Menunggu Konfirmasi"

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang, users.no_telp AS no_telp_tukang, users_2.no_telp AS no_telp_customer, users.latitude AS latitude_tukang, users.longitude AS longitude_tukang").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_customer = ? AND status = ?", id_customer, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			CustomerID:           order.IDCustomer,
			TukangID:             order.IDTukang,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			LatitudeTukang:       order.LatitudeTukang,
			LongitudeTukang:      order.LongitudeTukang,
			TotalBiaya:           order.TotalBiaya,
			NoTelpTukang:         order.NoTelpTukang,
			NoTelpCustomer:       order.NoTelpCustomer,
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

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang, users.no_telp AS no_telp_tukang, users_2.no_telp AS no_telp_customer, users.latitude AS latitude_tukang, users.longitude AS longitude_tukang").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_customer = ? AND status = ?", id_customer, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			CustomerID:           order.IDCustomer,
			TukangID:             order.IDTukang,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			LatitudeTukang:       order.LatitudeTukang,
			LongitudeTukang:      order.LongitudeTukang,
			TotalBiaya:           order.TotalBiaya,
			NoTelpTukang:         order.NoTelpTukang,
			NoTelpCustomer:       order.NoTelpCustomer,
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

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang, users.no_telp AS no_telp_tukang, users_2.no_telp AS no_telp_customer, users.latitude AS latitude_tukang, users.longitude AS longitude_tukang").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_customer = ? AND (status = ? OR status = ? OR status = ?)", id_customer, STATUS_SELESAI, STATUS_BATAL, STATUS_TOLAK).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			CustomerID:           order.IDCustomer,
			TukangID:             order.IDTukang,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			LatitudeTukang:       order.LatitudeTukang,
			LongitudeTukang:      order.LongitudeTukang,
			TotalBiaya:           order.TotalBiaya,
			NoTelpTukang:         order.NoTelpTukang,
			NoTelpCustomer:       order.NoTelpCustomer,
			AlasanTolakBatal:     order.AlasanTolakBatal,
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

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang, users.no_telp AS no_telp_tukang, users_2.no_telp AS no_telp_customer").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_tukang = ? AND status = ?", id_tukang, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			CustomerID:           order.IDCustomer,
			TukangID:             order.IDTukang,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			TotalBiaya:           order.TotalBiaya,
			NoTelpTukang:         order.NoTelpTukang,
			NoTelpCustomer:       order.NoTelpCustomer,
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

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang, users.no_telp AS no_telp_tukang, users_2.no_telp AS no_telp_customer").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_tukang = ? AND status = ?", id_tukang, STATUS).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			CustomerID:           order.IDCustomer,
			TukangID:             order.IDTukang,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			TotalBiaya:           order.TotalBiaya,
			NoTelpTukang:         order.NoTelpTukang,
			NoTelpCustomer:       order.NoTelpCustomer,
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

	if err := models.DB.Table("orders").Select("orders.*, users.nama AS tukang_name, users_2.nama AS customer_name, users.kategori AS kategori_tukang, users.no_telp AS no_telp_tukang, users_2.no_telp AS no_telp_customer").Joins("JOIN users ON orders.id_tukang = users.id").Joins("JOIN users AS users_2 ON orders.id_customer = users_2.id").Where("orders.id_tukang = ? AND (status = ? OR status = ? OR status = ?)", id_tukang, STATUS_SELESAI, STATUS_BATAL, STATUS_TOLAK).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:                   order.ID,
			CustomerID:           order.IDCustomer,
			TukangID:             order.IDTukang,
			Status:               order.Status,
			DetailPerbaikan:      order.DetailPerbaikan,
			JadwalPerbaikanAwal:  order.JadwalPerbaikanAwal,
			JadwalPerbaikanAkhir: order.JadwalPerbaikanAkhir,
			Alamat:               order.Alamat,
			CustomerName:         order.CustomerName,
			TukangName:           order.TukangName,
			KategoriTukang:       order.KategoriTukang,
			LatitudeCustomer:     order.LatitudeCustomer,
			LongitudeCustomer:    order.LongitudeCustomer,
			TotalBiaya:           order.TotalBiaya,
			NoTelpTukang:         order.NoTelpTukang,
			NoTelpCustomer:       order.NoTelpCustomer,
			AlasanTolakBatal:     order.AlasanTolakBatal,
		}
		orderResponse.CustomerName = order.CustomerName
		ordersResponse = append(ordersResponse, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   ordersResponse,
		"status": "success",
	})
}

func UbahWaktu(c *gin.Context) {
	var order models.Orders
	if err := models.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// Validate input
	var input UpdateWaktuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Model(&order).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   order,
	})
}
