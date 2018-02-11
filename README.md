YAML support for Terraform
===

# Introduction
Terrayaml is a tool that let you write your Terraform files using YAML.  
It supports HCL and JSON outputs.

# Installation
`go get -u github.com/hjdr4/terrayaml`

# Usage
Output is written to stdout.  

- `-f` flag let you choose between `hcl` or `json` formats
- `-c` flag let you define the paths where your YAML files are located. Those paths can be either directories or files. If a path is a directory, files ending with `.yml` extension within the directory will be parsed. Multiple paths are separated by a comma

ex:  
`terrayaml convert -f hcl -c terrayaml/templates,terrayaml/main.yml`

# Anchors
The main feature this tool provides is YAML anchors.    
YAML anchors let you reuse pieces of definitions and override them if necessary.  
To let you define anchors that are not rendered in the final output, you can name your top level keys with `_template` in the name of the keys. Those will be removed from the final output.

ex:
```yaml
instance_template:
  instance: &instance_template
    ami: ${var.ami}
    instance_type: ${var.instance_type}

chef_provisionner_template:
  provisioner: &chef_provisioner
    chef:
      server_url: ${var.chef_server_url}
      environment: ${var.env}
      node_name: ${var.node_name}
      run_list: ${var.run_list}
      user_name: ${var.chef_username}

instance_with_provisioner_template:
  instance: &instance_with_provisioner
    <<: *instance_template
    provisioner: *chef_provisioner

resource:
  - aws_instance:
      instance1:
        <<: *instance_template
        ami: blah
        instance_type: t2.nano
        provisioner: *chef_provisioner
        aws_block_device:
          - device: /dev/xvdb
            size: 64
          - device: /dev/xvdc
            size: 128
  - aws_instance:
      instance2: *instance_with_provisioner
```

# YAML anchors vs Terraform modules
Terraform supports code reuse through modules but those are very limited. The main limitation is that you cannot use blocks in modules.  

As an example, when you define a resource representing an instance that only uses a root device, you cannot reuse any code to use the same instance definition and just add additionnal disks if needed. You have to duplicate code and you will rapidly get troubles managing this at scale.

# License
MIT
