package getResource

import (
	"errors"
	"strconv"
)

var baseNumMap = map[string]string{
	"1": "一",
	"2": "二",
	"3": "三",
	"4": "四",
	"5": "五",
	"6": "六",
	"7": "七",
	"8": "八",
	"9": "九",
	"0": "零",
}

var baseDanWei = map[int]string{
	1: "",
	2: "十",
	3: "百",
	4: "千",
}

func GetChinaNum(num int) (string, error) {

	index := 1

	var resultString string

	var resultString2 string

	var danWei string

	if (num > 999999999 || num < 0) {
		return "数字不合法", errors.New("数字不合法")
	}

	numOldString := strconv.Itoa(num)

	numOldStringSince := []byte(numOldString)

	if len(numOldString) > 8 {
		resultString = baseNumMap[string(numOldStringSince[0])] + "亿"
		numOldStringSince = numOldStringSince[1:]
	}

	for i := len(numOldStringSince) - 1; i >= 0; i-- {

		if len(numOldStringSince)-i == 5 {
			danWei = "万"
		} else {
			danWei = ""
		}

		if len(numOldStringSince)-i == 5 && string(numOldStringSince[i]) == "0"{

			resultString2 = danWei + baseNumMap[string(numOldStringSince[i])] + baseDanWei[index]  + resultString2

		} else if len(numOldStringSince)-i > 1 && string(numOldStringSince[i]) == "0" &&  string(numOldStringSince[i+1]) != "0" {

			resultString2 = baseNumMap[string(numOldStringSince[i])] + resultString2

		} else if len(numOldStringSince)-i > 1 && string(numOldStringSince[i]) == "0" && string(numOldStringSince[i+1]) == "0" {

		} else if len(numOldStringSince)-i == 1 && string(numOldStringSince[i]) == "0"{

		}else {

			resultString2 = baseNumMap[string(numOldStringSince[i])] + baseDanWei[index] + danWei + resultString2

		}

		if len(numOldStringSince)-i == 4 {
			index = 1
		} else {
			index ++
		}

	}

	return resultString + resultString2, nil

}
