FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum continentes/names.txt ./

RUN go mod download && go mod verify

CMD ["go", "run", "dataNodes/datanode1/datanode1.go"]
#CMD ["go", "run", "DataNodes/datanode2.go"]

#CMD ["go", "run", "OMS/oms.go"]
#CMD ["go", "run", "ONU/onu.go"]

#CMD ["go", "run", "Continentes/asia/asia.go"]
# CMD ["go", "run", "Continentes/australia/australia.go"]
# CMD ["go", "run", "Continentes/europa/europa.go"]
# CMD ["go", "run", "Continentes/latinoamerica/latinoamerica.go"]