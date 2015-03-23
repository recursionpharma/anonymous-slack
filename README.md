# anonymous-slack
A Heroku service to send anonymous messages to colleagues on Slack. Your message shows up under the guise of a random animal: `an anonymous aardvark: [your message]`.

or As a slack administrator:

Clone this repository, and create a new heroku app with the Go buildpack: https://github.com/kr/heroku-buildpack-go

    git clone git@github.com:recursionpharma/anonymous-slack.git
    cd anonymous-slack
    heroku create -b https://github.com/kr/heroku-buildpack-go.git
    
In Slack integrations, add a Slash command, for example, `/anon` . Set the URL in Slack to your Heroku website URL. The resulting slack "token" should be set as a Heroku environment variable:

    heroku config:set INCOMING_SLACK_TOKEN=XXX

Then, in Slack integrations, add a Slash Incoming webhook. The resulting "webhook url" should be set as the Heroku environment variable:

    heroku config:set INCOMING_SLACK_WEBHOOK=https://hooks.slack.com/services/BLAH/BLAH/BLAH

Deploy to heroku.

    git push heroku master

Success! Now if you send a message in any channel, public or private, like the following:

    /anon @somebodyelse hey, guess who?

That message will be suppressed, and @somebodyelse gets a message like this:

    an anonymous capybara says: hey, guess who?
    
Be nice!
