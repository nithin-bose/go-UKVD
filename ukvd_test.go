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

func TestVDICheckBasicItemNotFound(t *testing.T) {
	_, err := vd.VDICheckBasic("KLM123")
	if err != ItemNotFoundError {
		t.Error(err)
	}
}

func TestVDICheckFinanceItemNotFound(t *testing.T) {
	_, err := vd.VDICheckFinance("KLM123")
	if err != ItemNotFoundError {
		t.Error(err)
	}
}

func TestVDICheckFullItemNotFound(t *testing.T) {
	_, err := vd.VDICheckFull("KLM123")
	if err != ItemNotFoundError {
		t.Error(err)
	}
}

func TestVDICheckBasic(t *testing.T) {
	_, err := vd.VDICheckBasic("KM12AKK")
	if err != nil {
		t.Error(err)
	}
}

func TestVDICheckFinance(t *testing.T) {
	_, err := vd.VDICheckFinance("KM12AKK")
	if err != nil {
		t.Error(err)
	}
}

func TestVDICheckFull(t *testing.T) {
	_, err := vd.VDICheckFull("KM12AKK")
	if err != nil {
		t.Error(err)
	}
}
