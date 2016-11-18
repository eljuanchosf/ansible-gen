# ansible-gen

A simple tool to scaffold Ansible projects according to the best practices described in the [Ansible documentation](http://docs.ansible.com/ansible/playbooks_best_practices.html).

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
