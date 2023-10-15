package main

import (
	"os"
	"runtime/pprof"
	"runtime"
	"fmt"
	"parsermods/enronmailparser"
	"parsermods/email"
	"parsermods/zincsearchparser"
)

func cpuTracking() func() {
	cpuf, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(cpuf)
	return func() {
		pprof.StopCPUProfile()
		cpuf.Close()
	}
}

func memorySnapshot(id string) {
	memf, err := os.Create(fmt.Sprintf("memory_%s.pprof", id))
	if err != nil {
		panic(err)
	}
	if err := pprof.WriteHeapProfile(memf); err != nil {
		memf.Close()
		panic(err)
	}
	memf.Close()
}

func resetEmailsSlice(emails *[]*email.Email) {
	emails = nil
}

func main() {
	fmt.Println(runtime.NumCPU())

	stopCPUProfile := cpuTracking()
	defer stopCPUProfile()
	emails := make([]*email.Email, 517549+1)
	enronPath := "../data/maildir"
	enronmailparser.LoadEnronEmailsIntoMemory(enronPath, emails)
	memorySnapshot("post_load")
	zincsearchparser.WriteEmailsToIndexerFile(&emails, "../data/enron_indexer_data.ndjson")
	resetEmailsSlice(&emails)
	runtime.GC()
	memorySnapshot("post_write")
	zincsearchparser.SendRequestToIndexer("../data/enron_indexer_data.ndjson")
}