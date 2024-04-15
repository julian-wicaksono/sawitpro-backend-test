package handler

import (
	"errors"
	"regexp"

	"github.com/SawitProRecruitment/UserService/generated"
)

type ValidationRegister struct {
	PhoneNumberValidation string `json:"phone_number"`
	FullNameValidation    string `json:"full_name"`
	PasswordValidaton     string `json:"password"`
}

func ValidateRegister(params generated.RegisterRequest) (message ValidationRegister, err error) {
	validationRegister := ValidationRegister{}
	tempErr := validationRegister.validatePhoneNumber(params)
	if tempErr != nil {
		err = tempErr
	}
	tempErr = validationRegister.validateFullName(params)
	if tempErr != nil {
		err = tempErr
	}
	tempErr = validationRegister.validatePassword(params)
	if tempErr != nil {
		err = tempErr
	}

	return validationRegister, err
}

func (validation *ValidationRegister) validatePhoneNumber(params generated.RegisterRequest) error {
	regex, _ := regexp.Compile(`^\+62`)
	containCountryCode := regex.MatchString(params.PhoneNumber)

	var err error
	if len(params.PhoneNumber) < 10 {
		validation.PhoneNumberValidation = "Phone number must be at least 10 characters"
		err = errors.New("fail validation")
	} else if len(params.PhoneNumber) > 13 {
		validation.PhoneNumberValidation = "Phone number must be less than 13 characters"
		err = errors.New("fail validation")
	} else if !containCountryCode {
		validation.PhoneNumberValidation = "Phone number must start with +62" + params.PhoneNumber
		err = errors.New("fail validation")
	}
	return err
}

func (validation *ValidationRegister) validateFullName(params generated.RegisterRequest) error {
	var err error
	if len(params.FullName) < 3 {
		validation.FullNameValidation = "Full name must be at least 3 characters"
		err = errors.New("fail validation")
	} else if len(params.FullName) > 60 {
		validation.FullNameValidation = "Full name must be less than 60 characters"
		err = errors.New("fail validation")
	}
	return err
}

func (validation *ValidationRegister) validatePassword(params generated.RegisterRequest) error {

	var err error
	if len(params.Password) < 6 {
		validation.PasswordValidaton = "Password must be at least 6 characters"
		err = errors.New("fail validation")
	} else if len(params.Password) > 64 {
		validation.PasswordValidaton = "Password must be less than 64 characters"
		err = errors.New("fail validation")
	} else {
		regex, _ := regexp.Compile(`[A-Z]+`)
		passwordValidaton := regex.MatchString(params.Password)
		if !passwordValidaton {
			err = errors.New("fail validation")
			validation.PasswordValidaton = "Password must contain at least 1 capital characters"
		}
		regex, _ = regexp.Compile(`\d`)
		passwordValidaton = regex.MatchString(params.Password)
		if !passwordValidaton {
			err = errors.New("fail validation")
			validation.PasswordValidaton = "Password must contain at least 1 number"
		}
		regex, _ = regexp.Compile(`[\W_]`)
		passwordValidaton = regex.MatchString(params.Password)
		if !passwordValidaton {
			err = errors.New("fail validation")
			validation.PasswordValidaton = "Password must contain at least 1 special (non alpha-numeric) characters"
		}
	}

	return err
}
