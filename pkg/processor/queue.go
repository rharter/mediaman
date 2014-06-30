package processor

import (
	"log"

	. "github.com/rharter/mediaman/pkg/model"
)

type Queue struct {
	tasks chan<- *FetchMetadataTask
}

type FetchMetadataTask struct {
	Movie   *Movie
	Series  *Series
	Episode *Episode
}

func Start(workers int) *Queue {
	tasks := make(chan *FetchMetadataTask)

	q := &Queue{tasks: tasks}

	log.Printf("Starting queue with %d workers.", workers)

	for i := 0; i < workers; i++ {
		go func(queue <-chan *FetchMetadataTask) {
			var task *FetchMetadataTask
			for {
				// get work item (pointer) from the queue
				task = <-queue
				if task == nil {
					continue
				}

				log.Printf("Received task")

				// Execute the task
				if task.Movie != nil {
					FetchMetadataForMovie(task.Movie)
				} else if task.Series != nil {
					FetchMetadataForSeries(q, task.Series)
				} else if task.Episode != nil {
					FetchMetadataForEpisode(task.Episode)
				}
			}
		}(tasks)
	}

	return q
}

func (q *Queue) Add(task *FetchMetadataTask) {
	q.tasks <- task
}
