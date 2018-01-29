package main

import (
	"crypto/sha256"
	"fmt"
)

func hashwork(message []byte, rounds int) []byte {
	var output []byte
	for i := 0; i < rounds; i++ {
		o := sha256.Sum256(message)
		//fmt.Printf("%d %x\n", i, o[:])
		message = o[:]
		output = o[:]
	}
	return output
}

func main() {
	d := []byte("Hello World")
	//h := sha256.Sum256(d)
	//fmt.Printf("%x\n",h[:])
	//s := fmt.Sprintf("%x", h)
	//o := sha256.Sum256([]byte(s))
	//fmt.Printf("%x\n", o)
	//fmt.Printf("\n")
	//m := sha256.Sum256(d)
	//fmt.Printf("%x\n",m[:])
	//n := sha256.Sum256(m[:])
	//fmt.Printf("%x\n",n[:])
	g := hashwork(d, 2)
	fmt.Printf("%x\n", g)
	g = hashwork(d, 3)
	fmt.Printf("%x\n", g)
	g = hashwork(d, 1000)
	fmt.Printf("%x\n", g)
	g = hashwork(d, 1000000)
	fmt.Printf("%x\n", g)
	// g = hashwork(d, 1000000000)
	// fmt.Printf("%x\n", g)
}

// workpool

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker", id, "started  job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished job", j)
// 		results <- j * 2
// 	}
// }
//
// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
//
// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
//
// 	for a := 1; a <= 5; a++ {
// 		<-results
// 	}
// }

// workerpool
// type Job struct {
// 	id       int
// 	randomno int
// }
// type Result struct {
// 	job         Job
// 	sumofdigits int
// }
//
// var jobs = make(chan Job, 10)
// var results = make(chan Result, 10)
//
// func digits(number int) int {
// 	sum := 0
// 	no := number
// 	for no != 0 {
// 		digit := no % 10
// 		sum += digit
// 		no /= 10
// 	}
// 	time.Sleep(2 * time.Second)
// 	return sum
// }
// func worker(wg *sync.WaitGroup) {
// 	for job := range jobs {
// 		output := Result{job, digits(job.randomno)}
// 		results <- output
// 	}
// 	wg.Done()
// }
// func createWorkerPool(noOfWorkers int) {
// 	var wg sync.WaitGroup
// 	for i := 0; i < noOfWorkers; i++ {
// 		wg.Add(1)
// 		go worker(&wg)
// 	}
// 	wg.Wait()
// 	close(results)
// }
// func allocate(noOfJobs int) {
// 	for i := 0; i < noOfJobs; i++ {
// 		randomno := rand.Intn(999)
// 		job := Job{i, randomno}
// 		jobs <- job
// 	}
// 	close(jobs)
// }
// func result(done chan bool) {
// 	for result := range results {
// 		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
// 	}
// 	done <- true
// }
// func main() {
// 	startTime := time.Now()
// 	noOfJobs := 100
// 	go allocate(noOfJobs)
// 	done := make(chan bool)
// 	go result(done)
// 	noOfWorkers := 10
// 	createWorkerPool(noOfWorkers)
// 	<-done
// 	endTime := time.Now()
// 	diff := endTime.Sub(startTime)
// 	fmt.Println("total time taken ", diff.Seconds(), "seconds")
// }
