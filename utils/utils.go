package utils

func Nillable[T any](input string, element T) *T {
	if input == "" {
		return nil
	}

	return &element
}

func NillableString(input string) *string {
	if input == "" {
		return nil
	}

	return &input
}
