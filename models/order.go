package models

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	IDCustomer           string  `json:"id_customer"`
	IDTukang             string  `json:"id_tukang"`
	DetailPerbaikan      string  `json:"detail_perbaikan" gorm:"type:text"`
	JadwalPerbaikanAwal  string  `json:"jadwal_perbaikan_awal"`
	JadwalPerbaikanAkhir string  `json:"jadwal_perbaikan_akhir"`
	Status               string  `json:"status" gorm:"default:'Menunggu Konfirmasi'"`
	Alamat               string  `json:"alamat"`
	CustomerName         string  `json:"nama_customer"`
	TukangName           string  `json:"nama_tukang"`
	KategoriTukang       string  `json:"kategori_tukang"`
	LatitudeCustomer     float64 `json:"latitude_customer"`
	LongitudeCustomer    float64 `json:"longitude_customer"`
	TotalBiaya           float64 `json:"total_biaya"`
}

// type StatusOrderCustomer struct {
// 	gorm.Model
// 	NamaTukang     string `json:"nama_tukang"`
// 	TotalBiaya     string `json:"total_biaya"`
// 	Status         string `json:"status"`
// 	JadwalPerbaikanAwal string `json:"jadwal_perbaikan_awal"`
// }

type OrderResponse struct {
	ID                   uint    `json:"ID"`
	CustomerID           string  `json:"id_customer"`
	TukangID             string  `json:"id_tukang"`
	DetailPerbaikan      string  `json:"detail_perbaikan" gorm:"type:text"`
	JadwalPerbaikanAwal  string  `json:"jadwal_perbaikan_awal"`
	JadwalPerbaikanAkhir string  `json:"jadwal_perbaikan_akhir"`
	Status               string  `json:"status"`
	Alamat               string  `json:"alamat"`
	TukangName           string  `json:"nama_tukang"`
	CustomerName         string  `json:"nama_customer"`
	KategoriTukang       string  `json:"kategori_tukang"`
	LatitudeCustomer     float64 `json:"latitude_customer"`
	LongitudeCustomer    float64 `json:"longitude_customer"`
	TotalBiaya           float64 `json:"total_biaya"`
}
