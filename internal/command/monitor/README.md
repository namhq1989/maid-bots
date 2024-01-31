The standard format for this command:

### `/monitor $action $target $value`

Actions:
- check - perform a one-time check on the target
- register - register a new target for continuous monitoring according to a predefined schedule
- list - display a list of all registered monitoring targets
- remove - remove a specific target by its id
- stats - get statistic of a specific target by its id

Targets:
- domain - domain name
- http - http/https url
- tcp - domain/ip with port
- icmp - domain/ip without port

---

Users have the flexibility to input various arguments with the `/monitor` command, resulting in different outcomes based on the number and nature of the provided arguments:

### `/monitor` && `/monitor $arg1`
- In cases where there is insufficient arguments, the bot will respond the content of `/help monitor` command

### `/monitor $arg1 $arg2`
- If `$arg1` is equal to `list`, the bot will validate the correctness of `$arg2` and respond accordingly.
- For other scenarios with insufficient arguments, the bot will respond the content of `/example monitor` command

### `/monitor $arg1 $arg2 $arg3`
- If `$arg1` is equal to `list`, the bot will respond "invalid arguments" message
- If either `$arg1` and `$arg2` is not valid, the bot will respond with the content of `/example monitor` command
- For valid inputs, the bot will set up the necessary data and route the command to the respective handler for further processing
