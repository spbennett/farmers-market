# FarmersMarket

## About
The **farmers-market** app is webserver implementing a grocery checkout with support for processing discounts.  It is implemented in Golang.

## Build
This app can be distributed as a docker container.  Use the included build script to run the app in a container.
```shell
./build.sh
```

If successful, the docker container will be launched by the docker daemon on the localhost.  Try to hit the endpoint with the usage instructions below.

## Usage

Once the app is running, you can use the following two functions.

### Add

Add items to your shopping basket with the **/add** path.
```shell
curl localhost:8080/add -d '{"id": "CF1"}'
```

### Checkout

When you are done adding items, you can checkout and see your discounts applied using the **/checkout** path.
```shell
curl localhost:8080/checkout
Item		Price
----		----
CF1 	 	 11.23
	  	
-----------------------------------
Total:		 11.23

```

## Testing

### Units Tests
Units tests are implemented for the main *Market* class.  Execute them with the following command:
```shell
go test
```

### Functionality Testing

Tested on Ubuntu 18.04 and OSX 10.13.5 with the following:
- Golang v1.10
- Docker v18.03
