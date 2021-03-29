package file

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
)

// verifies that a file exists, can be read and is a valid HTML
func InspectFile(path string) error {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	r := bytes.NewReader(content)
	d := xml.NewDecoder(r)

	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity

	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return nil
		case nil:
		default:
			return err
		}
	}
}
