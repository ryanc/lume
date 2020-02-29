package lifx

const (
	OK       Status = "ok"
	TimedOut Status = "timed_out"
	Offline  Status = "offline"
)

func (s Status) Success() bool {
	return s == OK
}
