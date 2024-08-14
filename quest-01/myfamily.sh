#!/usr/bin/bash
curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r --arg hid $HERO_ID ' .[] | select( (.id|tostring) == $hid ) | (.connections.relatives | gsub("\n"; "\\n")) '
