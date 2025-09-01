//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func work(w int, i int) {
//	fmt.Println(w, "I work", i)
//}
//
//func worker(f func(w int, i int), queue chan int, i int) {
//	fmt.Println("run worker ", i)
//	for val := range queue {
//		f(i, val)
//		runtime.Gosched()
//	}
//}
//
//func main() {
//	workerNums := 3
//	queue := make(chan int)
//
//	for i := 0; i < workerNums; i++ {
//		go worker(work, queue, i)
//	}
//
//	for i := 0; i < 10; i++ {
//		queue <- i
//	}
//	fmt.Println("Done!")
//}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Воркеры будут получать задания из этого канала
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Имитация работы
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Воркер %d обработал задачу %d\n", id, job)
		results <- job * 2 // Например, удваиваем число
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numWorkers = 4 // Количество воркеров
	const numJobs = 10   // Количество задач

	jobs := make(chan int, numJobs)    // Канал задач
	results := make(chan int, numJobs) // Канал результатов

	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Отправляем задачи
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Закрываем канал задач, чтобы воркеры знали, что новых задач не будет
	// Читаем результаты

	// Ждём завершения всех воркеров
	wg.Wait()
	close(results) // Закрываем канал результатов
	for res := range results {
		fmt.Println("Результат:", res)
	}
}
