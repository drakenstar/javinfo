package javinfo

type JAVInfo interface {
	GetTitleByCode(Code) *Title
}

func New() JAVInfo {
	return &JAVLibraryBackend{}
}
