FROM alpine
ADD agent agent
EXPOSE 10700
ENTRYPOINT ["./agent"]