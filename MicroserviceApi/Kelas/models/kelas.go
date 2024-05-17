package models

type Kelas struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	NamaKelas string `json:"nama_kelas" gorm:"type:varchar(100);not null"`
}
