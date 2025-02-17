
```sh
use menter_mini
db.createUser({user: "mcenter", pwd: "123456", roles: [{ role: "dbOwner", db: "menter_mini" }]})
```