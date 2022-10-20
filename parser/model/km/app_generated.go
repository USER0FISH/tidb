package km

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

var _Null = make([]byte, 120)
var _NullReader = karmem.NewReader(_Null)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierCIStr      = 7937507240537370775
	PacketIdentifierKV         = 6793174473685253598
	PacketIdentifierColumnInfo = 13095150776460547620
)

type CIStr struct {
	O string
	L string
}

func NewCIStr() CIStr {
	return CIStr{}
}

func (x *CIStr) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierCIStr
}

func (x *CIStr) Reset() {
	x.Read((*CIStrViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *CIStr) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *CIStr) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__OSize := uint(1 * len(x.O))
	__OOffset, err := writer.Alloc(__OSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__OOffset))
	writer.Write4At(offset+0+4, uint32(__OSize))
	writer.Write4At(offset+0+4+4, 1)
	__OSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.O)), __OSize, __OSize}
	writer.WriteAt(__OOffset, *(*[]byte)(unsafe.Pointer(&__OSlice)))
	__LSize := uint(1 * len(x.L))
	__LOffset, err := writer.Alloc(__LSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__LOffset))
	writer.Write4At(offset+12+4, uint32(__LSize))
	writer.Write4At(offset+12+4+4, 1)
	__LSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.L)), __LSize, __LSize}
	writer.WriteAt(__LOffset, *(*[]byte)(unsafe.Pointer(&__LSlice)))

	return offset, nil
}

func (x *CIStr) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewCIStrViewer(reader, 0), reader)
}

func (x *CIStr) Read(viewer *CIStrViewer, reader *karmem.Reader) {
	__OString := viewer.O(reader)
	if x.O != __OString {
		__OStringCopy := make([]byte, len(__OString))
		copy(__OStringCopy, __OString)
		x.O = *(*string)(unsafe.Pointer(&__OStringCopy))
	}
	__LString := viewer.L(reader)
	if x.L != __LString {
		__LStringCopy := make([]byte, len(__LString))
		copy(__LStringCopy, __LString)
		x.L = *(*string)(unsafe.Pointer(&__LStringCopy))
	}
}

type KV struct {
	K string
	V bool
}

func NewKV() KV {
	return KV{}
}

func (x *KV) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierKV
}

func (x *KV) Reset() {
	x.Read((*KVViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *KV) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *KV) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(16)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__KSize := uint(1 * len(x.K))
	__KOffset, err := writer.Alloc(__KSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__KOffset))
	writer.Write4At(offset+0+4, uint32(__KSize))
	writer.Write4At(offset+0+4+4, 1)
	__KSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.K)), __KSize, __KSize}
	writer.WriteAt(__KOffset, *(*[]byte)(unsafe.Pointer(&__KSlice)))
	__VOffset := offset + 12
	writer.Write1At(__VOffset, *(*uint8)(unsafe.Pointer(&x.V)))

	return offset, nil
}

func (x *KV) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewKVViewer(reader, 0), reader)
}

func (x *KV) Read(viewer *KVViewer, reader *karmem.Reader) {
	__KString := viewer.K(reader)
	if x.K != __KString {
		__KStringCopy := make([]byte, len(__KString))
		copy(__KStringCopy, __KString)
		x.K = *(*string)(unsafe.Pointer(&__KStringCopy))
	}
	x.V = viewer.V()
}

type ColumnInfo struct {
	ID                     int64
	Name                   CIStr
	Offset                 int64
	OriginDefaultValueType byte
	OriginDefaultValue     string
	OriginDefaultValueBit  []byte
	DefaultValueType       byte
	DefaultValue           string
	DefaultIsExpr          bool
	GeneratedExprString    string
	Dependences            []KV
}

func NewColumnInfo() ColumnInfo {
	return ColumnInfo{}
}

func (x *ColumnInfo) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierColumnInfo
}

