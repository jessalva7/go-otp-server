# Go OTP Server

## Create a file and set the environment variables  

```
TWILIO_URL=https://api.twilio.com/2010-04-01/Accounts/
TWILIO_SID=:)
PORT=8080
TWILIO_NUMBER=:P
TWILIO_AUTH_TOKEN=B)
```

---

## Use docker run and set environment variable using the above defined file for running in port 8000

```
docker run --publish 8000:8080 --env-file ./GoOTPEnv.list --detach --name go-otp jessalva35/go-otp-server:latest
```
