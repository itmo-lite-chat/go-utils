package closer

import "github.com/pkg/errors"

type CloseFunc = func() error

type Closer struct {
	closers map[string]CloseFunc
}

func NewCloser() *Closer {
	return &Closer{
		closers: map[string]CloseFunc{},
	}
}

func (c *Closer) AddCloseFunc(name string, closeFunc CloseFunc) {
	if _, ok := c.closers[name]; ok {
		panic("closer with selected name already exists")
	}
	c.closers[name] = closeFunc
}

func (c *Closer) Close() (err error) {
	for name, closeFunc := range c.closers {
		if err = closeFunc(); err != nil {
			err = errors.Wrapf(err, "can't close %s", name)
		}
	}
	return
}
