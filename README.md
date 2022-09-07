## Quick Start
+ Execute command below to install this library.
```
~$ go get github.com/capybara-alt/libdrawio
```
## Usage
### MxGraphModel
MxGraphModel struct manage mxGraphModel (xml)
```
	// Decompress compressed MxGraphModel string, then map to MxGraphModel struct.
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	// Compress MxGraphModel struct.
	compressed, err := mxGraphModel.Compress()
	if err != nil {
		t.Fail()
	}
```
### MxLibrary
MxLibrary struct manage mxlibrary (xml file).
```
	// Assign to Mxlibrary.MxGraphModels
	mxlibrary.MxGraphModels = mxGraphModels
	// Then you can make mxlibrary string from MxModels.
	err := mxlibrary.MakeMxLibrary(func(m *libdrawio.MxLibObject) *libdrawio.MxLibObject {
		m.H = 24
		m.W = 24
		m.Title = "test"
		return m
	})
	if err != nil {
		t.Fail()
	}
```
### MxFile
MxFile struct manage mxfile (drawio file).
```
	// Load drawio file, then map to MxFile struct.
	mxfile, err := libdrawio.NewMxFile("./mxfile_test_1.drawio")
	if err != nil {
		t.Fail()
	}

	// Do something with MxFile struct and MxGraphModel struct...

	// Write drawio file.
	if err = mxfile.Write("./out.drawio"); err != nil {
		t.Fail()
	}
```