func (x *ColumnInfo) Reset() {
	x.Read((*ColumnInfoViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *ColumnInfo) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *ColumnInfo) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(120)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(115))
	__IDOffset := offset + 4
	writer.Write8At(__IDOffset, *(*uint64)(unsafe.Pointer(&x.ID)))
	__NameOffset := offset + 12
	if _, err := x.Name.Write(writer, __NameOffset); err != nil {
		return offset, err
	}
	__OffsetOffset := offset + 44
	writer.Write8At(__OffsetOffset, *(*uint64)(unsafe.Pointer(&x.Offset)))
	__OriginDefaultValueTypeOffset := offset + 52
	writer.Write1At(__OriginDefaultValueTypeOffset, *(*uint8)(unsafe.Pointer(&x.OriginDefaultValueType)))
	__OriginDefaultValueSize := uint(1 * len(x.OriginDefaultValue))
	__OriginDefaultValueOffset, err := writer.Alloc(__OriginDefaultValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+53, uint32(__OriginDefaultValueOffset))
	writer.Write4At(offset+53+4, uint32(__OriginDefaultValueSize))
	writer.Write4At(offset+53+4+4, 1)
	__OriginDefaultValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.OriginDefaultValue)), __OriginDefaultValueSize, __OriginDefaultValueSize}
	writer.WriteAt(__OriginDefaultValueOffset, *(*[]byte)(unsafe.Pointer(&__OriginDefaultValueSlice)))
	__OriginDefaultValueBitSize := uint(1 * len(x.OriginDefaultValueBit))
	__OriginDefaultValueBitOffset, err := writer.Alloc(__OriginDefaultValueBitSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+65, uint32(__OriginDefaultValueBitOffset))
	writer.Write4At(offset+65+4, uint32(__OriginDefaultValueBitSize))
	writer.Write4At(offset+65+4+4, 1)
	__OriginDefaultValueBitSlice := *(*[3]uint)(unsafe.Pointer(&x.OriginDefaultValueBit))
	__OriginDefaultValueBitSlice[1] = __OriginDefaultValueBitSize
	__OriginDefaultValueBitSlice[2] = __OriginDefaultValueBitSize
	writer.WriteAt(__OriginDefaultValueBitOffset, *(*[]byte)(unsafe.Pointer(&__OriginDefaultValueBitSlice)))
	__DefaultValueTypeOffset := offset + 77
	writer.Write1At(__DefaultValueTypeOffset, *(*uint8)(unsafe.Pointer(&x.DefaultValueType)))
	__DefaultValueSize := uint(1 * len(x.DefaultValue))
	__DefaultValueOffset, err := writer.Alloc(__DefaultValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+78, uint32(__DefaultValueOffset))
	writer.Write4At(offset+78+4, uint32(__DefaultValueSize))
	writer.Write4At(offset+78+4+4, 1)
	__DefaultValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.DefaultValue)), __DefaultValueSize, __DefaultValueSize}
	writer.WriteAt(__DefaultValueOffset, *(*[]byte)(unsafe.Pointer(&__DefaultValueSlice)))
	__DefaultIsExprOffset := offset + 90
	writer.Write1At(__DefaultIsExprOffset, *(*uint8)(unsafe.Pointer(&x.DefaultIsExpr)))
	__GeneratedExprStringSize := uint(1 * len(x.GeneratedExprString))
	__GeneratedExprStringOffset, err := writer.Alloc(__GeneratedExprStringSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+91, uint32(__GeneratedExprStringOffset))
	writer.Write4At(offset+91+4, uint32(__GeneratedExprStringSize))
	writer.Write4At(offset+91+4+4, 1)
	__GeneratedExprStringSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.GeneratedExprString)), __GeneratedExprStringSize, __GeneratedExprStringSize}
	writer.WriteAt(__GeneratedExprStringOffset, *(*[]byte)(unsafe.Pointer(&__GeneratedExprStringSlice)))
	__DependencesSize := uint(16 * len(x.Dependences))
	__DependencesOffset, err := writer.Alloc(__DependencesSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+103, uint32(__DependencesOffset))
	writer.Write4At(offset+103+4, uint32(__DependencesSize))
	writer.Write4At(offset+103+4+4, 16)
	for i := range x.Dependences {
		if _, err := x.Dependences[i].Write(writer, __DependencesOffset); err != nil {
			return offset, err
		}
		__DependencesOffset += 16
	}

	return offset, nil
}

func (x *ColumnInfo) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewColumnInfoViewer(reader, 0), reader)
}

func (x *ColumnInfo) Read(viewer *ColumnInfoViewer, reader *karmem.Reader) {
	x.ID = viewer.ID()
	x.Name.Read(viewer.Name(), reader)
	x.Offset = viewer.Offset()
	x.OriginDefaultValueType = viewer.OriginDefaultValueType()
	__OriginDefaultValueString := viewer.OriginDefaultValue(reader)
	if x.OriginDefaultValue != __OriginDefaultValueString {
		__OriginDefaultValueStringCopy := make([]byte, len(__OriginDefaultValueString))
		copy(__OriginDefaultValueStringCopy, __OriginDefaultValueString)
		x.OriginDefaultValue = *(*string)(unsafe.Pointer(&__OriginDefaultValueStringCopy))
	}
	__OriginDefaultValueBitSlice := viewer.OriginDefaultValueBit(reader)
	__OriginDefaultValueBitLen := len(__OriginDefaultValueBitSlice)
	if __OriginDefaultValueBitLen > cap(x.OriginDefaultValueBit) {
		x.OriginDefaultValueBit = append(x.OriginDefaultValueBit, make([]byte, __OriginDefaultValueBitLen-len(x.OriginDefaultValueBit))...)
	}
	if __OriginDefaultValueBitLen > len(x.OriginDefaultValueBit) {
		x.OriginDefaultValueBit = x.OriginDefaultValueBit[:__OriginDefaultValueBitLen]
	}
	copy(x.OriginDefaultValueBit, __OriginDefaultValueBitSlice)
	x.OriginDefaultValueBit = x.OriginDefaultValueBit[:__OriginDefaultValueBitLen]
	x.DefaultValueType = viewer.DefaultValueType()
	__DefaultValueString := viewer.DefaultValue(reader)
	if x.DefaultValue != __DefaultValueString {
		__DefaultValueStringCopy := make([]byte, len(__DefaultValueString))
		copy(__DefaultValueStringCopy, __DefaultValueString)
		x.DefaultValue = *(*string)(unsafe.Pointer(&__DefaultValueStringCopy))
	}
	x.DefaultIsExpr = viewer.DefaultIsExpr()
	__GeneratedExprStringString := viewer.GeneratedExprString(reader)
	if x.GeneratedExprString != __GeneratedExprStringString {
		__GeneratedExprStringStringCopy := make([]byte, len(__GeneratedExprStringString))
		copy(__GeneratedExprStringStringCopy, __GeneratedExprStringString)
		x.GeneratedExprString = *(*string)(unsafe.Pointer(&__GeneratedExprStringStringCopy))
	}
	__DependencesSlice := viewer.Dependences(reader)
	__DependencesLen := len(__DependencesSlice)
	if __DependencesLen > cap(x.Dependences) {
		x.Dependences = append(x.Dependences, make([]KV, __DependencesLen-len(x.Dependences))...)
	}
	if __DependencesLen > len(x.Dependences) {
		x.Dependences = x.Dependences[:__DependencesLen]
	}
	for i := 0; i < __DependencesLen; i++ {
		x.Dependences[i].Read(&__DependencesSlice[i], reader)
	}
	x.Dependences = x.Dependences[:__DependencesLen]
}

