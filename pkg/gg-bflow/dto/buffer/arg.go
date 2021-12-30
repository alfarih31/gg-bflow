package buffer_dto

type WriteArg struct {
	Key  string `validate:"required,min=0,max=250"`
	Data []byte `validate:"required"`
}
