package swiss_kit

func ToPoint[T int64 | int32 | int16 | int8 | float64 | float32 | string](v T) *T {
	return &v
}

func FromNumberPoint[T int64 | int32 | int16 | int8 | float64](v *T) T {
	return *v
}

func FromStringPoint(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
