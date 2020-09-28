package main

func main() {
	worker_threads := NewWorkers(4)
	event_loop := NewWorkers(1)

  worker_threads.Post({
    // Task A
    fmt.Println("Hello from A")
  })

  worker_threads.Post({
    // Task B
    // Might run in parallel with task A
  })

  event_loop.Post({
    // Task C
    // Might run in parallel with task A and B
  })

  event_loop.Post({
    // Task D
    // Will run after task C
    // Might run in parallel with task A and B
  })

}
