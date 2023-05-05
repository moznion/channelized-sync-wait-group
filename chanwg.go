package chanwg

import "sync"

type ChannelizedWaitGroup struct {
	wg sync.WaitGroup
}

func (c *ChannelizedWaitGroup) Add(delta int) {
	c.wg.Add(delta)
}

func (c *ChannelizedWaitGroup) Done() {
	c.wg.Done()
}

func (c *ChannelizedWaitGroup) Wait() {
	c.wg.Wait()
}

func (c *ChannelizedWaitGroup) Await() <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		c.wg.Wait()
		ch <- struct{}{}
	}()
	return ch
}
