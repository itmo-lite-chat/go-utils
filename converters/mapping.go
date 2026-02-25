package converters

// Mapping converts one model to another with mapping func.
func Mapping[I, O any](models []I, mapping func(I) O) []O {
	if len(models) == 0 {
		return nil
	}

	var mappedModels = make([]O, 0, len(models))
	for _, m := range models {
		mappedModels = append(mappedModels, mapping(m))
	}

	return mappedModels
}

// MappingWithError converts one model to another with mapping func.
// If mapping func returns error stop mapping and return mapping func error.
func MappingWithError[I, O any](models []I, mapping func(I) (O, error)) ([]O, error) {
	if len(models) == 0 {
		return nil, nil
	}

	var mappedModels = make([]O, 0, len(models))
	for _, m := range models {
		mappedModel, err := mapping(m)
		if err != nil {
			return nil, err
		}

		mappedModels = append(mappedModels, mappedModel)
	}

	return mappedModels, nil
}

// Mapping converts one model to another with mapping func.
func MappingWithErrorIgnoring[I, O any](models []I, mapping func(I) (O, error)) []O {
	if len(models) == 0 {
		return nil
	}

	var mappedModels = make([]O, 0, len(models))
	for _, m := range models {
		if mappedModel, err := mapping(m); err == nil {
			mappedModels = append(mappedModels, mappedModel)
		}
	}

	return mappedModels
}
