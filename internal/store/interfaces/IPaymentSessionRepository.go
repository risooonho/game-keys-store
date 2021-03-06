package interfaces

import "github.com/rdsalakhov/game-keys-store/internal/model"

type IPaymentSessionRepository interface {
	Create(seller *model.PaymentSession) error
	Find(ID int) (*model.PaymentSession, error)
	DeleteByID(int) error
	GetPaymentInfo(sessionID int) (*model.PaymentInfo, error)
	UpdateNotifiedStatus(sessionID int, newStatus bool) error
	UpdatePerformedStatus(sessionID int, newStatus bool) error
}
