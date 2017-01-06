package plugin

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const Archive = "UEsDBAoAAAAAALEGJ0oAAAAAAAAAAAAAAAAMABwAdGVzdF9wbHVnaW4vVVQJAAN9vW9Yfb1vWHV4CwABBOgDAAAE6AMAAFBLAwQUAAAACACuBidKt7cTCnkAAACdAAAAGQAcAHRlc3RfcGx1Z2luLy5taWt1dHRlci55bWxVVAkAA3i9b1h4vW9YdXgLAAEE6AMAAAToAwAATY3BCgIxDETv/YrccrHFRbzkO/YmshQ3dEvdbWlSv9+gIMIchsc8xnvv5DkSASmLLs16PtzKjY9VyAHsuQxV7gSXcA2Tke+G4IZpZDwBJi14dy/ukqtxnMIZXRy6VbOqxKW4I+5MMNsF/C7k0XPTjzJvWcASQf8mb1BLAwQUAAAACACJBidKd84vQnUAAACRAAAAGgAcAHRlc3RfcGx1Z2luL3Rlc3RfcGx1Z2luLnJiVVQJAAMyvW9YMr1vWHV4CwABBOgDAAAE6AMAAE2NMQrDMBAEe7/icCWB8QPuDwZDypBCWIs5opyEdCLk97ZJihRbzDLsrqnvovNWEQyODc08xTwQZS2okuOXiNafGVJy3Es8/YlU0kT3RZ7dDJX59mmGF/OC1sKOWfF2EW2rUkyyMo3XA/1Vo3/4cx8ahyvDAVBLAQIeAwoAAAAAALEGJ0oAAAAAAAAAAAAAAAAMABgAAAAAAAAAEADtQQAAAAB0ZXN0X3BsdWdpbi9VVAUAA329b1h1eAsAAQToAwAABOgDAABQSwECHgMUAAAACACuBidKt7cTCnkAAACdAAAAGQAYAAAAAAABAAAApIFGAAAAdGVzdF9wbHVnaW4vLm1pa3V0dGVyLnltbFVUBQADeL1vWHV4CwABBOgDAAAE6AMAAFBLAQIeAxQAAAAIAIkGJ0p3zi9CdQAAAJEAAAAaABgAAAAAAAEAAACkgRIBAAB0ZXN0X3BsdWdpbi90ZXN0X3BsdWdpbi5yYlVUBQADMr1vWHV4CwABBOgDAAAE6AMAAFBLBQYAAAAAAwADABEBAADbAQAAAAA="

func TestLoadSpec(t *testing.T) {
	f, err := ioutil.TempFile("", "mpa-test")
	if err != nil {
		t.Fatalf("Failed to create tempfile: %v", err.Error())
	}
	defer os.Remove(f.Name())

	data, err := base64.StdEncoding.DecodeString(Archive)
	if err != nil {
		t.Fatalf("Broken payload: %v", err.Error())
	}
	_, err = f.Write(data)
	if err != nil {
		t.Fatalf("Failed to write: %v", err.Error())
	}
	f.Close()

	spec, err := LoadSpec(f.Name())
	if err != nil {
		t.Errorf("Parse error: %v", err.Error())
	}

	expected := Spec{
		Slug:        ":test_plugin",
		Name:        "Test plugin",
		Description: "This is a test plugin",
		Version:     "1.0",
		Dependency: Dependency{
			MikutterVersion: "3.5.1",
			Plugins:         []string{"gui", "gtk"},
		},
	}
	if !reflect.DeepEqual(spec, expected) {
		t.Errorf("Wrong spec: %v", spec)
	}
}
