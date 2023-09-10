package objects

import (
	"chain_counter/util"
	"strconv"
	"strings"
)

type HitObject struct {
	Time       int64
	ObjectType []bool
}

const (
	HitCircle = 0
	Slider    = 1
	NewCombo  = 2
	Spinner   = 3
	Color1    = 4
	Color2    = 5
	Color3    = 6
	LongNote  = 7
)

func GetObjectType(arg string) []bool {
	var x8 uint8
	x64, err := strconv.ParseInt(arg, 0, 8)
	util.CheckError(err)

	x8 = uint8(x64)
	objectType := make([]bool, 8)

	for i := 0; i < 8; i++ {
		if uint8(x8/128) > 0 {
			objectType[7-i] = true
		}
		x8 = x8 << 1
	}

	return objectType
}

func NewHitObject(hitObjectTxt string) HitObject {
	args := strings.Split(hitObjectTxt, ",")

	time, err := strconv.ParseInt(args[2], 0, 64)
	util.CheckError(err)

	objectType := GetObjectType(args[3])

	return HitObject{Time: time, ObjectType: objectType}
}
