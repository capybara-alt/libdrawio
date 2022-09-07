package libdrawio_test

import (
	"os"
	"testing"

	"github.com/capybara-alt/libdrawio"
)

func TestNewMxFile(t *testing.T) {
	mxfile, err := libdrawio.NewMxFile("./mxfile_test_1.drawio")
	if err != nil {
		t.Fail()
	}

	if mxfile.Etag != "gzH0kVSIBYVuorkLmoRh" {
		t.Fail()
	}
}

func TestNewMxFileNotFound(t *testing.T) {
	mxfile, err := libdrawio.NewMxFile("./notfound.drawio")
	if err == nil {
		t.Fail()
	}

	if mxfile != nil {
		t.Fail()
	}
}

func TestNewMxFileFilePathNil(t *testing.T) {
	mxfile, err := libdrawio.NewMxFile("")
	if err == nil {
		t.Fail()
	}

	if mxfile != nil {
		t.Fail()
	}
}

func TestWrite(t *testing.T) {
	mxfile, err := libdrawio.NewMxFile("./mxfile_test_1.drawio")
	if err != nil {
		t.Fail()
	}

	if err = mxfile.Write("./ut_result.drawio"); err != nil {
		t.Fail()
	}

	_, err = os.Open("./ut_result.drawio")
	if err != nil {
		t.Fail()
	}

	os.Remove("./ut_result.drawio")
}
