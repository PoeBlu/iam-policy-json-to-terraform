# iam-policy-json-to-terraform

[![Build Status](https://travis-ci.org/flosell/iam-policy-json-to-terraform.svg?branch=master)](https://travis-ci.org/flosell/iam-policy-json-to-terraform)[![Release](https://img.shields.io/github/release/flosell/iam-policy-json-to-terraform.svg)](https://github.com/flosell/iam-policy-json-to-terraform/releases)

Small tool to convert an IAM Policy in JSON format into a Terraform [`aws_iam_policy_document`](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)

## Installation

### OSX

```bash
$ brew install iam-policy-json-to-terraform
```
    
### Other

Download the latest binary from the [releases page](https://github.com/flosell/iam-policy-json-to-terraform/releases) and put it into your `PATH` under the name `iam-policy-json-to-terraform`

### Developer

If you're a go developer and have your `GOPATH` defined and have added your `$GOPATH/bin` directory to your path, you can simply run this command.
```bash testcase=usage
$ go get github.com/flosell/iam-policy-json-to-terraform
```

## Usage

From raw JSON

```bash testcase=usage
$ echo '{"Statement":[{"Effect":"Allow","Action":["ec2:Describe*"],"Resource":"*"}]}' | iam-policy-json-to-terraform
data "aws_iam_policy_document" "policy" {
  statement {
    sid       = ""
    effect    = "Allow"
    resources = ["*"]
    actions   = ["ec2:Describe*"]
  }
}
```

From a JSON policy file

```bash
$ iam-policy-json-to-terraform < some-policy.json
```

## Local development

### Prerequisites

* Install [`dep`](https://golang.github.io/dep/): 
  ```bash testcase=building
  $ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
  ```
  
* Clone the repository to the right location in the`$GOPATH`: 
  ```bash testcase=building
  $ mkdir -p $GOPATH/src/github.com/flosell/
  $ cd $GOPATH/src/github.com/flosell/
  $ git clone  git@github.com:flosell/iam-policy-json-to-terraform.git
  ```

* Install dependencies and tools: 
  ```bash testcase=building
  $ cd $GOPATH/src/github.com/flosell/iam-policy-json-to-terraform
  $ make vendor tools
  ```
  
### Development

`make` is your primary point of entry for any development activity. Call it without arguments to learn more: 

```bash testcase=building
$ make
build                          Test and build the whole application
clean                          Remove build artifacts and vendored dependencies
fmt                            Format code
fmtcheck                       Run linter
test                           Run all tests
tools                          Install additional required tooling
vendor                         Install dependencies into ./vendor 
```
