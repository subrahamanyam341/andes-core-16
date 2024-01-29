// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smartContractResult.proto

package smartContractResult

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_subrahamanyam341_andes_core_16_data "github.com/subrahamanyam341/andes-core-16/data"
	github_com_subrahamanyam341_andes_core_16_data_vm "github.com/subrahamanyam341/andes-core-16/data/vm"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SmartContractResult struct {
	Nonce          uint64                                                     `protobuf:"varint,1,opt,name=Nonce,proto3" json:"nonce"`
	Value          *math_big.Int                                              `protobuf:"bytes,2,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/subrahamanyam341/andes-core-16/data.BigIntCaster" json:"value"`
	RcvAddr        []byte                                                     `protobuf:"bytes,3,opt,name=RcvAddr,proto3" json:"receiver"`
	SndAddr        []byte                                                     `protobuf:"bytes,4,opt,name=SndAddr,proto3" json:"sender"`
	RelayerAddr    []byte                                                     `protobuf:"bytes,5,opt,name=RelayerAddr,proto3" json:"relayer"`
	RelayedValue   *math_big.Int                                              `protobuf:"bytes,6,opt,name=RelayedValue,proto3,casttypewith=math/big.Int;github.com/subrahamanyam341/andes-core-16/data.BigIntCaster" json:"relayedValue"`
	Code           []byte                                                     `protobuf:"bytes,7,opt,name=Code,proto3" json:"code,omitempty"`
	Data           []byte                                                     `protobuf:"bytes,8,opt,name=Data,proto3" json:"data,omitempty"`
	PrevTxHash     []byte                                                     `protobuf:"bytes,9,opt,name=PrevTxHash,proto3" json:"prevTxHash"`
	OriginalTxHash []byte                                                     `protobuf:"bytes,10,opt,name=OriginalTxHash,proto3" json:"originalTxHash"`
	GasLimit       uint64                                                     `protobuf:"varint,11,opt,name=GasLimit,proto3" json:"gasLimit"`
	GasPrice       uint64                                                     `protobuf:"varint,12,opt,name=GasPrice,proto3" json:"gasPrice"`
	CallType       github_com_subrahamanyam341_andes_core_16_data_vm.CallType `protobuf:"varint,13,opt,name=CallType,proto3,casttype=github.com/subrahamanyam341/andes-core-16/data/vm.CallType" json:"callType"`
	CodeMetadata   []byte                                                     `protobuf:"bytes,14,opt,name=CodeMetadata,proto3" json:"codeMetadata,omitempty"`
	ReturnMessage  []byte                                                     `protobuf:"bytes,15,opt,name=ReturnMessage,proto3" json:"returnMessage,omitempty"`
	OriginalSender []byte                                                     `protobuf:"bytes,16,opt,name=OriginalSender,proto3" json:"originalSender,omitempty"`
}

func (m *SmartContractResult) Reset()      { *m = SmartContractResult{} }
func (*SmartContractResult) ProtoMessage() {}
func (*SmartContractResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1605de0d3d805, []int{0}
}
func (m *SmartContractResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartContractResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SmartContractResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContractResult.Merge(m, src)
}
func (m *SmartContractResult) XXX_Size() int {
	return m.Size()
}
func (m *SmartContractResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContractResult.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContractResult proto.InternalMessageInfo

func (m *SmartContractResult) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SmartContractResult) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SmartContractResult) GetRcvAddr() []byte {
	if m != nil {
		return m.RcvAddr
	}
	return nil
}

func (m *SmartContractResult) GetSndAddr() []byte {
	if m != nil {
		return m.SndAddr
	}
	return nil
}

func (m *SmartContractResult) GetRelayerAddr() []byte {
	if m != nil {
		return m.RelayerAddr
	}
	return nil
}

func (m *SmartContractResult) GetRelayedValue() *math_big.Int {
	if m != nil {
		return m.RelayedValue
	}
	return nil
}

func (m *SmartContractResult) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *SmartContractResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SmartContractResult) GetPrevTxHash() []byte {
	if m != nil {
		return m.PrevTxHash
	}
	return nil
}

func (m *SmartContractResult) GetOriginalTxHash() []byte {
	if m != nil {
		return m.OriginalTxHash
	}
	return nil
}

