# IP Lookup

A simple Go utility for retrieving geolocation information from a public IP address using the ipinfo.io API.

## Features

* Lookup public IP addresses
* Display country, region, city, organization, and timezone
* Detect private/bogon IP addresses
* Built using the Go standard library

## Installation

```bash
git clone <repository-url>
cd ip-lookup
go run . 8.8.8.8
```

## Usage

```bash
go run . <ip_address>
```

Example:

```bash
go run . 8.8.8.8
```

Output:

```text
IP       : 8.8.8.8
Country  : US
Region   : California
City     : Mountain View
Org      : AS15169 Google LLC
Timezone : America/Los_Angeles
Location : 37.4056,-122.0775
```

## Private IP Addresses

Private IP addresses are detected and rejected.

Example:

```bash
go run . 192.168.1.10
```

Output:

```text
bogon IP address: 192.168.1.10
```

## Project Structure

```text
.
├── lookup.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Notes

* Geolocation results are approximate.
* VPNs and proxies may affect accuracy.
* Data is provided by ipinfo.io.
* Private IP addresses cannot be geolocated.

