# Increment

This program increments a Consul key value

Presumptions:
- Consul is running on localhost
- Consul does not require auth
- Key value is parsable as an int
- Increment file is parasable as an int

```
Start:
  spew [location of a file with an int]

Env Vars:
  INCREMENT_KEY - the consul key to increment
```
