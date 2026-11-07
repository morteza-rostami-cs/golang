```go


// Buffer

// a temp memory space -- we put the data their to wait to be processed

func main() {

	// a buffer can be a slice, array or whatever! basically a piece of memory that we put stuff their to wait
	var buffer []int

	// producer adds data
	buffer = append(buffer, 10)
	buffer = append(buffer, 20)
	buffer = append(buffer, 30)

	println("Buffer: ", buffer)

	// consumer processes data
	for _, v := range buffer{
		println("processing: ", v)
	}

}

```

```go

func main() {

	// create a file
	file, err := os.Create("output.txt")

	if err != nil {
		panic(err)
	}

	// defer closing the file
	defer file.Close()

	// create a buffered writer
	writer := bufio.NewWriter(file)

	// write data into the buffer
	writer.WriteString("hi, ")
	// write data into the buffer
	writer.WriteString("ali!")

	// send the buffered data into the underlying file writer (we write to permanent memory, disk)
	writer.Flush()
}
```

```go

// buffered and unbuffered channels

// an unbuffered channel
// there is not capacity to store data here, while waiting!
// sender and receiver -> they have to wait for the the job to be done => then we send and receive data in one go.
var channel1 chan int

func main() {
	channel1 = make(chan int)

	go func() {
		// if no body is ready to receives, sender blocks
		channel1<-42
	}()

	// main -> goroutine -- receives data
	value := <-channel1 // wait for it

	println(value)
}
```

```go
// buffered and unbuffered channels

var buffered chan int

func main() {

	// a buffered channel
	// sender can buffer 3 values inside Buffer
	// as long as there is room in buffer, sender can put value is here (without waiting for a receiver)
	buffered = make(
		chan int,
		3,// buffered capacity of 3 ints
	)

	go func() {
		// we send all 3 values, first! before receiving any of them.
		buffered<- 1
		// an unbuffered channel would be blocked, and waiting for the first value to be received!! to send another one.
		buffered<- 2
		buffered<- 3

		// if the buffer is full => if there is no space inside buffer => and no one consuming!! this blocks!!
		// buffered<-99
		// buffered<-99

		// this does not run, if blocked
		println("done")
	}()

	println(<-buffered)
	println(<-buffered)
	println(<-buffered)
}
```
