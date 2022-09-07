package libdrawio

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

type Diagram struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type MxFile struct {
	XMLName  xml.Name  `xml:"mxfile"`
	Host     string    `xml:"host,attr"`
	Modified string    `xml:"modified,attr"`
	Agent    string    `xml:"agent,attr"`
	Version  string    `xml:"version,attr"`
	Type     string    `xml:"type,attr"`
	Etag     string    `xml:"etag,attr"`
	Diagram  []Diagram `xml:"diagram"`
}

// Load drawio file, then map to MxFile struct.
func NewMxFile(filepath string) (*MxFile, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	mxfile := new(MxFile)
	if err = xml.Unmarshal(b, &mxfile); err != nil {
		return nil, err
	}

	return mxfile, nil
}

// Write MxFile.
func (m *MxFile) Write(filepath string) error {
	b, err := xml.Marshal(m)
	if err != nil {
		return err
	}

	xmlStr := strings.ReplaceAll(string(b), "&#34;", "\"")
	err = ioutil.WriteFile(filepath, []byte(xmlStr), 0664)
	if err != nil {
		return err
	}

	return nil
}
