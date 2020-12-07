FROM frolvlad/alpine-gxx

# labels
LABEL maintainer="Yash Kumar Verma yk.verma2000@gmail.com"

# configure prod environment
ENV GIN_MODE=release

# copy project to working directory
WORKDIR /

# take built packge into container
COPY ./build/task-runner ./task-runner

# create directory to store downloaded codes
RUN mkdir codes

# grant permissions to execute
RUN chmod +x task-runner

# launch server
CMD ["./task-runner"]

# exporse listening port
EXPOSE 8000