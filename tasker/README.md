# Building a "Task Tracker" CLI Using Cobra
Cobra is a library for creating powerful, modern CLI (Command Line Interface) applications in Go
We will build a simple CLI tool called tasker. It will have commands to add tasks and list them.

/tasker
  ├── go.mod
  ├── main.go         (Entry point)
  └── /cmd
       ├── root.go    (The base command, e.g., just 'tasker')
       ├── add.go     (The 'add' sub-command)
       └── list.go    (The 'list' sub-command)