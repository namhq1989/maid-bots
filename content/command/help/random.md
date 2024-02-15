*RANDOM:*

The `/random` command allows you to generate random things with configurable options

`/random <parameters>`

*Parameters:*

_Required_

`type:`
  `number` \- generate random numbers
  `string` \- generate random strings

_Optional_

`number:`
  `format` \(string\) \- specify the type of number \(int or decimal\)\. Default is int
  `min` \(int\) \- set the minimum value for the random number\. Default is 0
  `max` \(int\) \- set the maximum value for the random number\. Default is 0
  `count` \(int\) \- specify the number of random numbers to generate\. Default is 1\. Max is 100
  `unique` \(bool\) \- ensure that generated numbers are unique within the specified range

`string:`
  `value`:
    `person` \- random person
    `email` \- random email address
    `phone` \- random phone number
    `username` \- random username
    `address` \- random address
    `latlon` \- random latitude and longitude with format lat,lon
    `sentence` \- random sentence
    `paragraph` \- random paragraph
    `quote` \- random quote
    `uuid` \- random UUID
    `hexcolor` \- random HEX color
    `rgbcolor` \- random RGB color
    `url` \- random url
    `imageurl` \- random image url
    `domain` \- random domain
    `ipv4` \- random IPv4 address
    `ipv6` \- random IPv6 address
    `ua` \- random User\-Agent address
    `date` \- random date
    `timezone` \- random timezone
    `creditcard` \- random credit card
    `walletaddress` \- random wallet address
    `pet` \- random pet
    `emoji` \- random emoji

*Examples:*

_Number_:
/random type\=number format\=int min\=1 max\=10 unique\=true
/random type\=number format\=decimal min\=0 max\=100 count\=5

_String_:
/random type\=string value\=person
/random type\=string value\=uuid
/random type\=string value\=ipv4
