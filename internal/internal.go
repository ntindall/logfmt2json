package internal

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/kr/logfmt"
)

type handler map[string]string

var _ logfmt.Handler = &handler{}

func (h *handler) HandleLogfmt(k, val []byte) error {
	keyStr, valStr := string(k), string(val)
	(*h)[keyStr] = valStr
	return nil
}

// Convert reads logfmt messages from the scanner passed as a parameter and
// prints json output to stdout.
func Convert(s *bufio.Scanner) error {
	for s.Scan() {
		h := handler{}
		if err := logfmt.Unmarshal(s.Bytes(), &h); err != nil {
			return err
		}
		bytes, err := json.Marshal(h)
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
	}

	return s.Err()
}
