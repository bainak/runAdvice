package runadvice

import "fmt"

type Run interface {
	GetAdvice() string
}

type run struct {
	distance int
}

func NewRun(d int) Run {
	return &run{
		distance: d,
	}
}

func (r *run) GetAdvice() string {

	return fmt.Sprint("To distance %d run twice a week", r.distance)

}
