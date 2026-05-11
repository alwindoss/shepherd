# https://github.com/GoogleContainerTools/distroless
# Stage 1: Build the VueJS project
FROM node:trixie AS ui-builder
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

FROM golang:1.26.2-bookworm AS go-builder
WORKDIR /go/src/app
COPY . .
RUN rm -rf web
COPY --from=ui-builder /app/public ./public 
RUN make setup
RUN make generate
RUN make build-backend

FROM gcr.io/distroless/static-debian13
WORKDIR /root/
COPY --from=go-builder /go/src/app/bin/shepherd .
CMD [ "./shepherd" ]