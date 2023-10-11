package entities

type Pasien struct {
	Id           int64
	NamaLengkap  string `json:"nama_lengkap" validate:"required" label:"nama lengkap"`
	NIK          string `json:"nik" validate:"required" label:"nik"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required" label:"jenis kelamin"`
	TempatLahir  string `json:"tempat_lahir" validate:"required" label:"tempat lajir"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required" label:"tanggal lahir"`
	Alamat       string `json:"alamat" validate:"required" label:"alamat"`
	NoHp         string `json:"no_hp" validate:"required" label:"no hp"`
}
