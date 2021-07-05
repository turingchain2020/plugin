FROM ubuntu:16.04

WORKDIR /root
COPY turingchain turingchain
COPY turingchain-cli turingchain-cli
COPY entrypoint.sh entrypoint.sh
COPY turingchain.toml turingchain*.toml ./


CMD ["/root/turingchain", "-f", "/root/turingchain.toml"]
