package storage

import (
	"time"

	"github.com/google/uuid"
)

type StorageItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name           string    `json:"name"`
	Quantity       int       `json:"quantity"`
	Unit           string    `json:"unit"`
	ExpirationDate time.Time `json:"expirationDate"`
	PurchaseDate   time.Time `gorm:"default:now()" json:"purchaseDate"`
}
