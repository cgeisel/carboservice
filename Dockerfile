FROM scratch
ADD ./carboservice /carboservice
EXPOSE 8080
CMD ["/carboservice"]
