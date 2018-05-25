/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package intstr

import (
	"encoding/json"
	"fmt"
	"math"
	"runtime/debug"
	"strconv"

	"github.com/golang/glog"
	"github.com/google/gofuzz"
)

// Int64OrString is a type that can hold an int64 or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
//
// +protobuf=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:openapi-gen=true
type Int64OrString struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,casttype=Type"`
	IntVal int64  `protobuf:"varint,2,opt,name=intVal"`
	StrVal string `protobuf:"bytes,3,opt,name=strVal"`
}

// FromInt64 creates an Int64OrString object with an int64 value. It is
// your responsibility not to call this method with a value greater
// than int64.
func FromInt64(val int64) Int64OrString {
	if val > math.MaxInt64 || val < math.MinInt64 {
		glog.Errorf("value: %d overflows int64\n%s\n", val, debug.Stack())
	}

	return Int64OrString{Type: Int64, IntVal: int64(val)}
}

// FromString64 creates an Int64OrString object with a string value.
func FromString64(val string) Int64OrString {
	return Int64OrString{Type: String, StrVal: val}
}

// Parse64 the given string and try to convert it to an integer before
// setting it as a string value.
func Parse64(val string) Int64OrString {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return FromString64(val)
	}
	return FromInt64(i)
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (intstr *Int64OrString) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		intstr.Type = String
		return json.Unmarshal(value, &intstr.StrVal)
	}
	intstr.Type = Int64
	return json.Unmarshal(value, &intstr.IntVal)
}

// String returns the string value, or the Itoa of the int value.
func (intstr *Int64OrString) String() string {
	if intstr.Type == String {
		return intstr.StrVal
	}

	return strconv.FormatInt(intstr.IntVal, 10)
}

// MarshalJSON implements the json.Marshaller interface.
func (intstr Int64OrString) MarshalJSON() ([]byte, error) {
	switch intstr.Type {
	case Int64:
		return json.Marshal(intstr.IntVal)
	case String:
		return json.Marshal(intstr.StrVal)
	default:
		return []byte{}, fmt.Errorf("impossible Int64OrString.Type")
	}
}

// OpenAPISchemaType is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
//
// See: https://github.com/kubernetes/kube-openapi/tree/master/pkg/generators
func (_ Int64OrString) OpenAPISchemaType() []string { return []string{"string"} }

// OpenAPISchemaFormat is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
func (_ Int64OrString) OpenAPISchemaFormat() string { return "int64-or-string" }

func (intstr *Int64OrString) Fuzz(c fuzz.Continue) {
	if intstr == nil {
		return
	}
	if c.RandBool() {
		intstr.Type = Int
		c.Fuzz(&intstr.IntVal)
		intstr.StrVal = ""
	} else {
		intstr.Type = String
		intstr.IntVal = 0
		c.Fuzz(&intstr.StrVal)
	}
}
