# anonymous-slack
A service you can upload to Heroku to be able to send anonymous messages to colleagues. As a slack administrator:

Create a new heroku app with the Go buildpack: https://github.com/kr/heroku-buildpack-go

Add a Slash command, for example, `/anon` . Set the URL in Slack to your Heroku URL. The resulting slack "token" should be set as a Heroku environment variable:

    heroku config:set INCOMING_SLACK_TOKEN=XXX

Add a Slash Incoming webhook. The resulting "webhook url" should be set as the Heroku environment variable:

    heroku config:set INCOMING_SLACK_WEBHOOK=https://hooks.slack.com/services/BLAH/BLAH/BLAH

Deploy to heroku.

    git push heroku master

Success! Now if you send a message in any channel, public or private, like the following:

    /anon @somebodyelse hey, guess who?

That message will be suppressed, and @somebodyelse gets a message like this:

    an anonymous capybara says: hey, guess who?
    
Be nice!
