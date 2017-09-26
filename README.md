# Welcome to Reddit Lite

This is a demo website created using Go and Revel framework for learning purpose.
For frontend, bootstrap 3 and jQuery 2.2 are used.

### Start the web server:

   revel run github.com/lanjiang/reddit-lite

### Deployment

Please go to http://13.228.105.37:9000/ to view the demo

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file [auto-generated]
        routes        Routes definition file [auto-generated]

    app/              App sources
        init.go       Interceptor registration [auto-generated]
        controllers/  App controllers go here
        models/       Models directory
        views/        Templates directory
        services/     Services directory
        routes        Route tables [auto-generated]
        tmp           tmp files [auto-generated]

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Revel framework

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).ß
