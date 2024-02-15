*MONITOR:*

The `/monitor` command allows you to manage and monitor various targets with different commands

`/monitor <parameters>` 

*Available types:*
  `domain` \- domain name
  `http` \- http/https url
  `tcp` \- domain/ip with port
  `icmp` \- domain/ip without port

*Parameters:* _\(with appropriate subsequent parameters\)_
`action:`
  `check` \- perform a one\-time check on the target
    • `type` \- target type
    • `target` \- target value

  `register` \- register a new target for continuous monitoring according to a predefined schedule
    • `type` \- target type
    • `target` \- target value

  `list` \- display a list of all registered monitoring targets
    • `type` \- target type \(omit to retrieve all targets\)
    • `keyword` \- keyword for result filtering
    • `page` \- pagination \(10 records per page\)

  `remove` \- remove a specific target by its id
    • `id` \- target id

  `stats` \- obtain statistics of a specific target by its ID
    • `id` \- target id

*Examples:*

/monitor action\=check type\=domain target\=google\.com
/monitor action\=register type\=domain target\=google\.com
/monitor action\=list type\=domain