package mongodb

var collectionNames = struct {
	User              string
	Monitor           string
	HealthCheckRecord string
}{
	User:              "users",
	Monitor:           "monitors",
	HealthCheckRecord: "health-check-records",
}
