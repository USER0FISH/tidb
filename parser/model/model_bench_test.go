// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"encoding/json"
	"testing"

	"capnproto.org/go/capnp/v3"
	capnpt "github.com/pingcap/tidb/parser/model/capnp"
	kmt "github.com/pingcap/tidb/parser/model/km"
	protot "github.com/pingcap/tidb/parser/model/proto"
	"github.com/stretchr/testify/require"
	karmem "karmem.org/golang"

	//"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
)

type TColumnInfo struct {
	ID                    int64           `json:"id"`
	Name                  CIStr           `json:"name"`
	Offset                int             `json:"offset"`
	OriginDefaultValue    interface{}     `json:"origin_default"`
	OriginDefaultValueBit []byte          `json:"origin_default_bit"`
	DefaultValue          interface{}     `json:"default"`
	DefaultValueBit       []byte          `json:"default_bit"`
	DefaultIsExpr         bool            `json:"default_is_expr"`
	GeneratedExprString   string          `json:"generated_expr_string"`
	Dependences           map[string]bool `json:"dependences"`
}

func makeTCol() *TColumnInfo {
	return &TColumnInfo{
		ID: 6,
		Name: CIStr{
			O: "5454",
			L: "$545",
		},
		Offset:                6,
		OriginDefaultValue:    "vcbbcvbvcbcv",
		OriginDefaultValueBit: []byte("bbbbbbbbb98345435"),
		DefaultIsExpr:         true,
		GeneratedExprString:   "vcbvcby56yrt",
		Dependences: map[string]bool{
			"xcvxcv": true,
			"x3vxcv": true,
			"x4vxcv": true,
			"x6vxcv": true,
		},
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	col := makeTCol()
	c, err := json.Marshal(col)
	require.Nil(b, err)
	ncol := &TColumnInfo{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.Nil(b, json.Unmarshal(c, ncol))
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	col := makeTCol()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(col)
		require.Nil(b, err)
	}
}

func makeTColCap(b *testing.B, msg *capnp.Message, seg *capnp.Segment) capnpt.ColumnInfo {
	col, err := capnpt.NewColumnInfo(seg)
	col.SetId(6)
	name, err := col.NewName()
	require.Nil(b, err)
	name.SetO("5454")
	name.SetL("$454")
	col.SetOffset(6)
	col.SetOriginDefaultValue("vcbbcvbvcbcv")
	col.SetOriginDefaultValueBit([]byte("bbbbbbbbb98345435"))
	col.SetDefaultIsExpr(true)
	col.SetGeneratedExprString("vcbvcby56yrt")
	deps, err := col.NewDependences(4)
	require.Nil(b, err)
	deps.Set(0, "xcvxcv")
	deps.Set(1, "44vxcv")
	deps.Set(2, "4cvxcv")
	deps.Set(3, "6cvxcv")
	return col
}

func BenchmarkCapnpUnmarshal(b *testing.B) {
	msg, seg := capnp.NewSingleSegmentMessage(nil)
	makeTColCap(b, msg, seg)
	c, err := msg.Marshal()
	require.Nil(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nmsg, err := capnp.Unmarshal(c)
		require.Nil(b, err)
		_, err = capnpt.ReadRootColumnInfo(nmsg)
		require.Nil(b, err)
	}
}

func BenchmarkCapnpMarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msg, seg := capnp.NewSingleSegmentMessage(nil)
		makeTColCap(b, msg, seg)
		_, err := msg.Marshal()
		require.Nil(b, err)
	}
}

/*
func BenchmarkMsgpackUnmarshal(b *testing.B) {
	col := makeTCol()
	c, err := msgpack.Marshal(col)
	require.Nil(b, err)
	ncol := &TColumnInfo{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.Nil(b, msgpack.Unmarshal(c, ncol))
	}
}
*/

func makeTColProto() *protot.ColumnInfo {
	return &protot.ColumnInfo{
		ID: 6,
		Name: &protot.CIStr{
			O: "5454",
			L: "$545",
		},
		Offset:                6,
		OriginDefaultValue:    "vcbbcvbvcbcv",
		OriginDefaultValueBit: []byte("bbbbbbbbb98345435"),
		DefaultIsExpr:         true,
		GeneratedExprString:   "vcbvcby56yrt",
		Dependences: map[string]bool{
			"xcvxcv": true,
			"x3vxcv": true,
			"x4vxcv": true,
			"x6vxcv": true,
		},
	}
}

func BenchmarkProtoUnmarshal(b *testing.B) {
	col := makeTColProto()
	c, err := proto.Marshal(col)
	require.Nil(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.Nil(b, proto.Unmarshal(c, col))
	}
}

func BenchmarkProtoMarshal(b *testing.B) {
	col := makeTColProto()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(col)
		require.Nil(b, err)
	}
}

func makeTColKM() *kmt.ColumnInfo {
	return &kmt.ColumnInfo{
		ID: 6,
		Name: kmt.CIStr{
			O: "5454",
			L: "$545",
		},
		Offset:                6,
		OriginDefaultValue:    "vcbbcvbvcbcv",
		OriginDefaultValueBit: []byte("bbbbbbbbb98345435"),
		DefaultIsExpr:         true,
		GeneratedExprString:   "vcbvcby56yrt",
		Dependences: []kmt.KV{
			{"xcvxcv", true},
			{"x2vxcv", true},
			{"x3vxcv", true},
			{"x4vxcv", true},
		},
	}
}

func BenchmarkKMUnmarshal(b *testing.B) {
	col := makeTColKM()
	writer := karmem.NewWriter(64)
	_, err := col.WriteAsRoot(writer)
	require.Nil(b, err)
	content := writer.Bytes()
	ncol := &kmt.ColumnInfo{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := karmem.NewReader(append([]byte{}, content...))
		ncol.ReadAsRoot(reader)
	}
}

func BenchmarkKMMarshal(b *testing.B) {
	col := makeTColKM()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writer := karmem.NewWriter(512)
		_, err := col.WriteAsRoot(writer)
		require.Nil(b, err)
	}
}

func BenchmarkRawAccess(b *testing.B) {
	col := makeTColKM()
	writer := karmem.NewWriter(64)
	_, err := col.WriteAsRoot(writer)
	require.Nil(b, err)
	content := writer.Bytes()
	reader := karmem.NewReader(append([]byte{}, content...))
	var a string
	viewer := kmt.NewColumnInfoViewer(reader, 0)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = viewer.Name().L(reader)
	}
	b.ReportMetric(0, a)
}

func BenchmarkKMAccess(b *testing.B) {
	col := makeTColKM()
	var a string
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = col.Name.L
	}
	b.ReportMetric(0, a)
}
