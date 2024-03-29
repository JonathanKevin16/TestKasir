package service

import (
	"TestKasir/internal/domain/auth/model/dto"
	"TestKasir/shared"
	"TestKasir/shared/failure"
	"github.com/rs/zerolog/log"
)

type CashierService interface {
	RegisterCashier(req dto.RegisterCashierRequest) (string, error)
	LoginCashier(req dto.LoginCashierRequest) (dto.LoginResponse, error)
}

func (s *AuthServiceImpl) RegisterCashier(req dto.RegisterCashierRequest) (string, error) {
	message := "Failed"

	hashedPassword, err := shared.HashPassword(req.Password)
	if err != nil {
		log.Error().Err(err).Msg("[RegisterCashier] Failed to hash password")
		return message, err
	}
	req.Password = hashedPassword
	cashierUser := req.ToModel()

	err = s.AuthRepository.CreateUser(&cashierUser)
	if err != nil {
		log.Error().Err(err).Msg("[RegisterCashier] Failed to create user cashier")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *AuthServiceImpl) LoginCashier(req dto.LoginCashierRequest) (dto.LoginResponse, error) {

	user, err := s.AuthRepository.GetUserByUsername(req.Username)
	if err != nil {
		log.Error().Err(err).Msg("[LoginCashier] Failed to retrieve user")
		return dto.LoginResponse{}, err
	}

	err = shared.VerifyPassword(req.Password, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("[LoginCashier] Password verification failed")
		err = failure.Unauthorized("Invalid credentials")
		return dto.LoginResponse{}, err
	}

	token, err := shared.SignJWTToken(user.Username, user.Role, []byte(s.CFG.App.JWTAccessKey))
	if err != nil {
		log.Error().Err(err).Msg("[LoginCashier] Failed to sign JWT token")
		return dto.LoginResponse{}, err
	}

	return dto.BuildLoginResponse(token), nil
}
