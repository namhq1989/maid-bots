*MONITOR:*

The /monitor command allows you to manage and monitor various targets with different commands

`/monitor $command $target $value` 

*Commands:*
  `check` \- perform a one\-time check on the target
  `register` \- register a new target for continuous monitoring according to a predefined schedule
  `list` \- display a list of all registered monitoring targets
  `remove` \- remove a specific target by its ID
  `example` \- list of examples

*Targets:*
  `domain` \- domain name
  `http` \- http/https url
  `tcp` \- domain/ip with port
  `icmp` \- domain/ip without port
