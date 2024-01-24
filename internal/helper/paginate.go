package helper

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page        int64 `json:"page,omitempty"`
	PageSize    int64 `json:"pageSize,omitempty"`
	TotalRecord int64 `json:"totalRecord,omitempty"`
}

func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p == nil {
			// If Pagination object is nil, return the original DB without pagination
			return db
		}

		var page, pageSize int64

		if p.Page > 0 {
			page = p.Page
		} else {
			// If Page is not provided or invalid, set default value
			page = 1
		}

		if p.PageSize > 0 && p.PageSize <= 100 {
			pageSize = p.PageSize
		} else {
			// If PageSize is not provided or invalid, set default value
			pageSize = 10
		}

		// Convert int64 to int for Offset
		offset := int((page - 1) * pageSize)
		size := int(pageSize)

		// Apply pagination
		return db.Offset(offset).Limit(size)
	}
}
