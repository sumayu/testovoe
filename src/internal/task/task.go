package task

import (
	"context"
	"sync"
	"time"
)

type AnswerJson struct {
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	WorkTime  float64   `json:"work_time"`
}

var (
	currentTask *TaskInstance
	mu          sync.Mutex
	startTime   time.Time
)

type TaskInstance struct {
	output chan *AnswerJson
	cancel context.CancelFunc
}

func Create() *AnswerJson {
	mu.Lock()
	defer mu.Unlock()

	if currentTask != nil {
		return &AnswerJson{
			Status:    "400 task already exists",
			CreatedAt: startTime,
			WorkTime:  time.Since(startTime).Seconds(),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	startTime = time.Now()
	output := make(chan *AnswerJson)
	currentTask = &TaskInstance{
		output: output,
		cancel: cancel,
	}

	go runTask(ctx, output)

	return &AnswerJson{
		Status:    "201 created",
		CreatedAt: startTime,
		WorkTime:  0,
	}
}

func Get() *AnswerJson {
	mu.Lock()
	defer mu.Unlock()

	if currentTask == nil {
		return nil
	}

	select {
	case res := <-currentTask.output:
		return res
	default:
		return &AnswerJson{
			Status:    "202 in progress",
			CreatedAt: startTime,
			WorkTime:  time.Since(startTime).Seconds(),
		}
	}
}

func Delete() *AnswerJson {
	mu.Lock()
	defer mu.Unlock()

	if currentTask == nil {
		return nil
	}

	currentTask.cancel()
	currentTask = nil
	return &AnswerJson{
		Status:    "200 deleted",
		CreatedAt: startTime,
		WorkTime:  time.Since(startTime).Seconds(),
	}
}

func runTask(ctx context.Context, output chan<- *AnswerJson) {
	defer close(output)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			output <- &AnswerJson{
				Status:    "200 complete",
				CreatedAt: startTime,
				WorkTime:  time.Since(startTime).Seconds(),
			}
			return
		case <-ticker.C:
		}
	}
}