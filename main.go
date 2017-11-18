package main

import (
        "fmt"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/credentials"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
        sess, err := session.NewSession(&aws.Config{
                Region:      aws.String("us-east-1"),
                Credentials: credentials.NewSharedCredentials("", ""),
        })
        if err != nil {
                fmt.Println(err)
        }
        svc := ec2.New(sess)
        input := &ec2.DescribeAvailabilityZonesInput{}
        result, err := svc.DescribeAvailabilityZones(input)
        if err != nil {
                if aerr, ok := err.(awserr.Error); ok {
                        switch aerr.Code() {
                        default:
                                fmt.Println(aerr.Error())
                        }
                } else {
                        // Print the error, cast err to awserr.Error to get the Code and
                        // Message from an error.
                        fmt.Println(err.Error())
                }
                return
        }
        fmt.Println(result)
}