# Bashtasks

Executing of the bash scripts as a tasks

## Getting started

Create basic.yaml file with tasks
```yaml
#loading of tasks
show_output: true
tasks:
  - title: whoami
    cmd: whoami
  
  - title: network
    cmd: ifconfig
```

Moving of this file to the BASHTASKS directory like
```
export BASHTASKS=$HOME/bashtasks
```

Execute the command
```
bashtasks basic
```

Executing of the task with downloading of bash script from remote
```yaml
#loading of tasks
show_output: true
tasks:
  - title: whoami
    cmd: whoami
  
  - title: network
    cmd: ifconfig

  - title: Get bash script
    path: https://raw.githubusercontent.com/ruanyf/simple-bash-scripts/master/scripts/Hello.sh
```