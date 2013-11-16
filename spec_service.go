package libnodespec

func (spec SpecService) Run(defaults PlatformDefaults) (err error) {
	switch defaults.OperatingSystem {
	case "darwin":
		err = __findProcessDarwin(spec.Name)
	}

	return err
}
