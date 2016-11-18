[![Go Report Card](https://goreportcard.com/badge/github.com/eljuanchosf/ansible-gen)](https://goreportcard.com/report/github.com/eljuanchosf/ansible-gen)
[![codebeat badge](https://codebeat.co/badges/ffe93ea7-4a4d-47ba-9e63-d248b491b1b1)](https://codebeat.co/projects/github-com-eljuanchosf-ansible-gen)
# ansible-gen

A simple tool to scaffold Ansible projects according to the best practices described in the [Ansible documentation](http://docs.ansible.com/ansible/playbooks_best_practices.html).

## Installation

If you want to compile from source, you have to have Golang installed. I used version 1.7.3 for development. Didn't try any other version.

I highly recomment [gvm](https://github.com/moovweb/gvm) to manage your Go versions. 

To build it, use the included `run.sh` script:

```
./run.sh build
```

**TODO**: Soon, compiled versions to download.  

## Usage

### Global commands and options

```
USAGE:
   ansible-gen [global options] command [command options] [arguments...]

COMMANDS:
     project, p  Creates a new Ansible project
     role, r     Creates a new Ansible role
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dry-run, -d  just print results, do not modify the filesystem
   --help, -h     show help
   --version, -v  print the version
```

### New projects

```
NAME:
   ansible-gen project - Creates a new Ansible project

USAGE:
   ansible-gen project [command options] [arguments...]

OPTIONS:
   -c value    A comma separated list of the custom roles for the project
   -g value    A comma separated list of the Ansible Galaxy roles for the project
   --skip-git  Do not initialize a Git repository for the project
```

#### Galaxy roles

When you specify a list of Galaxy roles, a `galaxy-roles.yml` will be created with the appropriate format to install the roles with:

```
$ ansible-galaxy install -r galaxy-roles.yml
```

#### Example

```
$ ansible-gen p my-project -c backend-server,my-custom-role -g geerlingguy.redis,nickhammond.logrotate 
```

### New roles

```
NAME:
   ansible-gen role - Creates a new Ansible role

USAGE:
   ansible-gen role [arguments...]
```

#### Example

```
$ ansible-gen r my-role 
```

## Rationale

The idea was to experiment with some Golang features while developing something useful.
The code can be better, I know, but this is learning project. Any suggestion is welcome!


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
