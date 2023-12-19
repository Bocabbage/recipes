package main

import (
	gbasic "bocabbage/concurrency-learn/goroutine_basic"
	mmodel "bocabbage/concurrency-learn/memory_model"
)

func main() {
	mmodel.AtomicTestMutex()
	// mmodel.SequenceConsistency()
	// gbasic.SpinnerTest()
	gbasic.PipelineTest()
}
