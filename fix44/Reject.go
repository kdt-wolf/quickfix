package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type Reject struct {
	message.Message
}

func (m *Reject) RefSeqNum() (*field.RefSeqNum, error) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) RefTagID() (*field.RefTagID, error) {
	f := new(field.RefTagID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) RefMsgType() (*field.RefMsgType, error) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) SessionRejectReason() (*field.SessionRejectReason, error) {
	f := new(field.SessionRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
