package handler_home

import (
	"log"
	"sync"
	"time"

	container_home "ms-genexis-pos-operaciones/context/home/presentation/container"
)

var pendingTransmissionsSchedulerOnce sync.Once

func StartPendingTransmissionsScheduler() {
	pendingTransmissionsSchedulerOnce.Do(func() {
		go func() {
			ticker := time.NewTicker(time.Minute)
			defer ticker.Stop()

			runPendingTransmissionsJob()
			for range ticker.C {
				runPendingTransmissionsJob()
			}
		}()
	})
}

func runPendingTransmissionsJob() {
	start := time.Now()
	response, err := container_home.ResolveProcessPendingTransmissionsContainer().Execute()
	if err != nil {
		log.Printf("[PendingTransmissionsScheduler] error: %v", err)
		return
	}

	var processed, synchronized, failed int
	if response != nil && response.Data != nil {
		processed = response.Data.Processed
		synchronized = response.Data.Synchronized
		failed = response.Data.Failed
	}

	log.Printf(
		"[PendingTransmissionsScheduler] processed=%d synchronized=%d failed=%d duration=%s",
		processed,
		synchronized,
		failed,
		time.Since(start),
	)
}
