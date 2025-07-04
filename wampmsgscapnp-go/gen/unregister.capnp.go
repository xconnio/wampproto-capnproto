// Code generated by capnpc-go. DO NOT EDIT.

package gen

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
)

type Unregister capnp.Struct

// Unregister_TypeID is the unique identifier for the type Unregister.
const Unregister_TypeID = 0xc8aba2ad7f35b274

func NewUnregister(s *capnp.Segment) (Unregister, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Unregister(st), err
}

func NewRootUnregister(s *capnp.Segment) (Unregister, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Unregister(st), err
}

func ReadRootUnregister(msg *capnp.Message) (Unregister, error) {
	root, err := msg.Root()
	return Unregister(root.Struct()), err
}

func (s Unregister) String() string {
	str, _ := text.Marshal(0xc8aba2ad7f35b274, capnp.Struct(s))
	return str
}

func (s Unregister) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Unregister) DecodeFromPtr(p capnp.Ptr) Unregister {
	return Unregister(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Unregister) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Unregister) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Unregister) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Unregister) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Unregister) RequestID() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Unregister) SetRequestID(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

func (s Unregister) RegistrationID() int64 {
	return int64(capnp.Struct(s).Uint64(8))
}

func (s Unregister) SetRegistrationID(v int64) {
	capnp.Struct(s).SetUint64(8, uint64(v))
}

// Unregister_List is a list of Unregister.
type Unregister_List = capnp.StructList[Unregister]

// NewUnregister creates a new list of Unregister.
func NewUnregister_List(s *capnp.Segment, sz int32) (Unregister_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return capnp.StructList[Unregister](l), err
}

// Unregister_Future is a wrapper for a Unregister promised by a client call.
type Unregister_Future struct{ *capnp.Future }

func (f Unregister_Future) Struct() (Unregister, error) {
	p, err := f.Future.Ptr()
	return Unregister(p.Struct()), err
}
