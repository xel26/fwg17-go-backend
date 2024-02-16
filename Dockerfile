FROM golang:latest

WORKDIR /app
COPY . .

# label informasi tambahan
LABEL author="Example"
LABEL company="Fazztrack" website="https://www.fazztrack.com"

# dijalankan saat build image
RUN go mod tidy


# add instruction = mengambil file local untuk di copy ke file di dalam docker image. https://pkg.go.dev/path/filepath#Match
# ADD path/local path/dockerImage


EXPOSE 8080

# dijalankan setiap kali docker container di jalankan 
CMD go run .            

# komentar