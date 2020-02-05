package enums

// Enabled ...
type Enabled uint

const (
	// No ...
	No Enabled = iota
	// Yes ...
	Yes
)

func (e Enabled) String() string {
	str := ""
	switch e {
	case Yes:
		str = "yes"
	default:
		str = "no"
	}
	return str
}

// Status 状态
type Status int

const (
	// Failed 失败
	Failed Status = iota
	// Success 失败
	Success
	// Await 等待
	Await
)

func (s Status) String() string {
	switch s {
	case Success:
		return "Success"
	case Await:
		return "Await"
	default:
		return "Failed"
	}
}
