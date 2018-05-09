<h1 align="center">
    <img src="https://cdn.rawgit.com/pedrolopesme/iss-notifier/34ac93b7/docs/iss_notifier.png" alt="ISS Notifier" width="400">
     <br /> ISS Notifier
</h1>

<h4 align="center"> An AWS Lambda function to notify when the International Space Station (ISS) pass over an Earth coordinate </h4>

This project aims to create a serverless structure to detects whenever the ISS 
pass over an Earth coordinate (i.e. latitude/longitude) and send notifications 
about it.

More details about the project and its implementation soon.


### Makefile
 
 This project provides a Makefile with all common operations need to test and build to .
 
 * build: generates AWS binary file, compress it and move to /dist dir
 * test: runs all tests
 * fmt: runs gofmt for all go files
 
 
### License
 
 [MIT](LICENSE.md)