FROM node:alpine3.11 as build

ARG API_URL="  apiUrl: 'http://localhost:8080/api/server'"
ARG FRONTEND_APP_SRC_URL=https://github.com/DanijelRadakovic/Servers-Front/archive/refs/tags/v1.1.1.tar.gz

WORKDIR /usr/src/server
RUN apk --update --no-cache add curl tar && \
    curl -L "${FRONTEND_APP_SRC_URL}" | tar -xz &&\
    mv "$(find . -maxdepth 1 -type d | tail -n 1)" app && \
    cd app && \
    npm install && \
    sed -i -e "/apiUrl:/c\\${API_URL}" \
      ./src/environments/environment.ts && \
    npm run build --prod  && \
    cd dist && \
    mv $(ls -d */|head -n 1) /app


FROM nginx:1.21-alpine as execute
COPY ./files/nginx.conf /etc/nginx/nginx.conf
COPY ./files/api_gateway.conf /etc/nginx/api_gateway.conf
COPY --from=build /app /usr/share/nginx/html/app
CMD sed -i \
        -e 's/PORT/'$PORT'/g' \
        -e 's/API_URL/'$API_URL'/g' \
        /etc/nginx/api_gateway.conf && \
    nginx -g 'daemon off;'