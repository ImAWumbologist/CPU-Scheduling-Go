package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// struct to hold all my info
type myStruct struct {
	ID int
	processName string
	arrival     int
	burst       int
	ogBurst int
	hasArrived bool
	finished bool
	selected bool
	wait int
	turnAround int
}

func main() {
	// This is all for the file reading/writing
	input := os.Args[1]
	output := os.Args[2]
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// We use an array of structs for this
	var processes []myStruct
	var tempStruct myStruct

	var processCount int
	var runfor int
	var use string
	var quantum int

	// Scan in and store the number of processes, number of unites to run for, and the scheduling algorithm (and quantum if applicable)
	for scanner.Scan() {
		temp := scanner.Text()

		if (temp == "processcount") {
			scanner.Scan()
			processCount, _ = strconv.Atoi(scanner.Text())
		}

		if (temp == "runfor") {
			scanner.Scan()
			runfor, _ = strconv.Atoi(scanner.Text())
		}

		if (temp == "use") {
			scanner.Scan()
			use = scanner.Text()
			if (use != "rr") {
				break
			}
		}

		if (temp == "quantum") {
			scanner.Scan()
			quantum, _ = strconv.Atoi(scanner.Text())
			break
		}
	}

	// scan in all processes and initialize relevant variables
	for scanner.Scan() {
		temp := scanner.Text()

		if(temp == "end") {
			break
		}

		tempStruct.hasArrived = false
		tempStruct.finished = false
		tempStruct.selected = false
		tempStruct.wait = 0
		tempStruct.turnAround = 0

		if(temp == "name") {
			scanner.Scan()
			tempStruct.processName = scanner.Text()
			tempStruct.ID = assignID(tempStruct.processName)
		}

		if(temp == "arrival") {
			scanner.Scan()
			tempStruct.arrival, _ = strconv.Atoi(scanner.Text())
		}
		 
		if(temp == "burst") {
			scanner.Scan()
			tempStruct.burst, _ = strconv.Atoi(scanner.Text())
			tempStruct.ogBurst = tempStruct.burst
			processes = append(processes, tempStruct)
		}
	}

	// schedule stuff
	if(use == "fcfs") {
		fcfs(processes, processCount, runfor, output)
	} else if(use == "sjf") {
		sjf(processes, processCount, runfor, output)
	} else if(use == "rr") {
		rr(processes, processCount, runfor, quantum, output)
	}	
}

// Terribly hardcoded function to assign a process ID for later sorting after scheduling completed
func assignID(processName string) (ID int) {
	if(processName == "P1" || processName == "P01") {
		ID = 1
	} else if (processName == "P2" || processName == "P02") {
		ID = 2
	} else if (processName == "P3" || processName == "P03") {
		ID = 3
	} else if (processName == "P4" || processName == "P04") {
		ID = 4
	} else if (processName == "P5" || processName == "P05") {
		ID = 5
	} else if (processName == "P6" || processName == "P06") {
		ID = 6
	} else if (processName == "P7" || processName == "P07") {
		ID = 7
	} else if (processName == "P8" || processName == "P08") {
		ID = 8
	} else if (processName == "P9" || processName == "P09") {
		ID = 9
	} else if (processName == "P10") {
		ID = 10
	}
	return ID
}

// Sorts by arrival time using bubble sort
func sortByArrival(processes []myStruct) (sorted []myStruct) {
	length := len(processes)
	sorted = make([]myStruct, length)
	sorted = processes

	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if(sorted[j].arrival > sorted[j+1].arrival) {
				temp := sorted[j]
				sorted[j] = sorted[j+1]
				sorted[j+1] = temp
			}
		}
	}
	return sorted
}

// Bubble Sort to sort by ID
func sortByID(processes []myStruct) (sorted []myStruct) {
	length := len(processes)
	sorted = make([]myStruct, length)
	sorted = processes

	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if(sorted[j].ID > sorted[j+1].ID) {
				temp := sorted[j]
				sorted[j] = sorted[j+1]
				sorted[j+1] = temp
			}
		}
	}
	return sorted
}

// Bubble sort for sorting by burst
func sortByBurst(processes []myStruct) (sorted []myStruct) {
	length := len(processes)
	sorted = make([]myStruct, length)
	sorted = processes

	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if(sorted[j].burst > sorted[j+1].burst) {
				temp := sorted[j]
				sorted[j] = sorted[j+1]
				sorted[j+1] = temp
			}
		}
	}
	return sorted
}

