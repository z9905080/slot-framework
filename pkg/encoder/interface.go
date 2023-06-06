package encoder

type Encoder interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	MarshalToString(v interface{}) (string, error)
	UnmarshalFromString(data string, v interface{}) error
}
