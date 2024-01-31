*MONITOR:*

The `/monitor` command allows you to manage and monitor various targets with different commands

`/monitor $action $target $value` 

*Actions:*
  `check` \- perform a one\-time check on the target
  `register` \- register a new target for continuous monitoring according to a predefined schedule
  `list` \- display a list of all registered monitoring targets
  `remove` \- remove a specific target by its id
  `stats` \- get statistic of a specific target by its id

*Targets:*
  `domain` \- domain name
  `http` \- http/https url
  `tcp` \- domain/ip with port
  `icmp` \- domain/ip without port
