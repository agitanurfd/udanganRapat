package module

import (
	"context"
	"fmt"
	"errors"
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

func InsertUndanganRapat(db *mongo.Database, col string, lokasi string, phone_number string, biodata model.Tamu, prodi model.Universitas) (insertedID primitive.ObjectID, err error) {
	undanganrapat := bson.M{
		"location":     lokasi,
		"phone_number": phone_number,
		"datetime":     primitive.NewDateTimeFromTime(time.Now().UTC()),
		"biodata" :     biodata,
		"prodi"   :     prodi,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), undanganrapat)
	if err != nil {
		fmt.Printf("InsertUndanganRapat: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func InsertUndanganRapat(db *mongo.Database, col string, lokasi string, phone_number string, biodata model.Tamu, prodi model.Universitas) (InsertedID interface{}) {
// 	var undanganrapat model.UndanganRapat
// 	undanganrapat.Location = lokasi
// 	undanganrapat.Phone_number = phone_number
// 	undanganrapat.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
// 	undanganrapat.Biodata = biodata
// 	undanganrapat.Prodi = prodi
// 	return InsertOneDoc(db, col, undanganrapat)
// }

func InsertTamu(db *mongo.Database, col string, nama string, phone_number string, jabatan string, jam_kerja string, hari_kerja string) (insertedID primitive.ObjectID, err error) {
	tamu := bson.M{
		"nama"        : nama,
		"phone_number": phone_number,
		"jabatan"     : jabatan,
		"jam_kerja"   : jam_kerja,
		"hari_kerja"  : hari_kerja,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), tamu)
	if err != nil {
		fmt.Printf("InsertTamu: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func InsertTamu(db *mongo.Database, col string, nama string, phone_number string, jabatan string, jam_kerja string, hari_kerja string) (InsertedID interface{}) {
// 	var tamu model.Tamu
// 	tamu.Nama = nama
// 	tamu.Phone_number = phone_number
// 	tamu.Jabatan = jabatan
// 	tamu.Jam_kerja = jam_kerja
// 	tamu.Hari_kerja = hari_kerja
// 	return InsertOneDoc(db, col, tamu)
// }


func InsertJamRapat(db *mongo.Database, col string, durasi string, jam_rapat string, hari string, tanggal string) (insertedID primitive.ObjectID, err error) {
	jamrapat := bson.M{
		"durasi"        : durasi,
		"jam_rapat"     : jam_rapat,
		"hari"          : hari,
		"tanggal"       : tanggal,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), jamrapat)
	if err != nil {
		fmt.Printf("InsertJamRapat: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func InsertJamRapat(db *mongo.Database, col string, durasi string, jam_rapat string, hari string, tanggal string) (InsertedID interface{}) {
// 	var jamrapat model.JamRapat
// 	jamrapat.Durasi = durasi
// 	jamrapat.Jam_rapat = jam_rapat
// 	jamrapat.Hari = hari
// 	jamrapat.Tanggal = tanggal
// 	return InsertOneDoc(db, col, jamrapat)
// }

func InsertUniversitas(db *mongo.Database, col string, jurusan string) (insertedID primitive.ObjectID, err error) {
	universitas := bson.M{
		"jurusan" : jurusan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), universitas)
	if err != nil {
		fmt.Printf("InsertUniversitas: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func InsertUniversitas(db *mongo.Database, col string, jurusan string) (InsertedID interface{}) {
// 	var universitas model.Universitas
// 	universitas.Jurusan= jurusan
// 	return InsertOneDoc(db, col, universitas)
// }

func InsertLokasi(db *mongo.Database, col string, nama string, alamat string) (insertedID primitive.ObjectID, err error) {
	lokasi:= bson.M{
		"nama"   : nama,
		"alamat" : alamat,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), lokasi)
	if err != nil {
		fmt.Printf("InsertLokasi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func InsertLokasi(db *mongo.Database, col string, nama string, alamat string) (InsertedID interface{}) {
// 	var lokasi model.Lokasi
// 	lokasi.Nama = nama
// 	lokasi.Alamat = alamat
// 	return InsertOneDoc(db, col, lokasi)
// }

func InsertRuangan(db *mongo.Database, col string, no_ruangan string) (insertedID primitive.ObjectID, err error) {
	ruangan:= bson.M{
		"no_ruangan": no_ruangan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), ruangan)
	if err != nil {
		fmt.Printf("InsertRuangan: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// func InsertRuangan(db *mongo.Database, col string, no_ruangan string) (InsertedID interface{}) {
// 	var ruangan model.Ruangan
// 	ruangan.No_ruangan = no_ruangan
// 	return InsertOneDoc(db, col, ruangan)
// }

func GetUndanganRapatFromNamaTamu(db *mongo.Database, nama string, col string) (datang model.UndanganRapat) {
	tamu := db.Collection(col)
	filter := bson.M{"biodata.nama": nama}
	err := tamu.FindOne(context.TODO(), filter).Decode(&datang)
	if err != nil {
		fmt.Printf("GetUndanganRapatFromNamaTamu: %v\n", err)
	}
	return datang
}

// get baru
func GetUndanganRapatFromID(_id primitive.ObjectID, db *mongo.Database, col string) (undang model.UndanganRapat, errs error) {
	undanganrapat := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := undanganrapat.FindOne(context.TODO(), filter).Decode(&undang)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return undang, fmt.Errorf("no data found for ID %s", _id)
		}
		return undang, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return undang, nil
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

func GetAllTamu(db *mongo.Database, col string) (data []model.Tamu) {
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

func GetAllJamRapat(db *mongo.Database, col string) (data []model.JamRapat) {
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

func GetAllUniversitas(db *mongo.Database, col string) (data []model.Universitas) {
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

func GetAllRuangan(db *mongo.Database, col string) (data []model.Ruangan) {
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

func GetAllLokasi(db *mongo.Database, col string) (data []model.Lokasi) {
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

// Update Data
func UpdateUndanganRapat(db *mongo.Database, col string, lokasi string, phone_number string, biodata model.Tamu, prodi model.Universitas) (insertedID primitive.ObjectID, err error) {
	undanganrapat := bson.M{
		"location":     lokasi,
		"phone_number": phone_number,
		"datetime":     primitive.NewDateTimeFromTime(time.Now().UTC()),
		"biodata" :     biodata,
		"prodi"   :     prodi,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), undanganrapat)
	if err != nil {
		fmt.Printf("UpdateUndanganRapat: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// Delete Data
func DeleteUndanganRapatByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	undanganrapat := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := undanganrapat.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}