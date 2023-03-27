package main

import (
	"github.com/phillipedwards/pulumi-notes"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}

		_ = sum.GetSum(1, 2)

		out := pulumi.JSONMarshal(bucket.Arn)

		_, err = random.NewRandomId(ctx, "rrr", &random.RandomIdArgs{
			ByteLength: pulumi.Int(4),
		})

		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		ctx.Export("json-out", out)
		return nil
	})
}
