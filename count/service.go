package count

type CounterService struct {
	count int
}

func NewCounterService() *CounterService {
	return &CounterService{}
}

func (c *CounterService) Increment() {
	c.count++
}
