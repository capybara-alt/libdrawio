package libdrawio_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/capybara-alt/libdrawio/libdrawio"
)

func TestNewMxLibrary(t *testing.T) {
	tests := libdrawio.NewMxLibrary()
	if tests.Value != "" {
		t.Fail()
	}
}

func TestMakeMxLibrary(t *testing.T) {
	tests := libdrawio.NewMxLibrary()
	mxGraphModels := []libdrawio.MxGraphModel{
		{
			Dx:         "1106",
			Dy:         "818",
			Grid:       "1",
			GridSize:   "10",
			Guides:     "1",
			Tooltips:   "1",
			Connect:    "1",
			Arrows:     "1",
			Fold:       "1",
			Page:       "1",
			PageScale:  "1",
			PageWidth:  "827",
			PageHeight: "1169",
			Math:       "0",
			Shadow:     "0",
			Content: libdrawio.Root{
				MxCells: []libdrawio.MxCell{
					{
						Id: "0",
					},
					{
						Id:     "1",
						Parent: "0",
					},
					{
						Id:     "eCHvXLgnRnpgWLnYiHm2-1",
						Style:  "",
						Parent: "1",
						Vertex: "1",
					},
				},
			},
		},
		{
			Dx:         "1106",
			Dy:         "818",
			Grid:       "1",
			GridSize:   "10",
			Guides:     "1",
			Tooltips:   "1",
			Connect:    "1",
			Arrows:     "1",
			Fold:       "1",
			Page:       "1",
			PageScale:  "1",
			PageWidth:  "827",
			PageHeight: "1169",
			Math:       "0",
			Shadow:     "0",
			Content: libdrawio.Root{
				MxCells: []libdrawio.MxCell{
					{
						Id: "0",
					},
					{
						Id:     "1",
						Parent: "0",
					},
					{
						Id:     "eCHvXLgnRnpgWLnYiHm2-2",
						Style:  "",
						Parent: "1",
						Vertex: "1",
					},
				},
			},
		},
	}

	mxlibObjs := make([]libdrawio.MxLibObject, len(mxGraphModels))
	for index, mxGraphModel := range mxGraphModels {
		mxlibObj, err := tests.MakeMxLibObj(&mxGraphModel, "test", 24, 24)
		if err != nil {
			t.Fail()
		}
		mxlibObjs[index] = *mxlibObj
	}

	if err := tests.MakeMxLibrary(mxlibObjs); err != nil {
		t.Fail()
	}

	result := []libdrawio.MxLibObject{}
	err := json.Unmarshal([]byte(tests.Value), &result)
	if err != nil {
		t.Fail()
	}

	mxGraphModels = []libdrawio.MxGraphModel{}
	for _, mxlibObj := range result {
		mxGraphModel, err := libdrawio.Decompress(mxlibObj.Xml)
		if err != nil {
			t.Fail()
		}
		mxGraphModels = append(mxGraphModels, *mxGraphModel)
	}

	if mxGraphModels[0].Content.MxCells[0].Id != "0" {
		t.Fail()
	}

	if mxGraphModels[0].Content.MxCells[1].Id != "1" {
		t.Fail()
	}

	if mxGraphModels[0].Content.MxCells[2].Id != "eCHvXLgnRnpgWLnYiHm2-1" {
		t.Fail()
	}

	if mxGraphModels[1].Content.MxCells[0].Id != "0" {
		t.Fail()
	}

	if mxGraphModels[1].Content.MxCells[1].Id != "1" {
		t.Fail()
	}

	if mxGraphModels[1].Content.MxCells[2].Id != "eCHvXLgnRnpgWLnYiHm2-2" {
		t.Fail()
	}
}

func TestMakeMxLibObjMxGraphModelNil(t *testing.T) {
	tests := libdrawio.NewMxLibrary()
	mxlibObj, err := tests.MakeMxLibObj(nil, "test", 25, 25)
	if err != nil {
		t.Fail()
	}

	if mxlibObj.Title != "test" {
		t.Fail()
	}

	if mxlibObj.H != 25 {
		t.Fail()
	}

	if mxlibObj.W != 25 {
		t.Fail()
	}

	mxGraphModel := &libdrawio.MxGraphModel{}
	mxGraphModel = nil
	compressed, _ := mxGraphModel.Compress()
	if mxlibObj.Xml != compressed  {
		t.Fail()
	}
}

func TestWriteMxLibrary(t *testing.T) {
	tests := libdrawio.NewMxLibrary()
	err := tests.Write("./ut_mxlibrary.xml")
	if err != nil {
		t.Fail()
	}

	if _, err = os.Open("./ut_mxlibrary.xml"); err != nil {
		t.Fail()
	}

	os.Remove("./ut_mxlibrary.xml")
}

