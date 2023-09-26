package entity

import "time"

type Orders struct {
	Id         int
	Id_User    int
	Order_Date time.Time
}
