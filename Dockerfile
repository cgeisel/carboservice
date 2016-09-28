FROM scratch
ADD bin/linux_amd64/carboservice /carboservice
EXPOSE 8080
CMD ["/carboservice"]
