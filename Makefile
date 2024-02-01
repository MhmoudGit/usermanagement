.PHONY: run gen

run:
	@go build -o ./tmp/main.exe ./main.go

gen: 
	@templ generate
