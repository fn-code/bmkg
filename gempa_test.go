package bmkg

import (
	"fmt"
	"net/http"
	"testing"
)

var GempaTest = []struct {
	input    string
	expected *Infogempa
	err      error
}{
	{BaseUrl + GempaTerkiniUrl, &Infogempa{}, nil},
	{"", nil, fmt.Errorf("error http bad request: %v", http.StatusBadRequest)},
}

func TestGempaTerkini(t *testing.T) {
	for _, tb := range GempaTest {
		info, err := GempaTerkini(tb.input)

		if info != tb.expected && err != tb.err {
			t.Errorf("GempaTerkini(): expected %v-%v, Get %v-%v ", tb.expected, tb.err, info, err)
		}

	}

}
