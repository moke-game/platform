package common

import (
	"testing"
)

func Test_DecryptCBC(t *testing.T) {
	str := "E468f0U/vS/+LX8Cuk0QOMImPbmoM0pUCRNKjMaRU1BNfEQY3BXioRV5BO29sXlkdFjQ9iK47Zw4zCMhFcBhdnkzmWLLVJH9R/eyW7iCIUrftSZyNglgjk4jOyKZGDs6Zk48F5gcqPSI4RgUmJRE/aAeBr6kTgI4gknmyf4D7mNl5fyqqeQQmL7ofkkQwcCYfOd+jaX74KH/uFdGb3c5GSW7zAuFSDTwzqZiopsuT+Z0RCz59ikjCMKFzQ0PPVe1b5XKSkbRsF2bWw2decdDpc2Knm92TP20o0TYqDXS3pBGJVe7Dxx1aywoK1IgKMW1gFmi6D81myaAgksf19xZbFqNWfq8xLYHLGmY/sbEv7c1RsI7/bC0XiwMHcBDzScg9l3pDrIDFFKfdUiX6bfGqA=="
	data, err := CBCDecrypt([]byte("CTeGahnbQWfAr5hW"), str)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func Test_Md5(t *testing.T) {
	data := "E468f0U/vS/+LX8Cuk0QOMImPbmoM0pUCRNKjMaRU1BNfEQY3BXioRV5BO29sXlkdFjQ9iK47Zw4zCMhFcBhdnkzmWLLVJH9R/eyW7iCIUrftSZyNglgjk4jOyKZGDs6Zk48F5gcqPSI4RgUmJRE/aAeBr6kTgI4gknmyf4D7mNl5fyqqeQQmL7ofkkQwcCYfOd+jaX74KH/uFdGb3c5GSW7zAuFSDTwzqZiopsuT+Z0RCz59ikjCMKFzQ0PPVe1b5XKSkbRsF2bWw2decdDpc2Knm92TP20o0TYqDXS3pBGJVe7Dxx1aywoK1IgKMW1gFmi6D81myaAgksf19xZbFqNWfq8xLYHLGmY/sbEv7c1RsI7/bC0XiwMHcBDzScg9l3pDrIDFFKfdUiX6bfGqA\\u003d\\u003d"
	key := "CTeGahnbQWfAr5hW"
	res, err := EncryptMD5(data, key)
	if err != nil {
		panic(err)
	}
	t.Log(res)
}
