package libdrawio

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

type MxLibrary struct {
	XMLName       xml.Name `xml:"mxlibrary"`
	Value         string   `xml:",chardata"`
	MxGraphModels []MxGraphModel
}

type MxLibObject struct {
	Xml   string `json:"xml"`
	W     int    `json:"w"`
	H     int    `json:"h"`
	Title string `json:"title"`
}

func NewMxLibrary() *MxLibrary {
	return new(MxLibrary)
}

// Make MxLibrary string from MxLibrary.MxGraphModels, then set to MxLibrary.Value.
func (m *MxLibrary) MakeMxLibrary(mxlibObjects []MxLibObject) error {
	for index, mxlibObj := range mxlibObjects {
		mxlibObjects[index] = mxlibObj
	}

	b, err := json.Marshal(mxlibObjects)
	if err != nil {
		return err
	}

	m.Value = string(b)
	return nil
}

// Make MxLibObj from MxGraphModel, title, height, width
func (m *MxLibrary) MakeMxLibObj(mxGraphModel *MxGraphModel, title string, height, width int) (*MxLibObject, error) {
	result, err := mxGraphModel.Compress()
	if err != nil {
		return nil, err
	}

	mxlibObj := new(MxLibObject)
	mxlibObj.Xml = result
	mxlibObj.H = height
	mxlibObj.W = width
	mxlibObj.Title = title

	return mxlibObj, nil
}

// Write MxLibrary to xml file.
func (m *MxLibrary) Write(filepath string) error {
	file, err := xml.Marshal(m)
	if err != nil {
		return err
	}

	xmlStr := strings.ReplaceAll(string(file), "&#34;", "\"")
	err = ioutil.WriteFile(filepath, []byte(xmlStr), 0644)

	return nil
}
