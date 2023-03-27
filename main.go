package main

import (
	"github.com/phillipedwards/pulumi-notes"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)

		_ = sum.GetSum(1, 2)

		_, err := random.NewRandomId(ctx, "rrr", &random.RandomIdArgs{
			ByteLength: pulumi.Int(4),
		})

		if err != nil {
			return err
		}

		// Export the name of the bucket
		return nil
	})
}
