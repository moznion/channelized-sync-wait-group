package chanwg

import "testing"

func TestChannelizedWaitGroup_Wait(t *testing.T) {
	cwg := &ChannelizedWaitGroup{}

	cwg.Add(3)

	go func() {
		cwg.Done()
	}()
	go func() {
		cwg.Done()
	}()
	go func() {
		cwg.Done()
	}()

	cwg.Wait()
}

func TestChannelizedWaitGroup_Await(t *testing.T) {
	cwg := &ChannelizedWaitGroup{}

	cwg.Add(3)

	go func() {
		cwg.Done()
	}()
	go func() {
		cwg.Done()
	}()
	go func() {
		cwg.Done()
	}()

	<-cwg.Await()
}
