package workpools

//var Jobs chan func()

type Jobs struct {
	Job chan func()
	Result chan int
}

func (jobs *Jobs) StartPools(num int) {

	for i := 0; i < num; i++ {
		go func() {
			for f := range jobs.Job{
				f()
			}
		}()
	}
}