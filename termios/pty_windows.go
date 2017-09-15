package termios

func open_pty_master() (uintptr, error) {
	return 0, nil
}

func Ptsname(fd uintptr) (string, error) {
	return "", nil
}

func grantpt(fd uintptr) error {
	return nil
}

func unlockpt(fd uintptr) error {
	return nil
}
