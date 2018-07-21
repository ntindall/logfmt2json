package internal

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/kr/logfmt"
)

type handler map[string]string

var _ logfmt.Handler = &handler{}

func (h *handler) HandleLogfmt(k, val []byte) error {
	keyStr, valStr := string(k), string(val)
	(*h)[keyStr] = valStr
	return nil
}

// Logfmt2JSON reads logfmt messages from the reader passed as a parameter and
// writes json passed to the writer
func Logfmt2JSON(reader io.Reader, w io.Writer) error {
	s := bufio.NewScanner(reader)
	writer := bufio.NewWriter(w)

	// Parse standard in, using newlines to delineate each message
	for s.Scan() {
		h := handler{}
		if err := logfmt.Unmarshal(s.Bytes(), &h); err != nil {
			return err
		}

		bytes, err := json.Marshal(h)
		if err != nil {
			return err
		}

		if _, err := writer.WriteString(string(bytes)); err != nil {
			// For now, we fail out hard
			return err
		}

		if _, err := writer.WriteRune('\n'); err != nil {
			return err
		}

		if err := writer.Flush(); err != nil {
			return err
		}
	}

	return s.Err()
}
