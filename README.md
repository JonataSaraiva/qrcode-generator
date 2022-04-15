# Go Script to generate QRCODE with url redirect

This app is using go-qrcode lib to create QRCODE with follow parameters:

```
host = URL of file or app
Id = resource's id
```
Example:
```
http://domain.s3-website-us-east-1.amazonaws.com/1112455.png
```

You need set up a output path to img, changing it using follow parameter:
```
filePath := "/home/user/Documentos/QR.png"
```