## NewRelic Go Agent

A small convenience layer that sits on top of [newrelic-go-agent](https://github.com/paulsmith/newrelic-go-agent), to make
it easy to create transactions for NewRelic in Go.

## Installing & Building

New Relic's Agent SDK only supports Linux, so by default the agent is disabled to allow development on non-Linux environments.  It will however build and run on Linux environments (including Heroku & Cloud Foundry).

## Run on Heroku/Cloud Foundry

You will need [godep](https://github.com/tools/godep) installed to manage dependencies.

1. Add the github.com/sky-uk/newrelic-go-agent to your app's imports section.
2. Run godep save -r ./...
3. Add newrelic.Init to your app.
4. Add Heroku's [Go buildpack](https://github.com/heroku/heroku-buildpack-go) to your manifest.yml (Cloud Foundry's should work too).
5. Grab the [New Relic Agent SDK](http://download.newrelic.com/agent_sdk/) and extract the binaries into your app's base directory:  
 tar zxvf nr_agent_sdk-v0.16.1.0-beta.x86_64.tar.gz --strip=2 -C . "nr_agent_sdk-v0.16.1.0-beta.x86_64/lib"  
 This could be added to CircleCI/Jenkins as part of a build.
6. Change your Procfile/manifest.yml start command to have LD_LIBRARY_PATH=. before your app's binary location.

 e.g.

 Add  
 [command](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#start-commands): LD_LIBRARY_PATH=. your-app-name
 to your CF's manifest.yml

 Or add  
 web: LD_LIBRARY_PATH=. ./your-app-name  
 to your Procfile.
7. Push your app!
8. If your app requires a firewall to communicate to New Relic, it will use the HTTP_PROXY env variable by default.

## Build Locally

As explained earlier this will not run on Darwin or any non-Linux environment.

Follow the above guide for Heroku/Cloud Foundry, but build with:
```
  go build -tags heroku ./...
```

Run your app with:

```
LD_LIBRARY_PATH=. ./your-app-name
```


## Example Usage

Check out the examples folder for a 'hello world'.  
Check the [original repo](https://github.com/remind101/newrelic/blob/master/example/main.go) for more.
