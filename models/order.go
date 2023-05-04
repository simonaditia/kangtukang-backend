package models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	IDCustomer      string `json:"id_customer"`
	IDTukang        string `json:"id_tukang"`
	DetailPerbaikan string `json:"detail_perbaikan" gorm:"type:text"`
	WaktuPerbaikan  string `json:"waktu_perbaikan"`
	Status          string `json:"status" gorm:"default:'Menunggu Konfirmasi'"`
	Alamat          string `json:"alamat"`
	CustomerName    string `json:"nama_customer"`
	TukangName      string `json:"nama_tukang"`
}

// type StatusOrderCustomer struct {
// 	gorm.Model
// 	NamaTukang     string `json:"nama_tukang"`
// 	TotalBiaya     string `json:"total_biaya"`
// 	Status         string `json:"status"`
// 	WaktuPerbaikan string `json:"waktu_perbaikan"`
// }

type OrderResponse struct {
	ID              uint   `json:"ID"`
	CustomerID      string `json:"id_customer"`
	TukangID        string `json:"id_tukang"`
	DetailPerbaikan string `json:"detail_perbaikan" gorm:"type:text"`
	WaktuPerbaikan  string `json:"waktu_perbaikan"`
	Status          string `json:"status"`
	Alamat          string `json:"alamat"`
	TukangName      string `json:"nama_tukang"`
	CustomerName    string `json:"nama_customer"`
}