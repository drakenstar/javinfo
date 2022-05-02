package javinfo

type JAVLibraryBackend struct{}

func (*JAVLibraryBackend) GetTitleByCode(Code) *Title {
	panic("unimplemented")
}
