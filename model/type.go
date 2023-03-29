package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tamu struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja    string           `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja   string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type JamRapat struct {
	Durasi     string      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_rapat  string   `bson:"jam_rapat,omitempty" json:"jam_rapat,omitempty"`
	Hari       string `bson:"hari,omitempty" json:"hari,omitempty"`
	Tanggal    string      `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type UndanganRapat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Location     string             `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Biodata      Tamu               `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Prodi        Universitas        `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

type Universitas struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jurusan    string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Alamat   string             `bson:"alamat,omitempty" json:"alamat,omitempty"`         
}

type Ruangan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	No_ruangan  string             `bson:"no_ruangan,omitempty" json:"no_ruangan,omitempty"`         
}