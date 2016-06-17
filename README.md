# configr8
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


#### To-do
+ the ability to modify `tmpl.Delims`
+ panic handiling 
+ linting (nice to have, might be more work than worth) 
