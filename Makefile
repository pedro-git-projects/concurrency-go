run:
	go run ./src/*.go 

catfiles:
	go run ./src/*.go ./src/exercises/ch2/assets/txt1 ./src/exercises/ch2/assets/txt2 ./src/exercises/ch2/assets/txt3 

grepfiles:
	go run ./src/*.go illness ./src/exercises/ch2/assets/txt1 ./src/exercises/ch2/assets/txt2 ./src/exercises/ch2/assets/txt3 

grepdir: 
	go run ./src/*.go txt1 ./src/exercises/ch2/assets/


