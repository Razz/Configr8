[![Go Report Card](https://goreportcard.com/badge/github.com/razz/configr8)](https://goreportcard.com/report/github.com/razz/configr8)
# Configr8
A light, dynamic config generator that utilizes commonsense extensions to Go's templating libary.

## Use
`configr8 -t =/tmp/nginx.tmpl -d /etc/nginx.conf`

`-j or -json: Supply a json file`

`-t or -tmpl: Template File`

`-d or -dest: Destination File, defaults to Stdout`

## Plugins
- Make a function
- create a file in the Plugin directory
- follow the boilerplate
- include it to the `FuncMap`
- recomplie
- include in template 


## Example 
You could run the template with or without the added json information.
`configr8 -t example.tmpl` or `configr8 -t example.tmpl -j example.json`

Built by [@Razz](http://github.com/Razz)
