package cli

func (c *CliApp) SelectedDefault() bool {
	if len(c.Args) == 1 {
		return true
	}
	return false
}

func (c *CliApp) SelectedPing() bool {
	if len(c.Args) == 2 && c.Args[1] == "ping" {
		return true
	}
	return false
}

func (c *CliApp) SelectedConfigure() bool {
	if len(c.Args) == 2 && c.Args[1] == "configure" {
		return true
	}
	return false
}
