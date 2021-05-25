package workflow

import (
	"fmt"
	"math/rand"
	"time"
)

type Activities struct{}

func (a *Activities) UsuallyTrue(args []string) (bool, error) {
	result := rand.Intn(10) > 2
	fmt.Println("USUALLY TRUE RESULT", result)
	return result, nil
}

func (a *Activities) Wait() (string, error) {
	time.Sleep(time.Second * 5)
	return "waited", nil
}
func (a *Activities) SendEmail() (string, error) {
	return "sended", nil
}

// Get Name Activity.
func (a *Activities) DoMetrics() (string, error) {
	return "metriced", nil
}

// Get Greeting Activity.
func (a *Activities) CreateTicket() (string, error) {
	return "ticketed", nil
}
