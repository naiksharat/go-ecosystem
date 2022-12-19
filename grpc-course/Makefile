calculator: clean_calc
	protoc -I calculator/proto --go_opt=module=grpc-course --go_out=. calculator/proto/* --go-grpc_opt=module=grpc-course --go-grpc_out=. calculator/proto/*
	go build -o calculator/bin/server calculator/server/*
	go build -o calculator/bin/client calculator/client/*

clean_calc:
	rm calculator/proto/*.go

run_calc:
	calculator/bin/server &
	calculator/bin/client

kill_calc:
	ps | grep "calculator/bin/server" | grep -v grep | awk '{print $1}' | xargs kill -9

#######################################################################################################

greet: clean_greet
	protoc -I greet/proto --go_opt=module=grpc-course --go_out=. greet/proto/* --go-grpc_opt=module=grpc-course --go-grpc_out=. greet/proto/*
	go build -o greet/bin/server greet/server/*
	go build -o greet/bin/client greet/client/*

clean_greet:
	@if ls greet/proto/*.go 1> /dev/null 2>&1; then\
		rm greet/proto/*.go;\
	fi


run_greet:
	greet/bin/server &
	greet/bin/client

kill_greet:
	ps | grep "greet/bin/server" | grep -v grep | awk '{print $1}' | xargs kill -9

####################################################################################################

blog: clean_blog
	protoc -I blog/proto --go_opt=module=grpc-course --go_out=. blog/proto/* --go-grpc_opt=module=grpc-course --go-grpc_out=. blog/proto/*
	go build -o blog/bin/server blog/server/*
	go build -o blog/bin/client blog/client/*

clean_blog:
	@if ls blog/proto/*.go 1> /dev/null 2>&1; then\
		rm blog/proto/*.go;\
	fi


run_blog:
	blog/bin/server &
	blog/bin/client

kill_blog:
	ps | grep "blog/bin/server" | grep -v grep | awk '{print $1}' | xargs kill -9



.PHONY: calculator greet help blog
