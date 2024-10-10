# profilerz

A tool for managing config profiles at org levels, etc.

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
   0 lrwxrwxrwx  1 jdvh jdvh      32 Oct 10 23:36 .aws -> /home/jdvh/.profilerz.d/cbas/aws
   0 lrwxrwxrwx  1 jdvh jdvh      36 Oct 10 23:36 .kube -> /home/jdvh/.profilerz.d/cbas/kubectl
   0 lrwxrwxrwx  1 jdvh jdvh      32 Oct 10 23:36 .ssh -> /home/jdvh/.profilerz.d/cbas/ssh



$ profilerz profile set personal     
Profile 'personal' is now active.



$ ls -als ~ | grep $HOME/.profilerz.d
   0 lrwxrwxrwx  1 jdvh jdvh      36 Oct 10 23:40 .aws -> /home/jdvh/.profilerz.d/personal/aws
   0 lrwxrwxrwx  1 jdvh jdvh      36 Oct 10 23:40 .gitconfig -> /home/jdvh/.profilerz.d/personal/git
   0 lrwxrwxrwx  1 jdvh jdvh      40 Oct 10 23:40 .kube -> /home/jdvh/.profilerz.d/personal/kubectl
   0 lrwxrwxrwx  1 jdvh jdvh      36 Oct 10 23:40 .ssh -> /home/jdvh/.profilerz.d/personal/ssh

   ```