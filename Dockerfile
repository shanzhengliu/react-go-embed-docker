# node build start
FROM node:20-slim AS NodeBuild
WORKDIR /app
COPY ./frontend-react/ /app
RUN npm install -g pnpm
RUN pnpm install
RUN pnpm build
# node build end

# go build start
FROM golang:1.20.6-alpine3.18 AS BuildStage
WORKDIR /app
COPY . .
COPY --from=NodeBuild /app/dist/ /app/frontend-build
RUN apk --no-cache add upx
RUN go mod download
RUN OS="$(uname | tr '[:upper:]' '[:lower:]')" && \
    ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/' -e 's/armv[0-9]*/arm/')" && \
    GOOS=${OS} GOARCH=${ARCH} go build -o /app/main .
RUN upx /app/main
# go build end

# final image
FROM alpine:latest
WORKDIR /app
COPY --from=GoBuild app/ app/
EXPOSE 8080
ENTRYPOINT [ "app/main" ]