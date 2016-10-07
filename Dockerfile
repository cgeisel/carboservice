FROM scratch
ADD carboservice /
EXPOSE 8080
CMD ["/carboservice"]
