package storage

import (
	"crypto/sha1"
	"fmt"
	"io"

	"Bot/lib/e"
)

const (
	errMsg = "cant hash"
	errMsgNoSave   = "no save page"
)


type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p *Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.WrapIfErr(errMsg, err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.WrapIfErr(errMsg, err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
