<h1 align="center">
    <img src="https://cdn.rawgit.com/pedrolopesme/iss-notifier/2e87fe31/docs/iss_notifier.png" alt="ISS Notifier" width="500">
     <br /> ISS Notifier
</h1>

<h4 align="center"> An AWS Lambda function to notify when the International Space Station (ISS) pass over an Earth coordinate </h4>

This project aims to create a serverless structure to detects whenever the ISS 
pass over an Earth coordinate (i.e. latitude/longitude) and send notifications 
about it.

More details about the project and its implementation soon.


### Makefile
 
 This project provides a Makefile with all common operations need to test and build to .
 
 * build-lambda: generates AWS binary file, compress it and move to /dist dir
 * test: runs all tests
 * fmt: runs gofmt for all go files
 
### AWS Credentials

This application expects the default credential provider chain, as documented at
http://docs.aws.amazon.com/sdk-for-go/api/aws/defaults/#CredChain. The default provider
chain looks for credentials in the following order:

* Environment variables.
* Shared credentials file.
* If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.

This project does not contain any sensitve information hardcoded.

### License
 
 [MIT](LICENSE.md)
