package message

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"testing"
)

type A struct {
	Name string
	Age  int
}

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
	spew.Dump(phm)
	//any
	anyy := Person{
		Id:   6,
		Name: "xxxxxx",
		Phones: []*Phone{{
			Type:   PhoneType_HOME,
			Number: "12345",
		}},
	}
	marshalany, _ := ptypes.MarshalAny(&anyy)
	ph.Data = []*any.Any{marshalany}
	t.Log("with any:", ph.String())
	phm, _ = proto.Marshal(&ph)
	t.Log("xx:", phm)
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
			Number: "看看中文行不，,,，，。。にほんご",
		}},
	}
	personp, _ := proto.Marshal(&person)
	t.Log("xx:", personp)
	t.Log("xx:", person.String())
	spew.Dump(personp)
}
