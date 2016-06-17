# configr8
A light, dynamic config generator that utilizes commonsense extensions to Go's templating libary.

## Use
`configr8 -tmpl=/tmp/nginx.tmpl > /etc/nginx.conf`

## Plugins
- Make a function
- create a file in the Plugin directory
- follow the boilerplate
- include it to the `FuncMap`
- recomplie
- include in template 


#### To-do
+ the ability to modify `tmpl.Delims`
+ reflect json into `DataMap` for use in a template 
+ panic handiling 
+ error handling
+ linting (nice to have, might be more work than worth) 
