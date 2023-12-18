package main

import (
	mmodel "bocabbage/concurrency-learn/memory_model"
)

func main() {
	mmodel.AtomicTestMutex()
	mmodel.SequenceConsistency()
}
