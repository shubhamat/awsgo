# awsuptime.go
awsuptime.go  is a utility that prints the time duration an instance has been running for.  It also prints the Public IP of the instance.

E.g.

sam@saminux ~/awsgo $ ./awsuptime  
PublicIpAddress         UpTime  
5x.y.z.5                11m46.65482121s  

# buckets.go
buckets.go is another hello-world isuqe utility that lists the S3 buckets.   
sam@saminux ~/awsgo $ go run buckets.go  
Buckets  
saminuxbucket1  

