package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type CollateralAssignment struct {
	message.Message
}

func (m *CollateralAssignment) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CollAsgnReason() (*field.CollAsgnReason, error) {
	f := new(field.CollAsgnReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) TotalNetValue() (*field.TotalNetValue, error) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CollAsgnRefID() (*field.CollAsgnRefID, error) {
	f := new(field.CollAsgnRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoTrades() (*field.NoTrades, error) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SettlDeliveryType() (*field.SettlDeliveryType, error) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CashOutstanding() (*field.CashOutstanding, error) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CollAsgnTransType() (*field.CollAsgnTransType, error) {
	f := new(field.CollAsgnTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MarginExcess() (*field.MarginExcess, error) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoMiscFees() (*field.NoMiscFees, error) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CollReqID() (*field.CollReqID, error) {
	f := new(field.CollReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CollAsgnID() (*field.CollAsgnID, error) {
	f := new(field.CollAsgnID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoDlvyInst() (*field.NoDlvyInst, error) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralAssignment) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
