package dto

import "errors"

type Contacto struct {
	Id string `json:"id"`
}

func (c *Contacto) Validate() error {
	if len(c.Id) == 0 {
		return errors.New("empty contact id")
	}

	return nil
}
