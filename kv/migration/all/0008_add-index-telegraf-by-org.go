package all

import (
	"github.com/influxdata/influxdb/v2/kv"
	"github.com/influxdata/influxdb/v2/telegraf"
)

// Migration0008_AddIndexTelegrafByOrg adds the index telegraf configs by organization ID
var Migration0008_AddIndexTelegrafByOrg = kv.NewIndexMigration(telegraf.ByOrganizationIndexMapping, kv.WithIndexMigrationCleanup)
