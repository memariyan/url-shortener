package worker

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"url-shortner/internal/config"
	"url-shortner/internal/model"
	"url-shortner/internal/repository"
)

type SaveWorkerPool struct {
	Jobs     chan *model.URLData
	poolSize int
	wg       sync.WaitGroup
}

var instance *SaveWorkerPool

func SaveWorker() *SaveWorkerPool {
	if instance == nil {
		bufferSize := 100
		instance = &SaveWorkerPool{
			Jobs:     make(chan *model.URLData, bufferSize),
			poolSize: config.App().Worker.Size,
		}
	}
	return instance
}

func (w *SaveWorkerPool) Start() {
	if w.poolSize <= 0 {
		log.Fatalf("Worker pool size must be greater than 0")
		return
	}
	// dispatching
	for i := range w.poolSize {
		go w.worker(i + 1)
	}
}

func (w *SaveWorkerPool) Stop() {
	close(w.Jobs)
	w.wg.Wait()
}

func (w *SaveWorkerPool) worker(number int) {
	for urlData := range w.Jobs {
		w.wg.Add(1)
		err := repository.URLData().Save(urlData)
		if err != nil {
			log.Errorf("Worker #%d failed to save data: %s", number, err.Error())
		}
		w.wg.Done()
	}
}
