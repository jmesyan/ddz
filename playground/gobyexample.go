package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"time"
)

func runExample(ex func()) {
	name := runtime.FuncForPC(reflect.ValueOf(ex).Pointer()).Name()
	fmt.Println("\n****** running example " + name + " ******")
	ex()
}

func ex1HelloWorld() {
	fmt.Println("Hello world!")
}

func ex2Values() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func ex3Variables() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

const s string = "constant"

func ex4Constants() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func ex5For() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func ex6IfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func ex7Switch() {
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func ex8Arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func ex9Slices() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func ex10Maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func ex11Range() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func ex12Functions() {
	plus := func(a int, b int) int {
		return a + b
	}

	plusPlus := func(a, b, c int) int {
		return a + b + c
	}

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

func ex13MultipleReturnValues() {
	vals := func() (int, int) {
		return 3, 7
	}

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)
}

func ex14VariadicFunctions() {
	sum := func(nums ...int) {
		fmt.Print(nums, " ")
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func ex15Closures() {
	intSeq := func() func() int {
		i := 0
		return func() int {
			i += 1
			return i
		}
	}

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func ex16Recursion() {
	var fact func(n int) int
	fact = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}

	fmt.Println(fact(7))
}

func ex17Pointers() {
	zeroval := func(ival int) {
		ival = 0
	}

	zeroptr := func(iptr *int) {
		*iptr = 0
	}

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func ex18Structs() {
	type person struct {
		name string
		age  int
	}

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

type rect19 struct {
	width, height int
}

func (r *rect19) area() int {
	return r.width * r.height
}

func (r rect19) perim() int {
	return 2*r.width + 2*r.height
}

func ex19Methods() {
	r := rect19{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func ex20Interfaces() {
	measure := func(g geometry) {
		fmt.Println(g)
		fmt.Println(g.area())
		fmt.Println(g.perim())
	}

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func ex21Errors() {
	f1 := func(arg int) (int, error) {
		if arg == 42 {
			return -1, errors.New("can't work with 42")
		}
		return arg + 3, nil
	}

	f2 := func(arg int) (int, error) {
		if arg == 42 {
			return -1, &argError{arg, "can't work with it"}
		}
		return arg + 3, nil
	}

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok { // type assertion via i.(T)
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func ex22Goroutines() {
	f := func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	}

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Millisecond) // trigger goroutines
}

func ex23Channels() {
	messages := make(chan string)

	go func() {
		fmt.Println("before sending ping")
		messages <- "ping"
		fmt.Println("ping sent")
	}()

	fmt.Println("goroutine go")
	msg := <-messages
	fmt.Println(msg)
}

func ex24ChannelBuffering() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func ex25ChannelSynchronization() {
	worker := func(done chan bool) {
		fmt.Print("working...")
		time.Sleep(time.Second)
		fmt.Println("Done")

		done <- true
	}

	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func ex26ChannelDirections() {
	ping := func(pings chan<- string, msg string) {
		pings <- msg
	}

	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ex27Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func ex28Timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

func ex29NonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func ex30ClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

func ex31RangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

func ex32Timers() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("Timer 1 created")
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	fmt.Println("Timer 2 created")
	go func() {
		// no chance going in here if exfunc is main
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	fmt.Println("Stopping timer 2")
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func ex33Tickers() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func ex34WorkerPool() {
	worker := func(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			fmt.Println("worker", id, "started job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", j)
			results <- j * 2
		}
	}

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("jobs sent")
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}

func ex35RateLimiting() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func main() {
	runExample(ex1HelloWorld)
	runExample(ex2Values)
	runExample(ex3Variables)
	runExample(ex4Constants)
	runExample(ex5For)
	runExample(ex6IfElse)
	runExample(ex7Switch)
	runExample(ex8Arrays)
	runExample(ex9Slices)
	runExample(ex10Maps)
	runExample(ex11Range)
	runExample(ex12Functions)
	runExample(ex13MultipleReturnValues)
	runExample(ex14VariadicFunctions)
	runExample(ex15Closures)
	runExample(ex16Recursion)
	runExample(ex17Pointers)
	runExample(ex18Structs)
	runExample(ex19Methods)
	runExample(ex20Interfaces)
	runExample(ex21Errors)
	runExample(ex22Goroutines)
	runExample(ex23Channels)
	runExample(ex24ChannelBuffering)
	runExample(ex25ChannelSynchronization)
	runExample(ex26ChannelDirections)
	runExample(ex27Select)
	runExample(ex28Timeouts)
	runExample(ex29NonBlockingChannelOperations)
	runExample(ex30ClosingChannels)
	runExample(ex31RangeOverChannels)
	runExample(ex32Timers)
	runExample(ex33Tickers)
	runExample(ex34WorkerPool)
	runExample(ex35RateLimiting)
}
