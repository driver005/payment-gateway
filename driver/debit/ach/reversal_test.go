package ach

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFileCreditReversal(t *testing.T) {
	file, err := ReadJSONFile(filepath.Join("test", "testdata", "ppd-valid.json"))
	require.NoError(t, err)

	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 := file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries := b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, CheckingDebit, entries[0].TransactionCode)
}

func TestFileDebitReversal(t *testing.T) {
	file, err := ReadFile(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 := file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries := b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, CheckingCredit, entries[0].TransactionCode)
}
