# annona [![Build Status][travis image]][travis]
![Annona Image][image]

A service for posting anonymous messages to a Slack channel. Your message shows up under the guise of a random character.

Based on and [Slacker] (http://slacker.nathanhoad.net/) and [anonymous-slack](https://github.com/recursionpharma/anonymous-slack)


## As a slack administrator:
    
In Slack integrations, add a Slash command, for example, `/anon` . Set the URL in Slack to your service URL. The resulting slack "token" should be set as a environment variable:

     INCOMING_SLACK_TOKEN=XXX

Then, in Slack integrations, add a Slash Incoming webhook. The resulting "webhook url" should be set as environment variable:

    INCOMING_SLACK_WEBHOOK=https://hooks.slack.com/services/BLAH/BLAH/BLAH

The channel that the messages will appear in is set set as environment variable:

	SLACK_CHANNEL_ID="#anon"

 master

Success! Now if you send a message in any channel, public or private, like the following:

    /anon Who am I ?
       
![Slack Image][slack_image]
That message will be suppressed, and a message will appear in the SLACK\_CHANNEL\_ID channel

## Avatars

Avatars are in the avatars.json file.

Here is an example for an avatar

	{
    "username": "Tina",
    "default_text": "Errrggggggg",
    "icon_url": "http://i.imgur.com/VxDC1dz.png"
    },
  

[Docker Image](https://hub.docker.com/r/rounds/10m-annona/) 

[image]: annona.jpg
[slack_image]: Slack.png
[travis image]: https://travis-ci.org/rounds/annona.svg
[travis]: https://travis-ci.org/rounds/annona

