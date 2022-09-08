package libdrawio

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	"io"
	"net/url"
)

type MxPoint struct {
	X  string `xml:"x,attr"`
	Y  string `xml:"y,attr"`
	As string `xml:"as,attr"`
}

type MxGeometry struct {
	X        string   `xml:"x,attr"`
	Y        string   `xml:"y,attr"`
	Width    string   `xml:"width,attr"`
	Height   string   `xml:"height,attr"`
	As       string   `xml:"as,attr"`
	Relative string   `xml:"relative,attr,omitempty"`
	MxPoint  *MxPoint `xml:"mxPoint"`
}

type MxCell struct {
	Id     string      `xml:"id,attr,omitempty"`
	Value  *string     `xml:"value,attr,omitempty"`
	Style  string      `xml:"style,attr,omitempty"`
	Parent string      `xml:"parent,attr,omitempty"`
	Source string      `xml:"source,attr,omitempty"`
	Target string      `xml:"target,attr,omitempty"`
	Edge   string      `xml:"edge,attr,omitempty"`
	Vertex string      `xml:"vertex,attr,omitempty"`
	Geo    *MxGeometry `xml:"mxGeometry"`
}

type Object struct {
	Id        string `xml:"id,attr,omitempty"`
	Lable     string `xml:"label,attr"`
	Component string `xml:"component,attr"`
	MxCell    MxCell `xml:"mxCell"`
}

type Root struct {
	Object  []Object `xml:"object"`
	MxCells []MxCell `xml:"mxCell"`
}

type MxGraphModel struct {
	XMLName    xml.Name `xml:"mxGraphModel"`
	Content    Root     `xml:"root"`
	Dx         string   `xml:"dx,attr"`
	Dy         string   `xml:"dy,attr"`
	Grid       string   `xml:"grid,attr"`
	GridSize   string   `xml:"gridSize,attr"`
	Guides     string   `xml:"guides,attr"`
	Tooltips   string   `xml:"tooltips,attr"`
	Connect    string   `xml:"connect,attr"`
	Arrows     string   `xml:"arrows,attr"`
	Fold       string   `xml:"fold,attr"`
	Page       string   `xml:"page,attr"`
	PageScale  string   `xml:"pageScale,attr"`
	PageWidth  string   `xml:"pageWidth,attr"`
	PageHeight string   `xml:"pageHeight,attr"`
	Math       string   `xml:"math,attr"`
	Shadow     string   `xml:"shadow,attr"`
	Title      *string
}

// Decompress compressed MxGraphModel string.
// Map to MxGraphModel struct after decompress string.
func Decompress(compressedStr string) (*MxGraphModel, error) {
	decoded, err := base64.StdEncoding.DecodeString(compressedStr)
	if err != nil {
		return nil, err
	}

	r := flate.NewReader(bytes.NewBuffer(decoded))
	defer r.Close()

	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, r); err != nil {
		return nil, err
	}

	unescaped, err := url.PathUnescape(buff.String())
	if err != nil {
		return nil, err
	}

	mxGraphModel := new(MxGraphModel)
	if err = xml.Unmarshal([]byte(unescaped), &mxGraphModel); err != nil {
		return nil, err
	}

	return mxGraphModel, nil
}

// Compress MxGraphModel string.
func (m *MxGraphModel) Compress() (string, error) {
	b, err := xml.Marshal(m)
	if err != nil {
		return "", err
	}

	pathEscaped := url.PathEscape(string(b))
	buffer := new(bytes.Buffer)
	writer, err := flate.NewWriter(buffer, flate.BestCompression)
	if err != nil {
		return "", err
	}

	if _, err := writer.Write([]byte(pathEscaped)); err != nil {
		return "", err
	}

	if err = writer.Flush(); err != nil {
		return "", err
	}
	writer.Close()

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// Get Object and MxCell ID list.
func (m *MxGraphModel) GetIds() []string {
	ids := []string{}
	if m.Content.MxCells != nil {
		for _, mxcell := range m.Content.MxCells {
			ids = append(ids, mxcell.Id)
		}
	}

	if m.Content.Object != nil {
		for _, obj := range m.Content.Object {
			ids = append(ids, obj.Id)
		}
	}

	return ids
}

// Find Object from MxGraphModel.Object by ID.
func (m *MxGraphModel) FindObject(id string) *Object {
	for _, obj := range m.Content.Object {
		if obj.Id != id {
			continue
		}

		return &obj
	}

	return nil
}

// Find MxCell from MxGraphModel.Mxcells by ID.
func (m *MxGraphModel) FindMxCell(id string) *MxCell {
	for _, mxcell := range m.Content.MxCells {
		if mxcell.Id != id {
			continue
		}

		return &mxcell
	}

	return nil
}
