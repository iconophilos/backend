package service

const (
	// Enumerate supported monument types

	monumentTypeMuseum  = "museum"
	monumentTypeTheatre = "theatre"
	monumentTypeChurch  = "church"
)

func validateCreateInput(m *Monument) error {
	if m.Name == "" {
		return ErrMissingName
	}

	if m.Type == "" {
		return ErrMissingType
	}

	if !isValidMonumentType(m.Type) {
		return ErrInvalidType
	}

	if m.Dating == "" {
		return ErrMissingDating
	}

	if m.ArchitecturalPlant == "" {
		return ErrMissingArchitecturalPlant
	}

	if m.Model3D == "" {
		return ErrMissingModel3D
	}

	if m.Country == "" {
		return ErrMissingCountry
	}

	if m.Region == "" {
		return ErrMissingRegion
	}

	if m.Latitude == 0 {
		return ErrMissingLatitude
	}

	if m.Longitude == 0 {
		return ErrMissingLongitude
	}
	return nil
}

func isValidMonumentType(t string) bool {
	switch t {
	case monumentTypeMuseum,
		monumentTypeTheatre,
		monumentTypeChurch:
		return true
	default:
		return false
	}
}
