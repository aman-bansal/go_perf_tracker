# go_perf_tracker
simple tool to autogenerate timing metrics using go ast

PreBuild Step
-------------
go_perf_tracker actually will be a prebuild step for your application. go_perf_tracker regenerates files after adding
```defer perf_tracker.CalculateFunctionPerf(time.Now(), "FuntionName")``` with appropriate imports at the start of each function provided.

For this to work you need to add dependency of go_perf_tracker in your application. We are working towards removing this constraint.

Installation
------------

To get the latest released version use:

```GO111MODULE=on go get github.com/aman-bansal/go_perf_tracker/@latest```

Run Instructions
----------------

You will need to provide the path for config file. Config must be in json format 
with file path vs functions mapping whose performance needs to be logged.
for example:
```
    {
        "path/to/file":["funcName1", "funcName2"]
    }
```

run this command```go_perf_tracker -config=test/config.json```
this will regenerate those files with performance logging line. 

Scope of Improvement
--------------------
1. Can think of taking performance function via configuration. Then no need to add dependency.
2. Can add client for statsd to emit metrics.
3. If the process fails in between, then there is possibility that few of the files are already generated which are now
in inconsistent state. Can save file initially in cache and recreate the original one if it fails.
4. Can add extra check if the function is already defined don't add it again and again.

Issue Reporting
----------------
If you found any issues or wanted to suggest any enhancement, do create an issue in the repo itself or shoot an email to bansalaman2905[at]gmail[dot]com.
