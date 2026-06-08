# stop-db.ps1
docker stop my-postgres
docker rm my-postgres
Write-Host "PostgreSQL stopped and removed"