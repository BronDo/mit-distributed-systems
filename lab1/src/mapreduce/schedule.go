package mapreduce

import "fmt"

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//
	nworkers := len(mr.workers)
	for i := 0; i < ntasks; i++ {
		select {
		case <-mr.registerChannel:
			nworkers = len(mr.workers)
		default:
			if nworkers <= 0 {
				select {
				case <-mr.registerChannel:
					nworkers = len(mr.workers)
				}
			}
		}

		arg := &DoTaskArgs{
			JobName:       mr.jobName,
			File:          mr.files[i],
			Phase:         phase,
			TaskNumber:    i,
			NumOtherPhase: nios,
		}
		worker := mr.workers[i%nworkers]
		ok := call(worker, "Worker.DoTask", arg, new(struct{}))
		if !ok {
			fmt.Printf("Schedule: %d task to worker %s fail", i, worker)

			//failure solution
			copy(mr.workers[i%nworkers:], mr.workers[i%nworkers+1:])
			nworkers--
			mr.workers = mr.workers[0:nworkers]
			i--
		}
	}

	fmt.Printf("Schedule: %v phase done\n", phase)
}
