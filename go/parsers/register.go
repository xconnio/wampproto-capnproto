package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-capnproto/go/gen"
	"github.com/xconnio/wampproto-go/messages"
)

type Register struct {
	gen *gen.Register
}

func NewRegisterFields(g *gen.Register) messages.RegisterFields {
	return &Register{gen: g}
}

func (r *Register) RequestID() uint64 {
	return r.gen.RequestID()
}

func (r *Register) Options() map[string]any {
	options := make(map[string]any)
	if r.gen.Match().String() != "" {
		setDetail(&options, "match", r.gen.Match().String())
	}
	if r.gen.Invoke().String() != "" {
		setDetail(&options, "invoke", r.gen.Invoke().String())
	}
	return options
}

func (r *Register) Procedure() string {
	proc, _ := r.gen.Procedure()
	return proc
}

func RegisterToCapnproto(m *messages.Register) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	register, err := gen.NewRootRegister(seg)
	if err != nil {
		return nil, err
	}

	register.SetRequestID(m.RequestID())

	if err = register.SetProcedure(m.Procedure()); err != nil {
		return nil, err
	}

	matchString, ok := m.Options()["match"].(string)
	if ok {
		register.SetMatch(gen.Register_MatchFromString(matchString))
	}

	invokeString, ok := m.Options()["invoke"].(string)
	if ok {
		register.SetInvoke(gen.Register_InvokeFromString(invokeString))
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeRegister, data, nil), nil
}

func CapnprotoToRegister(data []byte) (*messages.Register, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	register, err := gen.ReadRootRegister(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewRegisterWithFields(NewRegisterFields(&register)), nil
}
