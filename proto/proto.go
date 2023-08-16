package proto

import "github.com/golang/protobuf/proto"

func Decode(id uint32, buf []byte) (err error, m interface{}) {
	switch id {
	case 0x10000001:
		msg := &LoginReq{}
		err = proto.Unmarshal(buf, msg)
		return err, msg
	case 0x20000001:
		msg := &LoginRsp{}
		err = proto.Unmarshal(buf, msg)
		return err, msg
	}
	return
}
