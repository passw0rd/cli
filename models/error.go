package models

import "fmt"

func (m *HttpError) Error() string {
	return fmt.Sprintf("%s", m.Message)
}
