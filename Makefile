docker-ONU:
	docker run -e PORT=50053 -p 50053:50053 go-onu:latest

docker-OMS:
	docker run -e PORT=50050 -p 50050:50050 go-oms:latest

docker-datanode:
	docker run -e PORT=50051 -p 50051:50051 go-datanode:latest

docker-continentes:
	docker run -e PORT=50050 -p 50050:50050 go-continentes:latest