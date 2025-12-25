package yamlutil

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func TestReadWriteFile(t *testing.T) {
	as := assert.New(t)
	tmpFile := "test_yamlutil.yaml"
	defer os.Remove(tmpFile)

	data := TestStruct{Name: "Alice", Age: 30}

	// Test WriteFile
	err := WriteFile(tmpFile, data)
	as.NoError(err)

	// Test ReadFile
	var readData TestStruct
	err = ReadFile(tmpFile, &readData)
	as.NoError(err)
	as.Equal(data, readData)
}

func TestEncoding(t *testing.T) {
	as := assert.New(t)
	data := TestStruct{Name: "Bob", Age: 25}

	// MustString
	str := MustString(data)
	as.Contains(str, "name: Bob")
	as.Contains(str, "age: 25")

	// Encode
	bs, err := Encode(data)
	as.NoError(err)
	as.Contains(string(bs), "name: Bob")

	// EncodeString
	s, err := EncodeString(data)
	as.NoError(err)
	as.Contains(s, "name: Bob")

	// EncodeToWriter
	var buf bytes.Buffer
	err = EncodeToWriter(data, &buf)
	as.NoError(err)
	as.Contains(buf.String(), "name: Bob")
}

func TestDecoding(t *testing.T) {
	as := assert.New(t)
	yamlStr := "name: Charlie\nage: 40\n"

	// IsYAML
	as.True(IsYAML(yamlStr))
	as.False(IsYAML(": invalid yaml"))

	// DecodeString
	var t1 TestStruct
	err := DecodeString(yamlStr, &t1)
	as.NoError(err)
	as.Equal("Charlie", t1.Name)
	as.Equal(40, t1.Age)

	// Decode
	var t2 TestStruct
	err = Decode([]byte(yamlStr), &t2)
	as.NoError(err)
	as.Equal("Charlie", t2.Name)

	// DecodeReader
	var t3 TestStruct
	buf := bytes.NewBufferString(yamlStr)
	err = DecodeReader(buf, &t3)
	as.NoError(err)
	as.Equal("Charlie", t3.Name)
}
