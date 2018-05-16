all: x

x: Test.g4 main.go
	java -cp "/Users/shiquan/bin/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -o parser Test.g4
	go build

clean:
	go clean
	rm -rf parser

