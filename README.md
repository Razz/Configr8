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

