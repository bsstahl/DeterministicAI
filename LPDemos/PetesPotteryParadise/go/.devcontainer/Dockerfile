FROM golang:1.24-bookworm

ARG ORTOOLS_ARCHIVE=or-tools_amd64_debian-12_cpp_v9.15.6755.tar.gz

RUN apt-get update \
    && apt-get install --no-install-recommends -y build-essential ca-certificates curl git protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

RUN git clone --depth 1 --branch v9.15 https://github.com/google/or-tools.git /opt/or-tools \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 \
    && go install honnef.co/go/tools/cmd/staticcheck@2025.1.1 \
    && PATH="$PATH:/go/bin" protoc \
        --proto_path=/opt/or-tools \
        --go_out=/opt/or-tools \
        --go_opt=module=github.com/google/or-tools \
        ortools/sat/cp_model.proto \
        ortools/sat/sat_parameters.proto \
    && printf 'go 1.24.0\n\nuse (\n\t/workspace\n\t/opt/or-tools\n)\n' > /opt/go.work \
    && curl --fail --location --silent --show-error \
        "https://github.com/google/or-tools/releases/download/v9.15/${ORTOOLS_ARCHIVE}" \
        --output /tmp/ortools.tar.gz \
    && echo "b2c9870c8778eeb26c98742402da17da039c058fca7eca87be5c90832b04153c  /tmp/ortools.tar.gz" | sha256sum --check \
    && mkdir -p /opt/or-tools \
    && tar -xzf /tmp/ortools.tar.gz -C /opt/or-tools --strip-components=1 \
    && rm /tmp/ortools.tar.gz

ENV GOWORK="/opt/go.work"
ENV CGO_CFLAGS="-I/opt/or-tools"
ENV CGO_LDFLAGS="-L/opt/or-tools/lib -lortools"
ENV LD_LIBRARY_PATH="/opt/or-tools/lib"

WORKDIR /workspace
