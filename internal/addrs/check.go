package addrs

import "fmt"

// FIXME: comment
type CheckType int

const (
	CheckInvalid               CheckType = 0
	CheckResourcePrecondition  CheckType = 1
	CheckResourcePostcondition CheckType = 2
	CheckOutputPrecondition    CheckType = 3
)

func (c CheckType) String() string {
	switch c {
	case CheckResourcePrecondition:
		return "Resource precondition"
	case CheckResourcePostcondition:
		return "Resource postcondition"
	case CheckOutputPrecondition:
		return "Module output value precondition"
	default:
		// This should not happen
		return "Condition"
	}
}

// FIXME: comment
type Check struct {
	Container Checkable
	Type      CheckType
	Index     int
}

func (c Check) String() string {
	container := c.Container.String()
	switch c.Type {
	case CheckResourcePrecondition:
		return fmt.Sprintf("%s.preconditions[%d]", container, c.Index)
	case CheckResourcePostcondition:
		return fmt.Sprintf("%s.postconditions[%d]", container, c.Index)
	case CheckOutputPrecondition:
		return fmt.Sprintf("%s.preconditions[%d]", container, c.Index)
	default:
		// This should not happen
		return fmt.Sprintf("%s.conditions[%d]", container, c.Index)
	}
}

// FIXME: comment
type Checkable interface {
	checkableSigil()

	Check(CheckType, int) Check
	String() string
}

// The following are all of the possible Checkable address types:
var (
	_ Checkable = AbsResourceInstance{}
	_ Checkable = AbsOutputValue{}
)

type checkable struct {
}

func (c checkable) checkableSigil() {
}
