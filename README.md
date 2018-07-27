# tail, for all logs

ok, so initially, i'm playing with elasticsearch...

Goals:

to make watching the log files for a system as simple as `tail -f /var/log/*` was on simple UNIX hosts in the 90's


## TODO

* have a `--connect docker:elasticsearch_esnetwork:elasticsearch_server:9200` which will add an ambassador-sidcar to give `tail` access to the containerised elasticsearch container using that overlay network.
* create index
* delete index
* view multiple indexes using the same wildcard expressions as in creating an index