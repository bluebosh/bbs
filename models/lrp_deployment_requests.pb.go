// Code generated by protoc-gen-gogo.
// source: lrp_deployment_requests.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LRPDeploymentLifecycleResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *LRPDeploymentLifecycleResponse) Reset()      { *m = LRPDeploymentLifecycleResponse{} }
func (*LRPDeploymentLifecycleResponse) ProtoMessage() {}
func (*LRPDeploymentLifecycleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorLrpDeploymentRequests, []int{0}
}

func (m *LRPDeploymentLifecycleResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type CreateLRPDeploymentRequest struct {
	Definition *LRPDeploymentDefinition `protobuf:"bytes,1,opt,name=definition" json:"definition,omitempty"`
}

func (m *CreateLRPDeploymentRequest) Reset()      { *m = CreateLRPDeploymentRequest{} }
func (*CreateLRPDeploymentRequest) ProtoMessage() {}
func (*CreateLRPDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorLrpDeploymentRequests, []int{1}
}

func (m *CreateLRPDeploymentRequest) GetDefinition() *LRPDeploymentDefinition {
	if m != nil {
		return m.Definition
	}
	return nil
}

type UpdateLRPDeploymentRequest struct {
	Id     string               `protobuf:"bytes,1,opt,name=id" json:"id"`
	Update *LRPDeploymentUpdate `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
}

func (m *UpdateLRPDeploymentRequest) Reset()      { *m = UpdateLRPDeploymentRequest{} }
func (*UpdateLRPDeploymentRequest) ProtoMessage() {}
func (*UpdateLRPDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorLrpDeploymentRequests, []int{2}
}

func (m *UpdateLRPDeploymentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateLRPDeploymentRequest) GetUpdate() *LRPDeploymentUpdate {
	if m != nil {
		return m.Update
	}
	return nil
}

type RemoveLRPDeploymentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id"`
}

func (m *RemoveLRPDeploymentRequest) Reset()      { *m = RemoveLRPDeploymentRequest{} }
func (*RemoveLRPDeploymentRequest) ProtoMessage() {}
func (*RemoveLRPDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorLrpDeploymentRequests, []int{3}
}

func (m *RemoveLRPDeploymentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ActivateLRPDeploymentDefinitionRequest struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id"`
	DefinitionId string `protobuf:"bytes,2,opt,name=definition_id,json=definitionId" json:"definition_id"`
}

func (m *ActivateLRPDeploymentDefinitionRequest) Reset() {
	*m = ActivateLRPDeploymentDefinitionRequest{}
}
func (*ActivateLRPDeploymentDefinitionRequest) ProtoMessage() {}
func (*ActivateLRPDeploymentDefinitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorLrpDeploymentRequests, []int{4}
}

func (m *ActivateLRPDeploymentDefinitionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActivateLRPDeploymentDefinitionRequest) GetDefinitionId() string {
	if m != nil {
		return m.DefinitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*LRPDeploymentLifecycleResponse)(nil), "models.LRPDeploymentLifecycleResponse")
	proto.RegisterType((*CreateLRPDeploymentRequest)(nil), "models.CreateLRPDeploymentRequest")
	proto.RegisterType((*UpdateLRPDeploymentRequest)(nil), "models.UpdateLRPDeploymentRequest")
	proto.RegisterType((*RemoveLRPDeploymentRequest)(nil), "models.RemoveLRPDeploymentRequest")
	proto.RegisterType((*ActivateLRPDeploymentDefinitionRequest)(nil), "models.ActivateLRPDeploymentDefinitionRequest")
}
func (this *LRPDeploymentLifecycleResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LRPDeploymentLifecycleResponse)
	if !ok {
		that2, ok := that.(LRPDeploymentLifecycleResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Error.Equal(that1.Error) {
		return false
	}
	return true
}
func (this *CreateLRPDeploymentRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CreateLRPDeploymentRequest)
	if !ok {
		that2, ok := that.(CreateLRPDeploymentRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Definition.Equal(that1.Definition) {
		return false
	}
	return true
}
func (this *UpdateLRPDeploymentRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UpdateLRPDeploymentRequest)
	if !ok {
		that2, ok := that.(UpdateLRPDeploymentRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if !this.Update.Equal(that1.Update) {
		return false
	}
	return true
}
func (this *RemoveLRPDeploymentRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RemoveLRPDeploymentRequest)
	if !ok {
		that2, ok := that.(RemoveLRPDeploymentRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	return true
}
func (this *ActivateLRPDeploymentDefinitionRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ActivateLRPDeploymentDefinitionRequest)
	if !ok {
		that2, ok := that.(ActivateLRPDeploymentDefinitionRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.DefinitionId != that1.DefinitionId {
		return false
	}
	return true
}
func (this *LRPDeploymentLifecycleResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&models.LRPDeploymentLifecycleResponse{")
	if this.Error != nil {
		s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateLRPDeploymentRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&models.CreateLRPDeploymentRequest{")
	if this.Definition != nil {
		s = append(s, "Definition: "+fmt.Sprintf("%#v", this.Definition)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpdateLRPDeploymentRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.UpdateLRPDeploymentRequest{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	if this.Update != nil {
		s = append(s, "Update: "+fmt.Sprintf("%#v", this.Update)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RemoveLRPDeploymentRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&models.RemoveLRPDeploymentRequest{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ActivateLRPDeploymentDefinitionRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.ActivateLRPDeploymentDefinitionRequest{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "DefinitionId: "+fmt.Sprintf("%#v", this.DefinitionId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLrpDeploymentRequests(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LRPDeploymentLifecycleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LRPDeploymentLifecycleResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *CreateLRPDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateLRPDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Definition != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(m.Definition.Size()))
		n2, err := m.Definition.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *UpdateLRPDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateLRPDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(len(m.Id)))
	i += copy(dAtA[i:], m.Id)
	if m.Update != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(m.Update.Size()))
		n3, err := m.Update.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *RemoveLRPDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveLRPDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(len(m.Id)))
	i += copy(dAtA[i:], m.Id)
	return i, nil
}

func (m *ActivateLRPDeploymentDefinitionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivateLRPDeploymentDefinitionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(len(m.Id)))
	i += copy(dAtA[i:], m.Id)
	dAtA[i] = 0x12
	i++
	i = encodeVarintLrpDeploymentRequests(dAtA, i, uint64(len(m.DefinitionId)))
	i += copy(dAtA[i:], m.DefinitionId)
	return i, nil
}

func encodeFixed64LrpDeploymentRequests(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32LrpDeploymentRequests(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLrpDeploymentRequests(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LRPDeploymentLifecycleResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	}
	return n
}

func (m *CreateLRPDeploymentRequest) Size() (n int) {
	var l int
	_ = l
	if m.Definition != nil {
		l = m.Definition.Size()
		n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	}
	return n
}

func (m *UpdateLRPDeploymentRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	if m.Update != nil {
		l = m.Update.Size()
		n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	}
	return n
}

func (m *RemoveLRPDeploymentRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	return n
}

func (m *ActivateLRPDeploymentDefinitionRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	l = len(m.DefinitionId)
	n += 1 + l + sovLrpDeploymentRequests(uint64(l))
	return n
}

func sovLrpDeploymentRequests(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLrpDeploymentRequests(x uint64) (n int) {
	return sovLrpDeploymentRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LRPDeploymentLifecycleResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LRPDeploymentLifecycleResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateLRPDeploymentRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateLRPDeploymentRequest{`,
		`Definition:` + strings.Replace(fmt.Sprintf("%v", this.Definition), "LRPDeploymentDefinition", "LRPDeploymentDefinition", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdateLRPDeploymentRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateLRPDeploymentRequest{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Update:` + strings.Replace(fmt.Sprintf("%v", this.Update), "LRPDeploymentUpdate", "LRPDeploymentUpdate", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RemoveLRPDeploymentRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RemoveLRPDeploymentRequest{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ActivateLRPDeploymentDefinitionRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ActivateLRPDeploymentDefinitionRequest{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`DefinitionId:` + fmt.Sprintf("%v", this.DefinitionId) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLrpDeploymentRequests(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LRPDeploymentLifecycleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrpDeploymentRequests
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LRPDeploymentLifecycleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LRPDeploymentLifecycleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrpDeploymentRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
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
func (m *CreateLRPDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrpDeploymentRequests
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateLRPDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateLRPDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Definition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Definition == nil {
				m.Definition = &LRPDeploymentDefinition{}
			}
			if err := m.Definition.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrpDeploymentRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
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
func (m *UpdateLRPDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrpDeploymentRequests
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateLRPDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateLRPDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Update", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Update == nil {
				m.Update = &LRPDeploymentUpdate{}
			}
			if err := m.Update.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrpDeploymentRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
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
func (m *RemoveLRPDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrpDeploymentRequests
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveLRPDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveLRPDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrpDeploymentRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
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
func (m *ActivateLRPDeploymentDefinitionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrpDeploymentRequests
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivateLRPDeploymentDefinitionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivateLRPDeploymentDefinitionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefinitionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefinitionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrpDeploymentRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrpDeploymentRequests
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
func skipLrpDeploymentRequests(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLrpDeploymentRequests
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
					return 0, ErrIntOverflowLrpDeploymentRequests
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLrpDeploymentRequests
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLrpDeploymentRequests
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLrpDeploymentRequests
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLrpDeploymentRequests(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLrpDeploymentRequests = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLrpDeploymentRequests   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("lrp_deployment_requests.proto", fileDescriptorLrpDeploymentRequests) }

var fileDescriptorLrpDeploymentRequests = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x33, 0xe5, 0xfb, 0x0a, 0x4e, 0xed, 0x26, 0x74, 0x51, 0x22, 0x4e, 0x25, 0x82, 0x28,
	0x68, 0x0a, 0xf5, 0x01, 0xc4, 0xda, 0x2e, 0x84, 0x2e, 0x64, 0xc0, 0xa5, 0x14, 0x9b, 0xb9, 0x8d,
	0x03, 0x49, 0x26, 0x4e, 0x26, 0x85, 0xee, 0x7c, 0x04, 0x1f, 0xc3, 0x47, 0xe9, 0xb2, 0x4b, 0x57,
	0x62, 0xc7, 0x8d, 0xcb, 0x3e, 0x82, 0x38, 0xe9, 0xbf, 0x88, 0x0a, 0xee, 0xe6, 0xce, 0x3d, 0xe7,
	0x77, 0xcf, 0xc1, 0xbb, 0xa1, 0x4c, 0xfa, 0x0c, 0x92, 0x50, 0x8c, 0x23, 0x88, 0x55, 0x5f, 0xc2,
	0x7d, 0x06, 0xa9, 0x4a, 0xbd, 0x44, 0x0a, 0x25, 0xec, 0x72, 0x24, 0x18, 0x84, 0xa9, 0x73, 0x12,
	0x70, 0x75, 0x97, 0x0d, 0x3c, 0x5f, 0x44, 0xcd, 0x40, 0x04, 0xa2, 0x69, 0xd6, 0x83, 0x6c, 0x68,
	0x26, 0x33, 0x98, 0x57, 0x6e, 0x73, 0x6a, 0x45, 0xea, 0xe2, 0xb7, 0x02, 0x52, 0x0a, 0x99, 0x0f,
	0x6e, 0x17, 0x93, 0x1e, 0xbd, 0xea, 0xac, 0x34, 0x3d, 0x3e, 0x04, 0x7f, 0xec, 0x87, 0x40, 0x21,
	0x4d, 0x44, 0x9c, 0x82, 0xbd, 0x8f, 0xff, 0x1b, 0x43, 0x1d, 0xed, 0xa1, 0xc3, 0x4a, 0xab, 0xea,
	0xe5, 0x59, 0xbc, 0xee, 0xe7, 0x27, 0xcd, 0x77, 0xee, 0x0d, 0x76, 0x2e, 0x24, 0xdc, 0x2a, 0x28,
	0xc0, 0x68, 0xde, 0xc2, 0x3e, 0xc3, 0x98, 0xc1, 0x90, 0xc7, 0x5c, 0x71, 0x11, 0x2f, 0x38, 0x8d,
	0x25, 0xa7, 0xe0, 0xe8, 0xac, 0x64, 0x74, 0xc3, 0xe2, 0x06, 0xd8, 0xb9, 0x4e, 0xd8, 0x4f, 0xf8,
	0x1a, 0x2e, 0x71, 0x66, 0xb0, 0x5b, 0xed, 0x7f, 0x93, 0x97, 0x86, 0x45, 0x4b, 0x9c, 0xd9, 0xa7,
	0xb8, 0x9c, 0x19, 0x4f, 0xbd, 0x64, 0x0e, 0xee, 0x7c, 0x7b, 0x30, 0xc7, 0xd2, 0x85, 0xd4, 0x6d,
	0x61, 0x87, 0x42, 0x24, 0x46, 0x7f, 0x38, 0xe4, 0x72, 0x7c, 0x70, 0xee, 0x2b, 0x3e, 0xfa, 0x1a,
	0x6f, 0xa3, 0xcb, 0xaf, 0x41, 0x8f, 0x70, 0x75, 0x5d, 0xb5, 0xcf, 0x99, 0xc9, 0xbb, 0x14, 0x6c,
	0xaf, 0x57, 0x97, 0xac, 0x7d, 0x3c, 0x9d, 0x11, 0xeb, 0x79, 0x46, 0xac, 0xf9, 0x8c, 0xa0, 0x07,
	0x4d, 0xd0, 0x93, 0x26, 0x68, 0xa2, 0x09, 0x9a, 0x6a, 0x82, 0x5e, 0x35, 0x41, 0xef, 0x9a, 0x58,
	0x73, 0x4d, 0xd0, 0xe3, 0x1b, 0xb1, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x07, 0x27, 0x26, 0x69,
	0x55, 0x02, 0x00, 0x00,
}
