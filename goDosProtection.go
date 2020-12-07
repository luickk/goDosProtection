package goDosProtection

import (
  "fmt"
  "time"
)

// contains timer and timer state information
type timerRoutine struct {
  timer time.Timer
  expired bool
}

// contains map with timer and timeroutine struct to manage client ban status
type dosController struct {
  ClientMap map[string]*timerRoutine
  ReconnectTimeLimit int
}

// inits new dos controller struct
func New(reconnectTimeLimit int) dosController{
  return dosController { make(map[string]*timerRoutine), reconnectTimeLimit }
}

// checks wether client's time is expired
// returns true if client addr is banned, fals if it isn't
// registers new client if not in map
func (dC *dosController) Client(address string) bool {
  if tR, ok := dC.ClientMap[address]; !ok {
    dC.ClientMap[address] = TimerRoutine(time.Second * time.Duration(dC.ReconnectTimeLimit))
    return false
  } else if tR.expired {
    dC.ClientMap[address] = TimerRoutine(time.Second * time.Duration(dC.ReconnectTimeLimit))
    return false
  }
  return true
}

// routine that isolates timer and gives timer an object like handling
func TimerRoutine(expTime time.Duration) *timerRoutine {
  timer := time.NewTimer(expTime)
  tR := timerRoutine { *timer, false }
  go func(){
      <-timer.C
      tR.expired = true
  }()
  return &tR
}
