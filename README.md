# TaskRecommenderCLI

A simple command line program that recommend one task at a time

I find myself wasting too much time deciding what to do which make it easier to procrastinate
So I build this death simple CLI program that recommend a single at a time

### Usage
Add task:

$ task add <task description>

List all task:

$ task list

Done and remove a task:

$ task done <task id>

Recommend a task:

default will recommend task which stay longest in the list:

$ task do

or pick a random task:

$ task do -r

### Implementation
Use Cobra to handle command line argument
Use BoltDB for data storage


