FROM node:latest
RUN npm install -g fauna-shell
COPY .fauna-shell /root
ENTRYPOINT ["tail", "-f", "/dev/null"]
