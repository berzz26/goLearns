<h1>gRPC playground</h1>
Okay so this project is a playground for me to understand what exactly  is.

We will mimic a service of [computebay](https://computebay.online) to learn this protocol.

Imagine this

    Backend
      |
      |
      |
    Worker A (someone's PC)
    Worker B (someone's PC)
    Worker C (someone's PC)

The backend asks to those nodes:

- `How much ram do you have`<br>
- `Launch a docker container`<br>
- `Kill this job`<br>

so,<br>
*How do two programs running on different machines talk?*<br>

## *That's the entire problem space.*

## The naive way: *REST*

the worker exposes :

- POST /hearbeat
- POST /start-job
- POST /stop-job
- GET /logs

<br>Backend does:

<br>
`http.Post(...)`

<br>Worker receives JSON:

    {
    "workerId": "abc",
    "ram": 32000
    }

This works.

## So why was gRPC invented?

<br> at scale, say you have multiple microservices:
    Service A
    Service B
    Service C
    Service D
    Service E

Then, most systems spend time in:

- Writing JSON structs
- Writing handlers
- Writing clients
- Parsing JSON
- Validating JSON

so imagine 100 microservice and each service having has

- `type user Struct{}`
- `interface User{}`
- `Class User {}`

if someone forgets to update any of the following while implementing a new feature, everything breaks.

## Key idea of gRPC

Instead of:

- Write server
- Write client
- Write JSON
- Write docs

Hoping they match,

we write:

<strong>ONE CONTRACT FILE</strong>

and *generate* everything.

That contract file is called:

`.proto`

Proto = *Protocol Buffer definition.*

That's all it is.
A contract.

A `.proto` file can be taken as an interface in go

But language independant.
we write:

    message User {
    string id = 1;
    string name = 2;
    }

This says:

"A User object contains an id and a name."

Thats it, no servers, no networking no clients nothing.<br>
just a defination

We may think
Why not just use Go structs?

Because maybe:

Backend has:

`type User struct`

Frontend has:

`interface User`

Python worker defines:

`class User`

Now we have 3 definitions to maintain everytime.

They drift apart.

Instead:

    message User {
    string id = 1;
    string name = 2;
    }

Generate all 3 automatically.


### what is a `message`?
A message is basically a `struct` in go
This:

    message User {
    string id = 1;
    string name = 2;
    }


    what are these numbers?
    basically, this is how gRPC is faster. it doesnt send stuff like {"id" : "123","name":"bjp"} across netowrk.
    
    instead it sends a compact binary format(1,2,3..)
    which is then internally mapped as:
    field 1 -> 123
    field 2 -> bjp

    this is much much smaller than json
    
this later becomes:

    type User struct {
        Id   string
        Name string
    }

after **generation**.
hence,
`message ~= struct`

That's one reason gRPC is popular for service-to-service communication.

### what is a service: 

it kinda resonates with go interface

    type UserService interface {
        GetUser(id string) User
    }

**Proto version:**

    service UserService {
    rpc GetUser(GetUserRequest)
        returns (User);
    }

this means that there exist a remote function called `GetUser`

now, imagine a handler doing this:

`user := service.GetUser("123")`

usually that is a local method inside service

with gRPC

`user:= client.GetUser("123")`

looks local, but actually, 

    network request
        |
    internet
        |
    another machine
        |
    execute
        |
    response

everything happened

This is why it's called:

    Remote Procedure Call

`RPC`.

Calling a procedure on another computer.

## Implementation

here, in the /proto dir, we have a worker.proto file that creates a specification of how everything should look like, it has messages and services

`protoc` this is a **protocol buffer compiler**

when we run `protoc --go_out=. user.proto`

protoc then basically generates go code from the given schema.

# What are we building?

We're building this:

    +----------------------+
    |   Control Plane      |
    |  (gRPC Server)       |
    +----------+-----------+
            ^
            |
            | Heartbeat
            |
    +----------+-----------+
    |    Worker Agent      |
    |   (gRPC Client)      |
    +----------------------+

The worker will send:

worker-1 is alive
- cpu = 45%
- ram = 60%

every few seconds.

The control plane receives it.

This is basically the first building block of ComputeBay.