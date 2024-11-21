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

## Usage example
```go run main.go <year> <day>```

```go run main.go 2024 day01```

## adventofcode2023
The first few days of 2023 were created in prep for 2024.