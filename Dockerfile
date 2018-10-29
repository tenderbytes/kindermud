FROM alpine:latest

ENV DIST_NAME kindermud
ENV APP_HOME /opt/app
ENV KINDER_CONFIG_PATH $APP_HOME/$DIST_NAME.config.yaml

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY ./data $APP_HOME/data
COPY ./dist $APP_HOME/
COPY ./docker/* $APP_HOME/
RUN chmod +x $APP_HOME/$DIST_NAME && chmod +x $APP_HOME/*.sh
RUN mv $APP_HOME/$DIST_NAME /bin

ENTRYPOINT ["/opt/app/entrypoint.sh"]
