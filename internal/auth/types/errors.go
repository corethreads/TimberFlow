package types

import "errors"

var (
	NothinginFields = errors.New("No credentials provided above")
	AlreadyAdded    = errors.New("Info Already added")
)
