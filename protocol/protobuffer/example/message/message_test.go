package message

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"testing"
	"time"
)

type A struct {
	Name string
	Age  int
}

func TestPhone_ProtoMessage(t *testing.T) {
	ph := Phone{
		Type:     PhoneType_HOME,
		Number:   "12345678",
		KeyValue: "kkkvvvvv",
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
	marshaler := jsonpb.Marshaler{}
	d, _ := marshaler.MarshalToString(&ph)
	t.Log("jsonp:", d)
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

func TestMapTest_ProtoMessage(t *testing.T) {
	m := MapTest{
		Maptest: map[string]string{},
	}
	m.Maptest["1"] = "343"
	m.Maptest["xfdr"] = "dfsfds"
	phm, err := proto.Marshal(&m)
	if err != nil {
		t.Error("err:", err)
		return
	}
	t.Log("xx:", phm)
	spew.Dump(phm)
}

func TestOneofMessage_ProtoMessage(t *testing.T) {
	o := OneofMessage{
		TestOneof: &OneofMessage_Age{
			Age: 100,
		},
	}
	o.TestOneof = &OneofMessage_Name{
		Name: "wang",
	}
	switch x := o.TestOneof.(type) {
	case *OneofMessage_Age:
		// Load profile image based on URL
		// using x.ImageUrl
		t.Log("AGE")
	case *OneofMessage_Name:
		// Load profile image based on bytes
		// using x.ImageData
		t.Log("Name")
	case nil:
		// The field is not set.
	default:
		t.Errorf("Profile.Avatar has unexpected type %T", x)
	}
	pMarshal, _ := proto.Marshal(&o)
	spew.Dump(pMarshal)
	k := OneofMessage{}
	proto.Unmarshal(pMarshal, &k)
	switch x := k.TestOneof.(type) {
	case *OneofMessage_Age:
		// Load profile image based on URL
		// using x.ImageUrl
		t.Log("kAGE ", k.GetAge())
	case *OneofMessage_Name:
		// Load profile image based on bytes
		// using x.ImageData
		t.Log("kName ", k.GetName())
	case nil:
		// The field is not set.
	default:
		t.Errorf("Profile.Avatar has unexpected type %T", x)
	}
}

func TestJsonp(t *testing.T) {
	ph := Phone{
		Type:     PhoneType_HOME,
		Number:   "12345678",
		KeyValue: "kkkvvvvv",
	}
	ph.Data = []*any.Any{}
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
	marshaler := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "",
		AnyResolver:  nil,
	}
	d, _ := marshaler.MarshalToString(&ph)
	t.Log("jsonp:", d)
	//UnMarshal
	phd := Phone{}
	jsonpb.UnmarshalString(d, &phd)
	t.Log("Unmarshal:", phd.String())
}

func TestContactBook_ProtoMessage(t *testing.T) {
	contract := ContactBook{
		Persons:    nil,
		LastUpdate: ptypes.TimestampNow(),
	}
	cm, err := proto.Marshal(&contract)
	if err != nil {
		t.Error("err:", err)
		return
	}
	spew.Dump(cm)
	marshaler := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "",
		AnyResolver:  nil,
	}
	d, _ := marshaler.MarshalToString(&contract)
	t.Log("jsonp:", d)
	t.Log("XXXXXXXXX")
	t.Log(time.Now().UTC().Format(time.RFC3339Nano))
	s := "2021-01-06T09:49:37.111123000Z"
	sb, _ := time.Parse(time.RFC3339Nano, s)
	t.Log(sb.UTC().Format(time.RFC3339Nano))

}
