package venueModel

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"unsafe"
)

type TableTennisVenueState struct {
	ID   int    `gorm:"primary_key;auto_increment:false" json:"id"`
	Date string `gorm:"primary_key;type:char(8)" json:"date"`
	T9   bool   `json:"T9"`
	T10  bool   `json:"T10"`
	T11  bool   `json:"T11"`
	T12  bool   `json:"T12"`
	T13  bool   `json:"T13"`
	T14  bool   `json:"T14"`
	T15  bool   `json:"T15"`
	T16  bool   `json:"T16"`
	T17  bool   `json:"T17"`
	T18  bool   `json:"T18"`
	T19  bool   `json:"T19"`
	T20  bool   `json:"T20"`
	T21  bool   `json:"T21"`
}

func (T *TableTennisVenueState) TableTennisStructToSlice(structure []TableTennisVenueState) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(TableTennisVenueState{})
	for i := range structure {
		sliceMock := &model.SliceMock{
			Addr: uintptr(unsafe.Pointer(&structure[i])),
			Cap:  int(Len),
			Len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func (T *TableTennisVenueState) UpdateTableTennisVenueStateEveryday(venueNum int) {
	db := database.ConnectToGormDB()
	//1、将所有date=today的场地全变成空闲
	for i := 0; i < venueNum; i++ {
		vs := TableTennisVenueState{ID: i + 1, Date: "today"}
		db.Save(&vs)
	}
	//2、将所有date=today的场地变成date=houtian
	db.Model(&TableTennisVenueState{}).Where("date = ?", "today").Update("date", "houtian")
	//3、将所有date=tomorrow的场地变成date=today
	db.Model(&TableTennisVenueState{}).Where("date = ?", "tomorrow").Update("date", "today")
	//4、将所有date=houtian的场地变成date=tomorrow
	db.Model(&TableTennisVenueState{}).Where("date = ?", "houtian").Update("date", "tomorrow")
	err := db.Close()
	if err != nil {
		fmt.Println("断开数据库连接出错：", err)
	}
}

func (T *TableTennisVenueState) UpdateVenueStateEveryHour(timeField string) {
	db := database.ConnectToGormDB()
	db.Model(&TableTennisVenueState{}).Where("date = ?", "today").Update(timeField, true)
	err := db.Close()
	if err != nil {
		fmt.Println("断开数据库连接出错：", err)
	}
}
