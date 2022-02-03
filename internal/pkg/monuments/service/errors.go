package service

import "errors"

var (
	// Enumerate service errors

	ErrMissingName               error = errors.New("missing monument name")
	ErrMissingType               error = errors.New("missing monument type")
	ErrInvalidType               error = errors.New("invalid monument type")
	ErrMissingDating             error = errors.New("missing monument dating")
	ErrMissingArchitecturalPlant error = errors.New("missing monument architectural plant")
	ErrMissingModel3D            error = errors.New("missing monument 3d model")
	ErrMissingLocation           error = errors.New("missing monument location")
	ErrMissingCountry            error = errors.New("missing monument location country")
	ErrMissingRegion             error = errors.New("missing monument location region")
	ErrMissingLatitude           error = errors.New("missing monument location latitude")
	ErrMissingLongitude          error = errors.New("missing monument location longitude")
	ErrMonumentAlreadyExists     error = errors.New("monument already exists")
	ErrMonumentNotFound          error = errors.New("monument not found")
)
