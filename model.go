package javinfo

import "time"

type Code string

type Title struct {
	ID          string
	Code        Code
	Title       string
	Models      []*Model
	Tags        []*Tag
	ReleaseDate time.Time
}

type Model struct {
	ID      string
	Name    string
	Aliases []string
}

func (m *Model) String() string {
	return m.Name
}

type Tag struct {
	ID   string
	Name string
}

func (t *Tag) String() string {
	return t.Name
}

func (c Code) Studio() string {
	switch c {
	case "IPX":
		return "IDEA POCKET"
	default:
		return ""
	}
}
