package meta

type MetaRepo interface {
	Create
	Retrieve
	Update
}

type repo struct {
	Create
	Retrieve
	Update
}

var Instance MetaRepo = &repo{
	Create:   NewCreate(),
	Retrieve: NewRetrieve(),
	Update:   NewUpdate(),
}
