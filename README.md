# news-website

Scalable news website on AWS.

## Getting started

Create a code commit repository:

```bash
aws cloudformation \
  --profile training \
  --region eu-central-1 create-stack \
  --stack-name code-commit-playground \
  --template-body file:///${PWD}/cf.yaml \
  --parameters ParameterKey=RepositoryName,ParameterValue=playground
```

Store the code:

```bash
git clone https://git-codecommit.eu-central-1.amazonaws.com/v1/repos/playground
cd playground
for f in $(ls ../ | grep -v playground); do cp -r ../${f} .; done
cp ../.gitignore .
git config user.name hello
git config user.email world
git add .
git commit -m "helloworld"
git push origin master
```
