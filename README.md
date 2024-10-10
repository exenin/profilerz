# profilerz

A tool for managing config profiles at org levels, etc.

```bash
 ~  profilerz 
Profile manager for config directories (AWS, kubectl, DigitalOcean, etc.)

Usage:
  profilerz [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize profilerz by creating default profile with current configs
  profile     Manage profiles (add, set, list, delete)

Flags:
  -h, --help   help for profilerz

Use "profilerz [command] --help" for more information about a command.
 ~  profilerz profile -h
Manage profiles (add, set, list, delete)

Usage:
  profilerz profile [command]

Available Commands:
  add         Add a new profile
  delete      Delete a profile
  list        List all profiles
  set         Set a profile as active

Flags:
  -h, --help   help for profile

Use "profilerz profile [command] --help" for more information about a command.
```

## Installation

Run the following commands:


```bash 

$ make build install

Building profilerz...
go build -o profilerz ./cmd
Installing profilerz...
go install 
```

Initialize - and copy current configs to "default profile 
```bash 

$ profilerz init            
Initializing profilerz...



$ profilerz profile add personal
Profile 'personal' created.

$ ls -las ~/.profilerz.d/
cbas/      default/     personal/



$ ls -als ~ | grep $HOME/.profilerz.d
.aws -> /home/user/.profilerz.d/cbas/aws
.kube -> /home/user/.profilerz.d/cbas/kubectl
.ssh -> /home/user/.profilerz.d/cbas/ssh



$ profilerz profile set personal     
Profile 'personal' is now active.



$ ls -als ~ | grep $HOME/.profilerz.d
.aws -> /home/user/.profilerz.d/personal/aws
.kube -> /home/user/.profilerz.d/personal/kubectl
.ssh -> /home/user/.profilerz.d/personal/ssh

   ```