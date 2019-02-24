# Twuser
Twuser uses the twitter api to block/unblock mute/unmute users given in plain text lists as exported from twitter.

It is not much use for blocking since you can easily import blocklists, but unblocking seems to be a different story.

## build && example usage
- go get 
- go build
- setup twitterapi.yaml in the directory ~/.twuser/ (check the code for variable names)
- get api keys and access secrets for your twitter account and put them in the twitterapi.yaml file
- copy the twuser binary to a directory thats in your $PATH (or you do whatever you want)
- use like: twuser --mode unblock --file bla.csv
- modify the the sleep time in the doIterate(...) function but beware of rate limits.
