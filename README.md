# Estuary Metrics Dashboard and API

API that wraps all Estuary objects for Metrics purposes.

# Components
## Available 
- [Estuary Metrics Core](./core/README.md) - the core metrics has all the core components to run metrics. This includes the pull metrics directly from a live estuary postgres database, heartbeat checks of all the servers, tracer and profiler.
- [Estuary Metrics REST Module](./rest/README.md) - this is the rest api endpoint layer that wraps the core functions. 

## WIP
- [Estuary Simple Metrics Dashboard (WIP)](./dashboard/README.md)
- [Estuary Metrics Notification API with Postgres Triggers(WIP)](./notification/README.md)
- [Estuary Metric CLI (WIP)](./cmd/README.md)
- [Estuary tracer and client (WIP)](#) - this is a tracer that estuary can use to push time series logs.


# API usage
This API is now being used on https://estuary.tech

![image](https://user-images.githubusercontent.com/4479171/204695762-515975c6-b4fa-4d6e-9fab-1df567187781.png)

![image](https://user-images.githubusercontent.com/4479171/204695822-e744f776-90be-4f04-a2b9-b06ed87b3345.png)

# Specs
[Specification](https://www.notion.so/ecosystem-wg/Metrics-Tracking-ea3da497096e4e4580c38a6a057b274f)