func fcfs(processes []myStruct, processCount int, runfor int, output string) {
	// Print everything to a file
	outputFile, _ := os.Create(output)
	
	// Print out that we are using fcfs
	fmt.Fprintf(outputFile, "%3d processes\n", processCount)
	fmt.Fprintf(outputFile, "Using First-Come First-Served\n")
	
	// Sort the list by arrival time
	processes = sortByArrival(processes)

	// Length of our array
	length := len(processes)

	// This keeps track of which process we are on
	numArrived := 0
	selected := 0
	processes[selected].selected = true

	// Big loop to keep track of unit of time we are on
	for time := 0; time < runfor; time++ {
		
		if(selected == length && processes[selected-1].finished == true) {
			fmt.Fprintf(outputFile, "Time %3d : %s finished\n", time, processes[selected-1].processName)
			processes[selected-1].finished = false
			processes[selected-1].turnAround = time - processes[selected-1].arrival
			processes[selected-1].wait = processes[selected-1].turnAround - processes[selected-1].ogBurst
		}

		if(selected == length) {
			fmt.Fprintf(outputFile, "Time %3d : Idle\n", time)
			continue
		}

		// This loop simply allows us to keep track of when a process arrives
		for i := 0; i < length; i++ {
			if(processes[i].arrival == time) {
				fmt.Fprintf(outputFile, "Time %3d : %s arrived\n", time, processes[i].processName)
				processes[i].hasArrived = true
				numArrived++
			}
		}

		if(selected != 0 && processes[selected-1].finished == true) {
			fmt.Fprintf(outputFile, "Time %3d : %s finished\n", time, processes[selected-1].processName)
			numArrived--
			processes[selected-1].finished = false
			processes[selected-1].turnAround = time - processes[selected-1].arrival
			processes[selected-1].wait = processes[selected-1].turnAround - processes[selected-1].ogBurst
		}

		if(numArrived == 0) {
			fmt.Fprintf(outputFile, "Time %3d : Idle\n", time)
			continue
		}
		
		if(processes[selected].hasArrived == true && processes[selected].selected == true) {
			fmt.Fprintf(outputFile, "Time %3d : %s selected (burst %3d)\n", time, processes[selected].processName, processes[selected].burst)
			processes[selected].selected = false
		}

		if(processes[selected].burst > 0) {
			processes[selected].burst--
		}

		if(processes[selected].burst == 0) {
			processes[selected].finished = true
			selected++
			if(selected != length) {
				processes[selected].selected = true
			}
		}
	}

	fmt.Fprintf(outputFile, "Finished at time %3d\n", runfor)
	fmt.Fprintf(outputFile, "\n")

	processes = sortByID(processes)

	for i := 0; i < length; i++ {
		fmt.Fprintf(outputFile, "%s wait %3d turnaround %3d\n", processes[i].processName, processes[i].wait, processes[i].turnAround)
	}
}

func sjf(processes []myStruct, processCount int, runfor int, output string) {
	// Print everything to a file
	outputFile, _ := os.Create(output)

	// Print out that we are using fcfs
	fmt.Fprintf(outputFile, "%3d processes\n", processCount)
	fmt.Fprintf(outputFile, "Using preemptive Shortest Job First\n")

	// Length of our array
	length := len(processes)

	// We need to use a queue for this along with an array to keep track of done
	var qq []myStruct
	var done []myStruct

	// This keeps track of which process we are on
	numArrived := 0

	// Big loop to keep track of unit of time we are on
	for time := 0; time < runfor; time++ {
		

		// This loop simply allows us to keep track of when a process arrives
		for i := 0; i < length; i++ {
			if(processes[i].arrival == time) {
				fmt.Fprintf(outputFile, "Time %3d : %s arrived\n", time, processes[i].processName)
				processes[i].hasArrived = true
				numArrived++
				qq = append(qq, processes[i])
			}
		}

		if(len(done) == 1) {
			fmt.Fprintf(outputFile, "Time %3d : %s finished\n", time, done[0].processName)

			for j := 0; j < length; j++ {
				if(done[0].ID == processes[j].ID) {
					processes[j].turnAround = time - processes[j].arrival
					processes[j].wait = processes[j].turnAround - processes[j].ogBurst
				}
			}

			done = append(done[:0])
		}

		if(numArrived == 0) {
			fmt.Fprintf(outputFile, "Time %3d : Idle\n", time)
		}

		// Double check if qq is empty
		// Sort it by burst then select the shortest one 
		if(len(qq) != 0) {	
			qq = sortByBurst(qq)

			for i := 1; i < len(qq); i++ {
				qq[i].selected = false
			}

			if(qq[0].selected == false) {
				fmt.Fprintf(outputFile, "Time %3d : %s selected (burst %3d)\n", time, qq[0].processName, qq[0].burst)
				qq[0].selected = true
			}

			if(len(qq) > 0) {
				qq[0].burst--
			}
			
			if(qq[0].burst == 0) {
				numArrived--
				qq[0].finished = true
				if(len(qq) > 1) {
					done = append(done, qq[0])
					qq = append(qq[:0], qq[1:]...)
				} else {
					done = append(done, qq[0])
					qq = append(qq[:0])
				}
			}	

			
		}	
	}

	fmt.Fprintf(outputFile, "Finished at time %3d\n", runfor)
	fmt.Fprintf(outputFile, "\n")

	processes = sortByID(processes)

	for i := 0; i < length; i++ {
		fmt.Fprintf(outputFile, "%s wait %3d turnaround %3d\n", processes[i].processName, processes[i].wait, processes[i].turnAround)
	}
}

