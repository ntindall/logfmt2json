package internal

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_HandleLogfmt(t *testing.T) {
	h := &handler{}
	err := h.HandleLogfmt([]byte("foo"), []byte("bar"))
	require.NoError(t, err, "had unexpected error")

	assert.EqualValues(t, map[string]string{"foo": "bar"}, *h)
}

func TestLogfmt2JSON(t *testing.T) {
	testCases := []struct {
		desc               string
		in                 []string
		expectedOut        []string
		expectedErrMessage string
	}{
		{
			desc: "works (one line)",
			in: []string{
				`hello=world`,
			},
			expectedOut: []string{
				`{"hello":"world"}`,
			},
		},
		{
			desc: "works (multiple lines)",
			in: []string{
				`hello=world`,
				`peace="on earth"`,
				`foo=bar"`,
			},
			expectedOut: []string{
				`{"hello":"world"}`,
				`{"peace":"on earth"}`,
				`{"foo":"bar"}`,
			},
		},
		{
			desc: "works (multiple keys)",
			in: []string{
				`timestamp=2018-07-20T22:11:24.932Z message="hello world" dreams="reality"`,
			},
			expectedOut: []string{
				`{"dreams":"reality","message":"hello world","timestamp":"2018-07-20T22:11:24.932Z"}`,
			},
		},
		{
			desc: "treats extraneous newlines as an empty object",
			in: []string{
				"\n",
			},
			expectedOut: []string{
				`{}`,
				`{}`,
			},
		},
		{
			desc: "empty strings and spaces are the empty object",
			in: []string{
				"",
				" \t",
			},
			expectedOut: []string{
				`{}`,
				`{}`,
			},
		},
		{
			desc: "propagates invalid logfmt error messages",
			in: []string{
				`foo="`,
			},
			expectedErrMessage: "logfmt: unterminated string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			inBuffer := bytes.NewBufferString(strings.Join(tc.in, "\n"))
			outBuffer := &bytes.Buffer{}

			err := Logfmt2JSON(inBuffer, outBuffer)
			if tc.expectedErrMessage != "" {
				require.EqualError(t, err, tc.expectedErrMessage, "err message did not match expectation")
				return
			}

			s := bufio.NewScanner(outBuffer)
			idx := 0
			for s.Scan() {
				assert.Equal(t, tc.expectedOut[idx], s.Text(), "marshalled output did not meet expectation")
				idx++
			}

			require.NoError(t, s.Err(), "unexpected scanner error")
		})
	}

}