func (m *SmartContractResult) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *SmartContractResult) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *SmartContractResult) GetCallType() github_com_subrahamanyam341_andes_core_16_data_vm.CallType {
	if m != nil {
		return m.CallType
	}
	return 0
}

func (m *SmartContractResult) GetCodeMetadata() []byte {
	if m != nil {
		return m.CodeMetadata
	}
	return nil
}

func (m *SmartContractResult) GetReturnMessage() []byte {
	if m != nil {
		return m.ReturnMessage
	}
	return nil
}

func (m *SmartContractResult) GetOriginalSender() []byte {
	if m != nil {
		return m.OriginalSender
	}
	return nil
}

func init() {
	proto.RegisterType((*SmartContractResult)(nil), "proto.SmartContractResult")
}

func init() { proto.RegisterFile("smartContractResult.proto", fileDescriptor_edc1605de0d3d805) }

var fileDescriptor_edc1605de0d3d805 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0x6b, 0xb6, 0xae, 0x9d, 0xd7, 0x15, 0xe4, 0x49, 0x60, 0x06, 0xb2, 0x27, 0x84, 0xa6,
	0x1e, 0x68, 0xab, 0x69, 0x88, 0xc3, 0x90, 0x90, 0xd6, 0x4e, 0x62, 0x43, 0x0c, 0x26, 0x6f, 0x70,
	0x40, 0x5c, 0xdc, 0xc4, 0x4b, 0x23, 0x9a, 0xb8, 0x72, 0x9c, 0x8a, 0xde, 0xf6, 0x13, 0xf8, 0x19,
	0x88, 0x5f, 0xc2, 0x71, 0xc7, 0x9d, 0x02, 0xcb, 0x38, 0xa0, 0x9c, 0x76, 0xe6, 0x84, 0xe2, 0x74,
	0x6d, 0x32, 0xb8, 0x20, 0x71, 0x4a, 0xfc, 0xbe, 0xcf, 0xf7, 0x7e, 0xf1, 0x17, 0xcb, 0xf0, 0x6e,
	0xe0, 0x71, 0xa5, 0xbb, 0xd2, 0xd7, 0x8a, 0x5b, 0x9a, 0x89, 0x20, 0x1c, 0xe8, 0xd6, 0x50, 0x49,
	0x2d, 0x51, 0xd9, 0x3c, 0x56, 0x9b, 0x8e, 0xab, 0xfb, 0x61, 0xaf, 0x65, 0x49, 0xaf, 0xed, 0x48,
	0x47, 0xb6, 0x8d, 0xdc, 0x0b, 0x8f, 0xcd, 0xca, 0x2c, 0xcc, 0x5b, 0x56, 0xf5, 0xe0, 0x47, 0x05,
	0xae, 0x1c, 0xfe, 0x99, 0x89, 0x28, 0x2c, 0xbf, 0x92, 0xbe, 0x25, 0x30, 0x58, 0x03, 0x8d, 0xf9,
	0xce, 0x62, 0x12, 0xd1, 0xb2, 0x9f, 0x0a, 0x2c, 0xd3, 0xd1, 0x07, 0x58, 0x7e, 0xcb, 0x07, 0xa1,
	0xc0, 0x37, 0xd6, 0x40, 0xa3, 0xd6, 0x79, 0x93, 0x02, 0xa3, 0x54, 0xf8, 0xf2, 0x8d, 0xee, 0x7a,
	0x5c, 0xf7, 0xdb, 0x3d, 0xd7, 0x69, 0xed, 0xf9, 0xfa, 0x69, 0xee, 0x83, 0x82, 0xb0, 0xa7, 0x78,
	0x9f, 0x7b, 0xdc, 0x1f, 0x73, 0x6f, 0xf3, 0xf1, 0x46, 0x9b, 0xfb, 0xb6, 0x08, 0x9a, 0x96, 0x54,
	0xa2, 0xb9, 0xf1, 0xa4, 0x6d, 0x73, 0xcd, 0x5b, 0x1d, 0xd7, 0xd9, 0xf3, 0x75, 0x97, 0x07, 0x5a,
	0x28, 0x96, 0xf5, 0x40, 0xeb, 0xb0, 0xc2, 0xac, 0xd1, 0xb6, 0x6d, 0x2b, 0x3c, 0x67, 0xda, 0xd5,
	0x92, 0x88, 0x56, 0x95, 0xb0, 0x84, 0x3b, 0x12, 0x8a, 0x5d, 0x99, 0xe8, 0x21, 0xac, 0x1c, 0xfa,
	0xb6, 0xe1, 0xe6, 0x0d, 0x07, 0x93, 0x88, 0x2e, 0x04, 0xc2, 0xb7, 0x53, 0x6a, 0x62, 0xa1, 0x26,
	0x5c, 0x62, 0x62, 0xc0, 0xc7, 0x42, 0x19, 0xb2, 0x6c, 0xc8, 0xa5, 0x24, 0xa2, 0x15, 0x95, 0xc9,
	0x2c, 0xef, 0xa3, 0x13, 0x00, 0x6b, 0xd9, 0xda, 0xce, 0x76, 0xbc, 0x60, 0x0a, 0xde, 0x27, 0x11,
	0xad, 0xa9, 0x9c, 0xfe, 0x5f, 0x37, 0x5e, 0xe8, 0x88, 0xd6, 0xe1, 0x7c, 0x57, 0xda, 0x02, 0x57,
	0x4c, 0x67, 0x94, 0x44, 0xb4, 0x6e, 0x49, 0x5b, 0x3c, 0x92, 0x9e, 0xab, 0x85, 0x37, 0xd4, 0x63,
	0x66, 0xfc, 0x94, 0xdb, 0xe1, 0x9a, 0xe3, 0xea, 0x8c, 0x4b, 0xa3, 0xf3, 0x5c, 0xea, 0xa3, 0x16,
	0x84, 0x07, 0x4a, 0x8c, 0x8e, 0x3e, 0xee, 0xf2, 0xa0, 0x8f, 0x17, 0x0d, 0x5d, 0x4f, 0x22, 0x0a,
	0x87, 0x53, 0x95, 0xe5, 0x08, 0xb4, 0x05, 0xeb, 0xaf, 0x95, 0xeb, 0xb8, 0x3e, 0x1f, 0x4c, 0x6a,
	0xe0, 0xac, 0x83, 0x2c, 0x38, 0xec, 0x1a, 0x89, 0x1a, 0xb0, 0xfa, 0x9c, 0x07, 0x2f, 0x5d, 0xcf,
	0xd5, 0x78, 0xc9, 0x1c, 0x26, 0xf3, 0xf3, 0x9c, 0x89, 0xc6, 0xa6, 0xee, 0x84, 0x3c, 0x50, 0xae,
	0x25, 0x70, 0xad, 0x40, 0x1a, 0x8d, 0x4d, 0x5d, 0x74, 0x0c, 0xab, 0x5d, 0x3e, 0x18, 0x1c, 0x8d,
	0x87, 0x02, 0x2f, 0xaf, 0x81, 0xc6, 0x5c, 0xe7, 0x45, 0x4a, 0x5a, 0x13, 0xed, 0x57, 0x44, 0xb7,
	0xfe, 0x6d, 0xf8, 0xed, 0x91, 0xd7, 0xba, 0x4a, 0x64, 0xd3, 0x6c, 0xf4, 0x0c, 0xd6, 0xd2, 0xb9,
	0xee, 0x0b, 0xcd, 0x53, 0x0a, 0xd7, 0xcd, 0xae, 0x57, 0x93, 0x88, 0xde, 0xb6, 0x72, 0x7a, 0x6e,
	0xbe, 0x05, 0x1e, 0x6d, 0xc3, 0x65, 0x26, 0x74, 0xa8, 0xfc, 0x7d, 0x11, 0x04, 0xdc, 0x11, 0xf8,
	0xa6, 0x09, 0xb8, 0x97, 0x44, 0xf4, 0x8e, 0xca, 0x1b, 0xb9, 0x84, 0x62, 0x05, 0xda, 0x99, 0x8d,
	0xfe, 0xd0, 0x9c, 0x63, 0x7c, 0xcb, 0x64, 0xdc, 0x4f, 0x22, 0x8a, 0x65, 0xc1, 0xc9, 0x85, 0x5c,
	0xab, 0xe9, 0xec, 0x9f, 0x9e, 0x93, 0xd2, 0xd9, 0x39, 0x29, 0x5d, 0x9e, 0x13, 0x70, 0x12, 0x13,
	0xf0, 0x39, 0x26, 0xe0, 0x6b, 0x4c, 0xc0, 0x69, 0x4c, 0xc0, 0x59, 0x4c, 0xc0, 0xf7, 0x98, 0x80,
	0x9f, 0x31, 0x29, 0x5d, 0xc6, 0x04, 0x7c, 0xba, 0x20, 0xa5, 0xd3, 0x0b, 0x52, 0x3a, 0xbb, 0x20,
	0xa5, 0x77, 0x2b, 0x7f, 0xb9, 0x71, 0x7a, 0x0b, 0xe6, 0xf2, 0xd8, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x67, 0x06, 0x1b, 0x3e, 0x8f, 0x04, 0x00, 0x00,
}

