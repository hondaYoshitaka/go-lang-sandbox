## Usage
### docker
1. Build docker image
```bash
% docker build -t ubuntu-ssh/go-lang .
```

2. Run docker
```bash
$ docker run -d -P --name go-lang-sandbox ubuntu-ssh/go-lang
```

3. Test `ssh` command
```bash
$ docker port go-lang-sandbox 22
  0.0.0.0:32769
$ ssh root@localhost -p 32769
```