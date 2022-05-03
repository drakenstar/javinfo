package javinfo

type JAVInfo interface {
	FindByCode(Code) ([]*Title, error)
}

func New() JAVInfo {
	return &JAVLibraryBackend{}
}
