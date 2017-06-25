package ukvd

import (
	"os"
	"testing"
)

var vd *UKVD

func init() {
	vd = NewUKVD(os.Getenv("UKVD_API_KEY"))
	vd.SetLog(true)
}

func TestVDICheckBasicInvalidRegistration(t *testing.T) {
	_, err := vd.VDICheckBasic("KLM123")
	if err != KeyInvalidError {
		t.Error(err)
	}
}

func TestVDICheckFinanceInvalidRegistration(t *testing.T) {
	_, err := vd.VDICheckFinance("KLM123")
	if err != KeyInvalidError {
		t.Error(err)
	}
}

func TestVDICheckFullInvalidRegistration(t *testing.T) {
	_, err := vd.VDICheckFull("KLM123")
	if err != KeyInvalidError {
		t.Error(err)
	}
}
