# fix this later: FROM scratch
FROM golang
COPY carboservice /carboservice
EXPOSE 8080
CMD ["/carboservice"]
