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
	// Create MxLibObj from MxGraphModel
	mxLibObjs := make([]libdrawio.MxLibObj, len(mxGraphModels))
	for index, mxGraphModel := range mxGraphModels {
		mxLibObj, err := mxlibrary.MakeMxLibObj(mxGraphModel, "title", 24, 24)
		if err != nil {
			log.Fatal(err)
		}
		mxLibObjs[index] = mxLibObj
	}
	// Create MxLibrary from array of MxLibObj
	if err := mxlibrary.MakeMxLibrary(mxLibObjs); err != nil {
		log.Fatal(err)
	}
	// Write mxlibrary file
	if err := mxlibrary.Write("path_to_mxlibrary"); err != nil {
		log.Fatal(err)
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
