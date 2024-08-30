package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Task 表示一个可执行的任务
type Task func()

// Pool 表示线程池
type Pool struct {
	workers    sync.WaitGroup
	jobs       chan Task
	stop       chan struct{}
	numWorkers int
}

// NewPool 创建一个新的线程池
func NewPool(numWorkers int) *Pool {
	return &Pool{
		workers:    sync.WaitGroup{},
		jobs:       make(chan Task),
		stop:       make(chan struct{}),
		numWorkers: numWorkers,
	}
}

// Start 启动线程池
func (p *Pool) Start() {
	for i := 0; i < p.numWorkers; i++ {
		go p.worker()
	}
}

// worker 线程池中的工作线程
func (p *Pool) worker() {
	for {
		select {
		case job := <-p.jobs:
			job()
			p.workers.Done()
		case <-p.stop:
			return
		}
	}
}

// Stop 停止线程池
func (p *Pool) Stop() {
	close(p.stop)
	p.workers.Wait()
}

// AddJob 向线程池添加任务
func (p *Pool) AddJob(job Task) {
	p.workers.Add(1)
	p.jobs <- job
}

func TestThreadPool(t *testing.T) {
	pool := NewPool(3)
	pool.Start()

	// 添加测试任务
	for i := 0; i < 10; i++ {
		task := fmt.Sprintf("test_task_%d", i)
		pool.AddJob(func() {
			fmt.Println("Executing test:", task)
			// 这里可以添加测试逻辑
		})
	}

	// 停止线程池
	pool.Stop()
}

func main() {
	// 创建一个有3个工作线程的线程池
	pool := NewPool(3)
	pool.Start()

	// 添加任务到线程池
	for i := 0; i < 10; i++ {
		task := fmt.Sprintf("task_%d", i)
		pool.AddJob(func() {
			fmt.Println("Executing:", task)
			time.Sleep(1 * time.Second) // 模拟任务执行时间
		})
	}

	// 等待所有任务完成
	pool.Stop()
	fmt.Println("All tasks completed.")
}
