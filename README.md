```
docker-compose up
curl localhost:8085/api/pet/start --data "Joao"

recepcionist_1  | Service created 05103aa6-3f11-4f60-a17a-380781f95246 for pet Joao
specialist_1    | Caring for service {"Name":"Joao","ServiceId":"05103aa6-3f11-4f60-a17a-380781f95246","Status":"Brushed","History":[
    {"Job":"Received","Date":"2022-07-27 14:52:43.205125332 +0000 UTC m=+56.062910822"},
    {"Job":"Brushed","Date":"2022-07-27 14:52:43.21155541 +0000 UTC m=+53.856411479"}]}
specialist_1    | Caring for service {"Name":"Joao","ServiceId":"05103aa6-3f11-4f60-a17a-380781f95246","Status":"Washed","History":[
    {"Job":"Received","Date":"2022-07-27 14:52:43.205125332 +0000 UTC m=+56.062910822"},
    {"Job":"Brushed","Date":"2022-07-27 14:52:43.21155541 +0000 UTC m=+53.856411479"},
    {"Job":"Washed","Date":"2022-07-27 14:52:43.216638817 +0000 UTC m=+53.861494886"}]}
specialist_1    | Caring for service {"Name":"Joao","ServiceId":"05103aa6-3f11-4f60-a17a-380781f95246","Status":"Dried","History":[
    {"Job":"Received","Date":"2022-07-27 14:52:43.205125332 +0000 UTC m=+56.062910822"},
    {"Job":"Brushed","Date":"2022-07-27 14:52:43.21155541 +0000 UTC m=+53.856411479"},
    {"Job":"Washed","Date":"2022-07-27 14:52:43.216638817 +0000 UTC m=+53.861494886"},
    {"Job":"Dried","Date":"2022-07-27 14:52:43.221055181 +0000 UTC m=+53.865911250"}]}
specialist_1    | Caring for service {"Name":"Joao","ServiceId":"05103aa6-3f11-4f60-a17a-380781f95246","Status":"Complete","History":[
    {"Job":"Received","Date":"2022-07-27 14:52:43.205125332 +0000 UTC m=+56.062910822"},
    {"Job":"Brushed","Date":"2022-07-27 14:52:43.21155541 +0000 UTC m=+53.856411479"},
    {"Job":"Washed","Date":"2022-07-27 14:52:43.216638817 +0000 UTC m=+53.861494886"},
    {"Job":"Dried","Date":"2022-07-27 14:52:43.221055181 +0000 UTC m=+53.865911250"},
    {"Job":"Complete","Date":"2022-07-27 14:52:43.225620471 +0000 UTC m=+53.870476539"}]}
specialist_1    | Caring for service {"Name":"Joao","ServiceId":"05103aa6-3f11-4f60-a17a-380781f95246","Status":"Complete","History":[
    {"Job":"Received","Date":"2022-07-27 14:52:43.205125332 +0000 UTC m=+56.062910822"},
    {"Job":"Brushed","Date":"2022-07-27 14:52:43.21155541 +0000 UTC m=+53.856411479"},
    {"Job":"Washed","Date":"2022-07-27 14:52:43.216638817 +0000 UTC m=+53.861494886"},
    {"Job":"Dried","Date":"2022-07-27 14:52:43.221055181 +0000 UTC m=+53.865911250"},
    {"Job":"Complete","Date":"2022-07-27 14:52:43.225620471 +0000 UTC m=+53.870476539"}]}
```