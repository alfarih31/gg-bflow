package buffer_dto

type Stat struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
	Exp  int64  `json:"exp"`
}
