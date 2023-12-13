# Go NFL cli

## Description

This is a fun project I started to familiarize myself with golang. It's a command line interface that (so far) pulls up NFL team standings.  I decided to include my passion for sports as the driver for this app. 

## Usage
### Commands available
```golang
nfl <command>

Available Commands:
  help        Help about any command
  standings   NFL standings
```

### Command `standings`
Get current standings of current season
```
Usage:
  nfl standings [flags]

Flags:
  -h, --help       help for standings
  -p, --playoffs   Shows playoff standings

# returns
Getting NFL standings

AFC East
-----------
| Miami Dolphins 9-3-0
| Buffalo Bills 6-6-0
| New York Jets 4-8-0
| New England Patriots 3-10-0

# --playoffs returns
Getting NFL standings

AFC Conference
1. Baltimore Ravens
2. Miami Dolphins
3. Kansas City Chiefs
4. Jacksonville Jaguars
5. Cleveland Browns
6. Pittsburgh Steelers
7. Indianapolis Colts

...

```
