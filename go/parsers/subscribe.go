package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-capnproto/go/gen"
	"github.com/xconnio/wampproto-go/messages"
)

type Subscribe struct {
	gen *gen.Subscribe
}

func NewSubscribeFields(g *gen.Subscribe) messages.SubscribeFields {
	return &Subscribe{gen: g}
}

func (s *Subscribe) RequestID() uint64 {
	return s.gen.RequestID()
}

func (s *Subscribe) Options() map[string]any {
	options := make(map[string]any)
	if s.gen.Match().String() != "" {
		setDetail(&options, "match", s.gen.Match().String())
	}
	return options
}

func (s *Subscribe) Topic() string {
	topic, _ := s.gen.Topic()
	return topic
}

func SubscribeToCapnproto(m *messages.Subscribe) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	subscribe, err := gen.NewRootSubscribe(seg)
	if err != nil {
		return nil, err
	}

	subscribe.SetRequestID(m.RequestID())
	if err := subscribe.SetTopic(m.Topic()); err != nil {
		return nil, err
	}

	matchString, ok := m.Options()["match"].(string)
	if ok {
		subscribe.SetMatch(gen.Subscribe_MatchFromString(matchString))
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeSubscribe, data, nil), nil
}

func CapnprotoToSubscribe(data []byte) (*messages.Subscribe, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	subscribe, err := gen.ReadRootSubscribe(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewSubscribeWithFields(NewSubscribeFields(&subscribe)), nil
}
