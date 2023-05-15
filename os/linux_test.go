package os

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMe(t *testing.T) {

	var (
		layout     = "2006-01-02 15:04:05.999999999 -0700"
		inputUTC   = "2022-03-28 10:01:08.816922127 +0000"
		inputLocal = "2022-03-28 12:01:08.816922127 +0200"
		expected   = int64(1648461668)
		expectedNs = int(816922127)
	)

	statted, err := os.Stat("/home/twieczorek/Repos/rig/connection.go")
	require.NoError(t, err)
	t.Log(statted.ModTime().GoString())
	assert.Equal(t, expected, statted.ModTime().Unix())
	assert.Equal(t, expectedNs, statted.ModTime().Nanosecond())

	pUTC, err := time.Parse(layout, inputUTC)
	require.NoError(t, err)
	t.Log(pUTC.GoString())

	pLocal, err := time.Parse(layout, inputLocal)
	require.NoError(t, err)
	t.Log(pLocal.GoString())

	assert.Equal(t, expected, pUTC.Unix())
	assert.Equal(t, expected, pLocal.Unix())
	assert.Equal(t, expectedNs, pUTC.Nanosecond())
	assert.Equal(t, statted.ModTime(), pLocal)
	assert.Equal(t, pUTC.UTC(), pLocal.UTC())
	assert.Equal(t, expectedNs, pUTC.UTC().Nanosecond())
	// t.Fail()
}

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
