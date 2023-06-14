package repository

type ValidatorInterface interface {
	Validate(action string) error
}
