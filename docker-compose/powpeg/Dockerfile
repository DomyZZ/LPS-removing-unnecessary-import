FROM openjdk:8-jdk-slim-buster as builder

RUN apt-get update -y && \
    apt-get install -qq --no-install-recommends git curl gnupg

WORKDIR /code/powpeg

ARG POWPEG_RELEASE="HOP"
ARG POWPEG_VERSION="4.0.0.0"

RUN gitrev="${POWPEG_RELEASE}-${POWPEG_VERSION}" && \
    git init && \
    git remote add origin https://github.com/rsksmart/powpeg-node.git && \
    git fetch --depth 1 origin tag "$gitrev" && \
    git checkout "$gitrev"

RUN gpg --keyserver https://secchannel.rsk.co/SUPPORT.asc --recv-keys A6DBEAC640C5A14B && \
    gpg --finger A6DBEAC640C5A14B && \
    gpg --verify --output SHA256SUMS SHA256SUMS.asc && \
    sha256sum --check SHA256SUMS && \
    ./configure.sh && \
    ./gradlew --no-daemon clean build -x test

FROM --platform=linux/amd64 openjdk:8-jdk as runner

ARG UID=1000
ARG HOME="/home/powpeg"
RUN useradd -m -u "$UID" --home="$HOME" powpeg

COPY --from=builder --chown="$UID" "/code/powpeg/build/libs/federate-node-*-all.jar" "$HOME/powpeg.jar"

COPY --chown="$UID" reg1.key "$HOME/reg1.key"
COPY --chown="$UID" regtest-fed.conf "$HOME/regtest-fed.conf"

RUN chmod 400 "$HOME/reg1.key"

RUN mkdir -p "$HOME/.powpeg"; chown "$UID" "$HOME/.powpeg"
RUN mkdir -p "$HOME/logs"; chown "$UID" "$HOME/logs"

WORKDIR "$HOME"

USER powpeg
