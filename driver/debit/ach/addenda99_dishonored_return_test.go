package ach

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func mockAddenda99Dishonored() *Addenda99Dishonored {
	addenda99 := NewAddenda99Dishonored()
	addenda99.DishonoredReturnReasonCode = "R68"
	addenda99.OriginalEntryTraceNumber = "059999990000301"
	addenda99.OriginalReceivingDFIIdentification = "12391871"
	addenda99.ReturnTraceNumber = "123918710000001"
	addenda99.ReturnSettlementDate = "179"
	addenda99.ReturnReasonCode = "01"
	addenda99.AddendaInformation = "Untimely Return"
	addenda99.TraceNumber = "059999990000001"
	return addenda99
}

func TestAddenda99Dishonored__Fields(t *testing.T) {
	addenda99 := mockAddenda99Dishonored()

	// shorten some fields
	addenda99.OriginalEntryTraceNumber = "0599999900301"
	addenda99.ReturnTraceNumber = "123918710001"
	addenda99.TraceNumber = "05999900001"

	require.Equal(t, "R68", addenda99.DishonoredReturnReasonCodeField())
	require.Equal(t, "000599999900301", addenda99.OriginalEntryTraceNumberField())
	require.Equal(t, "12391871", addenda99.OriginalReceivingDFIIdentificationField())
	require.Equal(t, "000123918710001", addenda99.ReturnTraceNumberField())
	require.Equal(t, "179", addenda99.ReturnSettlementDateField())
	require.Equal(t, "01", addenda99.ReturnReasonCodeField())
	require.Equal(t, "Untimely Return      ", addenda99.AddendaInformationField())
	require.Equal(t, "000005999900001", addenda99.TraceNumberField())
}
