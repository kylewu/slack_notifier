#!/bin/sh

if [ -z "$SLACK_TOKEN" ]; then 
	echo "$SLACK_TOKEN is unset"
	exit 1
fi

if [ "$#" -eq 0 ]; then
	echo "Usage: $0 your message"
	exit 1
fi

SLACK_API="https://slack.com/api"
USER_ID=$(curl -G -s "${SLACK_API}/auth.test" --data-urlencode "token=${SLACK_TOKEN}" | jq -r .user_id)
CHANNEL_ID=$(curl -G -s "${SLACK_API}/im.open" --data-urlencode "token=${SLACK_TOKEN}" --data-urlencode "user=${USER}" | jq -r .channel.id)
curl -s -o /dev/null -G "${SLACK_API}/chat.postMessage" --data-urlencode "channel=$CHANNEL_ID" --data-urlencode "text=$*" --data-urlencode "token=$SLACK_TOKEN" --data-urlencode "username=bot"
