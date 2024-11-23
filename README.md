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
A code generator is in ```/gen/day_generator.go``` with templates as ```/gen/*.tmpl```.

To use:
```go run ./gen/day_generator.go -year=<year> -day=<day>```

2024 is the default year so to create for day01 use:

```go run ./gen/day_generator.go -day=day01```

Note: if either of the following exist then the generator will fail:
- ```/input/adventofcode<year>/<day>.txt```
- ```/challenge/adventofcode<year>/<day>/```

Add the generated command to the subcommands in the relevant ```/challenge/adventofcode<year>/aoc<year>.go``` file


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