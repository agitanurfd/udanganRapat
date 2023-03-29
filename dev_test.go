package NPM

import (
	"fmt"
	model "github.com/agitanurfd/undanganRapat/model"
	"github.com/agitanurfd/undanganRapat/Module"
	"testing"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)


func TestInsertUndanganRapat(t *testing.T) {
	lokasi := "ULBI"
	phone_number := "681234567891"
	biodata := model.Tamu{
		Nama : "Jaemin",
		Phone_number : "68987544914",
	}
	prodi := model.Universitas{
		Jurusan : "Teknik Informatika",
	}
	hasil:=module.InsertUndanganRapat(module.MongoConn, "undanganrapat", lokasi , phone_number, biodata, prodi)
	fmt.Println(hasil)
}

func TestInsertTamu(t *testing.T) {
	nama := "Haechan"
	phone_number := "680987654321"
	jabatan := "Dosen"
	jam_kerja := "8 pagi"
	hari_kerja := "Weekend"
	hasil:=module.InsertTamu(module.MongoConn, "tamu", nama, phone_number, jabatan, jam_kerja, hari_kerja)
	fmt.Println(hasil)
}

func TestInsertJamRapat(t *testing.T) {
	durasi := "1 jam"
	jam_rapat := "10 pagi"
	hari := "Jumat"
	tanggal := "10/03/2023"
	hasil:=module.InsertJamRapat(module.MongoConn, "jamrapat", durasi, jam_rapat, hari, tanggal)
	fmt.Println(hasil)
}

func TestInsertUniversitas(t *testing.T) {
	jurusan := "Teknik Informatika"
	hasil:=module.InsertUniversitas(module.MongoConn, "universitas", jurusan)
	fmt.Println(hasil)
}

func TestInsertLokasi(t *testing.T) {
	nama := "ULBI"
	alamat := "Bandung"
	hasil:=module.InsertLokasi(module.MongoConn, "lokasi", nama, alamat)
	fmt.Println(hasil)
}

func TestInsertRuangan(t *testing.T) {
	ruangan := "315"
	hasil:=module.InsertRuangan(module.MongoConn, "ruangan",ruangan)
	fmt.Println(hasil)
}

func TestGetUndanganRapatFromNamaTamu(t *testing.T) {
	nama := "Jaemin"
	biodata:=module.GetUndanganRapatFromNamaTamu(nama, module.MongoConn, "undanganrapat")
	fmt.Println(biodata)
}

func TestGetTamuFromJabatan(t *testing.T) {
	jabatan := "dosen"
	biodata :=module.GetUndanganRapatFromNamaTamu(nama, module.MongoConn, "tamu")
	fmt.Println(biodata)
}

func TestGetJamRapatFromDurasi(t *testing.T) {
	durasi  := "1 Jam"
	jamrapat:=module.JamRapatFromDurasi(durasi, module.MongoConn, "jamrapat")
	fmt.Println(jamrapat)
}

func TestGetUniversitasFromJurusan(t *testing.T) {
	jurusan := "Teknik Informatika"
	prodi   :=module.GetUniversitasFromJurusan(prodi, module.MongoConn, "jurusan")
	fmt.Println(prodi)
}

func TestGetRuanganFromNoRuangan(t *testing.T) {
	no_ruangan := "315"
	location:=module.GetRuanganFromNoRuangan(location, module.MongoConn, "no_ruangan")
	fmt.Println(biodata)
}