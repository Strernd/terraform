package plans

import (
	"sync"

	"github.com/hashicorp/terraform/internal/addrs"
)

// FIXME: comment
type Conditions map[string]*ConditionResult

type ConditionResult struct {
	Address      addrs.Checkable
	Result       bool
	Unknown      bool
	Type         addrs.CheckType
	ErrorMessage string
}

func NewConditions() Conditions {
	return make(Conditions)
}

func (c Conditions) SyncWrapper() *ConditionsSync {
	return &ConditionsSync{
		results: c,
	}
}

type ConditionsSync struct {
	lock    sync.Mutex
	results Conditions
}

func (cs *ConditionsSync) SetResult(addr addrs.Check, result *ConditionResult) {
	if cs == nil {
		panic("SetResult on nil Conditions")
	}
	cs.lock.Lock()
	defer cs.lock.Unlock()

	cs.results[addr.String()] = result
}
