package goyar

import (
	"bytes"
	"encoding/json"
	"errors"
	_ "fmt"
	"strings"
)

//4+2+4+4+32+32+4

type Header struct {
	Id       uint32
	Version  uint16
	MagicNum uint32
	Reserved uint32
	Provider [32]byte
	Token    [32]byte
	BodyLen  uint32
}

type Packager struct {
	Name string
	Data [8]byte
}

type Request struct {
	Id     uint64        `json:"i"`
	Method string        `json:"m"`
	Params []interface{} `json:"p"`
}

type Response struct {
	Id     uint64      `json:"i"`
	Status int32       `json:"s"`
	Out    string      `json:"o"`
	Retval interface{} `json:"r"`
	Err    string      `json:"e"`
}

type Yar struct {
	Header   *Header
	Packager *Packager
	Request  Request
	Response Response
}

func Pack(yar *Yar) ([]byte, error) {
	// jsonData := []byte{}
	// encoder := json.NewEncoder(bytes.NewBuffer(jsonData))
	// if err := encoder.Encode(yar.Response); err != nil {
	// 	return nil, err
	// }
	jsonData, err := json.Marshal(&yar.Response)
	if err != nil {
		return nil, err
	}
	//fmt.Println("json:", string(jsonData))

	jsonDataLen := len(jsonData)
	dataLen := (82 + 8 + jsonDataLen)
	data := make([]byte, dataLen)

	bodyLen := jsonDataLen + 8
	yar.Header.BodyLen = uint32(bodyLen)

	copy(data[0:4], Uint32ToBytes(yar.Header.Id))
	copy(data[4:6], Uint16ToBytes(yar.Header.Version))
	copy(data[6:10], Uint32ToBytes(yar.Header.MagicNum))
	copy(data[10:14], Uint32ToBytes(yar.Header.Reserved))
	copy(data[14:46], yar.Header.Provider[:32])
	copy(data[46:78], yar.Header.Token[:32])
	copy(data[78:82], Uint32ToBytes(yar.Header.BodyLen))

	copy(data[82:90], yar.Packager.Data[:8])

	copy(data[90:dataLen], jsonData)

	return data, nil
}

func Unpack(data []byte) (*Yar, error) {
	dataLen := len(data)
	if dataLen < 90 {
		return nil, errors.New("Parse post data error")
	}

	header := &Header{}
	header.Id = BytesToUint32(data[0:4])
	header.Version = BytesToUint16(data[4:6])
	header.MagicNum = BytesToUint32(data[6:10])
	header.Reserved = BytesToUint32(data[10:14])
	copy(header.Provider[:], data[14:46])
	copy(header.Token[:], data[46:78])
	header.BodyLen = BytesToUint32(data[78:82])

	if header.MagicNum != 0x80DFEC60 {
		return nil, errors.New("Magic num error")
	}

	bodyLen := (dataLen - 82)
	if int(header.BodyLen) != bodyLen {
		return nil, errors.New("Body length error")
	}

	packager := &Packager{}
	packName := strings.ToUpper(string(data[82:86]))
	if packName != "JSON" {
		return nil, errors.New("Only supper json package")
	}
	packager.Name = packName
	copy(packager.Data[:], data[82:90])

	jsonData := data[90:dataLen]
	var request Request
	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	if err := decoder.Decode(&request); err != nil {
		return nil, err
	}

	return &Yar{Header: header,
		Packager: packager,
		Request:  request,
	}, nil
}
