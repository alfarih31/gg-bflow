package meta_dto

import keyvalue "github.com/alfarih31/nb-go-keyvalue"

type WriteArg struct {
	Key  string            `validate:"required,min=0,max=250"`
	Meta keyvalue.KeyValue `validate:"required"`
}
