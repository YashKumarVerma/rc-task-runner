FROM frolvlad/alpine-gxx

# labels
LABEL maintainer="Yash Kumar Verma yk.verma2000@gmail.com"

# configure prod environment
ENV GIN_MODE=release

# copy project to working directory
WORKDIR /

# take built packge into container
COPY ./build/i-judge ./i-judge

# create directory to store downloaded codes
CMD ["mkdir codes"]

# run the server
CMD ["chmod +x i-judge"]
CMD ["i-judge"]

# exporse listening port
EXPOSE 8000