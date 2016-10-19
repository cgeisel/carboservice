FROM scratch
COPY carboservice /carboservice
EXPOSE 8080
CMD ["/carboservice"]
