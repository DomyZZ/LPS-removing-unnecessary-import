FROM --platform=linux/amd64 node:16

RUN apt-get update -y && apt-get install -qq --no-install-recommends jq

RUN npm install -g truffle

USER node

WORKDIR /home/node

ARG LBC_GIT_BRANCH="master"

RUN gitBranch=${LBC_GIT_BRANCH} && \
    git init && \
    git remote add origin https://github.com/rsksmart/liquidity-bridge-contract.git && \
    git fetch --depth 1 origin "$gitBranch" && \
    git checkout "$gitBranch"

COPY --chown=node truffle-config.patch ./

RUN git apply truffle-config.patch

RUN npm ci

RUN npx truffle compile

COPY --chown=node deploy-lbc.sh ./
