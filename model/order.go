package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Venue struct {
	Number int      `json:"number"`
	Time   []string `json:"time"`
	Price  int      `json:"price"`
}
type Order struct {
	ID              int64  `json:"id"`
	UserID          string `gorm:"type:char(19)" json:"userID"`
	State           string `gorm:"type:varchar(3)" json:"state"`
	VenueType       string `gorm:"type:varchar(10)" json:"venueType"`
	Count           int    `gorm:"type:tinyint" json:"count"`
	ReservationDate string `gorm:"type:varchar(10)" json:"reservationDate"`
	UseDate         string `gorm:"type:varchar(10)" json:"useDate"`
	FinishTime      int    `gorm:"type:tinyint" json:"finishTime"`
	Venue1          Venue  `gorm:"type:json" json:"venue1"`
	Venue2          Venue  `gorm:"type:json" json:"venue2"`
	CancelFlag      bool   `json:"cancelFlag"`
}

func (v *Venue) Scan(value interface{}) error {
	byteValue, _ := value.([]byte) //类型断言，断定为[]byte类型，我们在value方法中也是转换为[]byte类型输入到数据库中的
	var receiver Venue
	err := json.Unmarshal(byteValue, &receiver) //反序列化，将[]byte类型转化为我们需要的结构体
	if err != nil {
		return err
	}
	//fmt.Println(receiver)
	*v = receiver //将其内容传输给venue
	return nil
}

// Value 存入数据库，将json转换为数据库可接受类型数据，实现dirver.Valuer接口
func (v Venue) Value() (driver.Value, error) {
	return json.Marshal(v) //由结构体转换为json类型数据，返回[]byte

}
