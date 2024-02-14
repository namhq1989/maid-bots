*RANDOM:*

The `/random number` command allows you to generate random numbers with configurable options

`/random number <config>`

*Configurable options:*
  `type` \(string\) \- specify the type of number \(int or decimal\)\. Default is int
  `min` \(int\) \- set the minimum value for the random number\. Default is 0
  `max` \(int\) \- set the maximum value for the random number\. Default is 0
  `count` \(int\) \- specify the number of random numbers to generate\. Default is 1\. Max is 100
  `unique` \(bool\) \- ensure that generated numbers are unique within the specified range

*Examples:*
/random number type\=int min\=1 max\=10
/random number type\=decimal min\=0 max\=100 count\=5 unique\=true
