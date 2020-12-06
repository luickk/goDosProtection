package goDosProtection

import (
  "time"
)

type timerRoutine struct {
  timer time.Timer
  expired bool
}

type dosController struct {
  ClientMap map[string]*timerRoutine
  ReconnectTimeLimit int
}


func New(reconnectTimeLimit int) dosController{
  return dosController { make(map[string]*timerRoutine), reconnectTimeLimit }
}

// checks wether client's time is expired
// returns true if time is expired, false if it isn't
// registers new client if not in map
func (dC *dosController) Client(address string) bool {
  if client, ok := dC.ClientMap[address]; !ok {
    dC.ClientMap[address] = TimerRoutine(time.Second * 1)
    return true
  } else if client.expired {
    return true
  }
  return false
}

func TimerRoutine(expTime time.Duration) *timerRoutine {
  timer := time.NewTimer(expTime)
  tR := timerRoutine { *timer, false }
  go func(){
      <-timer.C
      tR.expired = true
  }()
  return &tR
}
