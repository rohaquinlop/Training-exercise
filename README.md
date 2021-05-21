# Training Exercise

## GraphDB

start server
```bash
cd backend
go run server.go
```

Buyers endpoint:
```
https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=$date
$date unix timestamp format, converter, hours don't matter
```

```bash
Replace ^@ character
sed 's/\x0/;/g' transactions.txt > newtransactions.txt
sed 's/;;/\n/g' newtransactions.txt > newtransactions1.txt
```
