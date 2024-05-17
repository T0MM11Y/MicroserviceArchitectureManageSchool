package models

type Roster struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	MataPelajaran string `json:"mata_pelajaran" gorm:"not null"`
	WaktuMulai    string `json:"waktu_mulai" gorm:"not null"`
	WaktuSelesai  string `json:"waktu_selesai" gorm:"not null"`
	Kelas         string `json:"kelas"`
	Hari          string `json:"hari" gorm:"not null"`
	Pengajar      string `json:"pengajar"`
}
