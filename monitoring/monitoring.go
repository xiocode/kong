/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          monitoring.go
 * Description:   monitoring
 */

package monitoring

import (
	"github.com/alouca/gosnmp"
)

type Monitor struct {
	Conf  map[string]interface{}
	Check Checker
}