func (this *SmartContractResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SmartContractResult)
	if !ok {
		that2, ok := that.(SmartContractResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	{
		__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.RcvAddr, that1.RcvAddr) {
		return false
	}
	if !bytes.Equal(this.SndAddr, that1.SndAddr) {
		return false
	}
	if !bytes.Equal(this.RelayerAddr, that1.RelayerAddr) {
		return false
	}
	{
		__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
		if !__caster.Equal(this.RelayedValue, that1.RelayedValue) {
			return false
		}
	}
	if !bytes.Equal(this.Code, that1.Code) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.PrevTxHash, that1.PrevTxHash) {
		return false
	}
	if !bytes.Equal(this.OriginalTxHash, that1.OriginalTxHash) {
		return false
	}
	if this.GasLimit != that1.GasLimit {
		return false
	}
	if this.GasPrice != that1.GasPrice {
		return false
	}
	if this.CallType != that1.CallType {
		return false
	}
	if !bytes.Equal(this.CodeMetadata, that1.CodeMetadata) {
		return false
	}
	if !bytes.Equal(this.ReturnMessage, that1.ReturnMessage) {
		return false
	}
	if !bytes.Equal(this.OriginalSender, that1.OriginalSender) {
		return false
	}
	return true
}
func (this *SmartContractResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 20)
	s = append(s, "&smartContractResult.SmartContractResult{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "RcvAddr: "+fmt.Sprintf("%#v", this.RcvAddr)+",\n")
	s = append(s, "SndAddr: "+fmt.Sprintf("%#v", this.SndAddr)+",\n")
	s = append(s, "RelayerAddr: "+fmt.Sprintf("%#v", this.RelayerAddr)+",\n")
	s = append(s, "RelayedValue: "+fmt.Sprintf("%#v", this.RelayedValue)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "PrevTxHash: "+fmt.Sprintf("%#v", this.PrevTxHash)+",\n")
	s = append(s, "OriginalTxHash: "+fmt.Sprintf("%#v", this.OriginalTxHash)+",\n")
	s = append(s, "GasLimit: "+fmt.Sprintf("%#v", this.GasLimit)+",\n")
	s = append(s, "GasPrice: "+fmt.Sprintf("%#v", this.GasPrice)+",\n")
	s = append(s, "CallType: "+fmt.Sprintf("%#v", this.CallType)+",\n")
	s = append(s, "CodeMetadata: "+fmt.Sprintf("%#v", this.CodeMetadata)+",\n")
	s = append(s, "ReturnMessage: "+fmt.Sprintf("%#v", this.ReturnMessage)+",\n")
	s = append(s, "OriginalSender: "+fmt.Sprintf("%#v", this.OriginalSender)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSmartContractResult(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SmartContractResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartContractResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartContractResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OriginalSender) > 0 {
		i -= len(m.OriginalSender)
		copy(dAtA[i:], m.OriginalSender)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.OriginalSender)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.ReturnMessage) > 0 {
		i -= len(m.ReturnMessage)
		copy(dAtA[i:], m.ReturnMessage)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.ReturnMessage)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.CodeMetadata) > 0 {
		i -= len(m.CodeMetadata)
		copy(dAtA[i:], m.CodeMetadata)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.CodeMetadata)))
		i--
		dAtA[i] = 0x72
	}
	if m.CallType != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.CallType))
		i--
		dAtA[i] = 0x68
	}
	if m.GasPrice != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x60
	}
	if m.GasLimit != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x58
	}
	if len(m.OriginalTxHash) > 0 {
		i -= len(m.OriginalTxHash)
		copy(dAtA[i:], m.OriginalTxHash)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.OriginalTxHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.PrevTxHash) > 0 {
		i -= len(m.PrevTxHash)
		copy(dAtA[i:], m.PrevTxHash)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.PrevTxHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x3a
	}
	{
		__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
		size := __caster.Size(m.RelayedValue)
		i -= size
		if _, err := __caster.MarshalTo(m.RelayedValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSmartContractResult(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.RelayerAddr) > 0 {
		i -= len(m.RelayerAddr)
		copy(dAtA[i:], m.RelayerAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.RelayerAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SndAddr) > 0 {
		i -= len(m.SndAddr)
		copy(dAtA[i:], m.SndAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.SndAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RcvAddr) > 0 {
		i -= len(m.RcvAddr)
		copy(dAtA[i:], m.RcvAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.RcvAddr)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSmartContractResult(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Nonce != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSmartContractResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovSmartContractResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SmartContractResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovSmartContractResult(uint64(m.Nonce))
	}
	{
		__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.RcvAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.SndAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.RelayerAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	{
		__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
		l = __caster.Size(m.RelayedValue)
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.PrevTxHash)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.OriginalTxHash)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovSmartContractResult(uint64(m.GasLimit))
	}
	if m.GasPrice != 0 {
		n += 1 + sovSmartContractResult(uint64(m.GasPrice))
	}
	if m.CallType != 0 {
		n += 1 + sovSmartContractResult(uint64(m.CallType))
	}
	l = len(m.CodeMetadata)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.ReturnMessage)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.OriginalSender)
	if l > 0 {
		n += 2 + l + sovSmartContractResult(uint64(l))
	}
	return n
}

func sovSmartContractResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSmartContractResult(x uint64) (n int) {
	return sovSmartContractResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SmartContractResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SmartContractResult{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`RcvAddr:` + fmt.Sprintf("%v", this.RcvAddr) + `,`,
		`SndAddr:` + fmt.Sprintf("%v", this.SndAddr) + `,`,
		`RelayerAddr:` + fmt.Sprintf("%v", this.RelayerAddr) + `,`,
		`RelayedValue:` + fmt.Sprintf("%v", this.RelayedValue) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`PrevTxHash:` + fmt.Sprintf("%v", this.PrevTxHash) + `,`,
		`OriginalTxHash:` + fmt.Sprintf("%v", this.OriginalTxHash) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`GasPrice:` + fmt.Sprintf("%v", this.GasPrice) + `,`,
		`CallType:` + fmt.Sprintf("%v", this.CallType) + `,`,
		`CodeMetadata:` + fmt.Sprintf("%v", this.CodeMetadata) + `,`,
		`ReturnMessage:` + fmt.Sprintf("%v", this.ReturnMessage) + `,`,
		`OriginalSender:` + fmt.Sprintf("%v", this.OriginalSender) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSmartContractResult(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SmartContractResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmartContractResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SmartContractResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartContractResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvAddr = append(m.RcvAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvAddr == nil {
				m.RcvAddr = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndAddr = append(m.SndAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SndAddr == nil {
				m.SndAddr = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerAddr = append(m.RelayerAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RelayerAddr == nil {
				m.RelayerAddr = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayedValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_subrahamanyam341_andes_core_16_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.RelayedValue = tmp
				}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevTxHash = append(m.PrevTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevTxHash == nil {
				m.PrevTxHash = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalTxHash = append(m.OriginalTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.OriginalTxHash == nil {
				m.OriginalTxHash = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallType", wireType)
			}
			m.CallType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallType |= github_com_subrahamanyam341_andes_core_16_data_vm.CallType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeMetadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeMetadata = append(m.CodeMetadata[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeMetadata == nil {
				m.CodeMetadata = []byte{}
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnMessage", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReturnMessage = append(m.ReturnMessage[:0], dAtA[iNdEx:postIndex]...)
			if m.ReturnMessage == nil {
				m.ReturnMessage = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalSender = append(m.OriginalSender[:0], dAtA[iNdEx:postIndex]...)
			if m.OriginalSender == nil {
				m.OriginalSender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmartContractResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSmartContractResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSmartContractResult
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSmartContractResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSmartContractResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSmartContractResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSmartContractResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSmartContractResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSmartContractResult = fmt.Errorf("proto: unexpected end of group")
)
