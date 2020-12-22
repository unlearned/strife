package strife

func Predict(desiredNum uint16, jsonPath string) (*phases, error) {
	r, err := readJSON(jsonPath)
	if err != nil {
		return nil, err
	}
	rs := r.Phases.Rates()
	ps, err := rs.predict(desiredNum)

	if err != nil {
		return nil, err
	}

	return ps, nil
}
