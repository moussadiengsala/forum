build:
	docker build --tag forum .
run: build 
	docker run -d --name forum -p 8080:8080 forum
stop:
	docker stop forum
	docker container rm forum
