package NPM

import (
	"fmt"
	model "github.com/agitanurfd/undanganRapat/model"
	"github.com/agitanurfd/undanganRapat/module"
	"testing"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)


func TestInsertUndanganRapat(t *testing.T) {
	lokasi := "ULBI"
	phone_number := "681234567891"
	biodata := model.Tamu{
		Nama : "Popo",
		Phone_number : "68093284827492",
	}
	prodi := model.Universitas{
		Jurusan : "Akuntansi",
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
	biodata:=module.GetUndanganRapatFromNamaTamu(module.MongoConn, nama, "undanganrapat")
	fmt.Println(biodata)
}

func TestGetTamuFromJabatan(t *testing.T) {
	jabatan := "Dosen"
	biodata :=module.GetTamuFromJabatan(module.MongoConn, jabatan,  "tamu")
	fmt.Println(biodata)
}

func TestGetJamRapatFromDurasi(t *testing.T) {
	durasi  := "1 jam"
	jamrapat:=module.GetJamRapatFromDurasi(module.MongoConn, durasi, "jamrapat")
	fmt.Println(jamrapat)
}

func TestGetUniversitasFromJurusan(t *testing.T) {
	jurusan := "Teknik Informatika"
	prodi   :=module.GetUniversitasFromJurusan(module.MongoConn, jurusan, "universitas")
	fmt.Println(prodi)
}

func TestGetRuanganFromNoRuangan(t *testing.T) {
	no_ruangan := "315"
	ruangan:=module.GetRuanganFromNoRuangan(module.MongoConn, no_ruangan, "ruangan")
	fmt.Println(ruangan)
}

func GetAllUndanganRapat(t *testing.T) {
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

func TestGetAllUndanganRapat(t *testing.T) {
	data:=module.GetAllUndanganRapat(module.MongoConn, "undanganrapat")
	fmt.Println(data)
}

func TestGetAllTamu(t *testing.T) {
	data:=module.GetAllTamu(module.MongoConn, "tamu")
	fmt.Println(data)
}

func TestGetAllJamRapat(t *testing.T) {
	data:=module.GetAllJamRapat(module.MongoConn, "jamrapat")
	fmt.Println(data)
}

func TestGetAllUniversitas(t *testing.T) {
	data:=module.GetAllUniversitas(module.MongoConn, "universitas")
	fmt.Println(data)
}

func TestGetAllLokasi(t *testing.T) {
	data:=module.GetAllLokasi(module.MongoConn, "lokasi")
	fmt.Println(data)
}

func TestGetAllRuangan(t *testing.T) {
	data:=module.GetAllRuangan(module.MongoConn, "ruangan")
	fmt.Println(data)
}