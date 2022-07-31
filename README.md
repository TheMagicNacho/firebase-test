# firebase-test
testing firebase workflows

# :fire: Firebase CLI
We are going to need the firebase CLI

*ref: https://firebase.google.com/docs/cli*

* install the cli: `curl -sL https://firebase.tools | bash`
* authenticate: `firebase login`

# :test_tube: Testing
We're using testify instead of ginkgo only because my local host is having issues trying to install ginkgo.
*ref: https://github.com/stretchr/testify*


# hosting
*ref: https://firebase.google.com/docs/hosting/github-integration?authuser=0*

Init first: `firebase init hosting`

Init for github: `firebase init hosting:github`