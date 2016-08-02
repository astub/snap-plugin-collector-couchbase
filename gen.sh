#!/bin/bash

metriclist="$(snapctl metric list | cut -f1 | cut -f3- -d"/" | tr "/" .)"

#echo -e "echo: \n$metriclist"
for line in $metriclist; do
	foo1="  - pattern: \"%app%.%host%.$line\""
	foo2="    metric_key: $line"
	echo $foo1 >> ./sad.yaml
	echo $foo2 >> ./sad.yaml
done