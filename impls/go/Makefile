APP := step0_repl.exe
SRC := src/step0_repl.go

.PHONY: build run clean

build:
	go build -o $(APP) $(SRC)

run: build
	./$(APP)

clean:
	rm -f $(APP)
