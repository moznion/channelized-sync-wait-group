# channelized-sync-wait-group

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