type CIStrViewer struct {
	_data [32]byte
}

func NewCIStrViewer(reader *karmem.Reader, offset uint32) (v *CIStrViewer) {
	if !reader.IsValidOffset(offset, 32) {
		return (*CIStrViewer)(unsafe.Pointer(&_Null))
	}
	v = (*CIStrViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *CIStrViewer) size() uint32 {
	return 32
}
func (x *CIStrViewer) O(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *CIStrViewer) L(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}

type KVViewer struct {
	_data [16]byte
}

func NewKVViewer(reader *karmem.Reader, offset uint32) (v *KVViewer) {
	if !reader.IsValidOffset(offset, 16) {
		return (*KVViewer)(unsafe.Pointer(&_Null))
	}
	v = (*KVViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *KVViewer) size() uint32 {
	return 16
}
func (x *KVViewer) K(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *KVViewer) V() (v bool) {
	return *(*bool)(unsafe.Add(unsafe.Pointer(&x._data), 12))
}

type ColumnInfoViewer struct {
	_data [120]byte
}

func NewColumnInfoViewer(reader *karmem.Reader, offset uint32) (v *ColumnInfoViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*ColumnInfoViewer)(unsafe.Pointer(&_Null))
	}
	v = (*ColumnInfoViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*ColumnInfoViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *ColumnInfoViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *ColumnInfoViewer) ID() (v int64) {
	if 4+8 > x.size() {
		return v
	}
	return *(*int64)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}
func (x *ColumnInfoViewer) Name() (v *CIStrViewer) {
	if 12+32 > x.size() {
		return (*CIStrViewer)(unsafe.Pointer(&_Null))
	}
	return (*CIStrViewer)(unsafe.Add(unsafe.Pointer(&x._data), 12))
}
func (x *ColumnInfoViewer) Offset() (v int64) {
	if 44+8 > x.size() {
		return v
	}
	return *(*int64)(unsafe.Add(unsafe.Pointer(&x._data), 44))
}
func (x *ColumnInfoViewer) OriginDefaultValueType() (v byte) {
	if 52+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 52))
}
func (x *ColumnInfoViewer) OriginDefaultValue(reader *karmem.Reader) (v string) {
	if 53+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 53))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 53+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *ColumnInfoViewer) OriginDefaultValueBit(reader *karmem.Reader) (v []byte) {
	if 65+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 65))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 65+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *ColumnInfoViewer) DefaultValueType() (v byte) {
	if 77+1 > x.size() {
		return v
	}
	return *(*byte)(unsafe.Add(unsafe.Pointer(&x._data), 77))
}
func (x *ColumnInfoViewer) DefaultValue(reader *karmem.Reader) (v string) {
	if 78+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 78))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 78+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *ColumnInfoViewer) DefaultIsExpr() (v bool) {
	if 90+1 > x.size() {
		return v
	}
	return *(*bool)(unsafe.Add(unsafe.Pointer(&x._data), 90))
}
func (x *ColumnInfoViewer) GeneratedExprString(reader *karmem.Reader) (v string) {
	if 91+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 91))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 91+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *ColumnInfoViewer) Dependences(reader *karmem.Reader) (v []KVViewer) {
	if 103+12 > x.size() {
		return []KVViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 103))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 103+4))
	if !reader.IsValidOffset(offset, size) {
		return []KVViewer{}
	}
	length := uintptr(size / 16)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]KVViewer)(unsafe.Pointer(&slice))
}
