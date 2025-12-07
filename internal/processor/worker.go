/*
  Filename: worker.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Implements the worker pool for concurrent file processing.
*/

package processor

import (
	"sync"

	"github.com/infologicmgmt/outgen/internal/log"
)

// worker is a function that processes jobs from a channel.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		err := job.Process()
		results <- Result{Job: job, Error: err}
	}
}

// Run starts the worker pool and processes the jobs.
func Run(jobs []Job, numWorkers int) {
	jobChan := make(chan Job, len(jobs))
	resultChan := make(chan Result, len(jobs))

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobChan, resultChan, &wg)
	}

	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		if result.Error != nil {
			log.Logger.Error().Err(result.Error).Str("inputFile", result.Job.InputFile).Msg("Failed to process file")
		}
	}
}
