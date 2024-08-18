# runlin
## A Go based programm making wine easier and helping installing dependencies, benchmarking tool etc.

    go run runlin.go install 
    go run runlin.go wine programme.exe
    go run runlin.go proton programme.exe
    go run detect.go programme.exe
    go run benchmark.go programme.exe
    go run runlin.go optimize programme.exe

## WINE-FINDER (easily find wine executables)

    go run finder.go list
    go run finder.go grep <search term>

### For now it only works on arch linux since it uses some aur packages. BTW. YAY IS REQUIRED!

### DON'T JUDGE MY MESSI CODE :) I am still learning go!
### What I am basically trynna to do is making the process of installing needed wine dependencies easier and launching them faster (no GUI planned)

