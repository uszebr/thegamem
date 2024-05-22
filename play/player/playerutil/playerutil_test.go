package playerutil

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateName(t *testing.T) {
	attemts := 1000
	for range attemts {
		name := GenerateRandomName()
		//fmt.Println(name)
		if len(name) < 6 {
			t.Errorf("name too short: %v", len(name))
		}
		postfix := name[len(name)-3:]
		_, err := strconv.Atoi(postfix)
		if err != nil {
			t.Errorf("last 3 chars are not digits")
		}
		if !strings.Contains(name, "-") {
			t.Errorf("name without - %v", name)
		}
	}
}
