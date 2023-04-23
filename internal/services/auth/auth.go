package auth

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"
	"two-factor-auth/internal/config"
	"two-factor-auth/internal/models/confirmations"
)

const limitAttempt = 3

var (
	ErrToManyAttempts = errors.New("to many attempts")
	ErrWrongCode      = errors.New("wrong auth code")
)

type Storage interface {
	AddConfirmation(context.Context, confirmations.Confirmation, time.Duration) error
	GetAttempts(context.Context, confirmations.UID) (string, error)
	GetCodeByRequestUID(context.Context, confirmations.UID) (confirmations.Code, error)
	DeleteConfirmationByRequestUID(context.Context, confirmations.UID) error
}

func NewService(cfg config.ServiceConfiguration, storage Storage) *Service {
	return &Service{
		codeTTL: cfg.AuthConfig.TTL,
		codeLen: cfg.AuthConfig.CodeLen,
		storage: storage,
	}
}

type Service struct {
	codeTTL time.Duration
	codeLen uint
	storage Storage
}

func (s *Service) SendCode(ctx context.Context, number confirmations.Number) (confirmations.Confirmation, error) {
	confirmation := confirmations.NewConfirmation(s.codeLen)

	err := s.storage.AddConfirmation(ctx, confirmation, s.codeTTL)
	if err != nil {
		log.Printf("SendCode: cant add confirmation: %v", err)
		return confirmations.Confirmation{}, err
	}

	// TODO: сделать отправку смс с кодом

	return confirmation, nil
}

func (s *Service) Verify(ctx context.Context, confirmation confirmations.Confirmation) (confirmations.CurrentUnixDate, error) {
	getCodeAttempts, err := s.storage.GetAttempts(ctx, confirmation.RequestUID)
	if err != nil {
		log.Printf("Verify: cant get attempts count: %v", err)
		return 0, err
	}

	cntAttempts, err := strconv.Atoi(getCodeAttempts)
	if err != nil {
		log.Printf("Verify: cant convert cntAttempts: %v", err)
		return 0, err
	}

	if cntAttempts >= limitAttempt {
		return 0, ErrToManyAttempts
	}

	code, err := s.storage.GetCodeByRequestUID(ctx, confirmation.RequestUID)
	if err != nil {
		log.Printf("Verify: cant get code by request uid: %v", err)
		return 0, err
	}

	if confirmation.AuthCode != code {
		return 0, ErrWrongCode
	}

	err = s.storage.DeleteConfirmationByRequestUID(ctx, confirmation.RequestUID)
	if err != nil {
		log.Printf("Verify: cant delete confirmation by request uid: %v", err)
		return 0, err
	}

	return confirmations.CurrentUnixDate(time.Now().Unix()), nil
}
