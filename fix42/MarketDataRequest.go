package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type MarketDataRequest struct {
	message.Message
}

func (m *MarketDataRequest) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MarketDepth() (*field.MarketDepth, error) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MDUpdateType() (*field.MDUpdateType, error) {
	f := new(field.MDUpdateType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) AggregatedBook() (*field.AggregatedBook, error) {
	f := new(field.AggregatedBook)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) NoMDEntryTypes() (*field.NoMDEntryTypes, error) {
	f := new(field.NoMDEntryTypes)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
