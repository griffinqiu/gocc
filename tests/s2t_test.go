package opencc

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/longbridge/opencc"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func readFile(filename string) string {
	pwd, _ := os.Getwd()
	data, err := os.ReadFile(path.Join(pwd, "fixtures", filename))
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func Test_s2t(t *testing.T) {
	raw := readFile("html-raw.txt")
	expected := readFile("html-s2t.txt")

	cc, _ := opencc.New("s2t")
	out, _ := cc.Convert(raw)

	fmt.Println(out)

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(expected, out, true)

		t.Errorf(dmp.DiffPrettyText(diffs))
	}

}

func TestFinance_s2hk_finance(t *testing.T) {
	raw := readFile("html-raw.txt")
	expected := readFile("html-s2hk-finance.txt")

	cc, _ := opencc.New("s2hk-finance")
	out, _ := cc.Convert(raw)

	fmt.Println(out)

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(expected, out, true)

		t.Errorf(dmp.DiffPrettyText(diffs))
	}

}
