package main

import (
	ctxusage "bocabbage/concurrency-learn/context_basic"
	examples "bocabbage/concurrency-learn/examples"
	gbasic "bocabbage/concurrency-learn/goroutine_basic"
	mmodel "bocabbage/concurrency-learn/memory_model"
)

func main() {
	mmodel.AtomicTestMutex()
	// mmodel.SequenceConsistency()
	// gbasic.SpinnerTest()
	// gbasic.PipelineTestV2()
	gbasic.WaitRoutineTest()
	ctxusage.CancelContextExample()
	examples.ChatServerMain() // Good example for CSP model
}
