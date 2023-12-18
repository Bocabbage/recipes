package memorymodel

func SequenceConsistency() {
	done := make(chan int)

	go func() {
		println("你好, 世界")
		done <- 1
	}()

	<-done
	println("搁这等着呢.jpg")
}
