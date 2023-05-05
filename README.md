# channelized-sync-wait-group [![.github/workflows/check.yml](https://github.com/moznion/channelized-sync-wait-group/actions/workflows/check.yml/badge.svg)](https://github.com/moznion/channelized-sync-wait-group/actions/workflows/check.yml) [![codecov](https://codecov.io/gh/moznion/channelized-sync-wait-group/branch/main/graph/badge.svg?token=KLX02F6M6Y)](https://codecov.io/gh/moznion/channelized-sync-wait-group)

This tiny library wraps the [sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup) and enables it to await by a channel.

## Synopsis

### Normal `#Wait()`

```go
import "github.com/moznion/channelized-sync-wait-group"

func DoSomething() {
	cwg := &ChannelizedWaitGroup{}

	cwg.Add(1)

	go func() {
		// do something asynchronous operation
		cwg.Done()
	}()

	cwg.Wait()
}
```

### With a channel powered by `#Await()`

```go
import "github.com/moznion/channelized-sync-wait-group"

func DoSomething() {
	cwg := &ChannelizedWaitGroup{}

	cwg.Add(1)

	go func() {
		// do something asynchronous operation
		cwg.Done()
	}()

	<-cwg.Await()
}
```

## Author

moznion(<moznion@mail.moznion.net>)

