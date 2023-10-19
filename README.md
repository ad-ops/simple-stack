# Simple Stack
Here is a simple web-stack that is simple, long term maintainable and yet powerful enough to do most things in.

- Frontend -> [htmx](https://htmx.org)
- Backend -> [Golang](https://go.dev)
- Database -> [SQLite](https://www.sqlite.org)

All these parts are some of the simplest in the technology stack among alternatives and have ambitions to keep backwards compatibility.

The bet is that in 10 years time (2034) it will still run, but also still make sense. Of course new technology has come, but it won't be strange the way it is done and someone should be able to work with the code.

## Use
```sh
go run
```

## Alternatives
For all the layers (front, back, db) in the stack there are alternatives built on the same technology. The most stable part is to use the most basic components to stay compatible, but I would argue that some parts will benefit from the added ergonomics while still having a path to migrate back when needed. For example instead of using Go's built in templating system you could switch to [templ]() which in my opinion offers much nicer developer experience, but at the cost of depending on something that might not be around after 10 years. This can be left to what innovation budget you have and if a project have specific requirements on for example performance.

#### Frontend
Htmx makes good use of templating to serve HTML and Go comes with built-in templating support. Although competent it is lacking in some areas for developer experience and typesafety. [Templ]() is a good alternative. Since it is built around HTML migrating to an alternative should not be too difficult.

Right now [Tailwind]() is a popular way of handling styling instead of pure CSS. This is also something that should be possible to integrate or swap out if needed.

#### Backend
Go comes with a built-in http server. There are many alternatives with different goals that might make them viable depending on the project.

There are a number of addons that might make sense depending on the project such as APIs, gRPC, error handling, observability, auth and more. The foundation offers this flexibility and could be migrated to or from when needed.

#### Database
For projects to get started quickly and get many features out of the box there is [PocketBase](https://github.com/pocketbase/pocketbase). It has a great number of features such as real-time, admin-console, APIs and auth. It is built on SQLite and can be used as a dependency in Go to create a single binary, making it very easy to get started and deploy anywhere.

When it comes to scalable solutions [Turso](https://github.com/tursodatabase) is a great alternative built on SQLite. It is a highly scalable service.

Both solutions create additional dependencies, but since they are open-sourced and are built on SQLite it is possible to migrate away from them if you find an alternative.

#### Developer Experience
Quick feedback-loop is important and with Go's fast compilation it is possible to restart and reload the website. This is for development and is not as important to be strict about for maintainability. 

The same goes for using [DevContainers] to quickly spin up and test. If it is part of a more complex setup projects such as [Skaffold] can be used.

htmx, Go, SQLite all are very simple to deploy since they all use single files and Go has great support for different platforms. Running locally, on a VM, in a container or as WebAssembly are all possible.