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
}

type StatusOrderCustomer struct {
	gorm.Model
	NamaTukang     string `json:"nama_tukang"`
	TotalBiaya     string `json:"total_biaya"`
	Status         string `json:"status"`
	WaktuPerbaikan string `json:"waktu_perbaikan"`
}