func rr(processes []myStruct, processCount int, runfor int, quantum int, output string) {
	// Print everything to a file
	outputFile, _ := os.Create(output)

	// Print out that we are using fcfs
	fmt.Fprintf(outputFile, "%3d processes\n", processCount)
	fmt.Fprintf(outputFile, "Using Round-Robin\n")
	fmt.Fprintf(outputFile, "Quantum %d\n", quantum)

	// additional blank line for formatting
	fmt.Fprintf(outputFile, "\n")

	// Length of our array
	length := len(processes)

	// We need to use a queue for this along with an array to keep track of done
	var qq []myStruct
	var done []myStruct

	// This keeps track of which process we are on
	numArrived := 0
	selected := 0
	counter := 0

	// Big loop to keep track of unit of time we are on
	for time := 0; time < runfor; time++ {
		
		// This loop simply allows us to keep track of when a process arrives
		for i := 0; i < length; i++ {
			if(processes[i].arrival == time) {
				fmt.Fprintf(outputFile, "Time %3d : %s arrived\n", time, processes[i].processName)
				processes[i].hasArrived = true
				numArrived++
				qq = append(qq, processes[i])
			}
		}

		if(len(done) == 1) {
			fmt.Fprintf(outputFile, "Time %3d : %s finished\n", time, done[0].processName)

			for j := 0; j < length; j++ {
				if(done[0].ID == processes[j].ID) {
					processes[j].turnAround = time - processes[j].arrival
					processes[j].wait = processes[j].turnAround - processes[j].ogBurst
				}
			}

			done = append(done[:0])
		}

		if(numArrived == 0) {
			fmt.Fprintf(outputFile, "Time %3d : Idle\n", time)
		}

		// Double check if qq is empty
		// Sort it by burst then select the shortest one 
		if(len(qq) != 0) {
			selected = selected%len(qq)

			if(qq[selected].selected == false) {
				fmt.Fprintf(outputFile, "Time %3d : %s selected (burst %3d)\n", time, qq[selected].processName, qq[selected].burst)
				qq[selected].selected = true
			}

			if(len(qq) > 0) {
				qq[selected].burst--
				counter++
				if(counter == quantum) {
					qq[selected].selected = false
				}
			}
			
			if(qq[selected].burst == 0) {
				numArrived--
				counter = 0
				qq[selected].selected = false
				qq[selected].finished = true
				if(len(qq) > 1) {
					done = append(done, qq[selected])
					qq = append(qq[:selected], qq[selected+1:]...)
				} else {
					done = append(done, qq[selected])
					qq = append(qq[:selected])
				}
			}
			
			if(counter == 0) {
				continue
			}	

			if(counter == quantum) {
				counter = 0
				selected++
			}								
		}	
	}

	fmt.Fprintf(outputFile, "Finished at time %3d\n", runfor)
	fmt.Fprintf(outputFile, "\n")

	processes = sortByID(processes)

	for i := 0; i < length; i++ {
		fmt.Fprintf(outputFile, "%s wait %3d turnaround %3d\n", processes[i].processName, processes[i].wait, processes[i].turnAround)
	}
}