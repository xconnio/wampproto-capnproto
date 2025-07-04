package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Published struct {
	gen *gen.Published
}

func NewPublishedFields(gen *gen.Published) messages.PublishedFields {
	return &Published{gen: gen}
}

func (p *Published) RequestID() int64 {
	return p.gen.RequestID()
}

func (p *Published) PublicationID() int64 {
	return p.gen.PublicationID()
}

func PublishedToCapnproto(published *messages.Published) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	pubed, err := gen.NewPublished(seg)
	if err != nil {
		return nil, err
	}

	pubed.SetRequestID(published.RequestID())
	pubed.SetPublicationID(published.PublicationID())

	var data bytes.Buffer
	if err = capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	byteValue := byte(messages.MessageTypePublished & 0xFF)
	return append([]byte{byteValue}, data.Bytes()...), nil
}

func CapnprotoToPublished(data []byte) (*messages.Published, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	published, err := gen.ReadRootPublished(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewPublishedWithFields(NewPublishedFields(&published)), nil
}
