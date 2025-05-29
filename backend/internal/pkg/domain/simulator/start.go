package simulator

import (
	"fmt"
	"os"
	"path/filepath"
	"sensor-simulator/internal/configs"
	"time"
)

func (s *Simulator) Start() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.isWorking {
		return
	}

	if configs.GetConfig().Simulator.LogsEnabled {
		logDir := "/root/logs"
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			logDir = ""
		}

		filename := fmt.Sprintf("%s_%s.log", s.name, time.Now().Format("2006-01-02 15:04:05"))

		logPath := filepath.Join(logDir, filename)

		logFile, err := os.Create(logPath)
		if err != nil {
			s.logFile = os.NewFile(0, os.DevNull)
		}

		s.logFile = logFile
	}

	s.wg.Add(1)

	go func() {
		for {
			select {
			case <-s.endChan:
				s.wg.Done()
				return
			case <-s.ticker.C:
				s.tick()
			}
		}
	}()

	timer := time.NewTimer(s.duration)
	go func() {
		for {
			select {
			case <-timer.C:
				if s.duration > 0 {
					s.mx.Lock()
					defer s.mx.Unlock()

					if !s.isWorking {
						return
					}

					s.endChan <- struct{}{}
					s.wg.Wait()

					s.isWorking = false
					return
				}
			case <-s.stopTimerChan:
				return
			}
		}
	}()

	s.isWorking = true
}
