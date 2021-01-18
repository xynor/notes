[参考](https://blog.csdn.net/carson_ho/article/details/70568606)
##  T-L-V数据存储方式
* 即Tag-Length-Value
* T-L-V表示单个字段，将所有数据拼装成一个字节流
* 其中Length是可选的，Varint不需要Length
* 减少了分割符的使用，紧凑，若字段没有设置值，序列化后是不存在的，不需要编码
* Tag，包括了字段标示号(field_number)和数据类型(wire_type)，占用一个字节(标示号超过16，那么占用多一个字节)，也是使用Varint编码
* Length如果有的话，也是使用Varint编码
## 编码方式
#### Varint&Zigzag
* 变长的编码方式，1-10个字节
* 用字节表示数字，值越小的数字，使用越小的字节表示，从而进行了压缩
* Varint中，每个字节的第一位都有意义，为1表示后续的字节也是该数字一部分，0表示最后一个字节，剩下的7位也是表示该数字
* 因此，小于128的数可以用1个字节表示，大于128的数，用两个字节表示
* 对于负数，sint32/sint64通过先采用Zigzag(将有符号转化为无符号)编码，再用Varint编码
#### 需要Length的，String,嵌套的Message,packed repeated fields
* String，Value使用UTF-8编码
* 嵌套消息，即在T-L-V里的V再嵌套一系列的T-L-V
* packed repeated，repeated可以看成数组，相同的Tag只存储一次，组成T-L-V-V-V，packed只能修饰repeated字段，或基本类型的repeated值
## 序列化
* Tag的number是唯一的，可以乱序，缺失和嵌套
* 序列化是根据Tag标识从小到大
* 其实协议没有保证序列化后各个平台的字节流都是一样的
## 使用建议
* Tag值尽量只使用1-15，且不要跳动使用
* 如果有负数，使用sint32/sint64，而不是int32/int64
* 对于repeated字段，尽量用packed=true修饰，T-L-V-V-V
## Any
* 有点类void，可以不定义具体字段的类型，但是需要二次序列化/反序列化
````
marshalany, _ := ptypes.MarshalAny(&anyy)
ph.Data = []*any.Any{marshalany}
phm, _ = proto.Marshal(&ph)
````
* any的类型也应该是定义好的，或者自己定义并能序列化(需要自己序列化url/[]byte)
## Oneof
````
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
````
* 字段中只有一个字段会被设置值，重复设置会清除前一个的设置，类union，不支持repeated
* 使用时需要用api获取到底是哪个字段被设置，Unmarshal是根据T-field决定oneof interface解析成具体的结构
````
switch x := o.TestOneof.(type)
````
## Map映射
````
map<key_type, value_type> map_field = N;
map<string, Project> projects = 3;
````
* 不支持repeated
* 序列化顺序不确定，key不能重复
## Json
````
import "github.com/golang/protobuf/jsonpb"
marshal := jsonpb.Marshaler{}
replyJson, _ := marshal.MarshalToString(reply)

type Marshaler struct {
	// OrigName specifies whether to use the original protobuf name for fields.
	OrigName bool

	// EnumsAsInts specifies whether to render enum values as integers,
	// as opposed to string values.
	EnumsAsInts bool

	// EmitDefaults specifies whether to render fields with zero values.
	EmitDefaults bool

	// Indent controls whether the output is compact or not.
	// If empty, the output is compact JSON. Otherwise, every JSON object
	// entry and JSON array value will be on its own line.
	// Each line will be preceded by repeated copies of Indent, where the
	// number of copies is the current indentation depth.
	Indent string

	// AnyResolver is used to resolve the google.protobuf.Any well-known type.
	// If unset, the global registry is used by default.
	AnyResolver AnyResolver
}
````
#### A proto3 JSON implementation may provide the following options:

* Emit fields with default values: Fields with default values are omitted by default in proto3 JSON output. An implementation may provide an option to override this behavior and output fields with their default values.
* Ignore unknown fields: Proto3 JSON parser should reject unknown fields by default but may provide an option to ignore unknown fields in parsing.
* Use proto field name instead of lowerCamelCase name: By default proto3 JSON printer should convert the field name to lowerCamelCase and use that as the JSON name. An implementation may provide an option to use proto field name as the JSON name instead. Proto3 JSON parsers are required to accept both the converted lowerCamelCase name and the proto field name.
* Emit enum values as integers instead of strings: The name of an enum value is used by default in JSON output. An option may be provided to use the numeric value of the enum value instead.
## GoGo库