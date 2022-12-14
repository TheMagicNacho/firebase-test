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
*note: This is only really useful for static hosting. For microservices refer to google run*

# 🐋 Containerization
*ref: https://firebase.google.com/docs/hosting/cloud-run*

pre-req: you'll need the google CLI https://cloud.google.com/sdk/docs/install-sdk

- Write the code.
- Build the container using google cloud run `gcloud builds submit --tag gcr.io/PROJECT_ID/CONTAINER_TAG`
- Store the container in the container registry 
- push the container to cloud run `gcloud run deploy --image gcr.io/PROJECT_ID/CONTAINER_TAG`