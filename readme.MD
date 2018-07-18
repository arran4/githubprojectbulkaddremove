Simple program for bulk github operations, mostly adding and removing users from an organisations project.

Note; invitations don't seem to get revoked by the api like this. Also there is a small delay between running and impact.

## Usage

```
go run main.go -ghtoken <GITHUB PERSONAL API TOKEN>
```

WHich will pickup the file:
```
userchange.csv
```

## userchange.csv file contents

Which the contents can be either:
```
add user to repo,org4/project1,arran4
```
To add a user to a project

```
remove user from repo,org1/project2,arran4
```
To remove a user from an project

The order is:
```
<operation>,<org/repo>,<user>
```

## Other arguments: Changing the file

You can specify a different file using:
```
go run main.go -csv <CSV file> -ghtoken....
```

You can generate the Github personal api token using:
* https://github.com/settings/tokens

