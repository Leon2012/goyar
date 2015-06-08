package goyar

import (
	"fmt"
	"testing"
)

func TestUnpackProtocol(t *testing.T) {

	//id = 118, 166, 117, 140 = 0x76A6758C
	//version = 0, 0
	//magic_num  =  128, 223, 236, 96 = 0x80DFEC60
	//reserved = 0, 0, 0, 0
	//provider = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	//token =  0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	//data := []byte{118, 166, 117, 140, 0, 0, 128, 223, 236, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 60, 74, 83, 79, 78, 0, 117, 110, 97, 123, 34, 105, 34, 58, 49, 57, 57, 48, 54, 50, 48, 53, 53, 54, 44, 34, 109, 34, 58, 34, 115, 111, 109, 101, 95, 109, 101, 116, 104, 111, 100, 34, 44, 34, 112, 34, 58, 91, 34, 112, 97, 114, 97, 109, 101, 116, 101, 114, 34, 93, 125}

	data := []byte{170, 240, 15, 216, 0, 0, 128, 223, 236, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 74, 74, 83, 79, 78, 0, 117, 110, 97, 123, 34, 105, 34, 58, 50, 56, 54, 55, 56, 53, 57, 52, 49, 54, 44, 34, 109, 34, 58, 34, 115, 111, 109, 101, 95, 109, 101, 116, 104, 111, 100, 34, 44, 34, 112, 34, 58, 91, 123, 34, 112, 49, 34, 58, 34, 112, 97, 114, 97, 109, 101, 116, 101, 114, 34, 125, 44, 34, 116, 101, 115, 116, 34, 93, 125}
	fmt.Println("data len : ", len(data))

	// dataLen := len(data)

	// header := &YarHeader{}
	// header.Id = BytesToUint32(data[0:4])
	// header.Version = BytesToUint16(data[4:6])
	// header.MagicNum = BytesToUint32(data[6:10])
	// header.Reserved = BytesToUint32(data[10:14])
	// //header.Provider = data[14:46]
	// copy(header.Provider[:], data[14:46])
	// //header.Token = data[46:78]
	// copy(header.Token[:], data[46:78])
	// header.BodyLen = BytesToUint32(data[78:82])

	// fmt.Println(header)

	// fmt.Println(data[82:90])
	// packagerName := string(data[82:90])
	// fmt.Println(packagerName) //data[82:86] JSON

	// body := string(data[90:dataLen])
	// fmt.Println(body) //{"i":1990620556,"m":"some_method","p":["parameter"]}

	yar, err := Unpack(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(yar.Request)
	}

}

func TestPackProtocol(t *testing.T) {
	data := []byte{170, 240, 15, 216, 0, 0, 128, 223, 236, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 74, 74, 83, 79, 78, 0, 117, 110, 97, 123, 34, 105, 34, 58, 50, 56, 54, 55, 56, 53, 57, 52, 49, 54, 44, 34, 109, 34, 58, 34, 115, 111, 109, 101, 95, 109, 101, 116, 104, 111, 100, 34, 44, 34, 112, 34, 58, 91, 123, 34, 112, 49, 34, 58, 34, 112, 97, 114, 97, 109, 101, 116, 101, 114, 34, 125, 44, 34, 116, 101, 115, 116, 34, 93, 125}
	//fmt.Println("data len : ", len(data))
	yar, err := Unpack(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(yar)
	}

	response := Response{
		Id:     yar.Request.Id,
		Status: 20,
		Out:    "",
		Retval: "",
		Err:    "Test Error",
	}

	yar.Response = response

	out, err := Pack(yar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
