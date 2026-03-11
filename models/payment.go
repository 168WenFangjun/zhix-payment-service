package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          uint           `json:"userId"`
	OrderID         string         `gorm:"uniqueIndex" json:"orderId"`
	Amount          float64        `json:"amount"`
	Currency        string         `json:"currency"`
	Status          string         `json:"status"` // pending, completed, failed, refunded
	PaymentMethod   string         `json:"paymentMethod"` // apple_pay
	ApplePayToken   string         `json:"applePayToken,omitempty"`
	TransactionID   string         `json:"transactionId,omitempty"`
	ProductType     string         `json:"productType"` // membership, article
	ProductID       uint           `json:"productId,omitempty"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type ApplePaySession struct {
	MerchantIdentifier string `json:"merchantIdentifier"`
	DisplayName        string `json:"displayName"`
	Initiative         string `json:"initiative"`
	InitiativeContext  string `json:"initiativeContext"`
}

type ApplePayRequest struct {
	Token     ApplePayToken `json:"token"`
	OrderID   string        `json:"orderId"`
	Amount    float64       `json:"amount"`
	ProductType string      `json:"productType"`
	ProductID uint          `json:"productId,omitempty"`
}

type ApplePayToken struct {
	PaymentData       PaymentData       `json:"paymentData"`
	PaymentMethod     PaymentMethod     `json:"paymentMethod"`
	TransactionIdentifier string        `json:"transactionIdentifier"`
}

type PaymentData struct {
	Version   string `json:"version"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
	Header    Header `json:"header"`
}

type Header struct {
	EphemeralPublicKey string `json:"ephemeralPublicKey"`
	PublicKeyHash      string `json:"publicKeyHash"`
	TransactionId      string `json:"transactionId"`
}

type PaymentMethod struct {
	DisplayName string `json:"displayName"`
	Network     string `json:"network"`
	Type        string `json:"type"`
}
