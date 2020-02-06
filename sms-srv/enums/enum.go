package enums

// Enabled ...
type Enabled int32

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

// State 状态
type State int32

const (
	// Close 关
	Close State = iota
	// Open 开
	Open
	// Await 等待
	Await
)

func (s State) String() string {
	switch s {
	case Open:
		return "open"
	case Await:
		return "Await"
	default:
		return "close"
	}
}
