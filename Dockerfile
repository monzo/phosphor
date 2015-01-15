FROM bankpossible/shared:latest

# Add Gopath to environment, and bin dir for generated go binaries
ENV GOPATH /code
ENV PATH /code/bin:/usr/src/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games

# Our precompiler will build the binary and move it into position before we
# start the container. We can then add this into the container.
ADD ./_workspace /code

# Execute our precompiled binary!
RUN /code/bin/phosphord
