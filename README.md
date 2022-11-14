# Estuary Metrics Dashboard and API

API that wraps all Estuary objects for Metrics purposes.

# Components
- [Estuary Metrics Core](./core/README.md)
- [Estuary Metrics REST Module](./rest/README.md)
- [Estuary Simple Metrics Dashboard](./dashboard/README.md)
- [Estuary Metrics Notification API with Postgres Triggers(WIP)](./notification/README.md)
- [Estuary Metric CLI (WIP)](./cmd/README.md)

# Access Key
In order to access the API, you need an elevated access (perm > 2) Estuary API key.

# Endpoints
## Objects
- These are individual object APIs that can be used to query for a specific estuary table. 

## Stats
- These are stats APIs that has some aggregation logic to query for a specific estuary table.
## System / Devices
- These are system/device APIs to look up Equinix/AWS environment specific information
## Blockstore (WIP)
- These are blockstore information
## Reporting (WIP)
- These are push notification APIs that can be used to send notifications to users

# Public Dashboard (WIP)
The public dashboard is available at [https://metrics.estuary.tech](https://metrics.estuary.tech)

# Specs
[Specification](https://www.notion.so/ecosystem-wg/Metrics-Tracking-ea3da497096e4e4580c38a6a057b274f)