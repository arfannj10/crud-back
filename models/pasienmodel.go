package models

import (
	"crud-back/config"
	"crud-back/entities"
	"database/sql"
	"fmt"
	"time"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) FindAll() ([]entities.Pasien, error) { // menampilkan semnua data dari database
	var dataPasien []entities.Pasien

	rows, err := p.conn.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id, &pasien.NamaLengkap, &pasien.NIK, &pasien.JenisKelamin, &pasien.TempatLahir, &pasien.TanggalLahir, &pasien.Alamat, &pasien.NoHp)

		//yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)
		//dd-mm-yyyy
		pasien.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataPasien = append(dataPasien, pasien)
	}

	return dataPasien, nil
}

func (p *PasienModel) Create(pasien entities.Pasien) entities.Pasien {
	stmt, err := p.conn.Exec("insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values (?,?,?,?,?,?,?)", pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)
	if err != nil {
		fmt.Println(err)
	}

	lastInsertId, _ := stmt.LastInsertId()
	pasien.Id = lastInsertId

	return pasien
}
