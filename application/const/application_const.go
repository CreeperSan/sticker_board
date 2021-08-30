package StickerBoard

const prefix = "application_"

const SPMySQLDatabaseName = prefix + "mysql_database_name"
const SPMySQLDatabaseUserName = prefix + "mysql_database_username"
const SPMySQLDatabasePassword = prefix + "mysql_database_password"
const SPMySQLDatabasePort = prefix + "mysql_database_port"
const SPMySQLDatabaseAddress = prefix + "mysql_database_address"

const SPMongoDBDatabaseName = prefix + "mongodb_database_name"
const SPMongoDBDatabaseAddress = prefix + "mongodb_database_address"
const SPMongoDBDatabasePort = prefix + "mongodb_database_port"
const SPMongoDBDatabaseUsername = prefix + "mongodb_database_username"
const SPMongoDBDatabasePassword = prefix + "mongodb_database_password"

const SPOSSAlicloudEndpoint = prefix + "oss_alicloud_endpoint"
const SPOSSAlicloudAccessKeyID = prefix + "oss_alicloud_access_key_id"
const SPOSSAlicloudAccessKeySecret = prefix + "oss_alicloud_access_key_secret"
const SPOSSAlicloudBucket = prefix + "oss_alicloud_bucket"

const ExitCodeNormal = 0
const ExitCodeDatabaseCreateClientConnectionFailed = 1
const ExitCodeDatabasePingFailed = 2
const ExitCodeDatabaseCreateCollectionFailed = 3
const ExitCodeAlicloudOSSCollectionFailed = 4
const ExitCodeAlicloudOSSBucketCollectionFailed = 5