FROM bankpossible/shared:latest

# Add Gopath to environment, and bin dir for generated go binaries
ENV GOPATH /code
ENV PATH /code/bin:/usr/src/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games

ADD . /code/src/github.com/bankpossible/iamdev/phosphord
WORKDIR /code/src/github.com/bankpossible/iamdev/phosphord
RUN go get -v
