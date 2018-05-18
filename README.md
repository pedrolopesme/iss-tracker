<h1 align="center">
    <img src="https://cdn.rawgit.com/pedrolopesme/iss-notifier/2e87fe31/docs/iss_notifier.png" alt="ISS Notifier" width="500">
     <br /> ISS Notifier
</h1>

<h4 align="center"> An AWS Lambda function to notify when the International Space Station (ISS) pass over an Earth coordinate </h4>

This project aims to create a serverless structure to detects whenever the ISS 
pass over an Earth coordinate (i.e. latitude/longitude) and send notifications 
about it.


### Implementation

More details about the project and its implementation soon.


### Dependency Manager

This project uses govendor as default dependency manager. So, whenever you
add or remove any dependency, you should do a `govendor add` or `govendor remove`.

There is a task in the Makefile to simplify this: `make vendoring`. 


### Makefile
 
This project provides a Makefile with all common operations need to test and build to .
 
 * setup: installs govendor package
 * vendoring: updates dependencies under vendor dir
 * test: runs all tests
 * fmt: runs gofmt for all go files
 
### Environment Vars

This application expects the default credential provider chain, as documented at
http://docs.aws.amazon.com/sdk-for-go/api/aws/defaults/#CredChain. The default provider
chain looks for credentials in the following order:

* Environment variables.
* Shared credentials file.
* If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.

This project does not contain any sensitive information hardcoded.

This project expects a env variable containing SQL url. If you're using 
a *nix, you should export the following env vars:

```shell
# AWS keys
export AWS_REGION=[VALUE]
export AWS_ACCESS_KEY_ID=[VALUE]
export AWS_SECRET_ACCESS_KEY=[VALUE]

# SQS URL
export ISS_SQS_URL=[VALUE]
```


### Credits

_soon_

### License
 
[MIT](LICENSE.md)
