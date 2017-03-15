# Build for linux
env GOOS=linux GOARCH=386 go build -v goletsencrypt

# Edit letsencrypt.secret and serve run on your server

# Transferring file to ec2
scp -i aws_key.pem file.ext ec2-user@ ec2-52-32-134-226.us-west-2.compute.amazonaws.com:~/data/[new_file_name]