FROM cm2network/steamcmd

ENV APP="pzserver" \
	STEAM_HOME="/home/steam" \
	APP_DIR="${STEAM_HOME}/${APP}" \
	APP_ID="380870" \
	SERVER_PORT=16261 \
	PLAYER_PORTS=16262-16272 \
	STEAM_PORT_1=8766 \
	STEAM_PORT_2=8767

USER root
WORKDIR ${STEAM_HOME}

COPY ./start.sh "${STEAM_HOME}"/start.sh

RUN mkdir -p "${APP_DIR}" && \
	chown -R steam:steam ${APP_DIR} && \
	mkdir -p "${STEAM_HOME}"/Zomboid/Server && \
	ln -s "${STEAM_HOME}"/Zomboid/Server /config && \
	chown -R steam:steam "${STEAM_HOME}"/Zomboid && \
	chown steam:steam ./start.sh && \
	chown -R steam:steam /config

USER steam

RUN ./steamcmd/steamcmd.sh \
	+force_install_dir "${APP_DIR}" \
	+login anonymous +app_update "${APP_ID}" \
	-beta b41multiplayer validate \
	+quit

EXPOSE ${SERVER_PORT}/udp ${PLAYER_PORTS} ${STEAM_PORT_1}/udp ${STEAM_PORT_2}/udp
ENTRYPOINT ["/home/steam/start.sh"]
