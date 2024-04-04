package mock

import "time"

// ----------------------------------------------- mock spyTime implement interface

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// ----------------------------------------------- mock spy write sleep implement interface

const write = "write"
const sleep = "sleep"

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
