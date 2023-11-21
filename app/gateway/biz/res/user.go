package res

import "time"

type UserInfo struct {
	Phone      string   
	UserId     int64
	UserName   string
	Sex        int8
	Status     int8
	Avatar     string
	Assets1Num int64
	Assets2Num int64
	TradeNum   int64
	TitleId    int8
	LoginTime  time.Time
}