provider "couchdb" {
url = "http://localhost:8080"
}
resource "couchdb_database" "foo" {
    name = "database"
}