package javinfo

type Code string

type Title struct {
	Code   Code
	Models []*Model
	Tags   []*Tag
}

type Model struct {
	Name    string
	Aliases []string
}

type Tag struct {
	Name string
}
