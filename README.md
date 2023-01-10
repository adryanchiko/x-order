# X-Order

## Stacks
There are several stacks we've been used for build this apps. Before running the app, make sure to have these programming language toolchains installed (and/or services running).

### Backend
- [Golang](https://go.dev/) : Main language to develop this app.
- [Echo](https://echo.labstack.com/) : Golang web framework.
- [PostgreSQL](https://www.postgresql.org/) : Main database for the application.

### Frontend
- [NodeJS](https://www.npmjs.com/get-npm) : For building the frontend app.
- [VueJS](https://vuejs.org/) : Javascript framework to develop frontend app.
- [Vuetify](https://vuetifyjs.com/en/) : CSS framework.
- [Yarn](https://classic.yarnpkg.com/en/) : As dependency management.

## Getting Started
Make sure already install all dependencies above to your machine.

Clone this project, locate to your `GOPATH` folder by running this commands:
```
cd $GOPATH/src/github.com/adryanchiko

git clone git@github.com:adryanchiko/x-order.git
```

### Setup Backend

Go to folder `x-order/service/order-app/` and then fetch golang packages by running `go mod tidy`

```
cd $GOPATH/src/github.com/adryanchiko/x-order/service/order-app

go mod tidy
```

Go to folder `cmd/order` to update the config file that is `order.yaml`, make sure PostgreSQL address and setup server is already ok.

Then, build the app!

```
cd cmd/order

...
-- update order.yaml
-- notes: for server port DO NOT EDIT! because in frontend app still using static url to call API service
...

go build
```

After build app successfully, initialize database schema using `init-db` command
```
./order init-db
```

Then migrate data from CSV file into database using `migrate` command
```
-- for company
./order migrate company --file ${path/file.csv}

-- for customer
./order migrate customer --file ${path/file.csv}

-- for order
./order migrate order --file ${path/file.csv}

-- for order item
./order migrate order-item --file ${path/file.csv}

-- for delivery
./order migrate delivery --file ${path/file.csv}
```

Then run the order app!
```
./order
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8000
```

### Setup Frontend

Go to folder `x-order/frontend` and then install dependencies using `yarn`

```
cd $GOPATH/src/github.com/adryanchiko/x-order/frontend

yarn install
```

Then run the frontend app!

```
-- dev mode
yarn dev

  VITE v3.2.5  ready in 626 ms

  ➜  Local:   http://localhost:3000/
  ➜  Network: use --host to expose

```