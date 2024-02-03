*RANDOM:*

The `/random number` command allows you to generate random numbers with configurable options

`/random number <config>`

*Configurable options:*
  `type` \- specify the type of number \(int or decimal\)\. Default is int
  `min` \- set the minimum value for the random number\. Default is 0
  `max` \- set the maximum value for the random number\. Default is 0
  `count` \- specify the number of random numbers to generate\. Default is 1. Max is 100
  `unique` \- ensure that generated numbers are unique within the specified range

*Examples:*
/random number type\=int min\=1 max\=10
/random number type\=decimal min\=0 max\=100 count\=5
