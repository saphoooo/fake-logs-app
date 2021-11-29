# Fake logs

This application don't do anything but generating random logs!

## How to use it

### Build locally

```bash
$ git clone https://github.com/saphoooo/fake-logs-app.git
$ cd fake-logs-app
$ go build
$ ./fake-logs-app
```

By default, it will produce a custom log every 5 seconds:

```bash
[INFO] 2021/11/25 16:40:00 172.17.0.4 POST http://example.com/status 200 -- 106.26.105.79 Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15
```

However, you can change this behavior with args:

```bash
$ ./fake-logs-app -h
Usage of ./fake-logs-app:
  -i int
    	log interval in seconds (default 5)
  -f string
    	log format syle (nginx|custom) (default "custom")
$ ./fake-logs-app -i 1 -f nginx
172.17.0.3 - - [2021/11/25:16:40:10 +0000] "PUT /api HTTP/1.1" 200 621 "http://example.com/" "curl/7.74.0" "89.77.53.67"
172.17.0.3 - - [2021/11/25:16:40:12 +0000] "POST /logout HTTP/1.1" 200 364 "http://example.com/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0" "89.121.60.251"
172.17.0.3 - - [2021/11/25:16:40:14 +0000] "GET /api HTTP/1.1" 301 200 "http://example.com/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0" "251.219.14.168"
172.17.0.3 - - [2021/11/25:16:40:16 +0000] "GET /status HTTP/1.1" 200 821 "http://example.com/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36" "242.48.39.127"
```

### Docker|Podman

```
$ docker run saphoooo/fake-logs-app
```

### Kubernetes

There is an example `deployment.yaml` file with two containers and different options:

```yaml
containers:
- name: custom
  image: saphoooo/fake-logs-app:latest
  args: ['-i', '1']
- name: nginx
  image: saphoooo/fake-logs-app:latest
  args: ['-f', 'nginx', '-i', '2']
```