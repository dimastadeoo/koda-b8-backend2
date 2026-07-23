package lib

import "time"

func TimeToWIB (t time.Time) time.Time{
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil{
		return t
	}

	timeLoc := t.In(loc)
	return  timeLoc
}