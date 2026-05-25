# StartTech Operations Runbook

## Restart Backend

SSH into EC2:

```bash
ssh -i starttech.pem ubuntu@<EC2_PUBLIC_IP>
```

Restart container:

```bash
docker restart app
```

---

## View Logs

```bash
docker logs app
```

CloudWatch Logs:
- /starttech/backend
- /aws/ec2/starttech-app

---

## Rollback Deployment

```bash
./scripts/rollback.sh
```

---

## Health Check

```bash
./scripts/health-check.sh
```

---

## Common Issues

### Docker Container Not Running

Check:

```bash
docker ps -a
docker logs app
```

---

### MongoDB Connection Error

Verify:
- MONGO_URI
- Security groups
- Atlas IP whitelist

---

### GitHub Actions Deployment Failure

Check:
- DockerHub credentials
- AWS credentials
- SSH key secrets

---

### CloudWatch Logs Missing

Restart CloudWatch agent:

```bash
sudo systemctl restart amazon-cloudwatch-agent
```
