# Estuary Metrics Dashboard

Estuary Metrics Dashboard. This will show the following:

### **System**

- [ ]  Total objects pinned (Query) 
- [ ]  Total TiBs uploaded (Query)
- [ ]  Total TiBs sealed data on Filecoin
- [ ]  Available free space
- [ ]  Total space capacity
- [ ]  Downtime
- [ ]  Performance 

### **Users**

- [ ]  Total number of Storage Providers.
- [ ]  Total number of users.
- [ ]  Ongoing user activity — DAUs, WAUs, MAUs etc. Are users coming back? (custom Grafana plugin)
- [ ]  For Storage/Retrieval deal metrics, in addition to aggregate, we also want the following breakdowns
    - [ ]  per day breakdown.
    - [ ]  per week breakdown.
    - [ ]  per provider breakdown.

### **Storage**

- [ ]  Storage Deal Success Rate (Success % / All Deals)
- [ ]  Storage Deal Acceptance Rate (Success % / Accepted Deals)
    - [ ]  Total number of storage deals proposed  (Total Deals / Proposed)
    - [ ]  Total number of storage deal proposals accepted (Total Deals / Accepted Deals)
    - [ ]  Total number of storage deal proposals rejected (Total Deals / Rejected Deals)
- [ ]  Total number of storage deals attempted
    - [ ]  Total number of successful deals
    - [ ]  Total number of failed deals
- [ ]  Distribution of data size uploaded per user
- [ ]  Performance metrics
    - [ ]  Time to a successful deal
        - [ ]  how does that scale with data size?

### **Retrieval**
- [ ]  Retrieval Deal Success Rate
- [ ]  Retrieval Deal Acceptance Rate
    - [ ]  Total number of retrieval deals proposed
    - [ ]  Total number of retrieval deal proposals accepted
    - [ ]  Total number of retrieval deal proposals rejected
- [ ]  Total number of retrieval deals attempted (per day and per week breakdown)
    - [ ]  total number of successful retrievals
    - [ ]  total number of failed retrievals
- [ ]  Deals Failed Because Of Undiablable Miners
- [ ]  Time To First Byte (retrieval deals)

## Build
```
npm run build
```

## Install
```
npm install
```

## Server
```
npm run start
```
