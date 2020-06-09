package main
import(
	"fmt"
	"sync"
	"runtime"
)
func main(){
	var wg sync.WaitGroup
	fmt.Println("No of CPU",runtime.NumCPU())
	fmt.Println("No of Goroutine",runtime.NumGoroutine())
	wg.Add(2)
    go func(){
		fmt.Println("I am func one")
		fmt.Println("No of Goroutine in func one",runtime.NumGoroutine())
		wg.Done()
	}()
	fmt.Println("No of Goroutine in main",runtime.NumGoroutine())

    go func(){
		fmt.Println("I am func two")
		fmt.Println("No of Goroutine in func two",runtime.NumGoroutine())
		wg.Done()
	}()
	fmt.Println("No of Goroutine in main",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("No of Goroutine",runtime.NumGoroutine())
    fmt.Println("I exit from main")
}
