// Code generated by capnpc-go. DO NOT EDIT.

package gen

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
)

type Unsubscribed capnp.Struct

// Unsubscribed_TypeID is the unique identifier for the type Unsubscribed.
const Unsubscribed_TypeID = 0xf4a36356cc57469e

func NewUnsubscribed(s *capnp.Segment) (Unsubscribed, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Unsubscribed(st), err
}

func NewRootUnsubscribed(s *capnp.Segment) (Unsubscribed, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Unsubscribed(st), err
}

func ReadRootUnsubscribed(msg *capnp.Message) (Unsubscribed, error) {
	root, err := msg.Root()
	return Unsubscribed(root.Struct()), err
}

func (s Unsubscribed) String() string {
	str, _ := text.Marshal(0xf4a36356cc57469e, capnp.Struct(s))
	return str
}

func (s Unsubscribed) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Unsubscribed) DecodeFromPtr(p capnp.Ptr) Unsubscribed {
	return Unsubscribed(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Unsubscribed) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Unsubscribed) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Unsubscribed) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Unsubscribed) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Unsubscribed) RequestID() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Unsubscribed) SetRequestID(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

// Unsubscribed_List is a list of Unsubscribed.
type Unsubscribed_List = capnp.StructList[Unsubscribed]

// NewUnsubscribed creates a new list of Unsubscribed.
func NewUnsubscribed_List(s *capnp.Segment, sz int32) (Unsubscribed_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[Unsubscribed](l), err
}

// Unsubscribed_Future is a wrapper for a Unsubscribed promised by a client call.
type Unsubscribed_Future struct{ *capnp.Future }

func (f Unsubscribed_Future) Struct() (Unsubscribed, error) {
	p, err := f.Future.Ptr()
	return Unsubscribed(p.Struct()), err
}
