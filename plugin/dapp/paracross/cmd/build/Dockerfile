FROM ubuntu:16.04

WORKDIR /root
COPY turingchain turingchain
COPY turingchain-cli turingchain-cli
COPY turingchain.toml turingchain*.toml ./
COPY entrypoint.sh entrypoint.sh

CMD ["/root/turingchain", "-f", "/root/turingchain.toml"]
