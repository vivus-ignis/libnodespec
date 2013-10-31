package libnodespec

import (
	"os"
	"testing"
)

func TestLoadSpec(t *testing.T) {
	fd, err := os.Open("example.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	stat, err := fd.Stat()
	if err != nil {
		t.Fatal(err)
	}

	var outBuf []byte
	outBuf = make([]byte, stat.Size())
	_, err = fd.Read(outBuf)
	if err != nil {
		t.Fatal("Reading toml file failed:", err)
	}

	tomlData := string(outBuf)

	if _, err := loadSpec(tomlData); err != nil {
		t.Fatal("Decoding failed:", err)
	}
}
