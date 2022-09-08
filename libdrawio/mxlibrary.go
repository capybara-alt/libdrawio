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

// Make MxLibrary string, then set to MxLibrary.Value.
// Implement mxlibObjectSetterFunc to set custom field value (except MxLibObject.Xml) to MxLibObject.
func (m *MxLibrary) MakeMxLibrary(mxlibObjSetterFunc func(m *MxLibObject) *MxLibObject) error {
	mxlibObjects := make([]MxLibObject, len(m.MxGraphModels))
	for index, mxGraphModel := range m.MxGraphModels {
		result, err := mxGraphModel.Compress()
		if err != nil {
			return err
		}
		mxlibObj := new(MxLibObject)
		mxlibObj.Xml = result
		mxlibObjects[index] = *mxlibObjSetterFunc(mxlibObj)
	}

	b, err := json.Marshal(mxlibObjects)
	if err != nil {
		return err
	}

	m.Value = string(b)
	return nil
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
