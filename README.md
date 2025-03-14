Simple web app to get user info, their action count, possible next actions, and all users' referral indexes.

There are 4 REST endpoints:

GET /users/{id}
GET /users/{id}/actions/count
GET /users/{id}/actions/{actionType}/next
GET /referralIndexes

Endpoints don't have any authentication. App has in-memory storage only.

<h2>Usage</h2>

<h3>How to run</h3>

```
git clone https://github.com/lidiyatrf/user-actions-test-task.git
cd user-actions-test-task
go run cmd/main.go [-p <port>]
```
The port can be specified but is not required; the default is 8080.

<h3>How to verify</h3>

<h4>Get user by user id</h4>
```
curl --header "Content-Type: application/json" http://localhost:8080/users/1
```

Result:
```
{"id":1,"name":"Ferdinande","createdAt":"2020-07-14T05:48:54.798Z"}
```

<h4>Get user actions count</h4>
```
curl --header "Content-Type: application/json" http://localhost:8080/users/1/actions/count
```

Result:
```
{"count":49}
```

<h4>Get user next possible actions</h4>
```
curl --header "Content-Type: application/json" http://localhost:8080/users/100/actions/EDIT_CONTACT/next
```

Result:
```
{"ADD_CONTACT":0.33,"EDIT_CONTACT":0.33,"VIEW_CONTACTS":0.34}
```

<h4>Get users' referral indexes</h4>
```
curl --header "Content-Type: application/json" http://localhost:8080/referralIndexes
```

Result:
```
{"0":0,"1":1,"10":1,"100":0,"101":0,"102":0,"103":0,"104":3,"105":0,"106":0,"107":0,"108":0,"109":0,"11":1,"110":4,"111":0, ...}
```