# Enter the Container w/ Bash

`docker-compose exec --user root db /bin/bash`

# Enter the Mongo Shell

`docker-compose exec --user root db mongo`

# Create a DB w/ User

While in the shell, run the following:

```
use admin;
db.auth("root", "password");
use myDatabase;
db.createUser({user: "root", pwd: "password", roles:[{role: "readWrite" , db:"myDatabase"}]});
```
