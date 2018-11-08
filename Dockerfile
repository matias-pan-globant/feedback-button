FROM alpine

COPY feedback /app/feedback

EXPOSE 8000

ENTRYPOINT ["/app/feedback"]
