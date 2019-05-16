package utils

// Future boilerplate method
func Future(f func() (interface{}, error)) func() (interface{}, error) {
	var result interface{}
	var err error

	c := make(chan struct{}, 1)
	// quit := make(chan bool)
	go func() {
		defer close(c)
		result, err = f()
		// for {
		// 	select {
		// 	case <-time.After(3 * time.Second):
		// 		fmt.Println("timeout -----")
		// 		quit <- true
		// 	}
		// }
	}()

	return func() (interface{}, error) {
		<-c
		return result, err
	}
}
