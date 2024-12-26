package yamlutil

import (
	"io"

	"gopkg.in/yaml.v3"
)

// MustString encodes data to YAML string, will panic on error.
func MustString(v any) string {
	bs, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// Encode encodes data to YAML bytes. alias of yaml.Marshal
func Encode(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

// EncodePretty encodes data to pretty YAML bytes with indentation.
func EncodePretty(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

// EncodeString encodes data to YAML string.
func EncodeString(v any) (string, error) {
	bs, err := yaml.Marshal(v)
	return string(bs), err
}

// EncodeToWriter encodes data to YAML and writes it to the writer.
func EncodeToWriter(v any, w io.Writer) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	return enc.Encode(v)
}

// Decode decodes YAML bytes to data pointer. alias of yaml.Unmarshal
func Decode(bts []byte, ptr any) error {
	return yaml.Unmarshal(bts, ptr)
}

// DecodeString decodes YAML string to data pointer.
func DecodeString(str string, ptr any) error {
	return yaml.Unmarshal([]byte(str), ptr)
}

// DecodeReader decodes YAML from io.Reader.
func DecodeReader(r io.Reader, ptr any) error {
	dec := yaml.NewDecoder(r)
	return dec.Decode(ptr)
}
