/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          kong.go
 * Description:   kong main
 */

package kong

import (
	log "github.com/golang/glog"
	"kong/common"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	MAX_UNPROCESSED_PACKETS = 1000
)

type Kong struct {
	Interval   int // flush interval, N seconds!
	Metrics    chan common.Metric
	signalchan chan os.Signal
}

// submit monitoring data and tries to buffer to graphite
func (k *Kong) submit(deadline time.Time) error {
	log.Infoln("call submit!")

	return nil
}

// chan multiplexer !!!
func (k *Kong) multiplexer() {
	period := time.Duration(k.Interval) * time.Second
	ticker := time.NewTicker(period)
	for {
		select {
		case sign := <-k.signalchan: // oops, os signal!!
			switch sign {
			case syscall.SIGTERM, syscall.SIGINT:
				log.Infof("!! Caught signal %d... shutting down\n", sign)
				if err := k.submit(time.Now().Add(period)); err != nil {
					log.Errorf("ERROR: %s", err)
				}
				goto EXITING
			default:
				log.Warningf("unknown signal %d, ignoring\n", sign)
			}
		case <-ticker.C:
			if err := k.submit(time.Now().Add(period)); err != nil {
				log.Errorf("ERROR: %s", err)
			}
		case metric := <-k.Metrics:
			log.Infoln(metric)
		}
	}

EXITING:
	log.Infoln("Catch os signal, Kong is exiting!")
}

func (k *Kong) monitor() {
	period := time.Duration(k.Interval) * time.Second
	ticker := time.NewTicker(period)
	for {
		select {
		case <-ticker.C:
			// do somting
		}
	}

}

func main() {
	kong := &Kong{
		Interval:   1,
		Metrics:    make(chan common.Metric, MAX_UNPROCESSED_PACKETS),
		signalchan: make(chan os.Signal, 1),
	}
	signal.Notify(kong.signalchan) // waiting os singal !
	go kong.monitor()              // start monitor !
	kong.multiplexer()             // waiting for done!
}
