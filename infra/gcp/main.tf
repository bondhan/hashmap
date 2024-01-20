# Bucket to store website
resource "google_storage_bucket" "website" {
  provider = google
  name = "hashmap-website-1234"
  location = "US"
}

# Upload html to bucket
resource "google_storage_object_access_control" "public_rule" {
  object = google_storage_bucket_object.static_site_source.output_name
  bucket = google_storage_bucket.website.name
  role   = "READER"
  entity = "allUsers"
}

# Object in the bucket
resource "google_storage_bucket_object" "static_site_source" {
  name = "index.html"
  source = "../../website/index.html"
  bucket = google_storage_bucket.website.name
}

# Reserve static external IP address
resource "google_compute_global_address" "website" {
  provider = google
  name = "website"
}

# Get the managed DNS Zone
data "google_dns_managed_zone" "gcp_coffeetime_dev" {
  provider = google
  name = "hashmap-website"
}

# Add the IP to the DNS
resource "google_dns_record_set" "website" {
  provider     = google
  name         = "website.${data.google_dns_managed_zone.gcp_coffeetime_dev.dns_name}"
  managed_zone = data.google_dns_managed_zone.gcp_coffeetime_dev.name
  rrdatas      = [google_compute_global_address.website.address]
  type         = "A"
  ttl          = "300"
}

# Add the bucket as a CDN backend
resource "google_compute_backend_bucket" "website-backend" {
  provider    = google
  name        = "website-backend"
  bucket_name = google_storage_bucket.website.name
  description = "Contains file needed for the website"
  enable_cdn = true
}

# Create URL Map
resource "google_compute_url_map" "website" {
  provider    = google
  name = "website-url-map"
  default_service = google_compute_backend_bucket.website-backend.self_link
  host_rule {
    hosts        = ["*"]
    path_matcher = "allpaths"
  }
  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_bucket.website-backend.self_link
  }
}

# GCP HTTP Proxy
resource "google_compute_target_http_proxy" "website" {
  provider    = google
  name    = "website-target-proxy"
  url_map = google_compute_url_map.website.self_link
}

# GCP Forwarding rule
resource "google_compute_global_forwarding_rule" "default" {
  provider    = google
  name   = "website-forwarding-rule"
  load_balancing_scheme = "EXTERNAL"
  ip_address = google_compute_global_address.website.address
  ip_protocol = "TCP"
  port_range = "80"
  target = google_compute_target_http_proxy.website.self_link
}