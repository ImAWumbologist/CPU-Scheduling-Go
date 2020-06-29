Simulates First-Come-First-Serve, Preemptive Shortest Job First, and Round Robin scheduling algorithms in Go

FCFS: Allocates each process in order in which they arrive

PSJF: Selects process for execution which has the smallest amount of time remaining until completion

RR: Treats the processes as a ciscular list. After a certain amount of time, it halts the current process and starts the next



Example input:

processcount 2 # Read 2 processes

runfor 15 # Run for 15 time units

use rr # Can be fcfs, sjf, or rr

quantum 2 # Time quantum, only if using rr

process name P1 arrival 3 burst 5

process name P2 arrival 0 burst 9

end



To test, run pa1Test.sh.
This program is currently a work in progress...
