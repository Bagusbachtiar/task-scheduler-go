package main 

import (
	"log"
	"time"
)

type ScheduleTask struct { 
	ID		  	int
	PhoneNumber string
	Message     string
	ExecutionTZ time.Time //the exact time this task should fire
	IsCompleted bool
}

func main() { 
	log.Println("⏰ Task Scheduler Engine Initializing...")

	// 1.0 Create an in memory Mock Database containing 3 Scheduled tasks 
	//We will schedule them relatively to the current time right now 
	now := time.Now()

	taskDatabase := []ScheduleTask{ 
		{
				ID: 101,
				PhoneNumber: "62121212121",
				Message:  " [ALERT] : System health check verification reminder.",
				ExecutionTZ: now.Add(4 * time.Second), // 4 seconds from now
				IsCompleted: false,
		},
				{
				ID: 102,
				PhoneNumber: "62133333331",
				Message:  " [FINANCE] : Inbound subscription renewal invoice notice.",
				ExecutionTZ: now.Add(10 * time.Second), // 10 seconds from now
				IsCompleted: false,
		},
				{
				ID: 103,
				PhoneNumber: "644444444444",
				Message:  " [GENERAL] : Routine data cache collection update packet.",
				ExecutionTZ: now.Add(20 * time.Second), // 20 seconds from now
				IsCompleted: false,
		},

	} 

	// 1.1 Create a native Go Ticker that triggers a pulse every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() //Cean up the ticker when the program exits

	log.Println("🚀 Time-Driven Engine Online. Scanning for pending records.")

	//2. an infinite lopp that blocks and listens to the tickers internal channel 
	// every time 2 second passes, the ticker sends a message through 'ticker.C'
	for pulseTime := range ticker.C {
		log.Printf("\n⏱️ [Heartbeat Clock tick]: %v", pulseTime.Format("15:04:05"))

		//2.1 Loop through our mock database array to scan for pending items
		for i := 0; i < len(taskDatabase); i++ {
		task := &taskDatabase[i] //fetch a memmory pointer to modify state fields

		//if the task is already finished, skip past it entirely
		if task.IsCompleted {
			continue
		}
		//Core Scheduling Condition: Check if the System clock has crossed the task exectuion deadline
		if pulseTime.After(task.ExecutionTZ) {
			log.Printf("🔥 [Task DUE]: Exectuion condition met for job #%d to %s", task.ID, task.PhoneNumber)

			//simulate exectuing the work action
			log.Printf("✉️ [Transmitting Message]: \"%s\"", task.Message)
			
			//Mark as complete so it never fires again(DB state mutation)
			task.IsCompleted = true
			log.Printf("✅ [State Updated]: Job #%d succesfully marked COMPLETED in memmory.", task.ID)
		}

	  }
	}

}