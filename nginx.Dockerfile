FROM nginx:alpine
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY /ssl /ssl
CMD ["nginx", "-g", "daemon off;"]
