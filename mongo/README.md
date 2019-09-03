# Enter the Container w/ Bash

`docker-compose exec --user root db /bin/bash`

# Enter the Mongo Shell

`docker-compose exec --user root db mongo`

# Create a DB w/ User

While in the shell, run the following:

```
use admin;
db.auth("user", "password");
use myDatabase;
db.createUser({user: "user", pwd: "password", roles:[{role: "readWrite" , db:"myDatabase"}]});
```
