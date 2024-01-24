package helper

import (
	"gorm.io/gorm"
)

func PrettyORMErr(err error) *HTTPError {
	switch err {
	case gorm.ErrRecordNotFound:
		return NotFoundError("Record not found")
	case gorm.ErrDuplicatedKey:
		return BadRequestError("Duplicated key")
	default:
		return InternalServerError("Internal server error")
	}
}
