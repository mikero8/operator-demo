FROM scratch
COPY rest-api /rest-api
EXPOSE 3333
CMD ["/rest-api"]