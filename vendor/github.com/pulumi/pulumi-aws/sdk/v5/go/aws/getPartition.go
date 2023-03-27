// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to lookup information about the current AWS partition in
// which the provider is working.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetPartition(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Actions: []string{
//							"s3:ListBucket",
//						},
//						Resources: []string{
//							fmt.Sprintf("arn:%v:s3:::my-bucket", current.Partition),
//						},
//						Sid: pulumi.StringRef("1"),
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPartition(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetPartitionResult, error) {
	var rv GetPartitionResult
	err := ctx.Invoke("aws:index/getPartition:getPartition", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getPartition.
type GetPartitionResult struct {
	// Base DNS domain name for the current partition (e.g., `amazonaws.com` in AWS Commercial, `amazonaws.com.cn` in AWS China).
	DnsSuffix string `pulumi:"dnsSuffix"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Identifier of the current partition (e.g., `aws` in AWS Commercial, `aws-cn` in AWS China).
	Partition string `pulumi:"partition"`
	// Prefix of service names (e.g., `com.amazonaws` in AWS Commercial, `cn.com.amazonaws` in AWS China).
	ReverseDnsPrefix string `pulumi:"reverseDnsPrefix"`
}
