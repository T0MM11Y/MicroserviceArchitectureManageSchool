package models

type Absensi struct {
	ID        uint    `gorm:"primarykey"`
	UserID    uint    `json:"user_id"`
	Nama      string  `json:"nama"`
	Tanggal   string  `json:"tanggal"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"` // Hadir, Izin, Sakit, dll.

}
