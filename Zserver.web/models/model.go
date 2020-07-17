package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `sql:"index" json:"deletedAt"`
	// *time.Time 在 json.Marshal 后是空值, 再json.UnMarshal时就会报错
}
