package message

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestPhone_ProtoMessage(t *testing.T) {
	ph := Phone{
		Type:   PhoneType_HOME,
		Number: "12345678",
	}
	phm, err := proto.Marshal(&ph)
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("xx:", phm)
	var phr Phone
	proto.Unmarshal(phm, &phr)
	t.Logf("%+v:", phr)
	j, _ := json.Marshal(&ph)
	t.Log("json:", j)
	t.Log("s:", ph.String())
}

func TestPerson_ProtoMessage(t *testing.T) {
	person := Person{
		Id:   1,
		Name: "wang",
		Phones: []*Phone{{
			Type:   PhoneType_HOME,
			Number: "12345",
		}, {
			Type:   PhoneType_WORK,
			Number: "67890",
		}},
	}
	personp, _ := proto.Marshal(&person)
	t.Log("xx:", personp)
	t.Log("xx:", person.String())
}
