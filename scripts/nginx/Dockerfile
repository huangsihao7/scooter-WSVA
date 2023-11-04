FROM nginx:1.25-alpine-slim

ARG HOST="172.22.121.53"
COPY nginx.conf /etc/nginx/nginx.conf
RUN  sed -i "s/172\.22\.121\.53/$HOST/g" /etc/nginx/nginx.conf