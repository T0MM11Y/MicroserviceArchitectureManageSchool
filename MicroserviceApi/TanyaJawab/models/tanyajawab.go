package models

type TanyaJawab struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Pertanyaan   string `json:"pertanyaan"`
	Jawaban      string `json:"jawaban"`
	TanggalTanya string `json:"tanggal_tanya"`
	TanggalJawab string `json:"tanggal_jawab"`
	Admin        string `json:"admin"`
	User         string `json:"user"`
}
