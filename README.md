# adventofcode/2024
This is a repo for [advent of code 2024](https://adventofcode.com/2024) solutions in Go.

Utilises [cobra](https://github.com/spf13/cobra) to create a cli.

## Structure
Solutions and tests are in ```/challenge/adventofcode2024/<day>/``` with
* ```<day>.go```
* ```<day>_test.go```
* ```<day>_cmd.go```

Input files are in ```/input/adventofcode2024/<day>.txt```

### Setup
```/main.go``` uses ```/cmd/cmd.go``` to create the cli and include each year (e.g. ```/challenge/adventofcode2024/aoc2024.go```) as a command.

The year command adds each day (e.g. ```/challenge/adventofcode2024/day01/day01_cmd.go```) as a subcommand.

### Prep for a new day
- Create a new directory inside the year. e.g. ```/challenge/adventofcode2024/day01```
- Copy the files from ```/challenge/template``` into that directory
- rename from ```<day_template>``` to ```day<number>``` e.g. ```day01.go```, ```day01_cmd.go```, and ```day01_test.go```
- change the package in those files to be ```day<number>``` e.g. ```day01```
- rename the ```newDayTemplate``` function in ```day01.go``` to ```newDay<number>```. e.g. ```newDay01```
- Create a text file for input. e.g. ```/input/adventofcode2024/day01.txt``` and copy your input in here
- Modify the  command e.g.```day01.go``` to point to the input file
- Add the command to the ```subCommands``` in the year (e.g. ```/challenge/adventofcode2024/aoc2024.go```)
- Remove the ```t.Skip``` line from the test
- Tests are table driven so insert the appropriate example data into them
- Solve the challenge!


## Usage example
### Single solution
```go run main.go <year> <day>```

```go run main.go 2024 day01```
### All solutions for a year
```go run main.go <year>```

```go run main.go 2024```

### All solutions
```go run main.go```

## adventofcode2023
The first few days of 2023 were created in prep for 2024.