# project-zomboid-dedicated
Dedicated Server + Discord bot for project Zomboid

Simple Project Zomboid Dedicated Game Server with the option to provide your own discord bot for server restarts

In order to get the discord bot working you will need to create one on the developer portal and provide the API key in the docker-compose file. Sample compose file

If you clone this repo you can build it directly like this - otherwise just change the build to image. See docker-compose-sample.yml for pulling image via docker hub

```YAML

---
version: '3.9'
services:
  zomboid:
    container_name: zomboid-dedicated
    build: .
    volumes:
      - ./config:/config
    environment:
      - ADMIN_USERNAME=admin
      - ADMIN_PASSWORD=secret
      - SERVERNAME=pz_server
    ports:
      - 8766:8766/udp
      - 8767:8767/udp
      - 16261:16261/udp
      - 16262-16272:16262-16272
      - 27015:27015
      - 9000:9000
    networks:
      - pz_bridge
  pz_bot:
    container_name: pz_bot
    build:
      context: .
      dockerfile: ./pz_bot/Dockerfile
    environment:
      - CLIENT_TOKEN=<discord-bot-api-key>
      - PZ_SERVER=zomboid-dedicated
    depends_on:
      - zomboid
    networks:
      - pz_bridge
networks:
  pz_bridge:
    driver: bridge
 ```
