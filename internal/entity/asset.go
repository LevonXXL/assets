package entity

import (
	"time"
)

type Asset struct {
	Name      string    `json:"name"`
	UId       uint32    `json:"uid"`
	Data      []byte    `json:"data,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
