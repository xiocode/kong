/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          checker.go
 * Description:   status checker
 */

package monitoring

import (
	"time"
)

type Checker interface {
	Check() (Status, err)
}

type Status struct {
	T     string // status type, e.g.: cpu,in,out,write,read
	Time  time.Time
	Value string
}

// SNMP
type StandardChecker struct {
	URL      string
	Interval int32
}

func (s *StandardChecker) Check() {

}
