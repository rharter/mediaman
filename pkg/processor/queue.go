package processor

import (
	"log"
	. "github.com/rharter/mediaman/pkg/model"
)

type Queue struct {
	tasks chan<- *FetchMetadataTask
}

type FetchMetadataTask struct {
	Movie *Movie
}

func Start(workers int) *Queue {
	tasks := make(chan *FetchMetadataTask)

	queue := &Queue{tasks: tasks}

	log.Printf("Starting queue with %d workers.", workers)

	for i := 0; i < workers; i++ {
		go func(queue <-chan *FetchMetadataTask) {
			var task *FetchMetadataTask
			for {
				// get work item (pointer) from the queue
				task = <- queue
				if task == nil {
					continue
				}

				log.Printf("Received task")

				// Execute the task
				FetchMetadataForMovie(task.Movie)
			}
		}(tasks)
	}

	return queue
}

func (q *Queue) Add(task *FetchMetadataTask) {
	q.tasks <- task
}