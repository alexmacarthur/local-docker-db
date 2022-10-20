FROM typesense/typesense:0.24.0.rcn21

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

EXPOSE 9108

ENTRYPOINT ["/opt/typesense-server"]
