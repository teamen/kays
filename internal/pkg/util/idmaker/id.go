package idmaker

import shortid "github.com/jasonsoft/go-short-id"

func GenId() (id string) {
	opt := shortid.Options{
		Number:        11,
		StartWithYear: true,
		EndWithHost:   false,
	}
	id = shortid.Generate(opt)
	return
}
