FROM node:12.16.1

WORKDIR /flownms_web/
COPY . .

RUN npm install
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="forever765@github.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /flownms_web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
