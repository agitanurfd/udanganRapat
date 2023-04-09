package module

import (
	"context"
	"fmt"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
	"github.com/agitanurfd/undanganRapat/model"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "undangan_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertUndanganRapat(db *mongo.Database, col string, lokasi string, phone_number string, biodata model.Tamu, prodi model.Universitas) (InsertedID interface{}) {
	var undanganrapat model.UndanganRapat
	undanganrapat.Location = lokasi
	undanganrapat.Phone_number = phone_number
	undanganrapat.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	undanganrapat.Biodata = biodata
	undanganrapat.Prodi = prodi
	return InsertOneDoc(db, col, undanganrapat)
}

func InsertTamu(db *mongo.Database, col string, nama string, phone_number string, jabatan string, jam_kerja string, hari_kerja string) (InsertedID interface{}) {
	var tamu model.Tamu
	tamu.Nama = nama
	tamu.Phone_number = phone_number
	tamu.Jabatan = jabatan
	tamu.Jam_kerja = jam_kerja
	tamu.Hari_kerja = hari_kerja
	return InsertOneDoc(db, col, tamu)
}

func InsertJamRapat(db *mongo.Database, col string, durasi string, jam_rapat string, hari string, tanggal string) (InsertedID interface{}) {
	var jamrapat model.JamRapat
	jamrapat.Durasi = durasi
	jamrapat.Jam_rapat = jam_rapat
	jamrapat.Hari = hari
	jamrapat.Tanggal = tanggal
	return InsertOneDoc(db, col, jamrapat)
}

func InsertUniversitas(db *mongo.Database, col string, jurusan string) (InsertedID interface{}) {
	var universitas model.Universitas
	universitas.Jurusan= jurusan
	return InsertOneDoc(db, col, universitas)
}

func InsertLokasi(db *mongo.Database, col string, nama string, alamat string) (InsertedID interface{}) {
	var lokasi model.Lokasi
	lokasi.Nama = nama
	lokasi.Alamat = alamat
	return InsertOneDoc(db, col, lokasi)
}

func InsertRuangan(db *mongo.Database, col string, no_ruangan string) (InsertedID interface{}) {
	var ruangan model.Ruangan
	ruangan.No_ruangan = no_ruangan
	return InsertOneDoc(db, col, ruangan)
}

func GetUndanganRapatFromNamaTamu(db *mongo.Database, nama string, col string) (datang model.UndanganRapat) {
	tamu := db.Collection(col)
	filter := bson.M{"biodata.nama": nama}
	err := tamu.FindOne(context.TODO(), filter).Decode(&datang)
	if err != nil {
		fmt.Printf("GetUndanganRapatFromNamaTamu: %v\n", err)
	}
	return datang
}

func GetTamuFromJabatan(db *mongo.Database, jabatan string, col string) (tm model.Tamu) {
	tamu := db.Collection(col)
	filter := bson.M{"jabatan": jabatan}
	err := tamu.FindOne(context.TODO(), filter).Decode(&tm)
	if err != nil {
		fmt.Printf("GetTamuFromJabatan: %v\n", err)
	}
	return tm
}

func GetJamRapatFromDurasi(db *mongo.Database, durasi string, col string) (rpt model.JamRapat) {
	jamrapat := db.Collection(col)
	filter := bson.M{"durasi": durasi}
	err := jamrapat.FindOne(context.TODO(), filter).Decode(&rpt)
	if err != nil {
		fmt.Printf("GetJamRapatFromDurasi: %v\n", err)
	}
	return rpt
}

func GetUniversitasFromJurusan(db *mongo.Database, jurusan string, col string) (unv model.Universitas) {
	universitas := db.Collection(col)
	filter := bson.M{"jurusan": jurusan}
	err := universitas.FindOne(context.TODO(), filter).Decode(&unv)
	if err != nil {
		fmt.Printf("GetUniversitasFromJurusan: %v\n", err)
	}
	return unv
}

func GetRuanganFromNoRuangan(db *mongo.Database, no_ruangan string, col string) (rg model.Ruangan) {
	ruangan := db.Collection(col)
	filter := bson.M{"no_ruangan": no_ruangan}
	err := ruangan.FindOne(context.TODO(), filter).Decode(&rg)
	if err != nil {
		fmt.Printf("GetRuanganFromNoRuangan: %v\n", err)
	}
	return rg
}

func GetAllUndanganRapat(db *mongo.Database, col string) (data []model.UndanganRapat) {
	undangan := db.Collection(col)
	filter := bson.M{}
	cursor, err := undangan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}