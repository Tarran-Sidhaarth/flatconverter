# Build stage
FROM golang:tip-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o buffman .

# Final stage
FROM alpine:latest

# Arguments for user and group IDs
ARG USERNAME=buffman
ARG USER_UID=1000
ARG USER_GID=1000

# Install dependencies and tools
RUN apk add --no-cache \
    ca-certificates \
    wget \
    unzip \
    libc6-compat \
    bash \
    shadow

# Create group and user with matching UID/GID
RUN groupadd -g $USER_GID $USERNAME && \
    useradd -u $USER_UID -g $USER_GID -m $USERNAME

# Download and install flatc binary
RUN wget -O /tmp/flatc.zip https://github.com/google/flatbuffers/releases/download/v25.2.10/Linux.flatc.binary.g++-13.zip && \
    unzip /tmp/flatc.zip -d /tmp && \
    mv /tmp/flatc /usr/local/bin/flatc && \
    chmod +x /usr/local/bin/flatc && \
    rm -rf /tmp/flatc.zip

# Copy the binary from builder stage
COPY --from=builder /app/buffman /usr/local/bin/buffman

# Set working directory and permissions
WORKDIR /buffman
RUN chown -R $USER_UID:$USER_GID /buffman

# Switch to non-root user
USER $USERNAME

# Set the entrypoint
ENTRYPOINT ["buffman"]
