package desktop

type Window interface {
	Iconify()
	Restore()
	Focus()
}
