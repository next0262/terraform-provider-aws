package rds_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/rds"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccRDSClustersDataSource_filter(t *testing.T) {
	var dbCluster rds.DBCluster
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_rds_clusters.test"
	resourceName := "aws_rds_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClustersDataSourceConfig_filter(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(resourceName, &dbCluster),
					resource.TestCheckResourceAttr(dataSourceName, "cluster_arns.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_arns.0", resourceName, "arn"),
					resource.TestCheckResourceAttr(dataSourceName, "cluster_identifiers.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cluster_identifiers.0", resourceName, "cluster_identifier"),
				),
			},
		},
	})
}

func testAccClustersDataSourceConfig_filter(rName string) string {
	return fmt.Sprintf(`
resource "aws_rds_cluster" "test" {
  cluster_identifier              = %[1]q
  database_name                   = "test"
  master_username                 = "tfacctest"
  master_password                 = "avoid-plaintext-passwords"
  skip_final_snapshot             = true
}

resource "aws_rds_cluster" "wrong" {
  cluster_identifier              = "wrong-%[1]s"
  database_name                   = "test"
  master_username                 = "tfacctest"
  master_password                 = "avoid-plaintext-passwords"
  skip_final_snapshot             = true
}

data "aws_rds_clusters" "test" {
  filter {
	name   = "db-cluster-id"
	values = [aws_rds_cluster.test.cluster_identifier]
  }
}
`, rName)
}
