#!/usr/local/bin/bash

set -e
set -x

# push origin
git push origin master

# overwrite nginx config
scp luhmn.sirodoht.com.conf root@95.217.223.96:/etc/nginx/sites-available/

# pull and reload on server
ssh root@95.217.223.96 'cd /opt/apps/luhmn \
    && git pull \
    && systemctl reload nginx'
