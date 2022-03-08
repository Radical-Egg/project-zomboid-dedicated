FROM ubuntu:latest

ENV APP=pzserver
ENV STEAM_HOME=/home/steam
ENV APP_DIR=${STEAM_HOME}/${APP}
ENV APP_ID=380870
ENV SERVER_PORT=16261
ENV PLAYER_PORTS=16262-16272
ENV STEAM_PORT_1=8766
ENV STEAM_PORT_2=8767

RUN useradd -m steam

RUN echo steam steam/question select "I AGREE" | debconf-set-selections && \
	echo steam steam/license note '' | debconf-set-selections

RUN apt-get update -y && \
	apt-get install -y software-properties-common && \
	add-apt-repository multiverse && \
	dpkg --add-architecture i386 && \
	apt-get update -y && \
	apt-get install -y lib32gcc-s1 steamcmd &&  \
	ln -s /usr/games/steamcmd /home/steam/steamcmd

RUN apt-get install -y supervisor && \
	mkdir -p /var/log/supervisor && \
	chown -R steam:steam /var/log/supervisor

COPY supervisord/conf.d/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY supervisord/supervisord.conf /etc/supervisor/supervisord.conf
COPY ./src/bootstrap_pz.sh ${STEAM_HOME}/bootstrap_pz.sh

RUN mkdir -p ${APP_DIR} && \
	mkdir -p ${STEAM_HOME}/Zomboid/Server && \
	mkdir -p ${STEAM_HOME}/var/run && \
	ln -s ${STEAM_HOME}/Zomboid/Server /config && \
	chown -R steam:steam ${STEAM_HOME} && \ 
	chmod -R 755 ${STEAM_HOME} && \
	chmod 775 /config && chown -R steam:steam /config && \
	chgrp steam /etc/supervisor/conf.d/supervisord.conf

USER steam

RUN ${STEAM_HOME}/steamcmd \
	+force_install_dir "${APP_DIR}" \
	+login anonymous +app_update "${APP_ID}" \
	-beta b41multiplayer validate \
	+quit

EXPOSE ${SERVER_PORT}/udp ${PLAYER_PORTS} ${STEAM_PORT_1}/udp ${STEAM_PORT_2}/udp
VOLUME [ "${APP_DIR}" ]

CMD ["/usr/bin/supervisord"]