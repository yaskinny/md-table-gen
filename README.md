Table MD Gen
===
You can use this tool to generate documents for ansible roles variables, helm values files and ... .  
You annotate your values files and it will generate table for your values in markdown.

Build
---
```bash
git clone https://github.com/yaskinny/md-table-gen --single-branch --branch=master --depth=1
go build -o bin/md-table-gen ./cmd/

## if you want to use it only for your user move it somewhere writable in your PATH for example ~/.local/bin
mv -v bin/md-table-gen ~/.local/bin/

## or for all system users
sudo mv -v bin/md-table-gen /usr/local/bin/
```
annotations
---
```bash
## optional option/property
@opt optName description

## mandatory option/property
@mand optName description

## becomes a section name in MD file like **sectionName:**
@section sectionName


## mandatory object 
@obj_mand objName description

## optional object
@obj objName description

## mandatory property/option in an object
@obj_mand objName.optName description

## optional property/option in an object
@obj objName.optName description


```


Example
---
Example with two different values files:  
`file1.yaml`:
```yaml
## @opt user user to run the command with
user: yaser

## @section database
## @mand db_address database address
db_address: 'localhost'
## @mand db_port database port
db_port: 3306

## @obj_mand db_creds object to set database credentials
## @obj_mand db_creds.user database username
## @obj_mand db_creds.password database password
## @obj db_creds.tls_path path to client tls bundle
db_creds:
  user: yaser
  password: '1234'
```

`file2.yaml`:
```yaml
## @opt hostname hostname to set
hostname: test

## @mand timezone timezone to set for the host
timezone: UTC
```

```bash
## -f files to generate tables from.
## -r file to write generated output.
## -n section name in md file to use for output.
md-table-gen -f file1.yaml -f file2.yaml -r README.md -n Values

```
  above command will create markdown output from `file1.yaml` and `file2.yaml` and write it to `README.md` file. Also it will create a backup file of your `README.md` in your os default temporary directory ( for example `/tmp` in linux ) in case something goes wrong and You are not using version control.

**output:**
```
## Values
### file1.yaml
| name | description |
| --- | --- |
| user | `user to run the command with` |

**database:**
| name | description |
| --- | --- |
| **db_address** | `database address` |
| **db_port** | `database port` |

***db_creds***: object to set database credentials
| name | description |
| --- | --- |
| **user** | `database username` |
| **password** | `database password` |

### file2.yaml
| name | description |
| --- | --- |
| hostname | `hostname to set` |
| **timezone** | `timezone to set for the host` |

```

## TODO
- [ ] Add output to stdout instead of forcing using a file
- [ ] Add more tests
