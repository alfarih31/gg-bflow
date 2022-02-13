package meta_dto

type WriteArg struct {
	Key  string `validate:"required,min=0,max=250"`
	Meta string `validate:"required,min=0"`
}
