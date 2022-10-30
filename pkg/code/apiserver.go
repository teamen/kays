package code

const (
	ErrUserNotFound = iota + 110001

	ErrUserAlreadyExist
)

const (
	ErrSettingNotFound = iota + 111001

	ErrSettingSlugDuplicated
)

const (
	ErrCustomerNotFound = iota + 111101
)

const (
	ErrProductNotFound = iota + 111201
)
