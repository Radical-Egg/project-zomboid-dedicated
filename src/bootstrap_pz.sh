#!/bin/bash
function launch() {
    /home/steam/pzserver/start-server.sh \
        -adminusername $ADMIN_USERNAME \
        -adminpassword $ADMIN_PASSWORD \
        -servername $SERVERNAME
}

launch