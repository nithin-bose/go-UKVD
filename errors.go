package ukvd

import (
	"errors"
)

var (
	APIFailureError                = errors.New("API returned 'Failure'")
	InvalidPackageIdError          = errors.New("API returned 'InvalidPackageId'")
	ItemNotFoundError              = errors.New("API returned 'ItemNotFound'")
	KeyInvalidError                = errors.New("API returned 'KeyInvalid'")
	ServiceUnavailableError        = errors.New("API returned 'ServiceUnavailable'")
	InvalidPackageAccessClassError = errors.New("API returned 'InvalidPackageAccessClass'")
)
