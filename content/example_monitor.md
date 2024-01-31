*Monitor examples:*

\# check
/monitor check domain `google.com`
/monitor check tcp `1.2.3.4:3000`

\# register
/monitor register domain `google.com`
/monitor register http `https://jsonplaceholder.typicode.com/todos/1`

\# list
/monitor list all \# all watching targets
/monitor list domain \# all watching domains

\# remove
\# collect id from "/monitor list" command
/monitor remove $id