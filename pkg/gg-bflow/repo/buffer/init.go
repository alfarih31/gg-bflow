package buffer

type BufferRepo interface {
	Create
	Retrieve
}

type repo struct {
	Create
	Retrieve
}

var Repo BufferRepo = &repo{
	Create:   NewCreate(),
	Retrieve: NewRetrieve(),
}
